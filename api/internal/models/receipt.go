package models

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"time"
)

type Receipt struct {
	BaseModel
	Name   string          `gorm:"not null" json:"name"`
	Amount decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"amount"`
	// DocumentAmount and DocumentCurrencyCode preserve the value printed on the
	// source evidence. Amount remains the effective group base-currency value so
	// existing dashboards, settlement, and reports keep their meaning.
	DocumentAmount       decimal.Decimal    `gorm:"type:decimal(20,4);not null;default:0" json:"documentAmount"`
	DocumentCurrencyCode string             `gorm:"type:char(3);not null;default:'AUD'" json:"documentCurrencyCode"`
	EstimatedBaseAmount  *decimal.Decimal   `gorm:"type:decimal(20,4)" json:"estimatedBaseAmount"`
	FxRate               *decimal.Decimal   `gorm:"type:decimal(20,10)" json:"fxRate"`
	FxDate               *time.Time         `json:"fxDate"`
	FxProvider           *string            `gorm:"type:varchar(100)" json:"fxProvider"`
	FxRetrievedAt        *time.Time         `json:"fxRetrievedAt"`
	FxStatus             FxStatus           `gorm:"type:varchar(20);not null;default:'DOMESTIC'" json:"fxStatus"`
	Date                 time.Time          `gorm:"not null" json:"date"`
	ResolvedDate         *time.Time         `json:"resolvedDate"`
	PaidByUserID         uint               `json:"paidByUserId"`
	PaidByUser           User               `json:"-"`
	Status               ReceiptStatus      `gorm:"default:'OPEN';not null" json:"status"`
	GroupId              uint               `gorm:"not null" json:"groupId"`
	Group                Group              `json:"-"`
	Categories           []Category         `gorm:"many2many:receipt_categories" json:"categories"`
	Tags                 []Tag              `gorm:"many2many:receipt_tags" json:"tags"`
	ImageFiles           []FileData         `json:"imageFiles"`
	ReceiptItems         []Item             `json:"receiptItems"`
	Comments             []Comment          `json:"comments"`
	CustomFields         []CustomFieldValue `json:"customFields"`
}

func (r *Receipt) ToString() (string, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
