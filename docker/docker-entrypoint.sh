#!/bin/bash
set -eo pipefail

echo "docker-entrypoint: $@"

exec "$@"
