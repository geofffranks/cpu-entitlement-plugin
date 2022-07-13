#!/usr/bin/env bash
set -e

pushd cpu-entitlement-plugin
go install github.com/onsi/ginkgo/ginkgo@latest
make integration-test
popd
