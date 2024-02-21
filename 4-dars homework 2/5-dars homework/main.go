package main

import (
	"database/sql"
	"fmt"

	// "fmt"
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID         int
	Name       string
	Price      float64
	CategoryID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	// Open a connection to the database
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=shahzod password=1 database=shahzod sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Insert a new category
	// fmt.Println("Input category name: ")
	// var categoryName string
	// fmt.Scan(&categoryName)
	// categoryID, err := insertCategory(db, categoryName)
	// if err != nil {
	//  panic(err)
	// }

	// Insert a new product
	fmt.Println("input productName: ")
	var productName string
	fmt.Scan(&productName)
	fmt.Println("input productPrice: ")
	var price float64
	fmt.Scan(&price)
	fmt.Println("input categoryId: ")
	var categoryId int
	fmt.Scan(&categoryId)

	createdAt := time.Now()
	updatedAt := time.Now()
	err = insertProduct(db, productName, price, categoryId, createdAt, updatedAt)
	if err != nil {
		panic(err)
	}

	// Update the product
	// newPrice := 1700.00
	// updatedAt = time.Now()
	// err = updateProduct(db, productName, newPrice, updatedAt)
	// if err != nil {
	//  panic(err)
	// }

	// Get a list of products
	// productList, err := getProductList(db)
	// if err != nil {
	//     panic(err)
	// }
	// fmt.Println("Product List:")
	// for _, product := range productList {
	//     fmt.Printf("ID: %d, Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)
	// }

	// Get a product by ID
	fmt.Println("input id you want to delete")
	var d int
	fmt.Scan(&d)
	product, err := getProductByID(db, d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product found by ID %d: Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)

	// Delete a product

	err = deleteProduct(db, d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product deleted successfully")
}

func insertCategory(db *sql.DB, name string) (int, error) {
	var categoryID int
	err := db.QueryRow("INSERT INTO Category (Name) VALUES ($1) RETURNING Id", name).Scan(&categoryID)
	if err != nil {
		return 0, err
	}
	return categoryID, nil
}

func insertProduct(db *sql.DB, name string, price float64, categoryID int, createdAt, updatedAt time.Time) error {
	_, err := db.Exec(`INSERT INTO Product 
 (Name, Price, Category_id, Created_at, Updated_at) 
 VALUES ($1, $2, $3, $4, $5)`, name, price, categoryID, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, name string, price float64, updatedAt time.Time) error {
	_, err := db.Exec("UPDATE Product SET Price = $1, Updated_at = $2 WHERE Name = $3", price, updatedAt, name)
	if err != nil {
		return err
	}
	return nil
}

func getProductList(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * FROM Product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productList []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		productList = append(productList, product)
	}
	return productList, nil
}

func getProductByID(db *sql.DB, id int) (Product, error) {
	var product Product
	err := db.QueryRow("SELECT * FROM Product WHERE Id = $1", id).Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func deleteProduct(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM Product WHERE Id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
