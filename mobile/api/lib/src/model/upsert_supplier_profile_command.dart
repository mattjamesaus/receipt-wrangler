//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_collection/built_collection.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'upsert_supplier_profile_command.g.dart';

/// UpsertSupplierProfileCommand
///
/// Properties:
/// * [name] - Canonical supplier display name
/// * [aliases] - Additional names that should match this profile
/// * [categoryIds] - Default category ids from the catalogue
/// * [tagIds] - Default tag ids from the catalogue
/// * [expectedDocumentCurrencyCode] - Optional expected ISO 4217 document currency
/// * [enabled] - Whether the profile participates in matching. Defaults to true on create.
@BuiltValue()
abstract class UpsertSupplierProfileCommand implements Built<UpsertSupplierProfileCommand, UpsertSupplierProfileCommandBuilder> {
  /// Canonical supplier display name
  @BuiltValueField(wireName: r'name')
  String get name;

  /// Additional names that should match this profile
  @BuiltValueField(wireName: r'aliases')
  BuiltList<String>? get aliases;

  /// Default category ids from the catalogue
  @BuiltValueField(wireName: r'categoryIds')
  BuiltList<int>? get categoryIds;

  /// Default tag ids from the catalogue
  @BuiltValueField(wireName: r'tagIds')
  BuiltList<int>? get tagIds;

  /// Optional expected ISO 4217 document currency
  @BuiltValueField(wireName: r'expectedDocumentCurrencyCode')
  String? get expectedDocumentCurrencyCode;

  /// Whether the profile participates in matching. Defaults to true on create.
  @BuiltValueField(wireName: r'enabled')
  bool? get enabled;

  UpsertSupplierProfileCommand._();

  factory UpsertSupplierProfileCommand([void updates(UpsertSupplierProfileCommandBuilder b)]) = _$UpsertSupplierProfileCommand;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(UpsertSupplierProfileCommandBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<UpsertSupplierProfileCommand> get serializer => _$UpsertSupplierProfileCommandSerializer();
}

class _$UpsertSupplierProfileCommandSerializer implements PrimitiveSerializer<UpsertSupplierProfileCommand> {
  @override
  final Iterable<Type> types = const [UpsertSupplierProfileCommand, _$UpsertSupplierProfileCommand];

  @override
  final String wireName = r'UpsertSupplierProfileCommand';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    UpsertSupplierProfileCommand object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    yield r'name';
    yield serializers.serialize(
      object.name,
      specifiedType: const FullType(String),
    );
    if (object.aliases != null) {
      yield r'aliases';
      yield serializers.serialize(
        object.aliases,
        specifiedType: const FullType(BuiltList, [FullType(String)]),
      );
    }
    if (object.categoryIds != null) {
      yield r'categoryIds';
      yield serializers.serialize(
        object.categoryIds,
        specifiedType: const FullType(BuiltList, [FullType(int)]),
      );
    }
    if (object.tagIds != null) {
      yield r'tagIds';
      yield serializers.serialize(
        object.tagIds,
        specifiedType: const FullType(BuiltList, [FullType(int)]),
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
  }

  @override
  Object serialize(
    Serializers serializers,
    UpsertSupplierProfileCommand object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required UpsertSupplierProfileCommandBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'name':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(String),
          ) as String;
          result.name = valueDes;
          break;
        case r'aliases':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(BuiltList, [FullType(String)]),
          ) as BuiltList<String>;
          result.aliases.replace(valueDes);
          break;
        case r'categoryIds':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(BuiltList, [FullType(int)]),
          ) as BuiltList<int>;
          result.categoryIds.replace(valueDes);
          break;
        case r'tagIds':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(BuiltList, [FullType(int)]),
          ) as BuiltList<int>;
          result.tagIds.replace(valueDes);
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
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  UpsertSupplierProfileCommand deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = UpsertSupplierProfileCommandBuilder();
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

