package config

import (
	"fmt"
	"time"

	"github.com/ShiranaiZo/api-contact-form-v2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	dbUser := GetEnv("DB_USER", "user")
	dbPassword := GetEnv("DB_PASSWORD", "password")
	dbHost := GetEnv("DB_HOST", "db")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "contactsdb")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Print(dsn)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// Singular table biar tabel nya dari users jadi user, jadi tidak dijadikan plural
			SingularTable: true,
		},
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database instance!")
	}

	// Wajib, lebih baik di aksih ke env untuk value valuenya
	sqlDB.SetMaxOpenConns(10) //Maksimal koneksi ytang boleh dibuka. misal api dapet 100 request bersaaan, nah 10 request jalan dulu, 90 nya antri
	//idle itu nganggur tapi siap digunakan. Koneksi yang disimpan saat idle (nganggur). misal 10 koneksi aktif di max opencons, nah 5 disimpen idle, 5 nya lagi ditutup.
	sqlDB.SetMaxIdleConns(5)            // lebih jelas nya ketika ada kasir 10, nah karna sepi, yg active ngelayani pelanggan 2, yang standby dan tidak ada pelanggan/idle yang disimpan ada 5, lalu kasir yang pulang/ditutup 3 koneksi ()sisanya
	sqlDB.SetConnMaxLifetime(time.Hour) //Umur maksimal koneksi. misal dibuat jam 10 pagi, lifetimenya 1 jam. maka pas jam 11 koneksi ditutup, terus koneksi baru dibuat lagi. takutnya si database memutus koneksi diam diam. makanya dengan ini kaya ngerestart gitu

	err = DB.AutoMigrate(&models.Contact{})
	if err != nil {
		panic(fmt.Sprintf("AutoMIgrate failed: %v", err))
	}
}
