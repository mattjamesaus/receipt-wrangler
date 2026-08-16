import { AfterViewInit, Component, computed, DestroyRef, inject, OnInit, signal, TemplateRef, viewChild } from "@angular/core";
import { takeUntilDestroyed } from "@angular/core/rxjs-interop";
import { FormControl } from "@angular/forms";
import { MatDialog } from "@angular/material/dialog";
import { MatTableDataSource } from "@angular/material/table";
import { ActivatedRoute } from "@angular/router";
import { Store } from "@ngxs/store";
import { take, tap } from "rxjs";
import { DEFAULT_DIALOG_CONFIG } from "src/constants";
import { ConfirmationDialogComponent } from "src/shared-ui/confirmation-dialog/confirmation-dialog.component";
import { TableColumn } from "src/table/table-column.interface";
import { TableComponent } from "src/table/table/table.component";
import { Permission, SupplierProfile, SupplierProfileService } from "../../open-api";
import { SnackbarService } from "../../services";
import { AuthState, GroupState } from "../../store";
import { SupplierProfileFormDialogComponent, SupplierProfileFormDialogData } from "./supplier-profile-form-dialog.component";

@Component({
  selector: "app-supplier-defaults-list",
  templateUrl: "./supplier-defaults-list.component.html",
  styleUrls: ["./supplier-defaults-list.component.scss"],
  standalone: false,
})
export class SupplierDefaultsListComponent implements OnInit, AfterViewInit {
  public readonly nameCell = viewChild.required<TemplateRef<any>>("nameCell");

  public readonly currencyCell = viewChild.required<TemplateRef<any>>("currencyCell");

  public readonly categoriesCell = viewChild.required<TemplateRef<any>>("categoriesCell");

  public readonly tagsCell = viewChild.required<TemplateRef<any>>("tagsCell");

  public readonly aliasesCell = viewChild.required<TemplateRef<any>>("aliasesCell");

  public readonly enabledCell = viewChild.required<TemplateRef<any>>("enabledCell");

  public readonly updatedCell = viewChild.required<TemplateRef<any>>("updatedCell");

  public readonly actionsCell = viewChild.required<TemplateRef<any>>("actionsCell");

  public readonly table = viewChild.required(TableComponent);

  protected readonly Permission = Permission;

  public search = new FormControl("");

  public searchQuery = signal("");

  public profiles = signal<SupplierProfile[]>([]);

  private readonly destroyRef = inject(DestroyRef);

  public dataSource = computed(() => {
    const query = this.searchQuery().trim().toLowerCase();
    const filtered = this.profiles().filter((profile) => {
      if (!query) {
        return true;
      }
      const haystack = [
        profile.name,
        profile.expectedDocumentCurrencyCode,
        ...(profile.aliases ?? []).map((alias) => alias.name),
        ...(profile.categories ?? []).map((category) => category.name),
        ...(profile.tags ?? []).map((tag) => tag.name),
      ]
        .filter(Boolean)
        .join(" ")
        .toLowerCase();
      return haystack.includes(query);
    });
    return new MatTableDataSource<SupplierProfile>(filtered);
  });

  public displayedColumns: string[] = [];

  public columns: TableColumn[] = [];

  public groupId = 0;

  public canManage = false;

  constructor(
    private activatedRoute: ActivatedRoute,
    private matDialog: MatDialog,
    private snackbarService: SnackbarService,
    private store: Store,
    private supplierProfileService: SupplierProfileService
  ) {}

  public ngOnInit(): void {
    this.groupId = Number(
      this.activatedRoute.snapshot.params["groupId"] ??
        this.store.selectSnapshot(GroupState.selectedGroupId)
    );
    this.canManage = this.store.selectSnapshot(
      AuthState.hasGroupPermission(this.groupId, Permission.GroupReceiptsUpdate)
    );
    this.search.valueChanges
      .pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe((value) => this.searchQuery.set(value ?? ""));
    this.loadProfiles();
  }

