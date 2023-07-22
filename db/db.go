package db

import (
	"leleshop/entity/product"
	users "leleshop/entity/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	connStr := "postgres://lelshop_user:FnDmS31yQOKUu11tBvnrEQUVTGSnjUBS@dpg-cisam618g3n42om0l36g-a.singapore-postgres.render.com/lelshop"
	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//	Connection success at this stage
	db.Debug().AutoMigrate(
		users.UserEntity{},
		product.ProductEntity{},
		product.ProductImageEntity{},
	)
}

func GetDB() *gorm.DB {
	return db
}
