/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTrackerConfigResponse struct {
	Channel  *ChannelConfigBody  `json:"channel,omitempty"`
	Selector *SelectorConfigBody `json:"selector,omitempty"`
	// IAM委托名称
	AgencyName *string `json:"agency_name,omitempty"`
}

func (o ShowTrackerConfigResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowTrackerConfigResponse", string(data)}, " ")
}
