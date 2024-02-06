// main.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", getUser)
	fmt.Println("UserService is running on :8081")
	http.ListenAndServe(":8081", nil)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User data")
}
