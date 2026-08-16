import { Component, computed, DestroyRef, inject, input, output, signal } from "@angular/core";
import { takeUntilDestroyed, toObservable } from "@angular/core/rxjs-interop";
import { FormArray, FormBuilder, FormGroup } from "@angular/forms";
import { MatDialog } from "@angular/material/dialog";
import { combineLatest, debounceTime, distinctUntilChanged, of, switchMap } from "rxjs";
import { DEFAULT_DIALOG_CONFIG } from "src/constants";
import { Category, SupplierProfile, SupplierProfileService, Tag } from "../../open-api";
import { SnackbarService } from "../../services";
import { ReviewSuggestionsDialogComponent, ReviewSuggestionsDialogData } from "./review-suggestions-dialog.component";
import { SupplierProfileFormDialogComponent, SupplierProfileFormDialogData } from "./supplier-profile-form-dialog.component";
import { profileHasVisibleDefaults, supplierApplyConfirmation } from "./supplier-profile.utils";

@Component({
  selector: "app-supplier-suggestions-row",
  templateUrl: "./supplier-suggestions-row.component.html",
  styleUrls: ["./supplier-suggestions-row.component.scss"],
  standalone: false,
})
export class SupplierSuggestionsRowComponent {
  public readonly form = input.required<FormGroup>();

  public readonly groupId = input.required<number>();

  public readonly receiptName = input("");

  public readonly canManage = input(false);

  public readonly canApply = input(false);

  public readonly readonly = input(false);

  public readonly currencyIsExtracted = input(false);

  public readonly profileSaved = output<void>();

  public readonly matchedProfile = signal<SupplierProfile | null>(null);

  public readonly resolving = signal(false);

  public readonly hasName = computed(() => this.receiptName().trim().length > 0);

  public readonly hasVisibleDefaults = computed(() => {
    const profile = this.matchedProfile();
    return !!profile && profileHasVisibleDefaults(profile);
  });

  private readonly destroyRef = inject(DestroyRef);

  constructor(
    private formBuilder: FormBuilder,
    private matDialog: MatDialog,
    private snackbarService: SnackbarService,
    private supplierProfileService: SupplierProfileService
  ) {
    combineLatest([toObservable(this.receiptName), toObservable(this.groupId)])
      .pipe(
        debounceTime(250),
        distinctUntilChanged(
          (previous, current) =>
            previous[0] === current[0] && previous[1] === current[1]
        ),
        switchMap(([name, groupId]) => {
          const trimmed = name.trim();
          if (!trimmed || !groupId) {
            this.matchedProfile.set(null);
            return of(null);
          }
          this.resolving.set(true);
          return this.supplierProfileService.resolveSupplierProfile(groupId, {
            name: trimmed,
          });
        }),
        takeUntilDestroyed(this.destroyRef)
      )
      .subscribe((response) => {
        this.resolving.set(false);
        this.matchedProfile.set(response?.profile ?? null);
      });
  }

  public reviewSuggestions(): void {
    const profile = this.matchedProfile();
    if (!profile || this.readonly()) {
      return;
    }

    const dialogRef = this.matDialog.open(ReviewSuggestionsDialogComponent, {
      ...DEFAULT_DIALOG_CONFIG,
      data: {
        profile,
        existingCategories: this.currentCategories(),
        existingTags: this.currentTags(),
        currentCurrency: this.form().get("documentCurrencyCode")?.value ?? "",
        currencyIsExtracted: this.currencyIsExtracted(),
      } satisfies ReviewSuggestionsDialogData,
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (!result) {
        return;
      }
      this.patchForm(result.categories, result.tags, result.documentCurrencyCode);
    });
  }

  public saveAsDefaults(): void {
    if (!this.canManage()) {
      return;
    }

    const dialogRef = this.matDialog.open(SupplierProfileFormDialogComponent, {
      ...DEFAULT_DIALOG_CONFIG,
      data: {
        groupId: this.groupId(),
        prefills: {
          name: this.receiptName().trim(),
          categories: this.currentCategories(),
          tags: this.currentTags(),
          expectedDocumentCurrencyCode: this.form().get("documentCurrencyCode")?.value,
        },
      } satisfies SupplierProfileFormDialogData,
    });

    dialogRef.afterClosed().subscribe((saved) => {
      if (saved) {
        this.profileSaved.emit();
        this.refreshMatch();
      }
    });
  }

  public manageProfile(): void {
    const profile = this.matchedProfile();
    if (!profile || !this.canManage()) {
      return;
    }

    const dialogRef = this.matDialog.open(SupplierProfileFormDialogComponent, {
      ...DEFAULT_DIALOG_CONFIG,
      data: {
        groupId: this.groupId(),
        profile,
      } satisfies SupplierProfileFormDialogData,
    });

    dialogRef.afterClosed().subscribe((saved) => {
      if (saved) {
        this.profileSaved.emit();
        this.refreshMatch();
      }
    });
  }

  private refreshMatch(): void {
    const name = this.receiptName().trim();
    const groupId = this.groupId();
    if (!name || !groupId) {
      return;
    }
    this.supplierProfileService
      .resolveSupplierProfile(groupId, { name })
      .subscribe((response) => this.matchedProfile.set(response?.profile ?? null));
  }

  private currentCategories(): Category[] {
    return ((this.form().get("categories") as FormArray)?.value ?? []) as Category[];
  }

  private currentTags(): Tag[] {
    return ((this.form().get("tags") as FormArray)?.value ?? []) as Tag[];
  }

  private patchForm(
    categories: Category[],
    tags: Tag[],
    documentCurrencyCode?: string
  ): void {
    const previousCategories = this.currentCategories();
    const previousTags = this.currentTags();
    const previousCurrency = this.form().get("documentCurrencyCode")?.value ?? "";

    this.replaceArray("categories", categories);
    this.replaceArray("tags", tags);
    if (documentCurrencyCode) {
      this.form().get("documentCurrencyCode")?.setValue(documentCurrencyCode);
    }

    const addedCategories = categories.filter(
      (category) => !previousCategories.some((existing) => existing.id === category.id)
    );
    const addedTags = tags.filter(
      (tag) => !previousTags.some((existing) => existing.id === tag.id)
    );
    const currencyKept = !documentCurrencyCode || documentCurrencyCode === previousCurrency;
    this.snackbarService.success(
      supplierApplyConfirmation(
        addedCategories,
        addedTags,
        currencyKept,
        currencyKept ? previousCurrency : documentCurrencyCode
      )
    );
  }

  private replaceArray(key: "categories" | "tags", items: Category[] | Tag[]): void {
    const array = this.form().get(key) as FormArray;
    array.clear();
    items.forEach((item) => array.push(this.formBuilder.control(item)));
  }
}
