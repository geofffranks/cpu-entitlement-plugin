#!/usr/bin/env bash
set -e

if [ -n "$CONFIG" ]; then
  export CF_API=$(cat ${CONFIG} | jq -r .api)
  export CF_USERNAME=$(cat ${CONFIG} | jq -r .admin_user)
  export CF_PASSWORD=$(cat ${CONFIG} | jq -r .admin_password)
  export ROUTER_CA_CERT=$(cat ${CONFIG} | jq -r .ca_cert)
fi

# sample CONFIG file
# {
#   "admin_password": "meow",
#   "admin_user": "admin",
#   "api": "meow.com",
#   "lb_cert": "BEGIN MEOW CERT"
# }

pushd cpu-entitlement-plugin
go install github.com/onsi/ginkgo/ginkgo@latest
make integration-test
popd
