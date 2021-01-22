/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateL7RuleRequest struct {
	L7policyId string                   `json:"l7policy_id"`
	L7ruleId   string                   `json:"l7rule_id"`
	Body       *UpdateL7RuleRequestBody `json:"body,omitempty"`
}

func (o UpdateL7RuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateL7RuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleRequest", string(data)}, " ")
}
