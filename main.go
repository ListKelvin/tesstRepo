package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/ListKelvin/book-store/pkg/middleware"
	"github.com/ListKelvin/book-store/pkg/routes"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

// func init() {
// 	config.Connect()

// }
// init is a function that is called automatically when the program starts
// it loads the environment variables from the .env file
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		log.Print("sad .env file found")
	}
}
var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func main(){
	var err error
	r := mux.NewRouter()
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting values from .env "+ os.Getenv("GOOGLE_REDIRECT_URL"))
	}
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	// r.Use(mux.CORSMethodMiddleware(r))
	// r.Use(LoggingMiddleware)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
	
}


