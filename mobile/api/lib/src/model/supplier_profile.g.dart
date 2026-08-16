// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'supplier_profile.dart';

// **************************************************************************
// BuiltValueGenerator
// **************************************************************************

class _$SupplierProfile extends SupplierProfile {
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
  final int? groupId;
  @override
  final String? name;
  @override
  final String? normalisedName;
  @override
  final String? expectedDocumentCurrencyCode;
  @override
  final bool? enabled;
  @override
  final BuiltList<Category>? categories;
  @override
  final BuiltList<Tag>? tags;
  @override
  final BuiltList<SupplierProfileAlias>? aliases;

  factory _$SupplierProfile([void Function(SupplierProfileBuilder)? updates]) =>
      (SupplierProfileBuilder()..update(updates))._build();

  _$SupplierProfile._(
      {this.id,
      this.createdAt,
      this.createdBy,
      this.createdByString,
      this.updatedAt,
      this.groupId,
      this.name,
      this.normalisedName,
      this.expectedDocumentCurrencyCode,
      this.enabled,
      this.categories,
      this.tags,
      this.aliases})
      : super._();
  @override
  SupplierProfile rebuild(void Function(SupplierProfileBuilder) updates) =>
      (toBuilder()..update(updates)).build();

  @override
  SupplierProfileBuilder toBuilder() => SupplierProfileBuilder()..replace(this);

  @override
  bool operator ==(Object other) {
    if (identical(other, this)) return true;
    return other is SupplierProfile &&
        id == other.id &&
        createdAt == other.createdAt &&
        createdBy == other.createdBy &&
        createdByString == other.createdByString &&
        updatedAt == other.updatedAt &&
        groupId == other.groupId &&
        name == other.name &&
        normalisedName == other.normalisedName &&
        expectedDocumentCurrencyCode == other.expectedDocumentCurrencyCode &&
        enabled == other.enabled &&
        categories == other.categories &&
        tags == other.tags &&
        aliases == other.aliases;
  }

  @override
  int get hashCode {
    var _$hash = 0;
    _$hash = $jc(_$hash, id.hashCode);
    _$hash = $jc(_$hash, createdAt.hashCode);
    _$hash = $jc(_$hash, createdBy.hashCode);
    _$hash = $jc(_$hash, createdByString.hashCode);
    _$hash = $jc(_$hash, updatedAt.hashCode);
    _$hash = $jc(_$hash, groupId.hashCode);
    _$hash = $jc(_$hash, name.hashCode);
    _$hash = $jc(_$hash, normalisedName.hashCode);
    _$hash = $jc(_$hash, expectedDocumentCurrencyCode.hashCode);
    _$hash = $jc(_$hash, enabled.hashCode);
    _$hash = $jc(_$hash, categories.hashCode);
    _$hash = $jc(_$hash, tags.hashCode);
    _$hash = $jc(_$hash, aliases.hashCode);
    _$hash = $jf(_$hash);
    return _$hash;
  }

  @override
  String toString() {
    return (newBuiltValueToStringHelper(r'SupplierProfile')
          ..add('id', id)
          ..add('createdAt', createdAt)
          ..add('createdBy', createdBy)
          ..add('createdByString', createdByString)
          ..add('updatedAt', updatedAt)
          ..add('groupId', groupId)
          ..add('name', name)
          ..add('normalisedName', normalisedName)
          ..add('expectedDocumentCurrencyCode', expectedDocumentCurrencyCode)
          ..add('enabled', enabled)
          ..add('categories', categories)
          ..add('tags', tags)
          ..add('aliases', aliases))
        .toString();
  }
}

class SupplierProfileBuilder
    implements Builder<SupplierProfile, SupplierProfileBuilder> {
  _$SupplierProfile? _$v;

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

  String? _expectedDocumentCurrencyCode;
  String? get expectedDocumentCurrencyCode =>
      _$this._expectedDocumentCurrencyCode;
  set expectedDocumentCurrencyCode(String? expectedDocumentCurrencyCode) =>
      _$this._expectedDocumentCurrencyCode = expectedDocumentCurrencyCode;

  bool? _enabled;
  bool? get enabled => _$this._enabled;
  set enabled(bool? enabled) => _$this._enabled = enabled;

  ListBuilder<Category>? _categories;
  ListBuilder<Category> get categories =>
      _$this._categories ??= ListBuilder<Category>();
  set categories(ListBuilder<Category>? categories) =>
      _$this._categories = categories;

  ListBuilder<Tag>? _tags;
  ListBuilder<Tag> get tags => _$this._tags ??= ListBuilder<Tag>();
  set tags(ListBuilder<Tag>? tags) => _$this._tags = tags;

  ListBuilder<SupplierProfileAlias>? _aliases;
  ListBuilder<SupplierProfileAlias> get aliases =>
      _$this._aliases ??= ListBuilder<SupplierProfileAlias>();
  set aliases(ListBuilder<SupplierProfileAlias>? aliases) =>
      _$this._aliases = aliases;

  SupplierProfileBuilder() {
    SupplierProfile._defaults(this);
  }

  SupplierProfileBuilder get _$this {
    final $v = _$v;
    if ($v != null) {
      _id = $v.id;
      _createdAt = $v.createdAt;
      _createdBy = $v.createdBy;
      _createdByString = $v.createdByString;
      _updatedAt = $v.updatedAt;
      _groupId = $v.groupId;
      _name = $v.name;
      _normalisedName = $v.normalisedName;
      _expectedDocumentCurrencyCode = $v.expectedDocumentCurrencyCode;
      _enabled = $v.enabled;
      _categories = $v.categories?.toBuilder();
      _tags = $v.tags?.toBuilder();
      _aliases = $v.aliases?.toBuilder();
      _$v = null;
    }
    return this;
  }

  @override
  void replace(SupplierProfile other) {
    _$v = other as _$SupplierProfile;
  }

  @override
  void update(void Function(SupplierProfileBuilder)? updates) {
    if (updates != null) updates(this);
  }

  @override
  SupplierProfile build() => _build();

  _$SupplierProfile _build() {
    _$SupplierProfile _$result;
    try {
      _$result = _$v ??
          _$SupplierProfile._(
            id: id,
            createdAt: createdAt,
            createdBy: createdBy,
            createdByString: createdByString,
            updatedAt: updatedAt,
            groupId: groupId,
            name: name,
            normalisedName: normalisedName,
            expectedDocumentCurrencyCode: expectedDocumentCurrencyCode,
            enabled: enabled,
            categories: _categories?.build(),
            tags: _tags?.build(),
            aliases: _aliases?.build(),
          );
    } catch (_) {
      late String _$failedField;
      try {
        _$failedField = 'categories';
        _categories?.build();
        _$failedField = 'tags';
        _tags?.build();
        _$failedField = 'aliases';
        _aliases?.build();
      } catch (e) {
        throw BuiltValueNestedFieldError(
            r'SupplierProfile', _$failedField, e.toString());
      }
      rethrow;
    }
    replace(_$result);
    return _$result;
  }
}

// ignore_for_file: deprecated_member_use_from_same_package,type=lint
