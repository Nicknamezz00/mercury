#!/bin/bash

# Usage: ./scripts/air.sh

set -e

cd "$(dirname "$0")/../"

air -c ./scripts/.air.toml
