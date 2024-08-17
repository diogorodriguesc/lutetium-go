package main

import (
    "fmt"
    "log"
    "net/http"
    "database/sql"

    _ "github.com/go-sql-driver/mysql"
)

type Product struct {
    ProductId int
    ProductName string
}

func dbConnect(dsn string) (*sql.DB) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

	return db
}

func dbQuery(db *sql.DB, queryString string) (*sql.Rows) {
	rows, err := db.Query(queryString)
	if err != nil {
        log.Fatal(err)
    }

	return rows
}

func main() {
	db := dbConnect("user:password@tcp(127.0.0.1:3306)/lutetium")
	defer db.Close()

	rows := dbQuery(db, "SELECT product_id, product_name FROM products")
    defer rows.Close()

    for rows.Next() {
        var productId int
        var productName string

        if err := rows.Scan(&productId, &productName); err != nil {
            log.Fatal(err)
        }

        product := Product{ProductId: productId, ProductName: productName}

        fmt.Printf("Id: %d Name: %s\n", product.ProductId, product.ProductName)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
