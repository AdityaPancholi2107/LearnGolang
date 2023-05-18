package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("Hello"))
})
fmt.Println("Server started on http://localhost:8080")
http.ListenAndServe(":8080", nil)
}


// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// fmt.Fprintf(w, "Hello, World!")
	// })

	// fmt.Println("Server started on http://localhost:8080")
	// http.ListenAndServe(":8080", nil)
// }