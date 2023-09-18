//Api classes for prism's golang SDK
package api

import (
	"encoding/json"
	"github.com/orichter/package-publishing-example-go/prism-go-client/v4/client"
	import2 "github.com/orichter/package-publishing-example-go/prism-go-client/v4/models/prism/v4/config"
	"net/http"
	"net/url"
	"strings"
)

type CategoriesApi struct {
	ApiClient     *client.ApiClient
	headersToSkip map[string]bool
}

func NewCategoriesApi(apiClient *client.ApiClient) *CategoriesApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &CategoriesApi{
		ApiClient: apiClient,
	}

	headers := []string{"authorization", "cookie", "ntnx-request-id", "host", "user-agent"}
	a.headersToSkip = make(map[string]bool)
	for _, header := range headers {
		a.headersToSkip[header] = true
	}

	return a
}

// ___Create a category taking the input from the request body___  At the minimum, the `key` and the `value` fields must be provided in the payload. The `description` field can also be provided. Creation of a category succeeds only if there is no other category in the system with the same key-value combination.  A sample request body: ``` {   \"key\":\"pet\",   \"value\": \"cat\",   \"description\": \"This is a sample category\" } ```  __Other fields__ * All other fields, if specified in the payload, are ignored. * The field `extId` is autogenerated as a random (type 4) UUID. It does not  depend on the contents of any other fields. * Re-creation of a category with the same key-value combination after deleting the existing one, will generate a different extId to the previous category. * The field `ownerUuid` is set as the UUID of the currently logged in user. * The field `type` is set to be 'USER'
func (api *CategoriesApi) CreateCategory(body *import2.Category, args ...map[string]interface{}) (*import2.CategoryCreateApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a2/config/categories"

	// verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}

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
	unmarshalledResp := new(import2.CategoryCreateApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// ___Deletes the category with the given external identifier.___  The user has to specify inside the path parameter - a valid external identifier (`extId`) of the category to be deleted.  A sample call would look like this: ``` /prism/v4.0.a2/config/categories/cafc8e9e-b595-46f8-8d43-f62746180a5b ```  __Notes__ * Categories that contain associations cannot be deleted. * Internal categories cannot be deleted.
func (api *CategoriesApi) DeleteCategoryByExtId(extId *string, args ...map[string]interface{}) (*import2.CategoryDeleteApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a2/config/categories/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.CategoryDeleteApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// ___Fetch a list of categories with pagination, filtering, sorting, selection and expansion___
func (api *CategoriesApi) GetAllCategories(page_ *int, limit_ *int, filter_ *string, orderby_ *string, expand_ *string, select_ *string, args ...map[string]interface{}) (*import2.CategoryListApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a2/config/categories"

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if page_ != nil {

		queryParams.Add("$page", client.ParameterToString(*page_, ""))
	}
	if limit_ != nil {

		queryParams.Add("$limit", client.ParameterToString(*limit_, ""))
	}
	if filter_ != nil {

		queryParams.Add("$filter", client.ParameterToString(*filter_, ""))
	}
	if orderby_ != nil {

		queryParams.Add("$orderby", client.ParameterToString(*orderby_, ""))
	}
	if expand_ != nil {

		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}
	if select_ != nil {

		queryParams.Add("$select", client.ParameterToString(*select_, ""))
	}
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.CategoryListApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// ___Get a category with the given external identifier___
func (api *CategoriesApi) GetCategoryByExtId(extId *string, expand_ *string, args ...map[string]interface{}) (*import2.CategoryGetApiResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a2/config/categories/{extId}"

	// verify the required parameter 'extId' is set
	if nil == extId {
		return nil, client.ReportError("extId is required and must be specified")
	}

	// Path Params
	uri = strings.Replace(uri, "{"+"extId"+"}", url.PathEscape(client.ParameterToString(*extId, "")), -1)
	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := url.Values{}

	// to determine the Content-Type header
	contentTypes := []string{}

	// to determine the Accept header
	accepts := []string{"application/json"}

	// Query Params
	if expand_ != nil {

		queryParams.Add("$expand", client.ParameterToString(*expand_, ""))
	}
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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.CategoryGetApiResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

// ___Updates the category with the given external identifier using the provided request body___  Updation needs the extId in the URL path, and also a payload representing the new state of the category.  An update is a 'set' operation, indicating that whatever is specified in the payload will be written to the database wholesome.  Fields which need to be erased/set to null can be omitted from the payload.  Update operation requires an _E-Tag_ , specified in the header as _If-Match_. The value of the ETag can be obtained through a getByExtId call, in the response header _E-Tag_.  The following fields are available for modification: * value - mandatory field * description - not providing the field is equivalent to setting the value to null * ownerUuid - mandatory if the response of the GET endpoint contains this field  A sample request body would look like this: ``` {   \"extId\": \"cafc8e9e-b595-46f8-8d43-f62746180a5b\",   \"key\":\"sample-category\",   \"value\": \"sample-category-value-updated\",   \"description\": \"This is the updated description\",   \"type\": \"USER\",   \"ownerUuid\": \"5afc8e9e-b595-46f8-8d43-f62746180a5f\" } ```  __Note__ * Update request will fail if the resulting key + updated value combination already points to an existing category. * Updation of `value` should be done with caution, as categories might be referenced by other services using the same `key:value` pair, and after updation,   those stored references will be dangling. * Updation of the field `ownerUuid` is allowed only for _super-admin/legacy-roles/local-users_. * Update request will fail if the new `ownerUuid` field does not  refer to a valid user UUID. * Update request will fail if the `ownerUuid` field is erased (i.e not given in the payload) * The `type` and the `extId` fields are ignored. * Internal categories cannot be updated. * To create an update request payload, it is advisable to copy the response of a get api call and modify fields that need to be edited.
func (api *CategoriesApi) UpdateCategory(extId *string, body *import2.Category, args ...map[string]interface{}) (*import2.CategoryPutResponse, error) {
	argMap := make(map[string]interface{})
	if len(args) > 0 {
		argMap = args[0]
	}

	uri := "/api/prism/v4.0.a2/config/categories/{extId}"

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

	responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
	if nil != err || nil == responseBody {
		return nil, err
	}
	unmarshalledResp := new(import2.CategoryPutResponse)
	json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}
