// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'supplier_profile_alias.dart';

// **************************************************************************
// BuiltValueGenerator
// **************************************************************************

class _$SupplierProfileAlias extends SupplierProfileAlias {
  @override
  final int? id;
  @override
  final String? createdAt;
  @override
  final int? createdBy;
  @override
  final String? createdByString;
  @override
  final String? updatedAt;
  @override
  final int? supplierProfileId;
  @override
  final int? groupId;
  @override
  final String? name;
  @override
  final String? normalisedName;

  factory _$SupplierProfileAlias(
          [void Function(SupplierProfileAliasBuilder)? updates]) =>
      (SupplierProfileAliasBuilder()..update(updates))._build();

  _$SupplierProfileAlias._(
      {this.id,
      this.createdAt,
      this.createdBy,
      this.createdByString,
      this.updatedAt,
      this.supplierProfileId,
      this.groupId,
      this.name,
      this.normalisedName})
      : super._();
  @override
  SupplierProfileAlias rebuild(
          void Function(SupplierProfileAliasBuilder) updates) =>
      (toBuilder()..update(updates)).build();

  @override
  SupplierProfileAliasBuilder toBuilder() =>
      SupplierProfileAliasBuilder()..replace(this);

  @override
  bool operator ==(Object other) {
    if (identical(other, this)) return true;
    return other is SupplierProfileAlias &&
        id == other.id &&
        createdAt == other.createdAt &&
        createdBy == other.createdBy &&
        createdByString == other.createdByString &&
        updatedAt == other.updatedAt &&
        supplierProfileId == other.supplierProfileId &&
        groupId == other.groupId &&
        name == other.name &&
        normalisedName == other.normalisedName;
  }

  @override
  int get hashCode {
    var _$hash = 0;
    _$hash = $jc(_$hash, id.hashCode);
    _$hash = $jc(_$hash, createdAt.hashCode);
    _$hash = $jc(_$hash, createdBy.hashCode);
    _$hash = $jc(_$hash, createdByString.hashCode);
    _$hash = $jc(_$hash, updatedAt.hashCode);
    _$hash = $jc(_$hash, supplierProfileId.hashCode);
    _$hash = $jc(_$hash, groupId.hashCode);
    _$hash = $jc(_$hash, name.hashCode);
    _$hash = $jc(_$hash, normalisedName.hashCode);
    _$hash = $jf(_$hash);
    return _$hash;
  }

  @override
  String toString() {
    return (newBuiltValueToStringHelper(r'SupplierProfileAlias')
          ..add('id', id)
          ..add('createdAt', createdAt)
          ..add('createdBy', createdBy)
          ..add('createdByString', createdByString)
          ..add('updatedAt', updatedAt)
          ..add('supplierProfileId', supplierProfileId)
          ..add('groupId', groupId)
          ..add('name', name)
          ..add('normalisedName', normalisedName))
        .toString();
  }
}

class SupplierProfileAliasBuilder
    implements Builder<SupplierProfileAlias, SupplierProfileAliasBuilder> {
  _$SupplierProfileAlias? _$v;

  int? _id;
  int? get id => _$this._id;
  set id(int? id) => _$this._id = id;

  String? _createdAt;
  String? get createdAt => _$this._createdAt;
  set createdAt(String? createdAt) => _$this._createdAt = createdAt;

  int? _createdBy;
  int? get createdBy => _$this._createdBy;
  set createdBy(int? createdBy) => _$this._createdBy = createdBy;

  String? _createdByString;
  String? get createdByString => _$this._createdByString;
  set createdByString(String? createdByString) =>
      _$this._createdByString = createdByString;

  String? _updatedAt;
  String? get updatedAt => _$this._updatedAt;
  set updatedAt(String? updatedAt) => _$this._updatedAt = updatedAt;

  int? _supplierProfileId;
  int? get supplierProfileId => _$this._supplierProfileId;
  set supplierProfileId(int? supplierProfileId) =>
      _$this._supplierProfileId = supplierProfileId;

  int? _groupId;
  int? get groupId => _$this._groupId;
  set groupId(int? groupId) => _$this._groupId = groupId;

  String? _name;
  String? get name => _$this._name;
  set name(String? name) => _$this._name = name;

  String? _normalisedName;
  String? get normalisedName => _$this._normalisedName;
  set normalisedName(String? normalisedName) =>
      _$this._normalisedName = normalisedName;

  SupplierProfileAliasBuilder() {
    SupplierProfileAlias._defaults(this);
  }

  SupplierProfileAliasBuilder get _$this {
    final $v = _$v;
    if ($v != null) {
      _id = $v.id;
      _createdAt = $v.createdAt;
      _createdBy = $v.createdBy;
      _createdByString = $v.createdByString;
      _updatedAt = $v.updatedAt;
      _supplierProfileId = $v.supplierProfileId;
      _groupId = $v.groupId;
      _name = $v.name;
      _normalisedName = $v.normalisedName;
      _$v = null;
    }
    return this;
  }

  @override
  void replace(SupplierProfileAlias other) {
    _$v = other as _$SupplierProfileAlias;
  }

  @override
  void update(void Function(SupplierProfileAliasBuilder)? updates) {
    if (updates != null) updates(this);
  }

  @override
  SupplierProfileAlias build() => _build();

  _$SupplierProfileAlias _build() {
    final _$result = _$v ??
        _$SupplierProfileAlias._(
          id: id,
          createdAt: createdAt,
          createdBy: createdBy,
          createdByString: createdByString,
          updatedAt: updatedAt,
          supplierProfileId: supplierProfileId,
          groupId: groupId,
          name: name,
          normalisedName: normalisedName,
        );
    replace(_$result);
    return _$result;
  }
}

// ignore_for_file: deprecated_member_use_from_same_package,type=lint
