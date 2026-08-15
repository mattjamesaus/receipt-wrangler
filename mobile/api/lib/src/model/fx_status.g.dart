// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'fx_status.dart';

// **************************************************************************
// BuiltValueGenerator
// **************************************************************************

const FxStatus _$DOMESTIC = const FxStatus._('DOMESTIC');
const FxStatus _$ESTIMATED = const FxStatus._('ESTIMATED');
const FxStatus _$CONFIRMED = const FxStatus._('CONFIRMED');
const FxStatus _$NEEDS_REVIEW = const FxStatus._('NEEDS_REVIEW');

FxStatus _$valueOf(String name) {
  switch (name) {
    case 'DOMESTIC':
      return _$DOMESTIC;
    case 'ESTIMATED':
      return _$ESTIMATED;
    case 'CONFIRMED':
      return _$CONFIRMED;
    case 'NEEDS_REVIEW':
      return _$NEEDS_REVIEW;
    default:
      throw ArgumentError(name);
  }
}

final BuiltSet<FxStatus> _$values = BuiltSet<FxStatus>(const <FxStatus>[
  _$DOMESTIC,
  _$ESTIMATED,
  _$CONFIRMED,
  _$NEEDS_REVIEW,
]);

class _$FxStatusMeta {
  const _$FxStatusMeta();
  FxStatus get DOMESTIC => _$DOMESTIC;
  FxStatus get ESTIMATED => _$ESTIMATED;
  FxStatus get CONFIRMED => _$CONFIRMED;
  FxStatus get NEEDS_REVIEW => _$NEEDS_REVIEW;
  FxStatus valueOf(String name) => _$valueOf(name);
  BuiltSet<FxStatus> get values => _$values;
}

abstract class _$FxStatusMixin {
  // ignore: non_constant_identifier_names
  _$FxStatusMeta get FxStatus => const _$FxStatusMeta();
}

Serializer<FxStatus> _$fxStatusSerializer = _$FxStatusSerializer();

class _$FxStatusSerializer implements PrimitiveSerializer<FxStatus> {
  static const Map<String, Object> _toWire = const <String, Object>{
    'DOMESTIC': 'DOMESTIC',
    'ESTIMATED': 'ESTIMATED',
    'CONFIRMED': 'CONFIRMED',
    'NEEDS_REVIEW': 'NEEDS_REVIEW',
  };
  static const Map<Object, String> _fromWire = const <Object, String>{
    'DOMESTIC': 'DOMESTIC',
    'ESTIMATED': 'ESTIMATED',
    'CONFIRMED': 'CONFIRMED',
    'NEEDS_REVIEW': 'NEEDS_REVIEW',
  };

  @override
  final Iterable<Type> types = const <Type>[FxStatus];
  @override
  final String wireName = 'FxStatus';

  @override
  Object serialize(Serializers serializers, FxStatus object,
          {FullType specifiedType = FullType.unspecified}) =>
      _toWire[object.name] ?? object.name;

  @override
  FxStatus deserialize(Serializers serializers, Object serialized,
          {FullType specifiedType = FullType.unspecified}) =>
      FxStatus.valueOf(
          _fromWire[serialized] ?? (serialized is String ? serialized : ''));
}

// ignore_for_file: deprecated_member_use_from_same_package,type=lint
