#!/usr/bin/env bash
# Copyright 2018 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

TESTINFRA_ROOT=$(git rev-parse --show-toplevel)

GENERATED_SECURITY_CONFIG="${TESTINFRA_ROOT}/config/jobs/kubernetes-security/generated-security-jobs.yaml"

rm "${GENERATED_SECURITY_CONFIG}"

bazel run //config/jobs/kubernetes-security:genjobs -- \
"--config=${TESTINFRA_ROOT}/prow/config.yaml" \
"--jobs=${TESTINFRA_ROOT}/config/jobs" \
"--output=${GENERATED_SECURITY_CONFIG}"
