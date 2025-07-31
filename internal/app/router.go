package app

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rslbn/blog/internal/app/handler"
	"github.com/rslbn/blog/internal/app/service"
	"github.com/rslbn/blog/internal/app/web"
	db "github.com/rslbn/blog/postgres"
)

func NewRouter(dbConn *pgxpool.Pool) http.Handler {
	mux := http.NewServeMux()

	sqlcQueries := db.New(dbConn)

	jwtService := service.NewJwtService()

	userService := service.NewUserService(sqlcQueries)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userService, jwtService)
	authHandler := handler.NewAuthHandler(authService)

	mux.HandleFunc("GET /users", web.HandlerAdapter(userHandler.GetAll))
	mux.HandleFunc("GET /users/{username}", web.HandlerAdapter(userHandler.GetUserByUsername))
	mux.HandleFunc("POST /auth/login", web.HandlerAdapter(authHandler.Login))
	mux.HandleFunc("POST /auth/register", web.HandlerAdapter(userHandler.Register))
	return mux
}
