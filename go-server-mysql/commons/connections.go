package commons

import (
	"log"
	"main/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnetion() *gorm.DB {

	//Variables de Entorno para coneccion con Docker
	// si se usa phpmyadmin o otra tipo de distribucion de mysql cambiar el password y el Database
	Driver := "mysql"        //Driver para poder conectarnos a la BD
	Usser := "root"          //Usuarion con el cual nos logeamos a la BD
	Password := "tickets123" //Contraseña de accesp a la base de datos del usuario
	Database := "db_tickets" //Nombre de la base de datos donde vamos a alojar los datos

	db, err := gorm.Open(Driver, Usser+":"+Password+"@/"+Database+"?parseTime=true")

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
