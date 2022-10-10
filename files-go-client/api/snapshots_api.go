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

type SnapshotsApi struct {
  ApiClient *client.ApiClient
}

func NewSnapshotsApi(apiClient *client.ApiClient) *SnapshotsApi {
	if apiClient == nil {
		apiClient = client.NewApiClient()
	}

	a := &SnapshotsApi{
		ApiClient: apiClient,
	}
	return a
}


/**
    Delete mount target snapshot
    Delete a mount target snapshot with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the mount target to which the snapshot belongs to and a valid external identifier (`extId`) of the snapshot to be deleted. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> snapshotExtId (string) (required) : The extId of the snapshot. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.DeleteMountTargetSnapshotApiResponse, error)
*/
func (api *SnapshotsApi) DeleteMountTargetSnapshot(mountTargetExtId *string, snapshotExtId *string, args ...map[string]interface{}) (*import1.DeleteMountTargetSnapshotApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/snapshots/{snapshotExtId}"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'snapshotExtId' is set
	if nil == snapshotExtId {
		return nil, client.ReportError("snapshotExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"snapshotExtId"+"}", url.PathEscape(client.ParameterToString(*snapshotExtId, "")), -1)
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
    unmarshalledResp := new(import1.DeleteMountTargetSnapshotApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    Get mount target snapshot for the provided extId
    Get mount target snapshot with the given external identifier.  The user has to specify - a valid external identifier (`extId`) of the mount target to which the snapshot belongs to and a valid external identifier (`extId`) of the snapshot to be fetched. 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> snapshotExtId (string) (required) : The extId of the snapshot. Example:48f78959-14a6-4c47-b5db-920460c4b668
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.MountTargetSnapshotApiResponse, error)
*/
func (api *SnapshotsApi) GetMountTargetSnapshotByExtId(mountTargetExtId *string, snapshotExtId *string, args ...map[string]interface{}) (*import1.MountTargetSnapshotApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/snapshots/{snapshotExtId}"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}
    // verify the required parameter 'snapshotExtId' is set
	if nil == snapshotExtId {
		return nil, client.ReportError("snapshotExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
    uri = strings.Replace(uri, "{"+"snapshotExtId"+"}", url.PathEscape(client.ParameterToString(*snapshotExtId, "")), -1)
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
    unmarshalledResp := new(import1.MountTargetSnapshotApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

/**
    List mount target snapshots
    Get a paginated list of mount target snapshots.  You can use filtering options to manage the results of your query. The `$filter` query parameter provides a flexible way to subset the resources from collections by combining comparison and other functions. There are several kinds of basic predicates and built-in functions for $filter, including logical operators and arithmetic operators. For more detailed information, refer to the  [OData V4 URL Conventions Document](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html). For more information, see [5.1.2 System Query Option $filter](https://docs.oasis-open.org/odata/odata/v4.01/odata-v4.01-part2-url-conventions.html#sec_SystemQueryOptionfilter). Following OData filters are supported:   - name   - creator  A sample request URL would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots?$filter=name eq 'test' ```  Example of supported query parameters for mount target snapshots LIST API: ```   - ?$page=0&$limit=1   - ?$select=name, createTime, totalSpaceBytes   - ?$orderby=totalSpaceBytes asc   - ?$filter=creator eq Schema.Enums.SnapshotCreator'SSR_SNAPSHOT'   - ?$filter=startswith(name,'test')   - ?$filter=endswith(name,'test')   - ?$filter=contains(name,'test')   - ?$select=name, creator, totalSpaceBytes&$filter=startswith(name,'test')   - ?$limit=5&$select=name, totalSpaceBytes&$orderby=totalSpaceBytes asc&$filter=contains(name,'test')\\n ```  The `$orderby` query parameter allows specifying attributes on which to sort the returned list of snapshots.  The following parameters support sorting in snapshots API   - createTime   - name   - reclaimableSpaceBytes   - totalSpaceBytes   - creator  A sample call would look like this:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots?$orderby=totalSpaceBytes asc ```  The `$select` query parameter allows specifying attributes which the user wants to fetch in the returned list of snapshots, other attributes will be returned as a null value.  The following attributes can be selected: ```   - mountTargetExtId   - fileServerExtId   - name   - createTime   - totalSpaceBytes   - reclaimableSpaceBytes   - creator ```  Some more examples are given below:  1. Filter by name:  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots?$filter=contains(name,'test') ```  2. Order by totalSpaceBytes in ascending order  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots?$orderby=totalSpaceBytes asc ```  3. Order by totalSpaceBytes in descending order  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots?$orderby=totalSpaceBytes desc ```  4. Select by name  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots?$select=name ```  5. Paginate the results  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots?$page=0&$limit=1 ```  6. Combination of queries  ``` /api/files/v4.0.a2/config/file-server/mount-targets/9c1e537d-6777-4c22-5d41-ddd0c3337aa9/snapshots$limit=5&$select=name, totalSpaceBytes&$orderby=totalSpaceBytes asc&$filter=contains(name,'test') ``` 

    parameters:-
    -> mountTargetExtId (string) (required) : The {extId} of the mount target. Example:9c1e537d-6777-4c22-5d41-ddd0c3337aa9.
    -> page_ (int) (optional) : A URL query parameter that specifies the page number of the result set.  Must be a positive integer between 0 and the maximum number of pages that are available for that resource.  Any number out of this range will be set to its nearest bound.  In other words, a page number of less than 0 would be set to 0 and a page number greater than the total available pages would be set to the last page. 
    -> limit_ (int) (optional) : A URL query parameter that specifies the total number of records returned in the result set.  Must be a positive integer between 0 and 100. Any number out of this range will be set to the default maximum number of records, which is 100. 
    -> filter_ (string) (optional) : A URL query parameter that allows clients to filter a collection of resources. The expression specified with $filter is evaluated for each resource in the collection, and only items where the expression evaluates to true are included in the response. Expression specified with the $filter must conform to the OData V4.01 URL conventions. The filter can be applied on the following fields: - creator - name 
    -> orderby_ (string) (optional) : A URL query parameter that allows clients to specify the sort criteria for the returned list of objects. Resources can be sorted in ascending order using asc or descending order using desc. If asc or desc are not specified the resources will be sorted in ascending order by default. For example, 'orderby=templateName desc' would get all templates sorted by templateName in desc order. The orderby can be applied to the following fields: - createTime - creator - name - reclaimableSpaceBytes - totalSpaceBytes 
    -> select_ (string) (optional) : A URL query parameter that allows clients to request a specific set of properties for each entity or complex type. Expression specified with the $select must conform to the OData V4.01 URL conventions. If a $select expression consists of a single select item that is an asterisk (i.e. *), then all properties on the matching resource will be returned. - createTime - creator - extId - fileServerExtId - links - mountTargetExtId - name - reclaimableSpaceBytes - tenantId - totalSpaceBytes 
    -> args (map[string]interface{}) (optional) : Additional Arguments

    returns: (*files.v4.config.MountTargetSnapshotListApiResponse, error)
*/
func (api *SnapshotsApi) GetMountTargetSnapshots(mountTargetExtId *string, page_ *int, limit_ *int, filter_ *string, orderby_ *string, select_ *string, args ...map[string]interface{}) (*import1.MountTargetSnapshotListApiResponse, error) {
    argMap := make(map[string]interface{})
	if len(args) > 0 {
        argMap = args[0]
    }

    uri := "/api/files/v4.0.a2/config/file-server/mount-targets/{mountTargetExtId}/snapshots"

    // verify the required parameter 'mountTargetExtId' is set
	if nil == mountTargetExtId {
		return nil, client.ReportError("mountTargetExtId is required and must be specified")
	}

    // Path Params
    uri = strings.Replace(uri, "{"+"mountTargetExtId"+"}", url.PathEscape(client.ParameterToString(*mountTargetExtId, "")), -1)
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
    unmarshalledResp := new(import1.MountTargetSnapshotListApiResponse)
    json.Unmarshal(responseBody, &unmarshalledResp)
	return unmarshalledResp, err
}

