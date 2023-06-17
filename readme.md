#### Initialize project
```bash
go mod init github.com/username/repo
```

#### Add dependencies
```bash
go get github.com/gorilla/mux
github.com/joho/godotenv 
github.com/jinzhu/gorm
github.com/jinzhu/gorm/dialects/mysql
```

#### Create .env file
```bash
DB_HOST=********
DB_NAME=********
DB_PASSWORD=********
DB_USER=********
```

#### Create database
```bash
mysql -u root -p
create database db_name;
```

#### Download dependencies
```bash
go mod tidy
```

##### start server
```bash
go run cmd/main/main.go 
```

