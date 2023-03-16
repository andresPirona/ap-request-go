package tests

import (
	ap_request_go "github.com/andresPirona/ap-request-go"
	"github.com/andresPirona/ap-request-go/pkg/app/domain"
	"net/http"
	"testing"
)

type ObjectTest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestSingleRequest(t *testing.T) {

	testList := []domain.Options{{
		Method:  http.MethodGet,
		Url:     "https://6413050a3b710647375ca7be.mockapi.io/api/v1/users/1",
		Body:    nil,
		Auth:    nil,
		Headers: nil,
		Params:  nil,
	}, {
		Method:  http.MethodGet,
		Url:     "https://6413050a3b710647375ca7be.mockapi.io/api/v1/users/2",
		Body:    nil,
		Auth:    nil,
		Headers: nil,
		Params:  nil,
	}}

	for i, testOptions := range testList {

		_, _, err := ap_request_go.NewRequest[ObjectTest](testOptions)
		if err != nil {
			t.Errorf("ERROR - request: %v, url: %s get error --> %s", i+1, testOptions.Url, err)
			continue
		}

		t.Logf("SUCCESS - request: %v, url: %s", i+1, testOptions.Url)

	}

}