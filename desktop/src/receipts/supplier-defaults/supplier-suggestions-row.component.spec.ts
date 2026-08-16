import { provideHttpClient, withInterceptorsFromDi } from "@angular/common/http";
import { provideHttpClientTesting } from "@angular/common/http/testing";
import { CUSTOM_ELEMENTS_SCHEMA, provideZonelessChangeDetection } from "@angular/core";
import { ComponentFixture, TestBed } from "@angular/core/testing";
import { FormBuilder, ReactiveFormsModule } from "@angular/forms";
import { MatDialog, MatDialogModule } from "@angular/material/dialog";
import { MatSnackBarModule } from "@angular/material/snack-bar";
import { of } from "rxjs";
import { ApiModule, SupplierProfileService } from "../../open-api";
import { SupplierSuggestionsRowComponent } from "./supplier-suggestions-row.component";

describe("SupplierSuggestionsRowComponent", () => {
  let component: SupplierSuggestionsRowComponent;
  let fixture: ComponentFixture<SupplierSuggestionsRowComponent>;
  let formBuilder: FormBuilder;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SupplierSuggestionsRowComponent],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
      imports: [ApiModule, MatDialogModule, MatSnackBarModule, ReactiveFormsModule],
      providers: [
        provideZonelessChangeDetection(),
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting(),
      ],
    }).compileComponents();

    formBuilder = TestBed.inject(FormBuilder);
    jest.spyOn(TestBed.inject(SupplierProfileService), "resolveSupplierProfile").mockReturnValue(
      of({ profile: { id: 1, name: "GitHub" } }) as any
    );
    fixture = TestBed.createComponent(SupplierSuggestionsRowComponent);
    component = fixture.componentInstance;
    fixture.componentRef.setInput(
      "form",
      formBuilder.group({
        name: ["GitHub, Inc."],
        documentCurrencyCode: ["AUD"],
        categories: formBuilder.array([{ id: 9, name: "Existing" }]),
        tags: formBuilder.array([]),
      })
    );
    fixture.componentRef.setInput("groupId", 4);
    fixture.componentRef.setInput("receiptName", "GitHub, Inc.");
    fixture.componentRef.setInput("canManage", true);
    fixture.componentRef.setInput("canApply", true);
  });

  it("resolves a matching profile from the receipt name", async () => {
    fixture.detectChanges();
    await new Promise((resolve) => setTimeout(resolve, 300));

    expect(component.matchedProfile()?.name).toEqual("GitHub");
  });

  it("opens the save-as-defaults dialog without changing the form", () => {
    const openSpy = jest.spyOn(TestBed.inject(MatDialog), "open").mockReturnValue({
      afterClosed: () => of(false),
    } as any);
    const originalName = component.form().get("name")?.value;

    component.saveAsDefaults();

    expect(openSpy).toHaveBeenCalled();
    expect(component.form().get("name")?.value).toEqual(originalName);
  });
});
