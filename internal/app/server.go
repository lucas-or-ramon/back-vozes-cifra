package app

import (
	"back-vozes-cifras/internal/app/adapters/database/mongodb"
	"back-vozes-cifras/internal/app/adapters/handler/apphttp"
	"back-vozes-cifras/internal/app/adapters/middleware"
	"back-vozes-cifras/internal/app/application/usecase"
	"back-vozes-cifras/internal/app/config"
	"github.com/spf13/viper"
	"net/http"
)

func StartServer() {
	mongoClient := config.ConnectToMongoDB()
	mongoRepo := mongodb.NewMongoRepository(mongoClient)
	useCase := usecase.NewUseCase(mongoRepo)
	songHandler := apphttp.NewHandler(useCase)

	http.Handle("/songs", middleware.VerifyApiKeyMiddleware(songHandler.GetSongByIDHandler()))

	port := viper.GetString("HTTP_PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
