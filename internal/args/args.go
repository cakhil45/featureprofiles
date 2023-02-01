// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package args define arguments for testing that depend on the available components
// and their naming on the device, if they cannot be enumerated easily from /components by type.
// Having these arguments at the project level help us run the whole suite of tests
// without defining them per test.
package args

import (
	"flag"
)

// Global test flags.
var (
	NumControllerCards = flag.Int("arg_num_controller_cards", -1, "The expected number of controller cards. Some devices with a single controller report 0, which is a valid expected value. Expectation is not checked for values < 0.")
	NumLinecards       = flag.Int("arg_num_linecards", -1, "The expected number of linecards. Some devices with a single linecard report 0, which is a valid expected value. Expectation is not checked for values < 0.")
	P4RTNodeName1      = flag.String("arg_p4rt_node_name_1", "", "The P4RT Node Name for the first FAP. Test that reserves ports in the same FAP should configure this P4RT Node. The value will only be used if deviation ExplicitP4RTNodeComponent is applied.")
	P4RTNodeName2      = flag.String("arg_p4rt_node_name_2", "", "The P4RT Node Name for the second FAP. Test that reserves ports in two different FAPs should configure this P4RT Node in addition to the Node defined in P4RTNodeName1. The value will only be used if deviation ExplicitP4RTNodeComponent is applied.")
)