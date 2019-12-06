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
docker run --rm -d -p 8080:8080 chaunceyt/legislators-api:v0.0.5
```

## Kubernetes pod

```
kubectl run legislator-api-pod --image chaunceyt/legislators-api:v0.0.5 --restart=Never
```

## Kubernetes deployment

```
kubectl run legislator-api --image chaunceyt/legislators-api:v0.0.5 --expose --port 8080
```
