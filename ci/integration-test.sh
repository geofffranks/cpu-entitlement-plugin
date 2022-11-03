#!/usr/bin/env bash
set -e

mkdir secrets
if [ -n "$CONFIG" ]; then
  export CF_API=$(cat ${CONFIG} | jq -r .api)
  export CF_USERNAME=$(cat ${CONFIG} | jq -r .admin_user)
  export CF_PASSWORD=$(cat ${CONFIG} | jq -r .admin_password)
  echo "from config" > secrets/secrets.txt
fi
echo $CF_API >> secrets/secrets.txt
echo $CF_USERNAME >> secrets/secrets.txt
echo $CF_PASSWORD >> secrets/secrets.txt

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
