# \ProductsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationsGet**](ProductsApi.md#ConfigurationsGet) | **Get** /configurations | Get system configurations.
[**ConfigurationsPut**](ProductsApi.md#ConfigurationsPut) | **Put** /configurations | Modify system configurations.
[**ConfigurationsResetPost**](ProductsApi.md#ConfigurationsResetPost) | **Post** /configurations/reset | Reset system configurations.
[**EmailPingPost**](ProductsApi.md#EmailPingPost) | **Post** /email/ping | Test connection and authentication with email server.
[**InternalSyncregistryPost**](ProductsApi.md#InternalSyncregistryPost) | **Post** /internal/syncregistry | Sync repositories from registry to DB.
[**JobsReplicationGet**](ProductsApi.md#JobsReplicationGet) | **Get** /jobs/replication | List filters jobs according to the policy and repository
[**JobsReplicationIdDelete**](ProductsApi.md#JobsReplicationIdDelete) | **Delete** /jobs/replication/{id} | Delete specific ID job.
[**JobsReplicationIdLogGet**](ProductsApi.md#JobsReplicationIdLogGet) | **Get** /jobs/replication/{id}/log | Get job logs.
[**JobsReplicationPut**](ProductsApi.md#JobsReplicationPut) | **Put** /jobs/replication | Update status of jobs. Only stop is supported for now.
[**JobsScanIdLogGet**](ProductsApi.md#JobsScanIdLogGet) | **Get** /jobs/scan/{id}/log | Get job logs.
[**LdapPingPost**](ProductsApi.md#LdapPingPost) | **Post** /ldap/ping | Ping available ldap service.
[**LdapUsersImportPost**](ProductsApi.md#LdapUsersImportPost) | **Post** /ldap/users/import | Import selected available ldap users.
[**LdapUsersSearchPost**](ProductsApi.md#LdapUsersSearchPost) | **Post** /ldap/users/search | Search available ldap users.
[**LogsGet**](ProductsApi.md#LogsGet) | **Get** /logs | Get recent logs of the projects which the user is a member of
[**PoliciesReplicationGet**](ProductsApi.md#PoliciesReplicationGet) | **Get** /policies/replication | List filters policies by name and project_id
[**PoliciesReplicationIdGet**](ProductsApi.md#PoliciesReplicationIdGet) | **Get** /policies/replication/{id} | Get replication policy.
[**PoliciesReplicationIdPut**](ProductsApi.md#PoliciesReplicationIdPut) | **Put** /policies/replication/{id} | Put modifies name, description, target and enablement of policy.
[**PoliciesReplicationPost**](ProductsApi.md#PoliciesReplicationPost) | **Post** /policies/replication | Post creates a policy
[**ProjectsGet**](ProductsApi.md#ProjectsGet) | **Get** /projects | List projects
[**ProjectsHead**](ProductsApi.md#ProjectsHead) | **Head** /projects | Check if the project name user provided already exists.
[**ProjectsPost**](ProductsApi.md#ProjectsPost) | **Post** /projects | Create a new project.
[**ProjectsProjectIdDelete**](ProductsApi.md#ProjectsProjectIdDelete) | **Delete** /projects/{project_id} | Delete project by projectID
[**ProjectsProjectIdGet**](ProductsApi.md#ProjectsProjectIdGet) | **Get** /projects/{project_id} | Return specific project detail infomation
[**ProjectsProjectIdLogsGet**](ProductsApi.md#ProjectsProjectIdLogsGet) | **Get** /projects/{project_id}/logs | Get access logs accompany with a relevant project.
[**ProjectsProjectIdMembersGet**](ProductsApi.md#ProjectsProjectIdMembersGet) | **Get** /projects/{project_id}/members/ | Return a project&#39;s relevant role members.
[**ProjectsProjectIdMembersPost**](ProductsApi.md#ProjectsProjectIdMembersPost) | **Post** /projects/{project_id}/members/ | Add project role member accompany with relevant project and user.
[**ProjectsProjectIdMembersUserIdDelete**](ProductsApi.md#ProjectsProjectIdMembersUserIdDelete) | **Delete** /projects/{project_id}/members/{user_id} | Delete project role members accompany with relevant project and user.
[**ProjectsProjectIdMembersUserIdGet**](ProductsApi.md#ProjectsProjectIdMembersUserIdGet) | **Get** /projects/{project_id}/members/{user_id} | Return role members accompany with relevant project and user.
[**ProjectsProjectIdMembersUserIdPut**](ProductsApi.md#ProjectsProjectIdMembersUserIdPut) | **Put** /projects/{project_id}/members/{user_id} | Update project role members accompany with relevant project and user.
[**ProjectsProjectIdMetadatasGet**](ProductsApi.md#ProjectsProjectIdMetadatasGet) | **Get** /projects/{project_id}/metadatas | Get project metadata.
[**ProjectsProjectIdMetadatasMetaNameDelete**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNameDelete) | **Delete** /projects/{project_id}/metadatas/{meta_name} | Delete metadata of a project
[**ProjectsProjectIdMetadatasMetaNameGet**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNameGet) | **Get** /projects/{project_id}/metadatas/{meta_name} | Get project metadata
[**ProjectsProjectIdMetadatasMetaNamePut**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNamePut) | **Put** /projects/{project_id}/metadatas/{meta_name} | Update metadata of a project.
[**ProjectsProjectIdMetadatasPost**](ProductsApi.md#ProjectsProjectIdMetadatasPost) | **Post** /projects/{project_id}/metadatas | Add metadata for the project.
[**ProjectsProjectIdPut**](ProductsApi.md#ProjectsProjectIdPut) | **Put** /projects/{project_id} | Update properties for a selected project.
[**ReplicationsPost**](ProductsApi.md#ReplicationsPost) | **Post** /replications | Trigger the replication according to the specified policy.
[**RepositoriesGet**](ProductsApi.md#RepositoriesGet) | **Get** /repositories | Get repositories accompany with relevant project and repo name.
[**RepositoriesRepoNameDelete**](ProductsApi.md#RepositoriesRepoNameDelete) | **Delete** /repositories/{repo_name} | Delete a repository.
[**RepositoriesRepoNamePut**](ProductsApi.md#RepositoriesRepoNamePut) | **Put** /repositories/{repo_name} | Update description of the repository.
[**RepositoriesRepoNameSignaturesGet**](ProductsApi.md#RepositoriesRepoNameSignaturesGet) | **Get** /repositories/{repo_name}/signatures | Get signature information of a repository
[**RepositoriesRepoNameTagsGet**](ProductsApi.md#RepositoriesRepoNameTagsGet) | **Get** /repositories/{repo_name}/tags | Get tags of a relevant repository.
[**RepositoriesRepoNameTagsTagDelete**](ProductsApi.md#RepositoriesRepoNameTagsTagDelete) | **Delete** /repositories/{repo_name}/tags/{tag} | Delete a tag in a repository.
[**RepositoriesRepoNameTagsTagGet**](ProductsApi.md#RepositoriesRepoNameTagsTagGet) | **Get** /repositories/{repo_name}/tags/{tag} | Get the tag of the repository.
[**RepositoriesRepoNameTagsTagManifestGet**](ProductsApi.md#RepositoriesRepoNameTagsTagManifestGet) | **Get** /repositories/{repo_name}/tags/{tag}/manifest | Get manifests of a relevant repository.
[**RepositoriesRepoNameTagsTagScanPost**](ProductsApi.md#RepositoriesRepoNameTagsTagScanPost) | **Post** /repositories/{repo_name}/tags/{tag}/scan | Scan the image.
[**RepositoriesRepoNameTagsTagVulnerabilityDetailsGet**](ProductsApi.md#RepositoriesRepoNameTagsTagVulnerabilityDetailsGet) | **Get** /repositories/{repo_name}/tags/{tag}/vulnerability/details | Get vulnerability details of the image.
[**RepositoriesTopGet**](ProductsApi.md#RepositoriesTopGet) | **Get** /repositories/top | Get public repositories which are accessed most.
[**SearchGet**](ProductsApi.md#SearchGet) | **Get** /search | Search for projects and repositories
[**StatisticsGet**](ProductsApi.md#StatisticsGet) | **Get** /statistics | Get projects number and repositories number relevant to the user
[**SysteminfoGet**](ProductsApi.md#SysteminfoGet) | **Get** /systeminfo | Get general system info
[**SysteminfoGetcertGet**](ProductsApi.md#SysteminfoGetcertGet) | **Get** /systeminfo/getcert | Get default root certificate under OVA deployment.
[**SysteminfoVolumesGet**](ProductsApi.md#SysteminfoVolumesGet) | **Get** /systeminfo/volumes | Get system volume info (total/free size).
[**TargetsGet**](ProductsApi.md#TargetsGet) | **Get** /targets | List filters targets by name.
[**TargetsIdDelete**](ProductsApi.md#TargetsIdDelete) | **Delete** /targets/{id} | Delete specific replication&#39;s target.
[**TargetsIdGet**](ProductsApi.md#TargetsIdGet) | **Get** /targets/{id} | Get replication&#39;s target.
[**TargetsIdPoliciesGet**](ProductsApi.md#TargetsIdPoliciesGet) | **Get** /targets/{id}/policies/ | List the target relevant policies.
[**TargetsIdPut**](ProductsApi.md#TargetsIdPut) | **Put** /targets/{id} | Update replication&#39;s target.
[**TargetsPingPost**](ProductsApi.md#TargetsPingPost) | **Post** /targets/ping | Ping validates target.
[**TargetsPost**](ProductsApi.md#TargetsPost) | **Post** /targets | Create a new replication target.
[**UsersCurrentGet**](ProductsApi.md#UsersCurrentGet) | **Get** /users/current | Get current user info.
[**UsersGet**](ProductsApi.md#UsersGet) | **Get** /users | Get registered users of Harbor.
[**UsersPost**](ProductsApi.md#UsersPost) | **Post** /users | Creates a new user account.
[**UsersUserIdDelete**](ProductsApi.md#UsersUserIdDelete) | **Delete** /users/{user_id} | Mark a registered user as be removed.
[**UsersUserIdGet**](ProductsApi.md#UsersUserIdGet) | **Get** /users/{user_id} | Get a user&#39;s profile.
[**UsersUserIdPasswordPut**](ProductsApi.md#UsersUserIdPasswordPut) | **Put** /users/{user_id}/password | Change the password on a user that already exists.
[**UsersUserIdPut**](ProductsApi.md#UsersUserIdPut) | **Put** /users/{user_id} | Update a registered user to change his profile.
[**UsersUserIdSysadminPut**](ProductsApi.md#UsersUserIdSysadminPut) | **Put** /users/{user_id}/sysadmin | Update a registered user to change to be an administrator of Harbor.


# **ConfigurationsGet**
> Configurations ConfigurationsGet(ctx, )
Get system configurations.

This endpoint is for retrieving system configurations that only provides for admin user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Configurations**](Configurations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsPut**
> ConfigurationsPut(ctx, configurations)
Modify system configurations.

This endpoint is for modifying system configurations that only provides for admin user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configurations** | [**Configurations**](Configurations.md)| The configuration map can contain a subset of the attributes of the schema, which are to be updated. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsResetPost**
> ConfigurationsResetPost(ctx, )
Reset system configurations.

Reset system configurations from environment variables. Can only be accessed by admin user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailPingPost**
> EmailPingPost(ctx, optional)
Test connection and authentication with email server.

Test connection and authentication with email server.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiEmailPingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiEmailPingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | [**optional.Interface of EmailServerSetting**](EmailServerSetting.md)| Email server settings, if some of the settings are not assigned, they will be read from system configuration. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalSyncregistryPost**
> InternalSyncregistryPost(ctx, )
Sync repositories from registry to DB.

This endpoint is for syncing all repositories of registry with database.  

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsReplicationGet**
> []JobStatus JobsReplicationGet(ctx, policyId, optional)
List filters jobs according to the policy and repository

This endpoint let user list filters jobs according to the policy and repository. (if start_time and end_time are both null, list jobs of last 10 days) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyId** | **int32**| The ID of the policy that triggered this job. | 
 **optional** | ***ProductsApiJobsReplicationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiJobsReplicationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **num** | **optional.Int32**| The return list length number. | 
 **endTime** | **optional.Int64**| The end time of jobs done. (Timestamp) | 
 **startTime** | **optional.Int64**| The start time of jobs. (Timestamp) | 
 **repository** | **optional.String**| The respond jobs list filter by repository name. | 
 **status** | **optional.String**| The respond jobs list filter by status. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]JobStatus**](JobStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsReplicationIdDelete**
> JobsReplicationIdDelete(ctx, id)
Delete specific ID job.

This endpoint is aimed to remove specific ID job from jobservice. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Delete job ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsReplicationIdLogGet**
> JobsReplicationIdLogGet(ctx, id)
Get job logs.

This endpoint let user search job logs filtered by specific ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Relevant job ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsReplicationPut**
> JobsReplicationPut(ctx, policyinfo)
Update status of jobs. Only stop is supported for now.

The endpoint is used to stop the replication jobs of a policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyinfo** | [**UpdateJobs**](UpdateJobs.md)| The policy ID and status. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JobsScanIdLogGet**
> JobsScanIdLogGet(ctx, id)
Get job logs.

This endpoint let user get scan job logs filtered by specific ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Relevant job ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapPingPost**
> LdapPingPost(ctx, optional)
Ping available ldap service.

This endpoint ping the available ldap service for test related configuration parameters.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiLdapPingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLdapPingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapconf** | [**optional.Interface of LdapConf**](LdapConf.md)| ldap configuration. support input ldap service configuration. If it&#39;s a empty request, will load current configuration from the system. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersImportPost**
> LdapUsersImportPost(ctx, uidList)
Import selected available ldap users.

This endpoint adds the selected available ldap users to harbor based on related configuration parameters from the system. System will try to guess the user email address and realname, add to harbor user information.  If have errors when import user, will return the list of importing failed uid and the failed reason. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uidList** | [**LdapImportUsers**](LdapImportUsers.md)| The uid listed for importing. This list will check users validity of ldap service based on configuration from the system. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersSearchPost**
> []LdapUsers LdapUsersSearchPost(ctx, optional)
Search available ldap users.

This endpoint searches the available ldap users based on related configuration parameters. Support searched by input ladp configuration, load configuration from the system and specific filter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiLdapUsersSearchPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLdapUsersSearchPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Registered user ID | 
 **ldapConf** | [**optional.Interface of LdapConf**](LdapConf.md)| ldap search configuration. ldapconf field can input ldap service configuration. If this item are blank, will load default configuration will load current configuration from the system. | 

### Return type

[**[]LdapUsers**](LdapUsers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsGet**
> []AccessLog LogsGet(ctx, optional)
Get recent logs of the projects which the user is a member of

This endpoint let user see the recent operation logs of the projects which he is member of  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiLogsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Username of the operator. | 
 **repository** | **optional.String**| The name of repository | 
 **tag** | **optional.String**| The name of tag | 
 **operation** | **optional.String**| The operation | 
 **beginTimestamp** | **optional.String**| The begin timestamp | 
 **endTimestamp** | **optional.String**| The end timestamp | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationGet**
> []RepPolicy PoliciesReplicationGet(ctx, optional)
List filters policies by name and project_id

This endpoint let user list filters policies by name and project_id, if name and project_id are nil, list returns all policies 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiPoliciesReplicationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiPoliciesReplicationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The replication&#39;s policy name. | 
 **projectId** | **optional.Int64**| Relevant project ID. | 
 **page** | **optional.Int32**| The page nubmer. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]RepPolicy**](RepPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationIdGet**
> RepPolicy PoliciesReplicationIdGet(ctx, id)
Get replication policy.

This endpoint let user search replication policy by specific ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| policy ID | 

### Return type

[**RepPolicy**](RepPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationIdPut**
> PoliciesReplicationIdPut(ctx, id, policyupdate)
Put modifies name, description, target and enablement of policy.

This endpoint let user update policy name, description, target and enablement. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| policy ID | 
  **policyupdate** | [**RepPolicy**](RepPolicy.md)| Updated properties of the replication policy. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PoliciesReplicationPost**
> PoliciesReplicationPost(ctx, policyinfo)
Post creates a policy

This endpoint let user creates a policy, and if it is enabled, the replication will be triggered right now. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyinfo** | [**RepPolicy**](RepPolicy.md)| Create new policy. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsGet**
> []Project ProjectsGet(ctx, optional)
List projects

This endpoint returns all projects created by Harbor, and can be filtered by project name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiProjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of project. | 
 **public** | **optional.Bool**| The project is public or private. | 
 **owner** | **optional.String**| The name of project owner. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsHead**
> ProjectsHead(ctx, projectName)
Check if the project name user provided already exists.

This endpoint is used to check if the project name user provided already exist. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectName** | **string**| Project name for checking exists. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsPost**
> ProjectsPost(ctx, project)
Create a new project.

This endpoint is for user to create a new project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**ProjectReq**](ProjectReq.md)| New created project. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdDelete**
> ProjectsProjectIdDelete(ctx, projectId)
Delete project by projectID

This endpoint is aimed to delete project by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID of project which will be deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdGet**
> Project ProjectsProjectIdGet(ctx, projectId)
Return specific project detail infomation

This endpoint returns specific project information by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID for filtering results. | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdLogsGet**
> []AccessLog ProjectsProjectIdLogsGet(ctx, projectId, optional)
Get access logs accompany with a relevant project.

This endpoint let user search access logs filtered by operations and date time ranges. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID | 
 **optional** | ***ProductsApiProjectsProjectIdLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdLogsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Username of the operator. | 
 **repository** | **optional.String**| The name of repository | 
 **tag** | **optional.String**| The name of tag | 
 **operation** | **optional.String**| The operation | 
 **beginTimestamp** | **optional.String**| The begin timestamp | 
 **endTimestamp** | **optional.String**| The end timestamp | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersGet**
> []User ProjectsProjectIdMembersGet(ctx, projectId)
Return a project's relevant role members.

This endpoint is for user to search a specified project's relevant role members. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersPost**
> ProjectsProjectIdMembersPost(ctx, projectId, optional)
Add project role member accompany with relevant project and user.

This endpoint is for user to add project role member accompany with relevant project and user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
 **optional** | ***ProductsApiProjectsProjectIdMembersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdMembersPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roles** | [**optional.Interface of RoleParam**](RoleParam.md)| Role members for adding to relevant project. Only one role is supported in the role list. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersUserIdDelete**
> ProjectsProjectIdMembersUserIdDelete(ctx, projectId, userId)
Delete project role members accompany with relevant project and user.

This endpoint is aimed to remove project role members already added to the relevant project and user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **userId** | **int32**| Relevant user ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersUserIdGet**
> []Role ProjectsProjectIdMembersUserIdGet(ctx, projectId, userId)
Return role members accompany with relevant project and user.

This endpoint is for user to get role members accompany with relevant project and user.  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID | 
  **userId** | **int32**| Relevant user ID | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersUserIdPut**
> ProjectsProjectIdMembersUserIdPut(ctx, projectId, userId, optional)
Update project role members accompany with relevant project and user.

This endpoint is for user to update current project role members accompany with relevant project and user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **userId** | **int32**| Relevant user ID. | 
 **optional** | ***ProductsApiProjectsProjectIdMembersUserIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiProjectsProjectIdMembersUserIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roles** | [**optional.Interface of RoleParam**](RoleParam.md)| Updates for roles and username. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasGet**
> ProjectMetadata ProjectsProjectIdMetadatasGet(ctx, projectId)
Get project metadata.

This endpoint returns metadata of the project specified by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 

### Return type

[**ProjectMetadata**](ProjectMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNameDelete**
> ProjectsProjectIdMetadatasMetaNameDelete(ctx, projectId, metaName)
Delete metadata of a project

This endpoint is aimed to delete metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNameGet**
> ProjectMetadata ProjectsProjectIdMetadatasMetaNameGet(ctx, projectId, metaName)
Get project metadata

This endpoint returns specified metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID for filtering results. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

[**ProjectMetadata**](ProjectMetadata.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNamePut**
> ProjectsProjectIdMetadatasMetaNamePut(ctx, projectId, metaName)
Update metadata of a project.

This endpoint is aimed to update the metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasPost**
> ProjectsProjectIdMetadatasPost(ctx, projectId, metadata)
Add metadata for the project.

This endpoint is aimed to add metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Selected project ID. | 
  **metadata** | [**ProjectMetadata**](ProjectMetadata.md)| The metadata of project. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdPut**
> ProjectsProjectIdPut(ctx, projectId, project)
Update properties for a selected project.

This endpoint is aimed to update the properties of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Selected project ID. | 
  **project** | [**Project**](Project.md)| Updates of project. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationsPost**
> ReplicationsPost(ctx, policyID)
Trigger the replication according to the specified policy.

This endpoint is used to trigger a replication. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policyID** | [**Replication**](Replication.md)| The ID of replication policy. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesGet**
> []Repository RepositoriesGet(ctx, projectId, optional)
Get repositories accompany with relevant project and repo name.

This endpoint let user search repositories accompanying with relevant project ID and repo name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int32**| Relevant project ID. | 
 **optional** | ***ProductsApiRepositoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRepositoriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| Repo name for filtering results. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameDelete**
> RepositoriesRepoNameDelete(ctx, repoName)
Delete a repository.

This endpoint let user delete a repository with name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNamePut**
> RepositoriesRepoNamePut(ctx, repoName, description)
Update description of the repository.

This endpoint is used to update description of the repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 
  **description** | [**RepositoryDescription**](RepositoryDescription.md)| The description of the repository. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameSignaturesGet**
> []RepoSignature RepositoriesRepoNameSignaturesGet(ctx, repoName)
Get signature information of a repository

This endpoint aims to retrieve signature information of a repository, the data is from the nested notary instance of Harbor. If the repository does not have any signature information in notary, this API will return an empty list with response code 200, instead of 404 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| repository name. | 

### Return type

[**[]RepoSignature**](RepoSignature.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsGet**
> []DetailedTag RepositoriesRepoNameTagsGet(ctx, repoName)
Get tags of a relevant repository.

This endpoint aims to retrieve tags from a relevant repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Relevant repository name. | 

### Return type

[**[]DetailedTag**](DetailedTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagDelete**
> RepositoriesRepoNameTagsTagDelete(ctx, repoName, tag)
Delete a tag in a repository.

This endpoint let user delete tags with repo name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 
  **tag** | **string**| Tag of a repository. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagGet**
> DetailedTag RepositoriesRepoNameTagsTagGet(ctx, repoName, tag)
Get the tag of the repository.

This endpoint aims to retrieve the tag of the repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Relevant repository name. | 
  **tag** | **string**| Tag of the repository. | 

### Return type

[**DetailedTag**](DetailedTag.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagManifestGet**
> Manifest RepositoriesRepoNameTagsTagManifestGet(ctx, repoName, tag, optional)
Get manifests of a relevant repository.

This endpoint aims to retreive manifests from a relevant repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 
 **optional** | ***ProductsApiRepositoriesRepoNameTagsTagManifestGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRepositoriesRepoNameTagsTagManifestGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version of manifest, valid value are \&quot;v1\&quot; and \&quot;v2\&quot;, default is \&quot;v2\&quot; | 

### Return type

[**Manifest**](Manifest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagScanPost**
> RepositoriesRepoNameTagsTagScanPost(ctx, repoName, tag)
Scan the image.

Trigger jobservice to call Clair API to scan the image identified by the repo_name and tag.  Only project admins have permission to scan images under the project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagVulnerabilityDetailsGet**
> []VulnerabilityItem RepositoriesRepoNameTagsTagVulnerabilityDetailsGet(ctx, repoName, tag)
Get vulnerability details of the image.

Call Clair API to get the vulnerability based on the previous successful scan. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 

### Return type

[**[]VulnerabilityItem**](VulnerabilityItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesTopGet**
> []Repository RepositoriesTopGet(ctx, optional)
Get public repositories which are accessed most.

This endpoint aims to let users see the most popular public repositories 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiRepositoriesTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiRepositoriesTopGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of the requested public repositories, default is 10 if not provided. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGet**
> []Search SearchGet(ctx, q)
Search for projects and repositories

The Search endpoint returns information about the projects and repositories offered at public status or related to the current logged in user. The response includes the project and repository list in a proper display order. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| Search parameter for project and repository name. | 

### Return type

[**[]Search**](Search.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsGet**
> StatisticMap StatisticsGet(ctx, )
Get projects number and repositories number relevant to the user

This endpoint is aimed to statistic all of the projects number and repositories number relevant to the logined user, also the public projects number and repositories number. If the user is admin, he can also get total projects number and total repositories number. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StatisticMap**](StatisticMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGet**
> interface{} SysteminfoGet(ctx, )
Get general system info

This API is for retrieving general system info, this can be called by anonymous request. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGetcertGet**
> SysteminfoGetcertGet(ctx, )
Get default root certificate under OVA deployment.

This endpoint is for downloading a default root certificate that only provides for admin user under OVA deployment. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoVolumesGet**
> interface{} SysteminfoVolumesGet(ctx, )
Get system volume info (total/free size).

This endpoint is for retrieving system volume info that only provides for admin user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsGet**
> []RepTarget TargetsGet(ctx, optional)
List filters targets by name.

This endpoint let user list filters targets by name, if name is nil, list returns all targets. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiTargetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiTargetsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The replication&#39;s target name. | 

### Return type

[**[]RepTarget**](RepTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdDelete**
> TargetsIdDelete(ctx, id)
Delete specific replication's target.

This endpoint is for to delete specific replication's target. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The replication&#39;s target ID. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdGet**
> RepTarget TargetsIdGet(ctx, id)
Get replication's target.

This endpoint is for get specific replication's target.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The replication&#39;s target ID. | 

### Return type

[**RepTarget**](RepTarget.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdPoliciesGet**
> []RepPolicy TargetsIdPoliciesGet(ctx, id)
List the target relevant policies.

This endpoint list policies filter with specific replication's target ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The replication&#39;s target ID. | 

### Return type

[**[]RepPolicy**](RepPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsIdPut**
> TargetsIdPut(ctx, id, repoTarget)
Update replication's target.

This endpoint is for update specific replication's target. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The replication&#39;s target ID. | 
  **repoTarget** | [**PutTarget**](PutTarget.md)| Updates of replication&#39;s target. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsPingPost**
> TargetsPingPost(ctx, target)
Ping validates target.

This endpoint is for ping validates whether the target is reachable and whether the credential is valid. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **target** | [**PingTarget**](PingTarget.md)| The target object. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TargetsPost**
> TargetsPost(ctx, reptarget)
Create a new replication target.

This endpoint is for user to create a new replication target. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reptarget** | [**RepTargetPost**](RepTargetPost.md)| New created replication target. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersCurrentGet**
> User UsersCurrentGet(ctx, )
Get current user info.

This endpoint is to get the current user infomation. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGet**
> []User UsersGet(ctx, optional)
Get registered users of Harbor.

This endpoint is for user to search registered users, support for filtering results with username.Notice, by now this operation is only for administrator. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductsApiUsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsApiUsersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Username for filtering results. | 
 **email** | **optional.String**| Email for filtering results. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPost**
> UsersPost(ctx, user)
Creates a new user account.

This endpoint is to create a user if the user does not already exist. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | [**User**](User.md)| New created user. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdDelete**
> UsersUserIdDelete(ctx, userId)
Mark a registered user as be removed.

This endpoint let administrator of Harbor mark a registered user as be removed.It actually won't be deleted from DB. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| User ID for marking as to be removed. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdGet**
> UsersUserIdGet(ctx, userId)
Get a user's profile.

Get user's profile with user id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPasswordPut**
> UsersUserIdPasswordPut(ctx, userId, password)
Change the password on a user that already exists.

This endpoint is for user to update password. Users with the admin role can change any user's password. Guest users can change only their own password. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID. | 
  **password** | [**Password**](Password.md)| Password to be updated. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPut**
> UsersUserIdPut(ctx, userId, profile)
Update a registered user to change his profile.

This endpoint let a registered user change his profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID | 
  **profile** | [**UserProfile**](UserProfile.md)| Only email, realname and comment can be modified. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdSysadminPut**
> UsersUserIdSysadminPut(ctx, userId, hasAdminRole)
Update a registered user to change to be an administrator of Harbor.

This endpoint let a registered user change to be an administrator of Harbor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID | 
  **hasAdminRole** | [**HasAdminRole**](HasAdminRole.md)| Toggle a user to admin or not. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

