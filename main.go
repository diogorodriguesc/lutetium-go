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

func main() {
    dsn := "user:password@tcp(127.0.0.1:3306)/lutetium"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Query to select data
    rows, err := db.Query("SELECT product_id, product_name FROM products")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Iterate over the rows
    for rows.Next() {
        var productId int
        var productName string

        if err := rows.Scan(&productId, &productName); err != nil {
            log.Fatal(err)
        }

        product := Product{ProductId: productId, ProductName: productName}

        fmt.Printf("Id: %d Name: %s\n", product.ProductId, product.ProductName)
    }

    // Check for errors after iterating over rows
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}
