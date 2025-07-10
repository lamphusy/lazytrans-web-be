package repository

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3" // hoặc driver khác nếu dùng DB khác
)

func NewDB(dataSource string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dataSource)
    if err != nil {
        return nil, err
    }
    return db, nil
}