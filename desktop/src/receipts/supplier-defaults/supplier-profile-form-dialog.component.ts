import { Component, Inject, OnInit } from "@angular/core";
import { FormArray, FormBuilder, FormControl, FormGroup, Validators } from "@angular/forms";
import { MAT_DIALOG_DATA, MatDialogRef } from "@angular/material/dialog";
import { Store } from "@ngxs/store";
import { catchError, EMPTY, take, tap } from "rxjs";
import {
  Category,
  SupplierProfile,
  SupplierProfileService,
  Tag,
  UpsertSupplierProfileCommand,
} from "../../open-api";
import { SnackbarService } from "../../services";
import { AuthState } from "../../store";

export interface SupplierProfileFormDialogData {
  groupId: number;
  profile?: SupplierProfile;
  prefills?: {
    name?: string;
    aliases?: string[];
    categories?: Category[];
    tags?: Tag[];
    expectedDocumentCurrencyCode?: string;
  };
}

@Component({
  selector: "app-supplier-profile-form-dialog",
  templateUrl: "./supplier-profile-form-dialog.component.html",
  styleUrls: ["./supplier-profile-form-dialog.component.scss"],
  standalone: false,
})
export class SupplierProfileFormDialogComponent implements OnInit {
  public form: FormGroup = new FormGroup({});

  public categories: Category[] = [];

  public tags: Tag[] = [];

  public aliasDraft = new FormControl("");

  public headerText = "";

  constructor(
    @Inject(MAT_DIALOG_DATA) public readonly data: SupplierProfileFormDialogData,
    private formBuilder: FormBuilder,
    private matDialogRef: MatDialogRef<SupplierProfileFormDialogComponent>,
    private snackbarService: SnackbarService,
    private store: Store,
    private supplierProfileService: SupplierProfileService
  ) {}

  public ngOnInit(): void {
    this.categories = this.store.selectSnapshot(
      AuthState.groupCategories(this.data.groupId)
    );
    this.tags = this.store.selectSnapshot(AuthState.groupTags(this.data.groupId));
    this.headerText = this.data.profile
      ? `Edit ${this.data.profile.name}`
      : "Save as supplier defaults";
    this.initForm();
  }

  public get aliases(): FormArray {
    return this.form.get("aliases") as FormArray;
  }

  public get categoryControls(): FormArray {
    return this.form.get("categories") as FormArray;
  }

  public get tagControls(): FormArray {
    return this.form.get("tags") as FormArray;
  }

  public addAlias(): void {
    const value = (this.aliasDraft.value ?? "").trim();
    if (!value) {
      return;
    }
    const exists = this.aliases.controls.some(
      (control) =>
        (control.value as string).trim().toLowerCase() === value.toLowerCase()
    );
    if (!exists) {
      this.aliases.push(this.formBuilder.control(value));
    }
    this.aliasDraft.setValue("");
  }

  public removeAlias(index: number): void {
    this.aliases.removeAt(index);
  }

  public submit(): void {
    if (this.form.invalid) {
      return;
    }

    const command: UpsertSupplierProfileCommand = {
      name: this.form.value.name,
      aliases: this.aliases.controls
        .map((control) => (control.value as string)?.trim())
        .filter((alias) => !!alias),
      categoryIds: this.categoryControls.controls
        .map((control) => control.value?.id)
        .filter((id): id is number => !!id),
      tagIds: this.tagControls.controls
        .map((control) => control.value?.id)
        .filter((id): id is number => !!id),
      expectedDocumentCurrencyCode:
        this.form.value.expectedDocumentCurrencyCode || undefined,
      enabled: this.form.value.enabled,
      autoApply: this.form.value.autoApply,
    };

    const request = this.data.profile?.id
      ? this.supplierProfileService.updateSupplierProfile(
          this.data.groupId,
          this.data.profile.id,
          command
        )
      : this.supplierProfileService.createSupplierProfile(
          this.data.groupId,
          command
        );

    request
      .pipe(
        take(1),
        tap(() => {
          this.snackbarService.success(
            this.data.profile
              ? "Supplier defaults updated"
              : "Supplier defaults saved"
          );
          this.matDialogRef.close(true);
        }),
        catchError(() => EMPTY)
      )
      .subscribe();
  }

  public closeDialog(): void {
    this.matDialogRef.close(false);
  }

  private initForm(): void {
    const profile = this.data.profile;
    const prefills = this.data.prefills;
    const categories = profile?.categories ?? prefills?.categories ?? [];
    const tags = profile?.tags ?? prefills?.tags ?? [];
    const aliases = profile?.aliases?.map((alias) => alias.name ?? "") ??
      prefills?.aliases ??
      [];

    this.form = this.formBuilder.group(
      {
        name: [
          profile?.name ?? prefills?.name ?? "",
          Validators.required,
        ],
        aliases: this.formBuilder.array(
          aliases.filter((alias) => !!alias).map((alias) => this.formBuilder.control(alias))
        ),
        categories: this.formBuilder.array(
          categories.map((category) => this.formBuilder.control(category))
        ),
        tags: this.formBuilder.array(
          tags.map((tag) => this.formBuilder.control(tag))
        ),
        expectedDocumentCurrencyCode: [
          profile?.expectedDocumentCurrencyCode ??
            prefills?.expectedDocumentCurrencyCode ??
            "",
          Validators.pattern(/^[A-Za-z]{3}$/),
        ],
        enabled: [profile?.enabled ?? true],
        autoApply: [profile?.autoApply ?? false],
      },
      { validators: [atLeastOneSupplierDefault] }
    );
  }
}

function atLeastOneSupplierDefault(group: FormGroup) {
  const categories = group.get("categories") as FormArray;
  const tags = group.get("tags") as FormArray;
  const currency = (group.get("expectedDocumentCurrencyCode")?.value ?? "").trim();
  if ((categories?.length ?? 0) > 0 || (tags?.length ?? 0) > 0 || currency) {
    return null;
  }
  return { defaultsRequired: true };
}
