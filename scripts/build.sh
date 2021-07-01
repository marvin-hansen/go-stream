# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Build all sources
bazel build //:build
