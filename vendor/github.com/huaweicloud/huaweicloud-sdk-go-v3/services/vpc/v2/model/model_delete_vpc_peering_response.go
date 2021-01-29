/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteVpcPeeringResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpcPeeringResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteVpcPeeringResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpcPeeringResponse", string(data)}, " ")
}
