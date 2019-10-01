package config

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var screatkey = []byte("decandTampan")

func DBInit() *gorm.DB {

	//db, err := gorm.Open("postgres", "host=dna-developer.com port=5432 user=postgres dbname=warnadev password=nanda")
	db, err := gorm.Open("postgres", "host=192.168.133.81 port=5432 user=postgres dbname=warnaprod password=nandaprod")
	//db, err := gorm.Open("postgres", "host=192.168.133.89 port=5432 user=postgres dbname=warnadev password=nanda")
	if err != nil {
		// panic("Gagal koneksi ke database")
		panic(err.Error())
	}

	return db
}

func JwtKey() []byte {

	return screatkey
}
