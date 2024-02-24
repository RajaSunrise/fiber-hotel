package database

import (
    "fmt"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
    const DATA = "root:indra@tcp(127.0.0.1:3306)/hotel?charset=utf8mb4&parseTime=True&loc=Local"
    DSN := DATA

    DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
    if err != nil {
        panic("gagal terhubung ke database")
    }
    fmt.Println("Berhasil connect database")
}
