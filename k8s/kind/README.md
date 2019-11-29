# Install Kubernetes IN Docker (Kind)

```
mkdir $HOME/bin
curl -Lo ./kind https://github.com/kubernetes-sigs/kind/releases/download/v0.5.1/kind-$(uname)-amd64
chmod +x ./kind
mv ./kind $HOME/bin/kind
kind version
```

# Create a local cluster

```
kind create cluster --config kind-cluster.yaml
./fix_sysctls.sh

export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"

kubectl apply -f canal.yaml
``

Install MetalLB
```
kubectl apply -f https://raw.githubusercontent.com/google/metallb/v0.8.3/manifests/metallb.yaml
kubectl apply -f https://git.io/km-config.yaml
```

Install Nginx ingress

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
kubectl get pods -n ingress-nginx
kubectl get svc -n ingress-nginx
```
