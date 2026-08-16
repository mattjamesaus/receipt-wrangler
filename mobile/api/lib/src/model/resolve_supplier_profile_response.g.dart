// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'resolve_supplier_profile_response.dart';

// **************************************************************************
// BuiltValueGenerator
// **************************************************************************

class _$ResolveSupplierProfileResponse extends ResolveSupplierProfileResponse {
  @override
  final SupplierProfile? profile;

  factory _$ResolveSupplierProfileResponse(
          [void Function(ResolveSupplierProfileResponseBuilder)? updates]) =>
      (ResolveSupplierProfileResponseBuilder()..update(updates))._build();

  _$ResolveSupplierProfileResponse._({this.profile}) : super._();
  @override
  ResolveSupplierProfileResponse rebuild(
          void Function(ResolveSupplierProfileResponseBuilder) updates) =>
      (toBuilder()..update(updates)).build();

  @override
  ResolveSupplierProfileResponseBuilder toBuilder() =>
      ResolveSupplierProfileResponseBuilder()..replace(this);

  @override
  bool operator ==(Object other) {
    if (identical(other, this)) return true;
    return other is ResolveSupplierProfileResponse && profile == other.profile;
  }

  @override
  int get hashCode {
    var _$hash = 0;
    _$hash = $jc(_$hash, profile.hashCode);
    _$hash = $jf(_$hash);
    return _$hash;
  }

  @override
  String toString() {
    return (newBuiltValueToStringHelper(r'ResolveSupplierProfileResponse')
          ..add('profile', profile))
        .toString();
  }
}

class ResolveSupplierProfileResponseBuilder
    implements
        Builder<ResolveSupplierProfileResponse,
            ResolveSupplierProfileResponseBuilder> {
  _$ResolveSupplierProfileResponse? _$v;

  SupplierProfileBuilder? _profile;
  SupplierProfileBuilder get profile =>
      _$this._profile ??= SupplierProfileBuilder();
  set profile(SupplierProfileBuilder? profile) => _$this._profile = profile;

  ResolveSupplierProfileResponseBuilder() {
    ResolveSupplierProfileResponse._defaults(this);
  }

  ResolveSupplierProfileResponseBuilder get _$this {
    final $v = _$v;
    if ($v != null) {
      _profile = $v.profile?.toBuilder();
      _$v = null;
    }
    return this;
  }

  @override
  void replace(ResolveSupplierProfileResponse other) {
    _$v = other as _$ResolveSupplierProfileResponse;
  }

  @override
  void update(void Function(ResolveSupplierProfileResponseBuilder)? updates) {
    if (updates != null) updates(this);
  }

  @override
  ResolveSupplierProfileResponse build() => _build();

  _$ResolveSupplierProfileResponse _build() {
    _$ResolveSupplierProfileResponse _$result;
    try {
      _$result = _$v ??
          _$ResolveSupplierProfileResponse._(
            profile: _profile?.build(),
          );
    } catch (_) {
      late String _$failedField;
      try {
        _$failedField = 'profile';
        _profile?.build();
      } catch (e) {
        throw BuiltValueNestedFieldError(
            r'ResolveSupplierProfileResponse', _$failedField, e.toString());
      }
      rethrow;
    }
    replace(_$result);
    return _$result;
  }
}

// ignore_for_file: deprecated_member_use_from_same_package,type=lint
