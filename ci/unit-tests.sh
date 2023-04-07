#!/usr/bin/env bash
set -euo pipefail

go version

cd repo

make test
