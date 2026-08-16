import { Category, SupplierProfile, Tag } from "../../open-api";

export interface AppliedSupplierSuggestions {
  categories: Category[];
  tags: Tag[];
  documentCurrencyCode?: string;
}

export function mergeCatalogItems<T extends { id?: number }>(
  existing: T[],
  additions: T[]
): T[] {
  const seen = new Set(existing.map((item) => item.id).filter((id) => id != null));
  const merged = [...existing];
  for (const item of additions) {
    if (item.id == null || seen.has(item.id)) {
      continue;
    }
    seen.add(item.id);
    merged.push(item);
  }
  return merged;
}

export function applySupplierSuggestions(
  existingCategories: Category[],
  existingTags: Tag[],
  currentCurrency: string,
  selectedCategories: Category[],
  selectedTags: Tag[],
  selectedCurrency?: string
): AppliedSupplierSuggestions {
  return {
    categories: mergeCatalogItems(existingCategories, selectedCategories),
    tags: mergeCatalogItems(existingTags, selectedTags),
    documentCurrencyCode: selectedCurrency || currentCurrency,
  };
}

export function supplierApplyConfirmation(
  addedCategories: Category[],
  addedTags: Tag[],
  currencyKept: boolean,
  currencyCode?: string
): string {
  const names = [
    ...addedCategories.map((item) => item.name).filter((name): name is string => !!name),
    ...addedTags.map((item) => item.name).filter((name): name is string => !!name),
  ];
  const parts: string[] = [];
  if (names.length > 0) {
    parts.push(`Added ${names.join(" and ")}`);
  }
  if (currencyKept && currencyCode) {
    parts.push(`kept receipt currency ${currencyCode}`);
  } else if (!currencyKept && currencyCode) {
    parts.push(`set document currency to ${currencyCode}`);
  }
  return parts.join("; ") + ".";
}

export function profileHasVisibleDefaults(profile: SupplierProfile): boolean {
  return (
    (profile.categories?.length ?? 0) > 0 ||
    (profile.tags?.length ?? 0) > 0 ||
    !!profile.expectedDocumentCurrencyCode
  );
}
