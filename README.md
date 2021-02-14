## Users API

To Run the Project:
* Go lang installed
* Put the project in 
```
$GOPATH/src/github.com/abmomin/evatix_users_api

```
### Dependecies: for details please refet to this page https://github.com/gin-gonic/gin
```
go get github.com/gin-gonic/gin
```

### Database: Used MySql Database, Please install mysql driver for Golang.

```
go get github.com/go-sql-driver/mysql
```
### Database configuaration: Lets set some environment variables.
please configure as your datasource.
```
mysqlUsersUserName=//user name
mysqlUsersPassword=//password
mysqlUsersHost= //localhost
mysqlUsersSchema=//users
```
Now with all that done please create a users table in users database.
```
create table users(userId int primarykey, firstName varchar(20), lastName varchar(20), email varchar(20), dateCreated varchar(20), status tinyint, password varchar(50)); 
```
After all done, Please go to project root directory and run with golang commad 
```
go run main.go
```
