import 'package:test/test.dart';
import 'package:openapi/openapi.dart';


/// tests for SupplierProfileApi
void main() {
  final instance = Openapi().getSupplierProfileApi();

  group(SupplierProfileApi, () {
    // Create a supplier profile
    //
    // Creates a group-scoped supplier profile with optional default categories, tags, and expected document currency. At least one default is required. Canonical names and aliases must be unique within the group after normalisation. Gated by group.receipts.update.
    //
    //Future<SupplierProfile> createSupplierProfile(int groupId, UpsertSupplierProfileCommand upsertSupplierProfileCommand) async
    test('test createSupplierProfile', () async {
      // TODO
    });

    // Delete a supplier profile
    //
    // Deletes a supplier profile and its aliases. Receipts and source files are unchanged. Gated by group.receipts.update.
    //
    //Future deleteSupplierProfile(int groupId, int profileId) async
    test('test deleteSupplierProfile', () async {
      // TODO
    });

    // Get a supplier profile
    //
    // Returns one supplier profile in the group. Categories and tags are filtered to those the caller is permitted to use. Gated by group.receipts.create.
    //
    //Future<SupplierProfile> getSupplierProfileById(int groupId, int profileId) async
    test('test getSupplierProfileById', () async {
      // TODO
    });

    // List supplier profiles for a group
    //
    // Returns every supplier profile in the group. Categories and tags on each profile are filtered to those the caller is permitted to use. Gated by group.receipts.create.
    //
    //Future<BuiltList<SupplierProfile>> getSupplierProfilesForGroup(int groupId) async
    test('test getSupplierProfilesForGroup', () async {
      // TODO
    });

    // Resolve a supplier profile for a receipt name
    //
    // Returns the single enabled profile whose canonical name or alias matches the supplied receipt name after normalisation. Ambiguous or disabled matches return a null profile. Never applies defaults. Gated by group.receipts.create.
    //
    //Future<ResolveSupplierProfileResponse> resolveSupplierProfile(int groupId, ResolveSupplierProfileCommand resolveSupplierProfileCommand) async
    test('test resolveSupplierProfile', () async {
      // TODO
    });

    // Update a supplier profile
    //
    // Replaces a supplier profile's name, aliases, defaults, expected currency, and enabled state. Enable/disable is part of this update. Gated by group.receipts.update.
    //
    //Future<SupplierProfile> updateSupplierProfile(int groupId, int profileId, UpsertSupplierProfileCommand upsertSupplierProfileCommand) async
    test('test updateSupplierProfile', () async {
      // TODO
    });

  });
}
