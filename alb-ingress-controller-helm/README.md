# alb-ingress-controller

[alb-ingress-controller](https://github.com/coreos/alb-ingress-controller) satisfies Kubernetes ingress resources by provisioning Application Load Balancers and Route53 Resource Record Sets.

## TL;DR:

```console
$ helm registry install quay.io/coreos/alb-ingress-controller-helm
```

## Introduction

This chart bootstraps an alb-ingress-controller deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites
  - Kubernetes 1.4+ with Beta APIs enabled
  - [Helm Registry plugin](https://github.com/app-registry/helm-plugin)

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm registry install quay.io/coreos/alb-ingress-controller-helm --name=my-release
```

The command deploys alb-ingress-controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
$ helm delete my-release
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following tables lists the configurable parameters of the alb-ingress-controller chart and their default values.

Parameter | Description | Default
--- | --- | ---
`aws.accessKeyId` | If provided, AWS_ACCESS_KEY_ID environment variable will be set to this value | `""`
`aws.secretAccessKey` | If provided, AWS_SECRET_ACCESS_KEY environment variable will be set to this value | `""`
`aws.debug` | If true, enables logging on all outbound AWS API requests | `false`
`aws.region` | (REQUIRED) AWS region in which this ingress controller will operate | `us-west-1`
`clusterName` | (REQUIRED) Resources created by the ALB Ingress controller will be prefixed with this string | `k8s`
`controller.image.repository` | controller container image repository | `quay.io/coreos/alb-ingress-controller`
`controller.image.tag` | controller container image tag | `0.8`
`controller.image.pullPolicy` | controller container image pull policy | `IfNotPresent`
`controller.nodeSelector` | node labels for controller pod assignment | `{}`
`controller.podAnnotations` | annotations to be added to controller pod | `{}`
`controller.resources` | controller pod resource requests & limits | `{}`
`defaultBackend.image.repository` | default backend container image repository | `gcr.io/google_containers/defaultbackend`
`defaultBackend.image.tag` | default backend container image tag | `1.2`
`defaultBackend.image.pullPolicy` | default backend container image pull policy | `IfNotPresent`
`defaultBackend.nodeSelector` | node labels for default backend pod assignment | `{}`
`defaultBackend.podAnnotations` | annotations to be added to default backend pod | `{}`
`defaultBackend.replicaCount` | desired number of default backend pods | `1`
`defaultBackend.resources` | default backend pod resource requests & limits | `{}`
`defaultBackend.service.annotations` | annotations to be added to default backend service | `{}`
`scope.ingressClass` | If provided, the ALB ingress controller will only act on Ingress resources annotated with this class | `alb`
`scope.singleNamespace` | If true, the ALB ingress controller will only act on Ingress resources in a single namespace | `false` (watch all namespaces)
`scope.watchNamespace` | If scope.singleNamespace=true, the ALB ingress controller will only act on Ingress resources in this namespace | `""` (namespace of the ALB ingress controller)

```console
$ helm registry install quay.io/coreos/alb-ingress-controller-helm --name=my-release --set clusterName=mycluster
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while installing the chart. For example,

```console
$ helm registry install quay.io/coreos/alb-ingress-controller-helm --name=my-release -f values.yaml
```

> **Tip**: You can use the default [values.yaml](values.yaml)
