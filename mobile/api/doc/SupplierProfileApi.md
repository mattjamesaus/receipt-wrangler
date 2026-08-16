# openapi.api.SupplierProfileApi

## Load the API package
```dart
import 'package:openapi/api.dart';
```

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createSupplierProfile**](SupplierProfileApi.md#createsupplierprofile) | **POST** /group/{groupId}/supplierProfile | Create a supplier profile
[**deleteSupplierProfile**](SupplierProfileApi.md#deletesupplierprofile) | **DELETE** /group/{groupId}/supplierProfile/{profileId} | Delete a supplier profile
[**getSupplierProfileById**](SupplierProfileApi.md#getsupplierprofilebyid) | **GET** /group/{groupId}/supplierProfile/{profileId} | Get a supplier profile
[**getSupplierProfilesForGroup**](SupplierProfileApi.md#getsupplierprofilesforgroup) | **GET** /group/{groupId}/supplierProfile | List supplier profiles for a group
[**resolveSupplierProfile**](SupplierProfileApi.md#resolvesupplierprofile) | **POST** /group/{groupId}/supplierProfile/resolve | Resolve a supplier profile for a receipt name
[**updateSupplierProfile**](SupplierProfileApi.md#updatesupplierprofile) | **PUT** /group/{groupId}/supplierProfile/{profileId} | Update a supplier profile


# **createSupplierProfile**
> SupplierProfile createSupplierProfile(groupId, upsertSupplierProfileCommand)

Create a supplier profile

Creates a group-scoped supplier profile with optional default categories, tags, and expected document currency. At least one default is required. Canonical names and aliases must be unique within the group after normalisation. Gated by group.receipts.update.

### Example
```dart
import 'package:openapi/api.dart';
// TODO Configure API key authorization: apiKeyAuth
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKey = 'YOUR_API_KEY';
// uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKeyPrefix = 'Bearer';

final api = Openapi().getSupplierProfileApi();
final int groupId = 56; // int | Group id that owns the supplier profiles
final UpsertSupplierProfileCommand upsertSupplierProfileCommand = ; // UpsertSupplierProfileCommand | 

try {
    final response = api.createSupplierProfile(groupId, upsertSupplierProfileCommand);
    print(response);
} catch on DioException (e) {
    print('Exception when calling SupplierProfileApi->createSupplierProfile: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int**| Group id that owns the supplier profiles | 
 **upsertSupplierProfileCommand** | [**UpsertSupplierProfileCommand**](UpsertSupplierProfileCommand.md)|  | 

### Return type

[**SupplierProfile**](SupplierProfile.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteSupplierProfile**
> deleteSupplierProfile(groupId, profileId)

Delete a supplier profile

Deletes a supplier profile and its aliases. Receipts and source files are unchanged. Gated by group.receipts.update.

### Example
```dart
import 'package:openapi/api.dart';
// TODO Configure API key authorization: apiKeyAuth
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKey = 'YOUR_API_KEY';
// uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKeyPrefix = 'Bearer';

final api = Openapi().getSupplierProfileApi();
final int groupId = 56; // int | Group id that owns the supplier profile
final int profileId = 56; // int | Supplier profile id

try {
    api.deleteSupplierProfile(groupId, profileId);
} catch on DioException (e) {
    print('Exception when calling SupplierProfileApi->deleteSupplierProfile: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int**| Group id that owns the supplier profile | 
 **profileId** | **int**| Supplier profile id | 

### Return type

void (empty response body)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getSupplierProfileById**
> SupplierProfile getSupplierProfileById(groupId, profileId)

Get a supplier profile

Returns one supplier profile in the group. Categories and tags are filtered to those the caller is permitted to use. Gated by group.receipts.create.

### Example
```dart
import 'package:openapi/api.dart';
// TODO Configure API key authorization: apiKeyAuth
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKey = 'YOUR_API_KEY';
// uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKeyPrefix = 'Bearer';

final api = Openapi().getSupplierProfileApi();
final int groupId = 56; // int | Group id that owns the supplier profile
final int profileId = 56; // int | Supplier profile id

try {
    final response = api.getSupplierProfileById(groupId, profileId);
    print(response);
} catch on DioException (e) {
    print('Exception when calling SupplierProfileApi->getSupplierProfileById: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int**| Group id that owns the supplier profile | 
 **profileId** | **int**| Supplier profile id | 

### Return type

[**SupplierProfile**](SupplierProfile.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getSupplierProfilesForGroup**
> BuiltList<SupplierProfile> getSupplierProfilesForGroup(groupId)

List supplier profiles for a group

Returns every supplier profile in the group. Categories and tags on each profile are filtered to those the caller is permitted to use. Gated by group.receipts.create.

### Example
```dart
import 'package:openapi/api.dart';
// TODO Configure API key authorization: apiKeyAuth
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKey = 'YOUR_API_KEY';
// uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKeyPrefix = 'Bearer';

final api = Openapi().getSupplierProfileApi();
final int groupId = 56; // int | Group id that owns the supplier profiles

try {
    final response = api.getSupplierProfilesForGroup(groupId);
    print(response);
} catch on DioException (e) {
    print('Exception when calling SupplierProfileApi->getSupplierProfilesForGroup: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int**| Group id that owns the supplier profiles | 

### Return type

[**BuiltList&lt;SupplierProfile&gt;**](SupplierProfile.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **resolveSupplierProfile**
> ResolveSupplierProfileResponse resolveSupplierProfile(groupId, resolveSupplierProfileCommand)

Resolve a supplier profile for a receipt name

Returns the single enabled profile whose canonical name or alias matches the supplied receipt name after normalisation. Ambiguous or disabled matches return a null profile. Never applies defaults. Gated by group.receipts.create.

### Example
```dart
import 'package:openapi/api.dart';
// TODO Configure API key authorization: apiKeyAuth
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKey = 'YOUR_API_KEY';
// uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKeyPrefix = 'Bearer';

final api = Openapi().getSupplierProfileApi();
final int groupId = 56; // int | Group id to resolve within
final ResolveSupplierProfileCommand resolveSupplierProfileCommand = ; // ResolveSupplierProfileCommand | 

try {
    final response = api.resolveSupplierProfile(groupId, resolveSupplierProfileCommand);
    print(response);
} catch on DioException (e) {
    print('Exception when calling SupplierProfileApi->resolveSupplierProfile: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int**| Group id to resolve within | 
 **resolveSupplierProfileCommand** | [**ResolveSupplierProfileCommand**](ResolveSupplierProfileCommand.md)|  | 

### Return type

[**ResolveSupplierProfileResponse**](ResolveSupplierProfileResponse.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateSupplierProfile**
> SupplierProfile updateSupplierProfile(groupId, profileId, upsertSupplierProfileCommand)

Update a supplier profile

Replaces a supplier profile's name, aliases, defaults, expected currency, and enabled state. Enable/disable is part of this update. Gated by group.receipts.update.

### Example
```dart
import 'package:openapi/api.dart';
// TODO Configure API key authorization: apiKeyAuth
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKey = 'YOUR_API_KEY';
// uncomment below to setup prefix (e.g. Bearer) for API key, if needed
//defaultApiClient.getAuthentication<ApiKeyAuth>('apiKeyAuth').apiKeyPrefix = 'Bearer';

final api = Openapi().getSupplierProfileApi();
final int groupId = 56; // int | Group id that owns the supplier profile
final int profileId = 56; // int | Supplier profile id
final UpsertSupplierProfileCommand upsertSupplierProfileCommand = ; // UpsertSupplierProfileCommand | 

try {
    final response = api.updateSupplierProfile(groupId, profileId, upsertSupplierProfileCommand);
    print(response);
} catch on DioException (e) {
    print('Exception when calling SupplierProfileApi->updateSupplierProfile: $e\n');
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupId** | **int**| Group id that owns the supplier profile | 
 **profileId** | **int**| Supplier profile id | 
 **upsertSupplierProfileCommand** | [**UpsertSupplierProfileCommand**](UpsertSupplierProfileCommand.md)|  | 

### Return type

[**SupplierProfile**](SupplierProfile.md)

### Authorization

[apiKeyAuth](../README.md#apiKeyAuth), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

