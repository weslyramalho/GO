package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/jwtauth"

	"github.com/go-chi/chi/middleware"
	"github.com/weslyramalho/GO/tree/main/API/configs"
	"github.com/weslyramalho/GO/tree/main/API/docs"

	"github.com/weslyramalho/GO/tree/main/API/internal/entity"
	"github.com/weslyramalho/GO/tree/main/API/internal/infra/database"
	"github.com/weslyramalho/GO/tree/main/API/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Primeira API em Go
// @version 1.0
// @description Product API
// @termsOfService http://swagger.io/terms

// @contact.name Wesly Ramalho

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, configs.TokenAuth, configs.JwtExperiesIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Route) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.Create)
	r.Post("/users/login", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}
