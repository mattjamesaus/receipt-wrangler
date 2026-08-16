// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'resolve_supplier_profile_command.dart';

// **************************************************************************
// BuiltValueGenerator
// **************************************************************************

class _$ResolveSupplierProfileCommand extends ResolveSupplierProfileCommand {
  @override
  final String name;

  factory _$ResolveSupplierProfileCommand(
          [void Function(ResolveSupplierProfileCommandBuilder)? updates]) =>
      (ResolveSupplierProfileCommandBuilder()..update(updates))._build();

  _$ResolveSupplierProfileCommand._({required this.name}) : super._();
  @override
  ResolveSupplierProfileCommand rebuild(
          void Function(ResolveSupplierProfileCommandBuilder) updates) =>
      (toBuilder()..update(updates)).build();

  @override
  ResolveSupplierProfileCommandBuilder toBuilder() =>
      ResolveSupplierProfileCommandBuilder()..replace(this);

  @override
  bool operator ==(Object other) {
    if (identical(other, this)) return true;
    return other is ResolveSupplierProfileCommand && name == other.name;
  }

  @override
  int get hashCode {
    var _$hash = 0;
    _$hash = $jc(_$hash, name.hashCode);
    _$hash = $jf(_$hash);
    return _$hash;
  }

  @override
  String toString() {
    return (newBuiltValueToStringHelper(r'ResolveSupplierProfileCommand')
          ..add('name', name))
        .toString();
  }
}

class ResolveSupplierProfileCommandBuilder
    implements
        Builder<ResolveSupplierProfileCommand,
            ResolveSupplierProfileCommandBuilder> {
  _$ResolveSupplierProfileCommand? _$v;

  String? _name;
  String? get name => _$this._name;
  set name(String? name) => _$this._name = name;

  ResolveSupplierProfileCommandBuilder() {
    ResolveSupplierProfileCommand._defaults(this);
  }

  ResolveSupplierProfileCommandBuilder get _$this {
    final $v = _$v;
    if ($v != null) {
      _name = $v.name;
      _$v = null;
    }
    return this;
  }

  @override
  void replace(ResolveSupplierProfileCommand other) {
    _$v = other as _$ResolveSupplierProfileCommand;
  }

  @override
  void update(void Function(ResolveSupplierProfileCommandBuilder)? updates) {
    if (updates != null) updates(this);
  }

  @override
  ResolveSupplierProfileCommand build() => _build();

  _$ResolveSupplierProfileCommand _build() {
    final _$result = _$v ??
        _$ResolveSupplierProfileCommand._(
          name: BuiltValueNullFieldError.checkNotNull(
              name, r'ResolveSupplierProfileCommand', 'name'),
        );
    replace(_$result);
    return _$result;
  }
}

// ignore_for_file: deprecated_member_use_from_same_package,type=lint
