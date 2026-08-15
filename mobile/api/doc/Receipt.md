# openapi.model.Receipt

## Load the model package
```dart
import 'package:openapi/api.dart';
```

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**amount** | **String** | Effective amount in the group's base currency |
**documentAmount** | **String** | Original total printed on the receipt evidence |
**documentCurrencyCode** | **String** | ISO 4217 currency printed on the receipt evidence |
**estimatedBaseAmount** | **String** | Historical-rate estimate in the group's base currency | [optional]
**fxRate** | **String** | Exact document-to-base exchange rate used for the estimate | [optional]
**fxDate** | [**DateTime**](DateTime.md) | Effective date returned by the FX provider | [optional]
**fxProvider** | **String** | Provider used for the historical estimate | [optional]
**fxRetrievedAt** | [**DateTime**](DateTime.md) | UTC instant at which the rate was retrieved | [optional]
**fxStatus** | [**FxStatus**](FxStatus.md) |  |
**categories** | [**BuiltList&lt;Category&gt;**](Category.md) | Categories associated to receipt |
**comments** | [**BuiltList&lt;Comment&gt;**](Comment.md) | Comments associated to receipt |
**customFields** | [**BuiltList&lt;CustomFieldValue&gt;**](CustomFieldValue.md) | Custom fields associated to receipt |
**createdAt** | **String** |  | [optional]
**createdBy** | **int** |  | [optional]
**date** | **String** | Receipt date |
**groupId** | **int** | Group foreign key |
**id** | **int** |  |
**imageFiles** | [**BuiltList&lt;FileData&gt;**](FileData.md) | Files associated to receipt | [optional]
**name** | **String** | Receipt name |
**paidByUserId** | **int** | User paid foreign key |
**receiptItems** | [**BuiltList&lt;Item&gt;**](Item.md) | Items associated to receipt |
**resolvedDate** | **String** | Date resolved | [optional]
**status** | [**ReceiptStatus**](ReceiptStatus.md) |  |
**tags** | [**BuiltList&lt;Tag&gt;**](Tag.md) | Tags associated to receipt |
**updatedAt** | **String** |  | [optional]
**createdByString** | **String** | Created by string, which is anything that is not a user | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


