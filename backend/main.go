package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"employee-api/config"
	"employee-api/handlers"
	"employee-api/routes"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}

	db := config.ConnectDatabase()
	defer db.Close()

	mux := http.NewServeMux()

	employeeHandler := handlers.NewEmployeeHandler(db)

	routes.RegisterRoutes(
		mux,
		employeeHandler,
	)

	// Enable CORS
	handler := enableCORS(mux)

	log.Println("Go API running on :8080")

	err = http.ListenAndServe(
		":8080",
		handler,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func enableCORS(next http.Handler) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		// Allow React development server
		w.Header().Set(
			"Access-Control-Allow-Origin",
			"http://localhost:5173",
		)

		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)

		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type",
		)

		// Handle browser preflight request
		if r.Method == http.MethodOptions {

			w.WriteHeader(http.StatusNoContent)

			return
		}

		next.ServeHTTP(w, r)
	})
}