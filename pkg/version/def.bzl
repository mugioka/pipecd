# Copyright 2017 The Kubernetes Authors.
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

def version_x_defs():
    stamp_pkgs = [
        "github.com/pipe-cd/pipecd/pkg/version",
    ]

    # This should match the list of vars set in hack/print-workspace-status.sh.
    stamp_vars = [
        "version",
        "gitCommit",
        "buildDate",
    ]

    # Generate the cross-product.
    x_defs = {}
    for pkg in stamp_pkgs:
        for var in stamp_vars:
            x_defs["%s.%s" % (pkg, var)] = "{%s}" % var
    return x_defs
