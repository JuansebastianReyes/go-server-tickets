# Servidor Tickets Golang

Api creada con el lenguaje de programación **go** usando las librerias [`gorilla/mux`](https://github.com/gorilla/mux) y [`jinzhu/gorm`](https://github.com/jinzhu/gorm) la cual se conecta a una base de datos **MySql**, este se encarga de crear, eliminar, editar y recuperar tickets. Además se crea una imagen de la base de datos en **Docker** para poder realizar pruebas sin necesidad de contar con herramientas como **phpMyAdmin**.


## Como usar este servidor 


Para poder usar este servidor de forma local lo primero que vamos a hacer es clonar este repositorio usado la siguiente línea en consola:

```sh
git clone https://github.com/JuansebastianReyes/go-server-tickets.git
```

Una vez colonado el repositorio pasamos a intalar las dependencias necesarias para que el servidor funcionne, estando ubicados en consola en la carperta `go-server-mysql`

* Libreria `gorilla/mux`

```sh
go get -u github.com/gorilla/mux
```

* Libreria `jinzhu/gorm`

```sh
go get -u  github.com/jinzhu/gorm
```

* Driver MySql 

```sh
go get -u  github.com/jinzhu/gorm/dialects/mysql
```

Ya teniendo las depencencias instaladas podemos pasar a la base de datos 

#Base de Datos 