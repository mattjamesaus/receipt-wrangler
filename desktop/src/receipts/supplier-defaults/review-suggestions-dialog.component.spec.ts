import { CUSTOM_ELEMENTS_SCHEMA, provideZonelessChangeDetection } from "@angular/core";
import { ComponentFixture, TestBed } from "@angular/core/testing";
import { ReactiveFormsModule } from "@angular/forms";
import { MAT_DIALOG_DATA, MatDialogRef } from "@angular/material/dialog";
import { ReviewSuggestionsDialogComponent } from "./review-suggestions-dialog.component";

describe("ReviewSuggestionsDialogComponent", () => {
  let component: ReviewSuggestionsDialogComponent;
  let fixture: ComponentFixture<ReviewSuggestionsDialogComponent>;
  let dialogRef: { close: jest.Mock };

  beforeEach(async () => {
    dialogRef = { close: jest.fn() };
    await TestBed.configureTestingModule({
      declarations: [ReviewSuggestionsDialogComponent],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
      imports: [ReactiveFormsModule],
      providers: [
        provideZonelessChangeDetection(),
        { provide: MatDialogRef, useValue: dialogRef },
        {
          provide: MAT_DIALOG_DATA,
          useValue: {
            profile: {
              name: "GitHub",
              expectedDocumentCurrencyCode: "USD",
              categories: [{ id: 2, name: "Software" }],
              tags: [{ id: 4, name: "Work" }],
            },
            existingCategories: [{ id: 1, name: "Food" }],
            existingTags: [{ id: 3, name: "Personal" }],
            currentCurrency: "AUD",
            currencyIsExtracted: true,
          },
        },
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(ReviewSuggestionsDialogComponent);
    component = fixture.componentInstance;
    component.ngOnInit();
    await fixture.whenStable();
  });

  it("leaves a conflicting currency unselected", () => {
    expect(component.currencyConflict).toBe(true);
    expect(component.currencyControl.value).toBe(false);
  });

  it("merges selected defaults without removing existing values", () => {
    component.apply();

    expect(dialogRef.close).toHaveBeenCalledWith({
      categories: [
        { id: 1, name: "Food" },
        { id: 2, name: "Software" },
      ],
      tags: [
        { id: 3, name: "Personal" },
        { id: 4, name: "Work" },
      ],
      documentCurrencyCode: "AUD",
    });
  });
});
