# Legislators API

A simple US Legislator API.

# Purpose

The purpose of the application is be used when DEMOing various K8s features.

- Pod
- Deployments
- Scaling of Deployments
- [HPA](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)

Another reason I'm allocating time to this application is to learn Go.

# Playing around with this API

## Docker

```
docker run --rm -d -p 8080:8080 chaunceyt/legislator:v0.0.4
```

## Kubernetes pod

```
kubectl run legislator-api-pod --image chaunceyt/legislator:v0.0.4 --restart=Never
```

## Kubernetes deployment

```
kubectl run legislator-api --image chaunceyt/legislator:v0.0.4 --expose --port 8080
```
