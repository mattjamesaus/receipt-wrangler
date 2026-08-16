//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_collection/built_collection.dart';
import 'package:openapi/src/model/category.dart';
import 'package:openapi/src/model/tag.dart';
import 'package:openapi/src/model/supplier_profile_alias.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'supplier_profile.g.dart';

/// Group-scoped supplier profile that stores optional receipt-review defaults. A profile is a suggestion, never an automatic classification.
///
/// Properties:
/// * [id] 
/// * [createdAt] 
/// * [createdBy] 
/// * [createdByString] 
/// * [updatedAt] 
/// * [groupId] 
/// * [name] - Canonical supplier display name
/// * [normalisedName] - Normalised canonical name used for matching
/// * [expectedDocumentCurrencyCode] - Optional expected ISO 4217 document currency
/// * [enabled] - Disabled profiles remain stored but do not match
/// * [autoApply] - When true, matching receipts created via email, quick scan, or the create-receipt API receive these defaults automatically. Extracted document currency is never overwritten.
/// * [categories] 
/// * [tags] 
/// * [aliases] 
@BuiltValue()
abstract class SupplierProfile implements Built<SupplierProfile, SupplierProfileBuilder> {
  @BuiltValueField(wireName: r'id')
  int? get id;

  @BuiltValueField(wireName: r'createdAt')
  String? get createdAt;

  @BuiltValueField(wireName: r'createdBy')
  int? get createdBy;

  @BuiltValueField(wireName: r'createdByString')
  String? get createdByString;

  @BuiltValueField(wireName: r'updatedAt')
  String? get updatedAt;

  @BuiltValueField(wireName: r'groupId')
  int? get groupId;

  /// Canonical supplier display name
  @BuiltValueField(wireName: r'name')
  String? get name;

  /// Normalised canonical name used for matching
  @BuiltValueField(wireName: r'normalisedName')
  String? get normalisedName;

  /// Optional expected ISO 4217 document currency
  @BuiltValueField(wireName: r'expectedDocumentCurrencyCode')
  String? get expectedDocumentCurrencyCode;

  /// Disabled profiles remain stored but do not match
  @BuiltValueField(wireName: r'enabled')
  bool? get enabled;

  /// When true, matching receipts created via email, quick scan, or the create-receipt API receive these defaults automatically. Extracted document currency is never overwritten.
  @BuiltValueField(wireName: r'autoApply')
  bool? get autoApply;

  @BuiltValueField(wireName: r'categories')
  BuiltList<Category>? get categories;

  @BuiltValueField(wireName: r'tags')
  BuiltList<Tag>? get tags;

  @BuiltValueField(wireName: r'aliases')
  BuiltList<SupplierProfileAlias>? get aliases;

  SupplierProfile._();

  factory SupplierProfile([void updates(SupplierProfileBuilder b)]) = _$SupplierProfile;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(SupplierProfileBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<SupplierProfile> get serializer => _$SupplierProfileSerializer();
}

class _$SupplierProfileSerializer implements PrimitiveSerializer<SupplierProfile> {
  @override
  final Iterable<Type> types = const [SupplierProfile, _$SupplierProfile];

  @override
  final String wireName = r'SupplierProfile';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    SupplierProfile object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    if (object.id != null) {
      yield r'id';
      yield serializers.serialize(
        object.id,
        specifiedType: const FullType(int),
      );
    }
    if (object.createdAt != null) {
      yield r'createdAt';
      yield serializers.serialize(
        object.createdAt,
        specifiedType: const FullType(String),
      );
    }
    if (object.createdBy != null) {
      yield r'createdBy';
      yield serializers.serialize(
        object.createdBy,
        specifiedType: const FullType(int),
      );
    }
    if (object.createdByString != null) {
      yield r'createdByString';
      yield serializers.serialize(
        object.createdByString,
        specifiedType: const FullType(String),
      );
    }
    if (object.updatedAt != null) {
      yield r'updatedAt';
      yield serializers.serialize(
        object.updatedAt,
        specifiedType: const FullType(String),
      );
    }
    if (object.groupId != null) {
      yield r'groupId';
      yield serializers.serialize(
        object.groupId,
        specifiedType: const FullType(int),
      );
    }
    if (object.name != null) {
      yield r'name';
      yield serializers.serialize(
        object.name,
        specifiedType: const FullType(String),
      );
    }
    if (object.normalisedName != null) {
      yield r'normalisedName';
      yield serializers.serialize(
        object.normalisedName,
        specifiedType: const FullType(String),
      );
    }
    if (object.expectedDocumentCurrencyCode != null) {
      yield r'expectedDocumentCurrencyCode';
      yield serializers.serialize(
        object.expectedDocumentCurrencyCode,
        specifiedType: const FullType(String),
      );
    }
    if (object.enabled != null) {
      yield r'enabled';
      yield serializers.serialize(
        object.enabled,
        specifiedType: const FullType(bool),
      );
    }
    if (object.autoApply != null) {
      yield r'autoApply';
      yield serializers.serialize(
        object.autoApply,
        specifiedType: const FullType(bool),
      );
    }
    if (object.categories != null) {
      yield r'categories';
      yield serializers.serialize(
        object.categories,
        specifiedType: const FullType(BuiltList, [FullType(Category)]),
      );
    }
    if (object.tags != null) {
      yield r'tags';
      yield serializers.serialize(
        object.tags,
        specifiedType: const FullType(BuiltList, [FullType(Tag)]),
      );
    }
    if (object.aliases != null) {
      yield r'aliases';
      yield serializers.serialize(
        object.aliases,
        specifiedType: const FullType(BuiltList, [FullType(SupplierProfileAlias)]),
      );
    }
  }

  @override
  Object serialize(
    Serializers serializers,
    SupplierProfile object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required SupplierProfileBuilder result,
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
        case r'createdAt':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.createdAt = valueDes;
          break;
        case r'createdBy':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(int),
          ) as int;
          result.createdBy = valueDes;
          break;
        case r'createdByString':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.createdByString = valueDes;
          break;
        case r'updatedAt':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.updatedAt = valueDes;
          break;
        case r'groupId':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(int),
          ) as int;
          result.groupId = valueDes;
          break;
        case r'name':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.name = valueDes;
          break;
        case r'normalisedName':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.normalisedName = valueDes;
          break;
        case r'expectedDocumentCurrencyCode':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.expectedDocumentCurrencyCode = valueDes;
          break;
        case r'enabled':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(bool),
          ) as bool;
          result.enabled = valueDes;
          break;
        case r'autoApply':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(bool),
          ) as bool;
          result.autoApply = valueDes;
          break;
        case r'categories':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(BuiltList, [FullType(Category)]),
          ) as BuiltList<Category>;
          result.categories.replace(valueDes);
          break;
        case r'tags':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(BuiltList, [FullType(Tag)]),
          ) as BuiltList<Tag>;
          result.tags.replace(valueDes);
          break;
        case r'aliases':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(BuiltList, [FullType(SupplierProfileAlias)]),
          ) as BuiltList<SupplierProfileAlias>;
          result.aliases.replace(valueDes);
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  SupplierProfile deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = SupplierProfileBuilder();
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

