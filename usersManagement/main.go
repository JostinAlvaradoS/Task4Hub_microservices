package main

import (
	"task.com/usersManagement/firebase"
	"task.com/usersManagement/router"
)

func main() {
	firebase.InitFirebase() // Initialize Firebase for local testing

	// Use the new InitRouter function for local tests
	router.InitRouter()

}

// For Cloud Functions deployment, you would have a separate function:

// package main
// import (
//    "context"
//    "net/http"
//
//   "task.com/usersManagement/router" // Import the router package
//)

// Entry point for Cloud Functions.
//func UsersManagement(w http.ResponseWriter, r *http.Request) {
//   router.NewHTTPHandler().ServeHTTP(w,r)
//}
