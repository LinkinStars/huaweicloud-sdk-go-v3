/*
 * kms
 *
 * KMS v1.0 API, open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DecryptDataRequest struct {
	VersionId string                  `json:"version_id"`
	Body      *DecryptDataRequestBody `json:"body,omitempty"`
}

func (o DecryptDataRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DecryptDataRequest", string(data)}, " ")
}
