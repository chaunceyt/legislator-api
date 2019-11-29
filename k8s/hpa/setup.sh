#!/bin/bash

helm template metrics-server \
--set rbac.create=true \
--set args[0]="--kubelet-insecure-tls=true" \
--set args[1]="--kubelet-preferred-address-types=InternalIP" \
--set args[2]="--v=2" \
--name metrics-server | kubectl apply -f -
