#!/usr/bin/env bash
set -euo pipefail

go version

cd src/code.cloudfoundry.org/cpu-entitlement-plugin

go install github.com/onsi/ginkgo/ginkgo@latest
make test
