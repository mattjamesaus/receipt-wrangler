import { provideHttpClient, withInterceptorsFromDi } from "@angular/common/http";
import { provideHttpClientTesting } from "@angular/common/http/testing";
import { CUSTOM_ELEMENTS_SCHEMA, provideZonelessChangeDetection } from "@angular/core";
import { ComponentFixture, TestBed } from "@angular/core/testing";
import { ReactiveFormsModule } from "@angular/forms";
import { MAT_DIALOG_DATA, MatDialogRef } from "@angular/material/dialog";
import { MatSnackBarModule } from "@angular/material/snack-bar";
import { NgxsModule } from "@ngxs/store";
import { of } from "rxjs";
import { PipesModule } from "src/pipes/pipes.module";
import { ApiModule, SupplierProfileService } from "../../open-api";
import { AuthState } from "../../store";
import { SupplierProfileFormDialogComponent } from "./supplier-profile-form-dialog.component";

describe("SupplierProfileFormDialogComponent", () => {
  let component: SupplierProfileFormDialogComponent;
  let fixture: ComponentFixture<SupplierProfileFormDialogComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SupplierProfileFormDialogComponent],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
      imports: [
        ApiModule,
        MatSnackBarModule,
        NgxsModule.forRoot([AuthState]),
        PipesModule,
        ReactiveFormsModule,
      ],
      providers: [
        provideZonelessChangeDetection(),
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting(),
        { provide: MatDialogRef, useValue: { close: jest.fn() } },
        {
          provide: MAT_DIALOG_DATA,
          useValue: {
            groupId: 3,
            prefills: {
              name: "GitHub",
              categories: [{ id: 1, name: "Software" }],
              expectedDocumentCurrencyCode: "USD",
            },
          },
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(SupplierProfileFormDialogComponent);
    component = fixture.componentInstance;
    component.ngOnInit();
  });

  it("prefills from the receipt without requiring an existing profile", () => {
    expect(component.form.value.name).toEqual("GitHub");
    expect(component.form.valid).toBe(true);
  });

  it("creates a profile from the prefilled form", () => {
    const createSpy = jest
      .spyOn(TestBed.inject(SupplierProfileService), "createSupplierProfile")
      .mockReturnValue(of({ id: 9, name: "GitHub" }) as any);

    component.submit();

    expect(createSpy).toHaveBeenCalledWith(3, expect.objectContaining({
      name: "GitHub",
      categoryIds: [1],
      expectedDocumentCurrencyCode: "USD",
      enabled: true,
    }));
  });
});
