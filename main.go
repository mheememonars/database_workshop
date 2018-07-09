package main

import (
    "fmt"
    "sck/database_workshop/users"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testsck")

    if err != nil{
        fmt.Println("connect fail")
    }

    fmt.Println("connect success")

    defer db.Close()
    users := users.Read(db)
   fmt.Println(users)
//     fmt.Println(readByCitizenId(db,"1552425252111"))
//    fmt.Println(read(db))
}


