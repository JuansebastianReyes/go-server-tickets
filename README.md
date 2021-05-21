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
    go get -u github.com/jinzhu/gorm
    ```

* Driver MySql de `jinzhu/gorm`

    ```sh
    go get -u github.com/jinzhu/gorm/dialects/mysql
    ```

Ya teniendo las depencencias instaladas podemos pasar a la base de datos 

## Base de Datos 

* phpMyAdmin 
    
    Para ese caso no necesitamos ejecutar ningún comando solo se necesita tener corriendo **phpMyAdmin** bien sea con `wamp`, `xampp` o cualquier otro servidor que nos despliegue esta herramienta.

    Teniendo creada la base de datos en **phpMyAdmin** lo unico que queda modificar es la conexión a la base de datos haciendolo de esta manera dentro del archivo [connections.go](https://github.com/JuansebastianReyes/go-server-tickets/blob/master/go-server-mysql/commons/connections.go)

    ```go
	Driver := "mysql"        //Driver para poder conectarnos a la BD
	Usser := "root"          //Usuarion con el cual nos logeamos a la BD
	Password := "xxxxx" //Contraseña de acceso a la base de datos del usuario
	Database := "prueba" //Nombre de la base de datos donde vamos a alojar los datos

	db, err := gorm.Open(Driver, Usser+":"+Password+"@/"+Database+"?parseTime=true")
    ```


* Docker 

    El servidor viene cargargado por defecto con las credenciales para la conexión con la imagen de la base de datos en **Docker**, de este modo ahora explicaremos como desplegar esta base de datos dokcerizada. 

    Estando ubicados dentro de la carpeta `docker` en consola colocamos la sigiente linea:

    ```sh
    docker-compose up -d 
    ```

    para comprobar que se creo correctamente se puede usar la linea:

    ```sh
    docker-compose ps
    ```

## Ejecución Servidor 

Teniendo en cuenta que las dependencias ya han sido instaladas y la base de datos ya está corriendo para poner en marcha el servidor nos volvemos a ubicar la carpeta `go-server-mysql` y ejecutamos la línea:

```sh
   go run mian
```

## Ejemplos 

Para los ejemplos se puede usar una herramienta como `postman` para pasarle las direcciones de las peticiones que le porporcionamos al servidor.

* Crear
    
    para crear un registro usamos la url `http://localhost:9000/tickets/api/save` con el metodo `POST` y le ingresamos el siguente `json`: 

    ```json 
    {
        "usuario": "Sebastian Reyes",
        "estado": true
    }
    ```
* Editar 

    Para editar denuevo en la misma direccion `http://localhost:9000/tickets/api/save` con el metodo `POST`, pero le agregamos el id del elemento que queremos editar 

    ```json 
    {
        "id": 1,
        "usuario": "Sebastian Reyes",
        "estado": false
    }
    ```

* Listar 

    Para listar los tickets usamos la direccion `http://localhost:9000/tickets/api/all` con el metodo `GET`.

    Para listar un unico ticket usamos la direccion `http://localhost:9000/tickets/api/fing/{id}` con el metodo `GET`, `{id}` es el id del elemento a buscar.


* Eliminar 

    Para listar los tickets usamos la direccion `http://localhost:9000/tickets/api/all` con el metodo `POST`.



## Atutor

Juan Sebastian Reyes Garcia  
Estudiante Ingenieria de Sistemas  
Univerisidad Distrital Francisco Jose de Caldas 