  public names(items?: Array<{ name?: string }>): string {
    return (items ?? [])
      .map((item) => item.name)
      .filter((name): name is string => !!name)
      .join(", ") || "—";
  }

  public aliasNames(profile: SupplierProfile): string {
    return this.names(profile.aliases);
  }

  public ngAfterViewInit(): void {
    this.columns = [
      { columnHeader: "Name", matColumnDef: "name", sortable: false, template: this.nameCell() },
      { columnHeader: "Expected currency", matColumnDef: "currency", sortable: false, template: this.currencyCell() },
      { columnHeader: "Categories", matColumnDef: "categories", sortable: false, template: this.categoriesCell() },
      { columnHeader: "Tags", matColumnDef: "tags", sortable: false, template: this.tagsCell() },
      { columnHeader: "Aliases", matColumnDef: "aliases", sortable: false, template: this.aliasesCell() },
      { columnHeader: "Enabled", matColumnDef: "enabled", sortable: false, template: this.enabledCell() },
      { columnHeader: "Updated", matColumnDef: "updatedAt", sortable: false, template: this.updatedCell() },
      { columnHeader: "Actions", matColumnDef: "actions", sortable: false, template: this.actionsCell() },
    ];
    this.displayedColumns = this.columns.map((column) => column.matColumnDef);
  }

  public openAddDialog(): void {
    this.openFormDialog();
  }

  public openEditDialog(profile: SupplierProfile): void {
    this.openFormDialog(profile);
  }

  public toggleEnabled(profile: SupplierProfile): void {
    if (!profile.id) {
      return;
    }
    this.supplierProfileService
      .updateSupplierProfile(this.groupId, profile.id, {
        name: profile.name ?? "",
        aliases: (profile.aliases ?? []).map((alias) => alias.name ?? "").filter(Boolean),
        categoryIds: (profile.categories ?? []).map((category) => category.id).filter((id): id is number => !!id),
        tagIds: (profile.tags ?? []).map((tag) => tag.id).filter((id): id is number => !!id),
        expectedDocumentCurrencyCode: profile.expectedDocumentCurrencyCode,
        enabled: !profile.enabled,
      })
      .pipe(
        take(1),
        tap(() => {
          this.snackbarService.success(
            profile.enabled ? "Supplier defaults disabled" : "Supplier defaults enabled"
          );
          this.loadProfiles();
        })
      )
      .subscribe();
  }

  public openDeleteDialog(profile: SupplierProfile): void {
    const dialogRef = this.matDialog.open(ConfirmationDialogComponent, DEFAULT_DIALOG_CONFIG);
    dialogRef.componentInstance.headerText = "Delete supplier defaults";
    dialogRef.componentInstance.dialogContent = `Delete defaults for ${profile.name}? Receipts are not changed.`;
    dialogRef.afterClosed().subscribe((confirmed) => {
      if (!confirmed || !profile.id) {
        return;
      }
      this.supplierProfileService
        .deleteSupplierProfile(this.groupId, profile.id)
        .pipe(
          take(1),
          tap(() => {
            this.snackbarService.success("Supplier defaults deleted");
            this.loadProfiles();
          })
        )
        .subscribe();
    });
  }

  private openFormDialog(profile?: SupplierProfile): void {
    const dialogRef = this.matDialog.open(SupplierProfileFormDialogComponent, {
      ...DEFAULT_DIALOG_CONFIG,
      data: { groupId: this.groupId, profile } satisfies SupplierProfileFormDialogData,
    });
    dialogRef.afterClosed().subscribe((refresh) => {
      if (refresh) {
        this.loadProfiles();
      }
    });
  }

  private loadProfiles(): void {
    this.supplierProfileService
      .getSupplierProfilesForGroup(this.groupId)
      .pipe(
        take(1),
        tap((profiles) => this.profiles.set(profiles ?? []))
      )
      .subscribe();
  }
}
