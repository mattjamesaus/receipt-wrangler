import 'package:test/test.dart';
import 'package:openapi/openapi.dart';

// tests for UpsertSupplierProfileCommand
void main() {
  final instance = UpsertSupplierProfileCommandBuilder();
  // TODO add properties to the builder and call build()

  group(UpsertSupplierProfileCommand, () {
    // Canonical supplier display name
    // String name
    test('to test the property `name`', () async {
      // TODO
    });

    // Additional names that should match this profile
    // BuiltList<String> aliases
    test('to test the property `aliases`', () async {
      // TODO
    });

    // Default category ids from the catalogue
    // BuiltList<int> categoryIds
    test('to test the property `categoryIds`', () async {
      // TODO
    });

    // Default tag ids from the catalogue
    // BuiltList<int> tagIds
    test('to test the property `tagIds`', () async {
      // TODO
    });

    // Optional expected ISO 4217 document currency
    // String expectedDocumentCurrencyCode
    test('to test the property `expectedDocumentCurrencyCode`', () async {
      // TODO
    });

    // Whether the profile participates in matching. Defaults to true on create.
    // bool enabled
    test('to test the property `enabled`', () async {
      // TODO
    });

  });
}
