#!/bin/bash
cd ..
docker build --progress=plain -t openshift-egress-tenant-policies .

docker tag k8qu ghcr.io/gijsvandulmen/openshift-egress-tenant-policies:latest
docker tag k8qu ghcr.io/gijsvandulmen/openshift-egress-tenant-policies:1.0

docker push ghcr.io/gijsvandulmen/openshift-egress-tenant-policies:latest
docker push ghcr.io/gijsvandulmen/openshift-egress-tenant-policies:1.0