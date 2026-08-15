package models

import (
	"database/sql/driver"
	"fmt"
)

// FxStatus describes how a receipt's effective base-currency Amount was
// obtained. Estimated values are advisory and deliberately separate from the
// receipt's workflow Status.
type FxStatus string

const (
	FX_DOMESTIC     FxStatus = "DOMESTIC"
	FX_ESTIMATED    FxStatus = "ESTIMATED"
	FX_CONFIRMED    FxStatus = "CONFIRMED"
	FX_NEEDS_REVIEW FxStatus = "NEEDS_REVIEW"
)

func (status *FxStatus) Scan(value interface{}) error {
	switch typed := value.(type) {
	case []byte:
		*status = FxStatus(string(typed))
	case string:
		*status = FxStatus(typed)
	default:
		return fmt.Errorf("cannot scan FX status from %T", value)
	}
	return nil
}

func (status FxStatus) Value() (driver.Value, error) {
	return string(status), nil
}

func IsValidFxStatus(status FxStatus) bool {
	switch status {
	case FX_DOMESTIC, FX_ESTIMATED, FX_CONFIRMED, FX_NEEDS_REVIEW:
		return true
	default:
		return false
	}
}
