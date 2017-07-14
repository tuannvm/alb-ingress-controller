package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/coreos/alb-ingress-controller/controller"
	"github.com/coreos/alb-ingress-controller/controller/config"
	"github.com/coreos/alb-ingress-controller/log"
	ingresscontroller "k8s.io/ingress/core/pkg/ingress/controller"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.CommandLine.Parse([]string{})

	logLevel := os.Getenv("LOG_LEVEL")
	log.SetLogLevel(logLevel)

	awsDebug, _ := strconv.ParseBool(os.Getenv("AWS_DEBUG"))

	conf := &config.Config{
		ClusterName: os.Getenv("CLUSTER_NAME"),
		AWSDebug:    awsDebug,
	}

	port := "8080"
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	ac := controller.NewALBController(&aws.Config{MaxRetries: aws.Int(15)}, conf)
	ic := ingresscontroller.NewIngressController(ac)

	ac.IngressClass = ic.IngressClass()
	if ac.IngressClass != "" {
		log.Infof("Ingress class set to %s", "controller", ac.IngressClass)
	}

	http.HandleFunc("/state", ac.StateHandler)

	if *ac.ClusterName == "" {
		glog.Exit("A cluster name must be defined")
	}

	if len(*ac.ClusterName) > 11 {
		glog.Exit("Cluster name must be 11 characters or less")
	}

	defer func() {
		glog.Infof("Shutting down ingress controller...")
		ic.Stop()
	}()
	ic.Start()
}
