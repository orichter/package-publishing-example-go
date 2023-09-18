//Api classes for vmm's golang SDK
package api

import (
	"encoding/json"
	"github.com/orichter/package-publishing-example-go/vmm-go-client/v4/client"
	import3 "github.com/orichter/package-publishing-example-go/vmm-go-client/v4/models/vmm/v4/ahv/config"
	"net/http"
	"net/url"
	"strings"
)

type VmDataProtectionApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewVmDataProtectionApi(apiClient *client.ApiClient) *VmDataProtectionApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &VmDataProtectionApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "ntnx-request-id", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// Revert VM identified by {extId}. This does an in-place VM restore from a specified VM Recovery Point.
func (api *VmDataProtectionApi) RevertVm(extId *string, body *import3.VmRevertParams, args ...map[string]interface{}) (*import3.RevertVmResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/vmm/v4.0.a1/ahv/config/vms/{extId}/$actions/revert"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}
	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{"application/json"}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Headers provided explicitly on operation takes precedence
	for headerKey, value := range argMap {
		// Skip platform generated headers
		if !api.headersToSkip[strings.ToLower(headerKey)] {
			if value != nil {
				if headerValue, headerValueOk := value.(string); headerValueOk {
					headerParams[headerKey] = headerValue
				}
			}
		}
	}

	authNames := []string{"basicAuthScheme"}

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import3.RevertVmResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
