import { provideHttpClient, withInterceptorsFromDi } from "@angular/common/http";
import { provideHttpClientTesting } from "@angular/common/http/testing";
import { CUSTOM_ELEMENTS_SCHEMA, provideZonelessChangeDetection } from "@angular/core";
import { ComponentFixture, TestBed } from "@angular/core/testing";
import { ReactiveFormsModule } from "@angular/forms";
import { MatDialog, MatDialogModule } from "@angular/material/dialog";
import { MatSnackBarModule } from "@angular/material/snack-bar";
import { ActivatedRoute } from "@angular/router";
import { NgxsModule } from "@ngxs/store";
import { of } from "rxjs";
import { ApiModule, SupplierProfileService } from "../../open-api";
import { AuthState } from "../../store";
import { SupplierDefaultsListComponent } from "./supplier-defaults-list.component";

describe("SupplierDefaultsListComponent", () => {
  let component: SupplierDefaultsListComponent;
  let fixture: ComponentFixture<SupplierDefaultsListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SupplierDefaultsListComponent],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
      imports: [
        ApiModule,
        MatDialogModule,
        MatSnackBarModule,
        NgxsModule.forRoot([AuthState]),
        ReactiveFormsModule,
      ],
      providers: [
        provideZonelessChangeDetection(),
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting(),
        {
          provide: ActivatedRoute,
          useValue: { snapshot: { params: { groupId: "7" } } },
        },
      ],
    }).compileComponents();

    jest.spyOn(TestBed.inject(SupplierProfileService), "getSupplierProfilesForGroup").mockReturnValue(
      of([
        { id: 1, name: "GitHub", expectedDocumentCurrencyCode: "USD", enabled: true, aliases: [] },
        { id: 2, name: "Officeworks", expectedDocumentCurrencyCode: "AUD", enabled: false, aliases: [] },
      ]) as any
    );

    fixture = TestBed.createComponent(SupplierDefaultsListComponent);
    component = fixture.componentInstance;
    component.ngOnInit();
  });

  it("loads profiles for the route group", () => {
    expect(TestBed.inject(SupplierProfileService).getSupplierProfilesForGroup).toHaveBeenCalledWith(7);
    expect(component.profiles().map((profile) => profile.name)).toEqual(["GitHub", "Officeworks"]);
  });

  it("filters the table by search text", () => {
    component.searchQuery.set("git");
    expect(component.dataSource().data.map((profile) => profile.name)).toEqual(["GitHub"]);
  });

  it("opens the add dialog", () => {
    const openSpy = jest.spyOn(TestBed.inject(MatDialog), "open").mockReturnValue({
      afterClosed: () => of(false),
    } as any);

    component.openAddDialog();

    expect(openSpy).toHaveBeenCalled();
  });
});
