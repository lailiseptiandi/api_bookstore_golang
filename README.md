# api_bookstore_golang
Fundamental API CRUD Golang + gorm mysql

Configuration Database :
Open file pkg/config/app.go and search syntax :

gorm.Open("mysql", "root:123qwe@/api_golang_bookstore?charset=utf8mb4&parseTime=True&loc=Local")

1. set root => name user db
2. set 123qwe => pass user db
3. api_golang_bookstore  => database name

# finaly
open directory : 
1. cd cmd
2. go run main.go
3. open url localhost:9000

# endpoint
Open POSTMAN etc... test API

1. GET ALL DATA(GET)  => http://localhost:9000/book
2. GET DATA BY ID(GET)  => http://localhost:9000/book/{id}
3. CREATE DATA(POST) => http://localhost:9000/book
4. UPDATE DATA (PUT)  => http://localhost:9000/book/{id}
5. DELETE DATA(DELETE)  => http://localhost:9000/book/{id}

Thank you. Don't Forget Follow
