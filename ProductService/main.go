// main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/product", getProduct)
	fmt.Println("ProductService is running on :8082")
	http.ListenAndServe(":8082", nil)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Product data")
}
