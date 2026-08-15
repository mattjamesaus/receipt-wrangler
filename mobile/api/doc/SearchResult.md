# openapi.model.SearchResult

## Load the model package
```dart
import 'package:openapi/api.dart';
```

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  |
**name** | **String** |  |
**type** | **String** |  |
**groupId** | **int** |  |
**date** | **String** |  |
**amount** | **String** | Effective amount in the group's base currency |
**documentAmount** | **String** | Original total printed on the receipt evidence |
**documentCurrencyCode** | **String** | ISO 4217 currency printed on the receipt evidence |
**fxStatus** | [**FxStatus**](FxStatus.md) |  |
**receiptStatus** | [**ReceiptStatus**](ReceiptStatus.md) |  | [optional]
**paidByUserId** | **int** |  | [optional]
**createdAt** | **String** |  |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


