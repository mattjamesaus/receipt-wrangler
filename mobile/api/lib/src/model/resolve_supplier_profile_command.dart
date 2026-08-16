//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'resolve_supplier_profile_command.g.dart';

/// ResolveSupplierProfileCommand
///
/// Properties:
/// * [name] - Receipt name to match against profiles and aliases
@BuiltValue()
abstract class ResolveSupplierProfileCommand implements Built<ResolveSupplierProfileCommand, ResolveSupplierProfileCommandBuilder> {
  /// Receipt name to match against profiles and aliases
  @BuiltValueField(wireName: r'name')
  String get name;

  ResolveSupplierProfileCommand._();

  factory ResolveSupplierProfileCommand([void updates(ResolveSupplierProfileCommandBuilder b)]) = _$ResolveSupplierProfileCommand;

  @BuiltValueHook(initializeBuilder: true)
  static void _defaults(ResolveSupplierProfileCommandBuilder b) => b;

  @BuiltValueSerializer(custom: true)
  static Serializer<ResolveSupplierProfileCommand> get serializer => _$ResolveSupplierProfileCommandSerializer();
}

class _$ResolveSupplierProfileCommandSerializer implements PrimitiveSerializer<ResolveSupplierProfileCommand> {
  @override
  final Iterable<Type> types = const [ResolveSupplierProfileCommand, _$ResolveSupplierProfileCommand];

  @override
  final String wireName = r'ResolveSupplierProfileCommand';

  Iterable<Object?> _serializeProperties(
    Serializers serializers,
    ResolveSupplierProfileCommand object, {
    FullType specifiedType = FullType.unspecified,
  }) sync* {
    yield r'name';
    yield serializers.serialize(
      object.name,
      specifiedType: const FullType(String),
    );
  }

  @override
  Object serialize(
    Serializers serializers,
    ResolveSupplierProfileCommand object, {
    FullType specifiedType = FullType.unspecified,
  }) {
    return _serializeProperties(serializers, object, specifiedType: specifiedType).toList();
  }

  void _deserializeProperties(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
    required List<Object?> serializedList,
    required ResolveSupplierProfileCommandBuilder result,
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
        default:
          unhandled.add(key);
          unhandled.add(value);
          break;
      }
    }
  }

  @override
  ResolveSupplierProfileCommand deserialize(
    Serializers serializers,
    Object serialized, {
    FullType specifiedType = FullType.unspecified,
  }) {
    final result = ResolveSupplierProfileCommandBuilder();
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

