package models

type SupplierProfile struct {
	BaseModel
	GroupId                      uint                   `gorm:"not null;index;uniqueIndex:idx_supplier_profile_group_norm" json:"groupId"`
	Name                         string                 `gorm:"not null" json:"name"`
	NormalisedName               string                 `gorm:"not null;uniqueIndex:idx_supplier_profile_group_norm" json:"normalisedName"`
	ExpectedDocumentCurrencyCode *string                `gorm:"type:char(3)" json:"expectedDocumentCurrencyCode"`
	Enabled                      bool                   `gorm:"not null;default:true" json:"enabled"`
	Categories                   []Category             `gorm:"many2many:supplier_profile_categories" json:"categories"`
	Tags                         []Tag                  `gorm:"many2many:supplier_profile_tags" json:"tags"`
	Aliases                      []SupplierProfileAlias `gorm:"constraint:OnDelete:CASCADE" json:"aliases"`
}

type SupplierProfileAlias struct {
	BaseModel
	SupplierProfileId uint   `gorm:"not null;index" json:"supplierProfileId"`
	GroupId           uint   `gorm:"not null;index;uniqueIndex:idx_supplier_alias_group_norm" json:"groupId"`
	Name              string `gorm:"not null" json:"name"`
	NormalisedName    string `gorm:"not null;uniqueIndex:idx_supplier_alias_group_norm" json:"normalisedName"`
}

// SupplierProfileCategory is the explicit join row so deleting a category
// unlinks the profile default without deleting the profile.
type SupplierProfileCategory struct {
	SupplierProfileID uint            `gorm:"primaryKey;autoIncrement:false" json:"supplierProfileId"`
	CategoryID        uint            `gorm:"primaryKey;autoIncrement:false;index" json:"categoryId"`
	Category          Category        `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"-"`
	SupplierProfile   SupplierProfile `gorm:"foreignKey:SupplierProfileID;constraint:OnDelete:CASCADE" json:"-"`
}

// SupplierProfileTag is the explicit join row so deleting a tag unlinks the
// profile default without deleting the profile.
type SupplierProfileTag struct {
	SupplierProfileID uint            `gorm:"primaryKey;autoIncrement:false" json:"supplierProfileId"`
	TagID             uint            `gorm:"primaryKey;autoIncrement:false;index" json:"tagId"`
	Tag               Tag             `gorm:"foreignKey:TagID;constraint:OnDelete:CASCADE" json:"-"`
	SupplierProfile   SupplierProfile `gorm:"foreignKey:SupplierProfileID;constraint:OnDelete:CASCADE" json:"-"`
}
