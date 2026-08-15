import 'package:test/test.dart';
import 'package:openapi/openapi.dart';

// tests for FxStatus
void main() {
  group(FxStatus, () {
    test('exposes every supported provenance state', () {
      expect(
        FxStatus.values.toSet(),
        equals({
          FxStatus.DOMESTIC,
          FxStatus.ESTIMATED,
          FxStatus.CONFIRMED,
          FxStatus.NEEDS_REVIEW,
        }),
      );
    });

    test('rejects unknown provenance states', () {
      expect(() => FxStatus.valueOf('UNKNOWN'), throwsArgumentError);
    });
  });
}
