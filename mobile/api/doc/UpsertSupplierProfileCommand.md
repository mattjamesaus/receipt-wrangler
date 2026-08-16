# openapi.model.UpsertSupplierProfileCommand

## Load the model package
```dart
import 'package:openapi/api.dart';
```

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Canonical supplier display name | 
**aliases** | **BuiltList&lt;String&gt;** | Additional names that should match this profile | [optional] 
**categoryIds** | **BuiltList&lt;int&gt;** | Default category ids from the catalogue | [optional] 
**tagIds** | **BuiltList&lt;int&gt;** | Default tag ids from the catalogue | [optional] 
**expectedDocumentCurrencyCode** | **String** | Optional expected ISO 4217 document currency | [optional] 
**enabled** | **bool** | Whether the profile participates in matching. Defaults to true on create. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


