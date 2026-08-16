import { applySupplierSuggestions, mergeCatalogItems, profileHasVisibleDefaults, supplierApplyConfirmation } from "./supplier-profile.utils";

describe("supplier-profile.utils", () => {
  it("merges additions without removing existing items", () => {
    const merged = mergeCatalogItems(
      [{ id: 1, name: "Existing" }],
      [{ id: 2, name: "Software" }, { id: 1, name: "Existing" }]
    );

    expect(merged.map((item) => item.id)).toEqual([1, 2]);
  });

  it("applies selected currency only when provided", () => {
    const result = applySupplierSuggestions(
      [{ id: 1, name: "Food" }],
      [{ id: 3, name: "Personal" }],
      "AUD",
      [{ id: 2, name: "Software" }],
      [{ id: 4, name: "Work" }],
      "USD"
    );

    expect(result.categories.map((item) => item.name)).toEqual(["Food", "Software"]);
    expect(result.tags.map((item) => item.name)).toEqual(["Personal", "Work"]);
    expect(result.documentCurrencyCode).toEqual("USD");
  });

  it("keeps the receipt currency when no currency is selected", () => {
    const result = applySupplierSuggestions([], [], "AUD", [], [], undefined);
    expect(result.documentCurrencyCode).toEqual("AUD");
  });

  it("builds a concise confirmation", () => {
    expect(
      supplierApplyConfirmation(
        [{ name: "Software" }],
        [{ name: "Work" }],
        true,
        "AUD"
      )
    ).toEqual("Added Software and Work; kept receipt currency AUD.");
  });

  it("treats a currency-only profile as having visible defaults", () => {
    expect(profileHasVisibleDefaults({ expectedDocumentCurrencyCode: "USD" })).toBe(true);
    expect(profileHasVisibleDefaults({ categories: [], tags: [] })).toBe(false);
  });
});
