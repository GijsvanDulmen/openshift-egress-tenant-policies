#!/usr/bin/env bash
source settings.sh

kubectl --kubeconfig ${CONFIG} apply -f ../helm/crds/base.yaml
kubectl --kubeconfig ${CONFIG} apply -f ../helm/crds/policy.yaml
LOG_LEVEL=debug go run ../ -kubeconfig=${CONFIG} -egressfirewall=true -egressnetworkpolicy=true
