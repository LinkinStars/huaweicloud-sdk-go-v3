/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowVolumeResponse struct {
	Volume *VolumeDetail `json:"volume,omitempty"`
}

func (o ShowVolumeResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowVolumeResponse", string(data)}, " ")
}
