//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:openapi/src/model/fx_status.dart';
import 'package:openapi/src/model/receipt_status.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'search_result.g.dart';

/// SearchResult
///
/// Properties:
/// * [id]
/// * [name]
/// * [type]
/// * [groupId]
/// * [date]
/// * [amount] - Effective amount in the group's base currency
/// * [documentAmount] - Original total printed on the receipt evidence
/// * [documentCurrencyCode] - ISO 4217 currency printed on the receipt evidence
/// * [fxStatus]
/// * [receiptStatus]
/// * [paidByUserId]
/// * [createdAt]
@BuiltValue()
abstract class SearchResult implements Built<SearchResult, SearchResultBuilder> {
  @BuiltValueField(wireName: r'id')
  int get id;

  @BuiltValueField(wireName: r'name')
  String get name;

  @BuiltValueField(wireName: r'type')
  String get type;

  @BuiltValueField(wireName: r'groupId')
  int get groupId;

  @BuiltValueField(wireName: r'date')
  String get date;

  /// Effective amount in the group's base currency
  @BuiltValueField(wireName: r'amount')
  String get amount;

  /// Original total printed on the receipt evidence
  @BuiltValueField(wireName: r'documentAmount')
  String get documentAmount;

  /// ISO 4217 currency printed on the receipt evidence
  @BuiltValueField(wireName: r'documentCurrencyCode')
  String get documentCurrencyCode;

  @BuiltValueField(wireName: r'fxStatus')
  FxStatus get fxStatus;
  // enum fxStatusEnum {  DOMESTIC,  ESTIMATED,  CONFIRMED,  NEEDS_REVIEW,  };

  @BuiltValueField(wireName: r'receiptStatus')
  ReceiptStatus? get receiptStatus;
  // enum receiptStatusEnum {  OPEN,  NEEDS_ATTENTION,  RESOLVED,  DRAFT,  ,  };

  @BuiltValueField(wireName: r'paidByUserId')
  int? get paidByUserId;

  @BuiltValueField(wireName: r'createdAt')
  String get createdAt;

  SearchResult._();

  factory SearchResult([void updates(SearchResultBuilder b)]) = _$SearchResult;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(SearchResultBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<SearchResult> get serializer => _$SearchResultSerializer();
}

class _$SearchResultSerializer implements PrimitiveSerializer<SearchResult> {
  @override
  final Iterable<Type> types = const [SearchResult, _$SearchResult];

  @override
  final String wireName = r'SearchResult';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    SearchResult object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    yield r'id';
    yield serializers.serialize(
      object.id,
      specifiedType: const FullType(int),
    );
    yield r'name';
    yield serializers.serialize(
      object.name,
      specifiedType: const FullType(String),
    );
    yield r'type';
    yield serializers.serialize(
      object.type,
      specifiedType: const FullType(String),
    );
    yield r'groupId';
    yield serializers.serialize(
      object.groupId,
      specifiedType: const FullType(int),
    );
    yield r'date';
    yield serializers.serialize(
      object.date,
      specifiedType: const FullType(String),
    );
    yield r'amount';
    yield serializers.serialize(
      object.amount,
      specifiedType: const FullType(String),
    );
    yield r'documentAmount';
    yield serializers.serialize(
      object.documentAmount,
      specifiedType: const FullType(String),
    );
    yield r'documentCurrencyCode';
    yield serializers.serialize(
      object.documentCurrencyCode,
      specifiedType: const FullType(String),
    );
    yield r'fxStatus';
    yield serializers.serialize(
      object.fxStatus,
      specifiedType: const FullType(FxStatus),
    );
    if (object.receiptStatus != null) {
      yield r'receiptStatus';
      yield serializers.serialize(
        object.receiptStatus,
        specifiedType: const FullType(ReceiptStatus),
      );
    }
    if (object.paidByUserId != null) {
      yield r'paidByUserId';
      yield serializers.serialize(
        object.paidByUserId,
        specifiedType: const FullType(int),
      );
    }
    yield r'createdAt';
    yield serializers.serialize(
      object.createdAt,
      specifiedType: const FullType(String),
    );
  }

  @override
  Object serialize(
    Serializers serializers,
    SearchResult object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required SearchResultBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'id':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(int),
          ) as int;
          result.id = valueDes;
          break;
        case r'name':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.name = valueDes;
          break;
        case r'type':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.type = valueDes;
          break;
        case r'groupId':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(int),
          ) as int;
          result.groupId = valueDes;
          break;
        case r'date':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.date = valueDes;
          break;
        case r'amount':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.amount = valueDes;
          break;
        case r'documentAmount':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.documentAmount = valueDes;
          break;
        case r'documentCurrencyCode':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.documentCurrencyCode = valueDes;
          break;
        case r'fxStatus':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(FxStatus),
          ) as FxStatus;
          result.fxStatus = valueDes;
          break;
        case r'receiptStatus':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(ReceiptStatus),
          ) as ReceiptStatus;
          result.receiptStatus = valueDes;
          break;
        case r'paidByUserId':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(int),
          ) as int;
          result.paidByUserId = valueDes;
          break;
        case r'createdAt':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.createdAt = valueDes;
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  SearchResult deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = SearchResultBuilder();
    final serializedList = (serialized as Iterable<Object?>).toList();
    final unhandled = <Object?>[];
    _deserializeProperties(
      serializers,
      serialized,
      specifiedType: specifiedType,
      serializedList: serializedList,
      unhandled: unhandled,
      result: result,
    );
    return result.build();
  }
}

