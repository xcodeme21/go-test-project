package migration

import (
	"gorm.io/gorm"
)

// SourceProduct :nodoc
type SourceProduct struct {
	gorm.Model
	ProductName  string `gorm:"type:varchar(150);"`
	Qty          int    `gorm:"default:0;"`
	SellingPrice int    `gorm:"default:0;"`
	PromoPrice   int    `gorm:"default:0;"`
}

// MigrateSourceProduct :nodoc
func MigrateSourceProduct(db *gorm.DB) {
	const tableName = "source_product"
	// todo to change if need to changing
	const version = "1.1"

	var migrateData Migrate
	var SourceProductData SourceProduct

	db.Where(&Migrate{Table: tableName}).First(&migrateData)

	// First create installer table
	if migrateData.Table == "" {
		if !db.Migrator().HasTable(&SourceProductData) {
			db.AutoMigrate(&SourceProduct{})
			db.Create(&Migrate{
				Table:   tableName,
				Version: version,
			})
		}
	}

	// Upgrade version
	if migrateData.Version == "1.0" {
		db.AutoMigrate(&SourceProduct{})
		migrateData.Version = version
		db.Save(&migrateData)
	}

}
