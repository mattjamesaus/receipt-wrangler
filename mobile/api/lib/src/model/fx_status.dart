//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//

// ignore_for_file: unused_element
import 'package:built_collection/built_collection.dart';
import 'package:built_value/built_value.dart';
import 'package:built_value/serializer.dart';

part 'fx_status.g.dart';

class FxStatus extends EnumClass {

  /// Provenance status for the receipt's effective base-currency amount
  @BuiltValueEnumConst(wireName: r'DOMESTIC')
  static const FxStatus DOMESTIC = _$DOMESTIC;
  /// Provenance status for the receipt's effective base-currency amount
  @BuiltValueEnumConst(wireName: r'ESTIMATED')
  static const FxStatus ESTIMATED = _$ESTIMATED;
  /// Provenance status for the receipt's effective base-currency amount
  @BuiltValueEnumConst(wireName: r'CONFIRMED')
  static const FxStatus CONFIRMED = _$CONFIRMED;
  /// Provenance status for the receipt's effective base-currency amount
  @BuiltValueEnumConst(wireName: r'NEEDS_REVIEW')
  static const FxStatus NEEDS_REVIEW = _$NEEDS_REVIEW;

  static Serializer<FxStatus> get serializer => _$fxStatusSerializer;

  const FxStatus._(String name): super(name);

  static BuiltSet<FxStatus> get values => _$values;
  static FxStatus valueOf(String name) => _$valueOf(name);
}

/// Optionally, enum_class can generate a mixin to go with your enum for use
/// with Angular. It exposes your enum constants as getters. So, if you mix it
/// in to your Dart component class, the values become available to the
/// corresponding Angular template.
///
/// Trigger mixin generation by writing a line like this one next to your enum.
abstract class FxStatusMixin = Object with _$FxStatusMixin;
