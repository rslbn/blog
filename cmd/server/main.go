package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rslbn/blog/internal/app"
	"github.com/rslbn/blog/pkg/database"
)

func main() {
	ctx := context.Background()
	dbConn, err := database.NewPostgresDB(ctx, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("failed to load the database: %v", err)
	}
	defer dbConn.Close()
	fmt.Println("Database pool connection is established.")
	router := app.NewRouter(dbConn)
	log.Fatal(http.ListenAndServe(":8080", router))
}
