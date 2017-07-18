# Ingress Resources

This document covers how ingress resources work in relation to The ALB Ingress Controller.

## Ingress Behavior

Periodically, ingress update events are seen by the controller. The controller retains a list of all ingress resources it knows about, along with the current state of AWS components that satisfy them. When an update event is fired, the controller re-scans the list of ingress resources known to the cluster and determines, by comparing the list to its previously stored one, the ingresses requiring deletion, creation or modification.

An example ingress, from `example/2048/2048-ingress.yaml` is as follows.

```yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: "nginx-ingress"
  namespace: "2048-game"
  annotations:
    alb.ingress.kubernetes.io/scheme: internal
    alb.ingress.kubernetes.io/subnets: subnet-1234
    alb.ingress.kubernetes.io/security-groups: sg-1234
  labels:
    app: 2048-nginx-ingress
spec:
  rules:
  - host: 2048.example.com
    http:
      paths:
      - path: /
        backend:
          serviceName: "service-2048"
          servicePort: 80
```

The host field specifies the eventual Route 53-managed domain that will route to this service. The service, service-2048, must be of type NodePort (see [../examples/echoservice/echoserver-service.yaml](../examples/echoservice/echoserver-service.yaml)) in order for the provisioned ALB to route to it. If no NodePort exists, the controller will not attempt to provision resources in AWS. For details on purpose of annotations seen above, see [Annotations](#annotations).

## Annotations

The ALB Ingress Controller is configured by Annotations on the `Ingress` resource object. Some are required and some are optional. All annotations use the namespace `alb.ingress.kubernetes.io/`.

### Required Annotations

```
alb.ingress.kubernetes.io/security-groups
alb.ingress.kubernetes.io/subnets
```

Required annotations are:

- **security-groups**: Required. [Security groups](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html) that should be applied to the ALB instance. These can be referenced by security group IDs or the name tag associated with each security group. Example ID values are `sg-723a380a,sg-a6181ede,sg-a5181edd`. Example tag values are `appSG, webSG`.

- **subnets**: Required. The subnets where the ALB instance should be deployed. Must include 2 subnets, each in a different [availability zone](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html). These can be referenced by subnet IDs or the name tag associated with the subnet.  Example values for subnet IDs are `subnet-a4f0098e,subnet-457ed533,subnet-95c904cd`. Example values for name tags are: `webSubnet,appSubnet`.

### Optional Annotations

```
alb.ingress.kubernetes.io/backend-protocol
alb.ingress.kubernetes.io/certificate-arn
alb.ingress.kubernetes.io/healthcheck-path
alb.ingress.kubernetes.io/listen-ports
alb.ingress.kubernetes.io/scheme
alb.ingress.kubernetes.io/successCodes
alb.ingress.kubernetes.io/tags
```

Optional annotations are:

- **backend-protocol**: Enables selection of protocol for ALB to use to connect to backend service. When omitted, `HTTP` is used.

- **certificate-arn**: Enables HTTPS and uses the certificate defined, based on arn, stored in your [AWS Certificate Manager](https://aws.amazon.com/certificate-manager).

- **healthcheck-path**: Defines the path ALB health checks will check. When omitted, `/` is used.

- **listen-ports**: Defines the ports the ALB will expose. When omitted, `80` is used for HTTP and `443` is used for HTTPS. Uses a format as follows '[{"HTTP":8080,"HTTPS": 443}]'.

- **scheme**: Defines whether an ALB should be `internal` or `internet-facing`. See [Load balancer scheme](http://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.html#load-balancer-scheme) in the AWS documentation for more details.

- **successCodes**: Defines the HTTP status code that should be expected when doing health checks against the defined `healthcheck-path`. When omitted, `200` is used.

- **tags**: Defines [AWS Tags](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) that should be applied to the ALB instance and Target groups.
