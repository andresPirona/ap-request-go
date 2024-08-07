package ap_request_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andresPirona/ap-request-go/pkg/app/domain"
	"github.com/andresPirona/ap-request-go/pkg/app/request"
	"net/http"
)

// NewRequest Request with response type STRUCT
//
// Testing upgrade v2
// *T* Struct to parse
// *opts* Request options
func NewRequest[T any](opts domain.Options) (*T, *http.Response, error) {
	// Type for the final response
	var apiResponse T
	// Response for the request
	bodyResult, response, errPet := request.HandleRequest(opts)
	if errPet != nil {
		return nil, nil, errPet
	}
	// Parse for the response
	errFormat := json.Unmarshal(bodyResult, &apiResponse)
	if errFormat != nil {
		msg := fmt.Sprintf("error_unmarshal_response  %s ", errFormat.Error())
		return nil, nil, errors.New(msg)
	}
	return &apiResponse, response, nil
}
