/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/apache/camel-k/pkg/apis/camel/v1"
)

// AddonTraitApplyConfiguration represents an declarative configuration of the AddonTrait type for use
// with apply.
type AddonTraitApplyConfiguration struct {
	v1.RawMessage `json:",inline"`
}

// AddonTraitApplyConfiguration constructs an declarative configuration of the AddonTrait type for use with
// apply.
func AddonTrait() *AddonTraitApplyConfiguration {
	return &AddonTraitApplyConfiguration{}
}
