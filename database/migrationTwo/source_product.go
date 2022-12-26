package migration

import (
	"gorm.io/gorm"
)

// DestinationProduct :nodoc
type DestinationProduct struct {
	gorm.Model
	ProductName  string `gorm:"type:varchar(150);"`
	Qty          int    `gorm:"default:0;"`
	SellingPrice int    `gorm:"default:0;"`
	PromoPrice   int    `gorm:"default:0;"`
}

// MigrateDestinationProduct :nodoc
func MigrateDestinationProduct(db *gorm.DB) {
	const tableName = "destination_product"
	// todo to change if need to changing
	const version = "1.1"

	var migrateData Migrate
	var DestinationProductData DestinationProduct

	db.Where(&Migrate{Table: tableName}).First(&migrateData)

	// First create installer table
	if migrateData.Table == "" {
		if !db.Migrator().HasTable(&DestinationProductData) {
			db.AutoMigrate(&DestinationProduct{})
			db.Create(&Migrate{
				Table:   tableName,
				Version: version,
			})
		}
	}

	// Upgrade version
	if migrateData.Version == "1.0" {
		db.AutoMigrate(&DestinationProduct{})
		migrateData.Version = version
		db.Save(&migrateData)
	}

}
