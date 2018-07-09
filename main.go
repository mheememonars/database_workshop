package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type UserData struct{
    Id int 
    CitizenId string
    Firstname string
    Lastname string
    BirthYear int
    FirstnameFather string
    LastnameFather string
    FirstnameMother string
    LastnameMother string
    SoldierId int
    AddressId int
}


func main() {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testsck")

    if err != nil{
        fmt.Println("connect fail")
    }

    fmt.Println("connect success")

    defer db.Close()
    fmt.Println(add(db))
}


func read(db *sql.DB) []UserData{
    results,_ := db.Query("SELECT * FROM user")
    var userDataList []UserData
    for results.Next() {
        var userData UserData
        err := results.Scan(
            &userData.Id, 
            &userData.CitizenId,
            &userData.Firstname,
            &userData.Lastname,
            &userData.BirthYear,
            &userData.FirstnameFather,
            &userData.LastnameFather,
            &userData.FirstnameMother,
            &userData.LastnameMother,
            &userData.SoldierId,
            &userData.AddressId,
        )
        if err != nil {
            panic(err.Error()) 
        }
        userDataList = append(userDataList, userData)
    }
    return userDataList
}

func add(db *sql.DB) bool{
    results,_ := db.Prepare(`INSERT INTO user 
    (citizen_id,firstname,lastname,birthyear
    ,firstname_father,lastname_father,
    firstname_mother,lastname_mother,soldier_id
    ,address_id) 
    VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ? ) `)

    _, err := results.Exec("1552425252111","ชาติชาย2","เพ็ชรเม็ด","1985","สุชาติ","เพ็ชรเม็ด","สุณี","เพ็ชรเม็ด","8","1")
		if err != nil {
            panic(err.Error()) 
            return false
		}
        return true
}