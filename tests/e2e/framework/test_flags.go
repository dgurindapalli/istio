// Copyright 2017,2018 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

import (
	"flag"
)

// TestFlags holds routing versions to test, and also if Ingress/Egress should be tested
type TestFlags struct {
	V1alpha1 bool
	V1alpha3 bool
	Ingress  bool
	Egress   bool
}

// Init initializes golang "flags" with the flags of TestFlags
func (tf *TestFlags) Init() {
	flag.BoolVar(&tf.V1alpha1, "v1alpha1", tf.V1alpha1, "Enable / disable v1alpha1 routing rules.")
	flag.BoolVar(&tf.V1alpha3, "v1alpha3", tf.V1alpha3, "Enable / disable v1alpha3 routing rules.")
	flag.BoolVar(&tf.Ingress, "ingress", tf.Ingress, "Enable / disable Ingress tests.")
	flag.BoolVar(&tf.Egress, "egress", tf.Egress, "Enable / disable Egress tests.")
}

// ConfigVersions returns a list of strings of the enabled config versions
func (tf *TestFlags) ConfigVersions() []string {
	versions := []string{}
	if tf.V1alpha1 {
		versions = append(versions, "v1alpha1")
	}
	if tf.V1alpha3 {
		versions = append(versions, "v1alpha3")
	}
	return versions
}
