#!/usr/bin/env bash
source settings.sh

cd ..
KUBECONFIG=${CONFIG} helm -n openshift-egress-tenant-policies install --set LOG_LEVEL=DEBUG --create-namespace openshift-egress-tenant-policies helm/