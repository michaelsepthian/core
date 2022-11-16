package models

import "gorm.io/gorm"

var ModelLists []interface{} = []interface{}{
	&Users{},
	&Products{},
	&Orders{},
	&Transactions{},
	&Payments{},
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelLists...)
}
