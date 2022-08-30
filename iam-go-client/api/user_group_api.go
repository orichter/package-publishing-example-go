//Api classes for iam's golang SDK
package api

import (
    "github.com/orichter/package-publishing-example-go/iam-go-client/v4/client"
	"strings"
	import3 "github.com/orichter/package-publishing-example-go/iam-go-client/v4/models/iam/v4/authn"
	"encoding/json"
	"net/http"
    "net/url"
)

type UserGroupApi struct {
  ApiClient *client.ApiClient
}

func NewUserGroupApi(apiClient *client.ApiClient) *UserGroupApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &UserGroupApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create user group
    Create a user group

    parameters:-
    -> body (iam.v4.authn.UserGroup) (required) : Create a user group
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (iam.v4.authn.CreateUserGroupApiResponse, error)
*/
func (api *UserGroupApi) CreateUserGroup(body *import3.UserGroup, args ...map[string]interface{}) (*import3.CreateUserGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/iam/v4.0.a1/authn/user-groups"

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

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPost, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import3.CreateUserGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get user group
    View a user group

    parameters:-
    -> extId (string) (required) : External Identifier of the user group
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (iam.v4.authn.GetUserGroupApiResponse, error)
*/
func (api *UserGroupApi) GetUserGroupById(extId *string, args ...map[string]interface{}) (*import3.GetUserGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/iam/v4.0.a1/authn/user-groups/{extId}"

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

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import3.GetUserGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List user group(s)
    List all user group(s)

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - createdBy - distinguishedName - extId - groupType - idpId - name 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - createdTime - distinguishedName - groupType - lastUpdatedTime - name 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - createdBy - createdTime - distinguishedName - extId - groupType - idpId - lastUpdatedTime - links - name - tenantId 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (iam.v4.authn.ListUserGroupApiResponse, error)
*/
func (api *UserGroupApi) GetUserGroupList(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import3.ListUserGroupApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/iam/v4.0.a1/authn/user-groups"


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
	if select_ != nil {
		queryParams.Add("$select", client.ParameterToString(*select_, ""))
	}

    // Header Params
    if ifMatch, ifMatchOk := argMap["If-Match"].(string); ifMatchOk {
        headerParams["If-Match"] = ifMatch
    }
    if ifNoneMatch, ifNoneMatchOk := argMap["If-None-Match"].(string); ifNoneMatchOk {
        headerParams["If-None-Match"] = ifNoneMatch
    }
    authNames := []string{"basicAuthScheme"}

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodGet, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import3.ListUserGroupApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

