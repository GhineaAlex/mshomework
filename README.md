# Kubernetes Cluster Deployment

The setup includes two services, Service A and Service B, running on a cluster with the latest stable version of Kubernetes and RBAC enabled.

## Requirements

- **Kubernetes Cluster:** Set up a Kubernetes cluster using either Pulumi or Terraform.
- **Kubernetes Version:** Use the latest stable version.
- **RBAC:** Role-Based Access Control (RBAC) must be enabled.

## Services Deployment

The cluster should have the following two services deployed:

### Service A

- **Language:** WebServer written in C# or Go.
- **Functionality:**
  - Displays the current value of Bitcoin in USD, updated every 10 seconds from an external API.
  - Shows the average value of Bitcoin over the last 10 minutes.

### Service B

- **Type:** REST API service.
- **Functionality:**
  - Exposes a single controller.
  - Responds with a 200 status code on GET requests.

## Ingress

- **NGINX Ingress Controller:** The cluster should have an NGINX Ingress controller deployed.
- **Ingress Rules:** Set up corresponding ingress rules for both Service A and Service B.

## Network Policy

- Ensure that Service A is not able to communicate with Service B.

## Getting Started

terraform init
terraform apply
docker build -t 
docker push 
kubectl apply -f policy-btc.yaml
kubectl apply -f policy-api.yaml

