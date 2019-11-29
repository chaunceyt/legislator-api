# Horizontal Pod Autoscaler 

Deploying Metrics Server


Download `helm fetch stable/metrics-server --untar` and install the Metrics server

```
helm template metrics-server \
--set rbac.create=true \
--set args[0]="--kubelet-insecure-tls=true" \
--set args[1]="--kubelet-preferred-address-types=InternalIP" \
--set args[2]="--v=2" \
--name metrics-server | kubectl apply -f -
```

Install a sample application 

```
kubectl run legislator-api --image=chaunceyt/legislators-api:v0.0.3 --requests=cpu=100m --expose --port=8080
kubectl autoscale deployment legislator-api --cpu-percent=50 --min=1 --max=10
```

Add some load to the system

```
kubectl run -it --rm load-generator --image=busybox --restart=Never /bin/sh
while true; do wget -q -O - http://legislator-api:8080; done
```

In different terminal

```
kubectl get hpa -w
```
