// Copyright 2016-2018 Authors of Cilium
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

package model

import (
	"github.com/cilium/cilium/api/v1/models"
	"github.com/cilium/cilium/pkg/labels"
)

// NewOplabelsFromModel creates new label from the model.
func NewOplabelsFromModel(base *models.LabelConfigurationStatus) *labels.OpLabels {
	if base == nil {
		return nil
	}
	oplabels := labels.NewOpLabelsFromModel(base.Realized.User, base.Disabled, base.SecurityRelevant, base.Derived)
	return &oplabels
}

func NewModel(o *labels.OpLabels) *models.LabelConfigurationStatus {
	return &models.LabelConfigurationStatus{
		Realized: &models.LabelConfigurationSpec{
			User: o.GetUserModel(),
		},
		SecurityRelevant: o.GetIdentityModel(),
		Derived:          o.GetInfoModel(),
		Disabled:         o.GetDisabledModel(),
	}
}
