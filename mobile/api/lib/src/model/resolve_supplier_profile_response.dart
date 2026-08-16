//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:openapi/src/model/supplier_profile.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'resolve_supplier_profile_response.g.dart';

/// ResolveSupplierProfileResponse
///
/// Properties:
/// * [profile] - Matching enabled profile, or null when none / ambiguous
@BuiltValue()
abstract class ResolveSupplierProfileResponse implements Built<ResolveSupplierProfileResponse, ResolveSupplierProfileResponseBuilder> {
  /// Matching enabled profile, or null when none / ambiguous
  @BuiltValueField(wireName: r'profile')
  SupplierProfile? get profile;

  ResolveSupplierProfileResponse._();

  factory ResolveSupplierProfileResponse([void updates(ResolveSupplierProfileResponseBuilder b)]) = _$ResolveSupplierProfileResponse;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(ResolveSupplierProfileResponseBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<ResolveSupplierProfileResponse> get serializer => _$ResolveSupplierProfileResponseSerializer();
}

class _$ResolveSupplierProfileResponseSerializer implements PrimitiveSerializer<ResolveSupplierProfileResponse> {
  @override
  final Iterable<Type> types = const [ResolveSupplierProfileResponse, _$ResolveSupplierProfileResponse];

  @override
  final String wireName = r'ResolveSupplierProfileResponse';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    ResolveSupplierProfileResponse object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    if (object.profile != null) {
      yield r'profile';
      yield serializers.serialize(
        object.profile,
        specifiedType: const FullType(SupplierProfile),
      );
    }
  }

  @override
  Object serialize(
    Serializers serializers,
    ResolveSupplierProfileResponse object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required ResolveSupplierProfileResponseBuilder result,
    required List<Object?> unhandled,
  }) {
    for (var i = 0; i < serializedList.length; i += 2) {
      final key = serializedList[i] as String;
      final value = serializedList[i + 1];
      switch (key) {
        case r'profile':
          final valueDes = serializers.deserialize(
            value,
            specifiedType: const FullType(SupplierProfile),
          ) as SupplierProfile;
          result.profile.replace(valueDes);
          break;
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  ResolveSupplierProfileResponse deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = ResolveSupplierProfileResponseBuilder();
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

