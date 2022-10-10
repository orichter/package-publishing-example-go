//Api classes for files's golang SDK
package api

import (
    "github.com/orichter/package-publishing-example-go/files-go-client/v4/client"
	"strings"
	import1 "github.com/orichter/package-publishing-example-go/files-go-client/v4/models/files/v4/config"
	"encoding/json"
	"net/http"
    "net/url"
)

type TierApi struct {
  ApiClient *client.ApiClient
}

func NewTierApi(apiClient *client.ApiClient) *TierApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &TierApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Create the tiering configuration.
    Create the tiering configuration using the provided request body.  The user needs to specify `capacityThreshold` in percentage. Tiering will only happen if the memory has exceeded the capacity threshold. Users can specify `schedule` for each day of the week, auto tiering can happen at that specified time.  A sample request body would look like this: ``` {   \"capacityThreshold\": 50,   \"schedule\": [     {         \"dayOfWeek\": 1,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },     {         \"dayOfWeek\": 2,         \"schedules\" :[             {                 \"startTimeHours\": 11,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 120             }         ]     },     {         \"dayOfWeek\": 3,         \"schedules\" :[             {                 \"startTimeHours\": 23,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },     {         \"dayOfWeek\": 4,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },      {         \"dayOfWeek\": 5,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },      {         \"dayOfWeek\": 6,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },      {         \"dayOfWeek\": 7,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     }   ] } ``` 

    parameters:-
    -> body (files.v4.config.TieringConfig) (required) : Tiering configuration model
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateTieringConfigApiResponse, error)
*/
func (api *TierApi) CreateTieringConfig(body *import1.TieringConfig, args ...map[string]interface{}) (*import1.CreateTieringConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-config"

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
    unmarshalledResp := new(import1.CreateTieringConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Create a tiering policy.
    Create the tiering policy using the provided request body.  User can specify list of shares, cool-off periods, minimum file size. This will decide the list of files which will be identified for tiering.  A sample request body would look like this:  ``` {   \"isIncludeFutureShares\": true,   \"extId\": \"234ee9c2-8f08-4276-9b4c-a084cd592111\",   \"mountTargetExtIds\": [     \"d89ee9c2-8f08-4276-9b4c-a084cd59271b\",     \"2e9fe9c1-9d08-6251-4b4c-9271ba084cd5\"   ],   \"cooloffPeriodInSeconds\": 86400,   \"minimumFileSizeInBytes\": 70000 } ``` 

    parameters:-
    -> body (files.v4.config.TieringPolicy) (required) : Tiering policy model
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.CreateTieringPolicyApiResponse, error)
*/
func (api *TierApi) CreateTieringPolicy(body *import1.TieringPolicy, args ...map[string]interface{}) (*import1.CreateTieringPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-policies"

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
    unmarshalledResp := new(import1.CreateTieringPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete the tiering configuration.
    Delete the tiering configuration with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the configuration to be deleted. After deletion auto tiering will be disabled.  How to use Etag  For performing delete, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/tier-config/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ````  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the DELETE request to the below URL  ``` /api/files/v4.0.a2/config/file-server/tier-config/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status 

    parameters:-
    -> extId (string) (required) : Tiering config extId.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteTieringConfigApiResponse, error)
*/
func (api *TierApi) DeleteTieringConfig(extId *string, args ...map[string]interface{}) (*import1.DeleteTieringConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-config/{extId}"

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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DeleteTieringConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Delete a tiering policy
    Delete the tiering policy identified with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the policy to be deleted.  How to use Etag  For performing delete, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/tier-policies/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ````  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the DELETE request to the below URL  ``` /api/files/v4.0.a2/config/file-server/tier-policies/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> extId (string) (required) : Tiering policy extId.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteTieringPolicyApiResponse, error)
*/
func (api *TierApi) DeleteTieringPolicy(extId *string, args ...map[string]interface{}) (*import1.DeleteTieringPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-policies/{extId}"

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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodDelete, nil, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.DeleteTieringPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get object store profile
    Get a single object store profile for the provided external identifier.  The user has to specify - a valid external identifier  (`profileExtId`) of the object store profile to be fetched.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> profileExtId (string) (required) : UUID of the tiering object store
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.GetObjectStoreApiResponse, error)
*/
func (api *TierApi) GetObjectStoreProfile(profileExtId *string, args ...map[string]interface{}) (*import1.GetObjectStoreApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/object-store/profiles/{profileExtId}"

    // verify the required parameter 'profileExtId' is set
	if nil == profileExtId {
		return nil, client.ReportError("profileExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"profileExtId"+"}", url.PathEscape(client.ParameterToString(*profileExtId, "")), -1)
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
    unmarshalledResp := new(import1.GetObjectStoreApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get the tiering configuration.
    Get tiering configuration with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the configuration to be fetched.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> extId (string) (required) : Tiering config extId.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.GetTieringConfigApiResponse, error)
*/
func (api *TierApi) GetTieringConfig(extId *string, args ...map[string]interface{}) (*import1.GetTieringConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-config/{extId}"

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
    unmarshalledResp := new(import1.GetTieringConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get list of the tiering configuration.
    Get list of the tiering configuration.

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.GetTieringConfigListApiResponse, error)
*/
func (api *TierApi) GetTieringConfigList(page_ *int, limit_ *int, args ...map[string]interface{}) (*import1.GetTieringConfigListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-config"


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
    unmarshalledResp := new(import1.GetTieringConfigListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List tiering policies
    Get a paginated list of the tiering policies.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported: - cooloffPeriodInSeconds - minimumFileSizeInBytes - isIncludeFutureShares  Example of supported query parameters for the tiering policy list API: ```   - ?$page=0&$limit=1   - ?$select=cooloffPeriodInSeconds, minimumFileSizeInBytes   - ?$limit=5&$select=cooloffPeriodInSeconds, minimumFileSizeInBytes   - ?$filter=cooloffPeriodInSeconds gt 50000   - ?$orderby=minimumFileSizeInBytes ``` The `$orderby` query parameter allows specifying attributes on which to sort the returned list of the tiering policies  The following parameters support sorting in the object store profile list: - cooloffPeriodInSeconds - minimumFileSizeInBytes - isIncludeFutureShares  A sample call would look like this: ``` /api/files/v4.0.a2/config/file-server/tier-policies?$orderby=cooloffPeriodInSeconds desc ```  The `$select` query parameter allows specifying attributes which the user wants to fetch in the returned list of the tiering policies, other attributes will be returned as a null value.  the following attributes can be selected: - cooloffPeriodInSeconds - minimumFileSizeInBytes - isIncludeFutureShares  Some more examples are given below: 1. Filter by cooloffPeriodInSeconds: ``` /api/files/v4.0.a2/config/file-server/tier-policies?$filter=cooloffPeriodInSeconds gt 50 ```  2. Order by cooloffPeriodInSeconds in ascending order ``` /api/files/v4.0.a2/config/file-server/tier-policies?$orderby=cooloffPeriodInSeconds asc ```  3. Order by cooloffPeriodInSeconds in descending order ``` /api/files/v4.0.a2/config/file-server/tier-policies?$orderby=cooloffPeriodInSeconds desc ```  4. Select by cooloffPeriodInSeconds ``` /api/files/v4.0.a2/config/file-server/tier-policies?$select=cooloffPeriodInSeconds ```  5. Paginate the returned tiering policy list ``` /api/files/v4.0.a2/config/file-server/tier-policies?$page=0&$limit=1 ```  6. Combination of queries ``` /api/files/v4.0.a2/config/file-server/tier-policies?$limit=5&$select=cooloffPeriodInSeconds &$orderby=cooloffPeriodInSeconds desc ```  If the user doesn't specify any search query parameters, a list of all tiering policies will be returned. 

    parameters:-
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - cooloffPeriodInSeconds - isIncludeFutureShares - minimumFileSizeInBytes 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - cooloffPeriodInSeconds - isIncludeFutureShares - minimumFileSizeInBytes 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - cooloffPeriodInSeconds - extId - isIncludeFutureShares - links - minimumFileSizeInBytes - mountTargetExtIds - tenantId 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.GetTieringPolicyListApiResponse, error)
*/
func (api *TierApi) GetTieringPolicies(page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.GetTieringPolicyListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-policies"


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
    unmarshalledResp := new(import1.GetTieringPolicyListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get a tiering policy by extId
    Get the tiering policy identified with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the policy to be fetched.  Etag  Etag is used to cache unchanged resources. When making a GET call to the above resource, `If-None-Match` header can be passed as shown in the example below.  ``` If-None-Match:        9 Content-Type:         application/json ```  The server compares the above Etag (sent with If-None-Match) with the Etag for its current version of the resource, and if both values match (that is, the resource has not changed), the server sends back a 304 Not Modified status, without a body. This tells the user that the cached version of the response is still good to use (fresh). If it doesn't match, it will send the response body of the latest resource with the updated value of Etag in the response headers as below:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  If this header is not passed, the server will send the full payload with the latest etag value in response headers. 

    parameters:-
    -> extId (string) (required) : Tiering policy extId.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.GetTieringPolicyApiResponse, error)
*/
func (api *TierApi) GetTieringPolicyById(extId *string, args ...map[string]interface{}) (*import1.GetTieringPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-policies/{extId}"

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
    unmarshalledResp := new(import1.GetTieringPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Initiate tiering data.
    Initiate the process of tiering data.  The user needs to provide duration in seconds. Tiering will start immediately and will run for the specified duration. If the duration is not provided the tiering will run until file server storage usage has reached the threshold capacity value.  A sample request body would look like this:  ``` {   \"durationInSeconds\": 86400 } ``` 

    parameters:-
    -> body (files.v4.config.TierData) (required) : Model for tier data.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.TierDataApiResponse, error)
*/
func (api *TierApi) TierData(body *import1.TierData, args ...map[string]interface{}) (*import1.TierDataApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/$actions/tier-data"

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
    unmarshalledResp := new(import1.TierDataApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update object store profile
    Update attributes of the object store profile with the given external identifier using the provided request body.  The user has to specify - a valid external identifier (`profileExtId`) of the policy to be updated. They also need to provide a request body for performing the update.  A sample request body would look like this:  ``` {   \"secretKey\": \"NaabhdrtaPDndqmvfdoaVqTPyotqz\",   \"fileServerExtId\": \"cc118c4e-bfe7-4cfd-a363-c2717983fb75\",   \"accessKey\": \"5--abgdrVrjL6zCKnAXfBNCUIPk_iEGpT\",   \"extId\": \"546b1234-af77-48a3-a1fb-f5cba068f037\",   \"tenantId\": \"aa123c4a-bfe7-4cfd-a363-c2717983f111\" } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/object-store/profiles/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/object-store/profiles/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.TieringObjectStoreProfile) (required) : Object store profile specification.
    -> profileExtId (string) (required) : UUID of the tiering object store
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateObjectStoreApiResponse, error)
*/
func (api *TierApi) UpdateObjectStoreProfile(body *import1.TieringObjectStoreProfile, profileExtId *string, args ...map[string]interface{}) (*import1.UpdateObjectStoreApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/object-store/profiles/{profileExtId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
    // verify the required parameter 'profileExtId' is set
	if nil == profileExtId {
		return nil, client.ReportError("profileExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"profileExtId"+"}", url.PathEscape(client.ParameterToString(*profileExtId, "")), -1)
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateObjectStoreApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update the tiering configuration.
    Update the tiering configuration with the given external identifier using the provided request body.  The user has to specify - a valid external identifier (`extId`) of the configuration to be updated. They also need to provide a request body for performing the update.  A sample request body would look like this:  ``` {   \"capacityThreshold\": 50,   \"schedule\": [     {         \"dayOfWeek\": 1,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },     {         \"dayOfWeek\": 2,         \"schedules\" :[             {                 \"startTimeHours\": 11,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 120             }         ]     },     {         \"dayOfWeek\": 3,         \"schedules\" :[             {                 \"startTimeHours\": 23,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },     {         \"dayOfWeek\": 4,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },      {         \"dayOfWeek\": 5,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },      {         \"dayOfWeek\": 6,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     },      {         \"dayOfWeek\": 7,         \"schedules\" :[             {                 \"startTimeHours\": 10,                 \"startTimeMinutes\": 10,                 \"durationMinutes\": 100             }         ]     }   ] } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/tier-config/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/tier-config/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.TieringConfig) (required) : Tiering configuration model
    -> extId (string) (required) : Tiering config extId.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateTieringConfigApiResponse, error)
*/
func (api *TierApi) UpdateTieringConfig(body *import1.TieringConfig, extId *string, args ...map[string]interface{}) (*import1.UpdateTieringConfigApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-config/{extId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateTieringConfigApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Update a tiering policy
    Update the tiering policy with the given external identifier using the provided request body.  The user has to specify - a valid external identifier (`extId`) of the policy to be updated. They also need to provide a request body for performing the update.  A sample request body would look like this:  ``` {   \"isIncludeFutureShares\": true,   \"mountTargetExtIds\": [     \"d89ee9c2-8f08-4276-9b4c-a084cd59271b\",     \"2e9fe9c1-9d08-6251-4b4c-9271ba084cd5\"   ],   \"cooloffPeriodInSeconds\": 86400,   \"minimumFileSizeInBytes\": 70000 } ```  It is always recommended to do a GET on a resource before doing an UPDATE.  How to pass Etag  For performing an update, the user needs an Etag value which is returned as a part of the response headers for the get operation.  A sample GET request url to get etag value would look like this:  ``` /api/files/v4.0.a2/config/file-server/tier-policies/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Response headers for the above call would look like this:  ``` Etag:              10 Ntnx-Request-Id:   fc6f929a-3ece-41d3-5a49-dd35f2610530 Content-Type:      application/json ```  The user needs to pass the above value of Etag to `If-Match` header in the PUT request to the below URL  ``` /api/files/v4.0.a2/config/file-server/tier-policies/bb118c5e-bfe7-4cfd-a363-c3717983fb75 ```  Request headers for the above call would look like this:  ``` If-Match:        10 Content-Type:    application/json ```  Etag is required for v4.0.a2 APIs. If this header is not passed or an incorrect value is passed, the request will fail with 412 precondition failed status. 

    parameters:-
    -> body (files.v4.config.TieringPolicy) (required) : Tiering policy model
    -> extId (string) (required) : Tiering policy extId.
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.UpdateTieringPolicyApiResponse, error)
*/
func (api *TierApi) UpdateTieringPolicy(body *import1.TieringPolicy, extId *string, args ...map[string]interface{}) (*import1.UpdateTieringPolicyApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/tier-policies/{extId}"

    // verify the required parameter 'body' is set
	if nil == body {
		return nil, client.ReportError("body is required and must be specified")
	}
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

    responseBody, err := api.ApiClient.CallApi(&uri, http.MethodPut, body, queryParams, headerParams, formParams, accepts, contentTypes, authNames)
    if nil != err || nil == responseBody{
    	return nil, err
	}
    unmarshalledResp := new(import1.UpdateTieringPolicyApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

