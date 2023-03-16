package auth

import (
	"encoding/json"
	"github.com/andresPirona/ap-request-go/pkg/app/domain"
	"net/http"
)

// Logic to set the properties for the auth method selected
func SetAuthMethod(req *http.Request, authMethod *domain.AuthMethod) {

	if authMethod != nil {
		// Switch to pick the auth method
		switch authMethod.GetName() {
		// Setting properties by each auth method
		case string(domain.Basic):
			var data domain.BasicAuth
			dataBytes, _ := json.Marshal(authMethod.Data)
			_ = json.Unmarshal(dataBytes, &data)
			req.SetBasicAuth(data.User, data.Pass)
		default:
		}
	}

}
