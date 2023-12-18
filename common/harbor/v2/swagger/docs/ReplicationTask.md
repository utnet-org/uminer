# ReplicationTask

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the task | [optional] [default to null]
**ExecutionId** | **int32** | The ID of the execution that the task belongs to | [optional] [default to null]
**Status** | **string** | The status of the task | [optional] [default to null]
**JobId** | **string** | The ID of the underlying job that the task related to | [optional] [default to null]
**Operation** | **string** | The operation of the task | [optional] [default to null]
**ResourceType** | **string** | The type of the resource that the task operates | [optional] [default to null]
**SrcResource** | **string** | The source resource that the task operates | [optional] [default to null]
**DstResource** | **string** | The destination resource that the task operates | [optional] [default to null]
**StartTime** | [**time.Time**](time.Time.md) | The start time of the task | [optional] [default to null]
**EndTime** | [**time.Time**](time.Time.md) | The end time of the task | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


