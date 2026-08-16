import { Component, Inject, OnInit } from "@angular/core";
import { FormBuilder, FormControl, FormGroup } from "@angular/forms";
import { MAT_DIALOG_DATA, MatDialogRef } from "@angular/material/dialog";
import { Category, SupplierProfile, Tag } from "../../open-api";
import { applySupplierSuggestions, AppliedSupplierSuggestions } from "./supplier-profile.utils";

export interface ReviewSuggestionsDialogData {
  profile: SupplierProfile;
  existingCategories: Category[];
  existingTags: Tag[];
  currentCurrency: string;
  currencyIsExtracted: boolean;
}

@Component({
  selector: "app-review-suggestions-dialog",
  templateUrl: "./review-suggestions-dialog.component.html",
  styleUrls: ["./review-suggestions-dialog.component.scss"],
  standalone: false,
})
export class ReviewSuggestionsDialogComponent implements OnInit {
  public form: FormGroup = new FormGroup({});

  public currencyConflict = false;

  constructor(
    @Inject(MAT_DIALOG_DATA) public readonly data: ReviewSuggestionsDialogData,
    private formBuilder: FormBuilder,
    private matDialogRef: MatDialogRef<
      ReviewSuggestionsDialogComponent,
      AppliedSupplierSuggestions | undefined
    >
  ) {}

  public ngOnInit(): void {
    const expected = (this.data.profile.expectedDocumentCurrencyCode ?? "").toUpperCase();
    const current = (this.data.currentCurrency ?? "").toUpperCase();
    this.currencyConflict = !!expected && expected !== current;

    const categoryGroup: Record<string, FormControl<boolean | null>> = {};
    for (const category of this.data.profile.categories ?? []) {
      if (category.id != null) {
        categoryGroup[String(category.id)] = this.formBuilder.control(true);
      }
    }

    const tagGroup: Record<string, FormControl<boolean | null>> = {};
    for (const tag of this.data.profile.tags ?? []) {
      if (tag.id != null) {
        tagGroup[String(tag.id)] = this.formBuilder.control(true);
      }
    }

    this.form = this.formBuilder.group({
      categories: this.formBuilder.group(categoryGroup),
      tags: this.formBuilder.group(tagGroup),
      applyCurrency: [!!expected && !this.currencyConflict],
    });
  }

  public categoryControl(id?: number): FormControl {
    return this.form.get(["categories", String(id)]) as FormControl;
  }

  public tagControl(id?: number): FormControl {
    return this.form.get(["tags", String(id)]) as FormControl;
  }

  public get currencyControl(): FormControl {
    return this.form.get("applyCurrency") as FormControl;
  }

  public apply(): void {
    const selectedCategories = (this.data.profile.categories ?? []).filter(
      (category) => !!this.categoryControl(category.id)?.value
    );
    const selectedTags = (this.data.profile.tags ?? []).filter(
      (tag) => !!this.tagControl(tag.id)?.value
    );
    const selectedCurrency = this.currencyControl.value
      ? this.data.profile.expectedDocumentCurrencyCode
      : undefined;

    this.matDialogRef.close(
      applySupplierSuggestions(
        this.data.existingCategories,
        this.data.existingTags,
        this.data.currentCurrency,
        selectedCategories,
        selectedTags,
        selectedCurrency
      )
    );
  }

  public cancel(): void {
    this.matDialogRef.close(undefined);
  }
}
