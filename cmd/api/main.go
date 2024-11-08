// @title           Unwind API
// @version         1.0
// @description     API for Unwind
// @contact.name    Al-Ameen Adeyemi
// @contact.email   adeyemialameen04@gmail.com
// @contact.url     https://github.com/adeyemialameen04
// @host            localhost:8080
// @BasePath        /

// @securityDefinitions.apikey AccessTokenBearer
// @in                           header
// @name                         Authorization
// @description                  AccessTokenBearer Authentication

// @securityDefinitions.apikey RefreshTokenBearer
// @in                           header
// @name                         Authorization
// @description                  RefreshTokenBearer Authentication

// @tag.name       Auth
// @tag.description Authentication endpoints
package main

import (
	"context"
	"log"

	"github.com/adeyemialameen04/unwind-be/core/router"
	"github.com/adeyemialameen04/unwind-be/core/server"
	"github.com/adeyemialameen04/unwind-be/internal/config"
	"github.com/adeyemialameen04/unwind-be/internal/projectpath"
	"github.com/jackc/pgx/v5"
)

func main() {
	cfg, err := config.Load(projectpath.Root)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	conn, err := pgx.Connect(context.Background(), cfg.DbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer conn.Close(context.Background())

	srv, err := server.NewServer(cfg, conn)
	if err != nil {
		log.Fatal(err)
	}

	router.SetupRouter(srv)
	server.RunServer(srv)
}
