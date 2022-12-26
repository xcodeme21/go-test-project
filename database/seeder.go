package database

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/xcodeme21/go-test-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SourceSeeder() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"), os.Getenv("DB_NAME_1"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db)

	var cek models.SourceProduct
	db.Table("source_products").First(&cek)
	if cek.ID != 0 {
		fmt.Println("Data found")
	} else {
		// Tambahkan data ke tabel
		for i := 1; i <= 500; i++ {
			name := fmt.Sprintf("Product %d", i)
			randomNumber := rand.Intn(100) + 1
			db.Create(&models.SourceProduct{ProductName: name, Qty: randomNumber, SellingPrice: 15000 + i, PromoPrice: 12000 + i, ID: i})
		}

		// Menampilkan data
		var products []models.SourceProduct
		db.Find(&products)
		fmt.Println("Total products:", len(products))
	}
}

func DestinationSeeder() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_TIMEZONE"), os.Getenv("DB_NAME_2"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(db)

	var cek models.DestinationProduct
	db.Table("destination_products").First(&cek)
	if cek.ID != 0 {
		fmt.Println("Data found")
	} else {
		// Tambahkan data ke tabel
		for i := 1; i <= 500; i++ {
			name := fmt.Sprintf("Product %d", i)
			db.Create(&models.DestinationProduct{ProductName: name, Qty: 0, SellingPrice: 0, PromoPrice: 0, ID: i})
		}

		// Menampilkan data
		var products []models.DestinationProduct
		db.Find(&products)
		fmt.Println("Total products:", len(products))
	}
}
