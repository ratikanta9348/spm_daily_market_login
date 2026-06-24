package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
type User struct {
	Password string
	Role     string
}

var users = map[string]User{
	"user1":  {Password: "pass123", Role: "user"},
	"admin":  {Password: "admin123", Role: "admin"},
	"admin2": {Password: "admin456", Role: "admin"},
}
var templates = template.Must(template.ParseGlob("templates/*.html"))

func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	err := templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		http.Error(w, "Unable to load page", http.StatusInternalServerError)
		log.Println(err)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.FormValue("userid")
	password := r.FormValue("password")

	w.Header().Set("Content-Type", "application/json")

	// Demo credentials
	if user, ok := users[userID]; ok && user.Password == password {

		resp := LoginResponse{
			Success: true,
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)

	resp := LoginResponse{
		Success: false,
		Message: "Invalid User ID or Password",
	}

	json.NewEncoder(w).Encode(resp)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application Running"))
}

func main() {

	// Static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/health", health)

	log.Println("Server Started")
	log.Println("URL: http://localhost:8080/login")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
