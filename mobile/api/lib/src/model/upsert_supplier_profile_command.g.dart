// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'upsert_supplier_profile_command.dart';

// **************************************************************************
// BuiltValueGenerator
// **************************************************************************

class _$UpsertSupplierProfileCommand extends UpsertSupplierProfileCommand {
  @override
  final String name;
  @override
  final BuiltList<String>? aliases;
  @override
  final BuiltList<int>? categoryIds;
  @override
  final BuiltList<int>? tagIds;
  @override
  final String? expectedDocumentCurrencyCode;
  @override
  final bool? enabled;
  @override
  final bool? autoApply;

  factory _$UpsertSupplierProfileCommand(
          [void Function(UpsertSupplierProfileCommandBuilder)? updates]) =>
      (UpsertSupplierProfileCommandBuilder()..update(updates))._build();

  _$UpsertSupplierProfileCommand._(
      {required this.name,
      this.aliases,
      this.categoryIds,
      this.tagIds,
      this.expectedDocumentCurrencyCode,
      this.enabled,
      this.autoApply})
      : super._();
  @override
  UpsertSupplierProfileCommand rebuild(
          void Function(UpsertSupplierProfileCommandBuilder) updates) =>
      (toBuilder()..update(updates)).build();

  @override
  UpsertSupplierProfileCommandBuilder toBuilder() =>
      UpsertSupplierProfileCommandBuilder()..replace(this);

  @override
  bool operator ==(Object other) {
    if (identical(other, this)) return true;
    return other is UpsertSupplierProfileCommand &&
        name == other.name &&
        aliases == other.aliases &&
        categoryIds == other.categoryIds &&
        tagIds == other.tagIds &&
        expectedDocumentCurrencyCode == other.expectedDocumentCurrencyCode &&
        enabled == other.enabled &&
        autoApply == other.autoApply;
  }

  @override
  int get hashCode {
    var _$hash = 0;
    _$hash = $jc(_$hash, name.hashCode);
    _$hash = $jc(_$hash, aliases.hashCode);
    _$hash = $jc(_$hash, categoryIds.hashCode);
    _$hash = $jc(_$hash, tagIds.hashCode);
    _$hash = $jc(_$hash, expectedDocumentCurrencyCode.hashCode);
    _$hash = $jc(_$hash, enabled.hashCode);
    _$hash = $jc(_$hash, autoApply.hashCode);
    _$hash = $jf(_$hash);
    return _$hash;
  }

  @override
  String toString() {
    return (newBuiltValueToStringHelper(r'UpsertSupplierProfileCommand')
          ..add('name', name)
          ..add('aliases', aliases)
          ..add('categoryIds', categoryIds)
          ..add('tagIds', tagIds)
          ..add('expectedDocumentCurrencyCode', expectedDocumentCurrencyCode)
          ..add('enabled', enabled)
          ..add('autoApply', autoApply))
        .toString();
  }
}

class UpsertSupplierProfileCommandBuilder
    implements
        Builder<UpsertSupplierProfileCommand,
            UpsertSupplierProfileCommandBuilder> {
  _$UpsertSupplierProfileCommand? _$v;

  String? _name;
  String? get name => _$this._name;
  set name(String? name) => _$this._name = name;

  ListBuilder<String>? _aliases;
  ListBuilder<String> get aliases => _$this._aliases ??= ListBuilder<String>();
  set aliases(ListBuilder<String>? aliases) => _$this._aliases = aliases;

  ListBuilder<int>? _categoryIds;
  ListBuilder<int> get categoryIds =>
      _$this._categoryIds ??= ListBuilder<int>();
  set categoryIds(ListBuilder<int>? categoryIds) =>
      _$this._categoryIds = categoryIds;

  ListBuilder<int>? _tagIds;
  ListBuilder<int> get tagIds => _$this._tagIds ??= ListBuilder<int>();
  set tagIds(ListBuilder<int>? tagIds) => _$this._tagIds = tagIds;

  String? _expectedDocumentCurrencyCode;
  String? get expectedDocumentCurrencyCode =>
      _$this._expectedDocumentCurrencyCode;
  set expectedDocumentCurrencyCode(String? expectedDocumentCurrencyCode) =>
      _$this._expectedDocumentCurrencyCode = expectedDocumentCurrencyCode;

  bool? _enabled;
  bool? get enabled => _$this._enabled;
  set enabled(bool? enabled) => _$this._enabled = enabled;

  bool? _autoApply;
  bool? get autoApply => _$this._autoApply;
  set autoApply(bool? autoApply) => _$this._autoApply = autoApply;

  UpsertSupplierProfileCommandBuilder() {
    UpsertSupplierProfileCommand._defaults(this);
  }

  UpsertSupplierProfileCommandBuilder get _$this {
    final $v = _$v;
    if ($v != null) {
      _name = $v.name;
      _aliases = $v.aliases?.toBuilder();
      _categoryIds = $v.categoryIds?.toBuilder();
      _tagIds = $v.tagIds?.toBuilder();
      _expectedDocumentCurrencyCode = $v.expectedDocumentCurrencyCode;
      _enabled = $v.enabled;
      _autoApply = $v.autoApply;
      _$v = null;
    }
    return this;
  }

  @override
  void replace(UpsertSupplierProfileCommand other) {
    _$v = other as _$UpsertSupplierProfileCommand;
  }

  @override
  void update(void Function(UpsertSupplierProfileCommandBuilder)? updates) {
    if (updates != null) updates(this);
  }

  @override
  UpsertSupplierProfileCommand build() => _build();

  _$UpsertSupplierProfileCommand _build() {
    _$UpsertSupplierProfileCommand _$result;
    try {
      _$result = _$v ??
          _$UpsertSupplierProfileCommand._(
            name: BuiltValueNullFieldError.checkNotNull(
                name, r'UpsertSupplierProfileCommand', 'name'),
            aliases: _aliases?.build(),
            categoryIds: _categoryIds?.build(),
            tagIds: _tagIds?.build(),
            expectedDocumentCurrencyCode: expectedDocumentCurrencyCode,
            enabled: enabled,
            autoApply: autoApply,
          );
    } catch (_) {
      late String _$failedField;
      try {
        _$failedField = 'aliases';
        _aliases?.build();
        _$failedField = 'categoryIds';
        _categoryIds?.build();
        _$failedField = 'tagIds';
        _tagIds?.build();
      } catch (e) {
        throw BuiltValueNestedFieldError(
            r'UpsertSupplierProfileCommand', _$failedField, e.toString());
      }
      rethrow;
    }
    replace(_$result);
    return _$result;
  }
}

// ignore_for_file: deprecated_member_use_from_same_package,type=lint
