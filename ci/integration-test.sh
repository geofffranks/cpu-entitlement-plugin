#!/usr/bin/env bash
set -e

if [ -n "$CONFIG" ]; then
  CF_API=$(cat ${CONFIG_PATH} | jq .api)
  CF_USERNAME=$(cat ${CONFIG_PATH} | jq .admin_user)
  CF_PASSWORD=$(cat ${CONFIG_PATH} | jq .admin_password)
fi

# sample CONFIG file
# {
#   "admin_password": "meow",
#   "admin_user": "admin",
#   "api": "meow.com"
# }

pushd cpu-entitlement-plugin
go install github.com/onsi/ginkgo/ginkgo@latest
make integration-test
popd
