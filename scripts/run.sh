
# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Run binary
command bazel run //:main