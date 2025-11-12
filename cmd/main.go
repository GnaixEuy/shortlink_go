package main

import (
	"fmt"
	"log"
	"shortLink/internal/api"
	"shortLink/internal/config"
	"shortLink/internal/repo"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "shortLink/docs"
)

// @title ShortLink Service API
// @version 1.0
// @description A simple short link service built with Go + Gin + MySQL + Redis
// @host localhost:8080
// @BasePath /
func main() {
	if err := config.Load("internal/config/config.yaml"); err != nil {
		log.Fatalf("load config error: %v", err)
	}

	if repo.InitMySQL() != nil {
		log.Fatal("mysql init error")
	}

	if err := repo.InitRedis(); err != nil {
		log.Fatal("redis init error")
	}

	g := gin.Default()
	api.RegisterShortLinkRoutes(g)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addr := fmt.Sprintf(":%d", config.C.Server.Port)
	log.Printf("server listening on %s\n", addr)
	if err := g.Run(addr); err != nil {
		log.Fatal(err)
	}
}
