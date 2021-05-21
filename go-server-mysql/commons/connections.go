package commons

import (
	"log"
	"main/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnetion() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/tickets?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnetion()
	defer db.Close()

	log.Println("Migrando....")

	db.AutoMigrate(&models.Ticket{})
}
