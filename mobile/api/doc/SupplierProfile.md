# openapi.model.SupplierProfile

## Load the model package
```dart
import 'package:openapi/api.dart';
```

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**createdAt** | **String** |  | [optional] 
**createdBy** | **int** |  | [optional] 
**createdByString** | **String** |  | [optional] 
**updatedAt** | **String** |  | [optional] 
**groupId** | **int** |  | [optional] 
**name** | **String** | Canonical supplier display name | [optional] 
**normalisedName** | **String** | Normalised canonical name used for matching | [optional] 
**expectedDocumentCurrencyCode** | **String** | Optional expected ISO 4217 document currency | [optional] 
**enabled** | **bool** | Disabled profiles remain stored but do not match | [optional] 
**categories** | [**BuiltList&lt;Category&gt;**](Category.md) |  | [optional] 
**tags** | [**BuiltList&lt;Tag&gt;**](Tag.md) |  | [optional] 
**aliases** | [**BuiltList&lt;SupplierProfileAlias&gt;**](SupplierProfileAlias.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


