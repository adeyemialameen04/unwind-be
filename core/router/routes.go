package router

import (
	"net/http"
	"path/filepath"

	"github.com/adeyemialameen04/unwind-be/core/server"
	"github.com/adeyemialameen04/unwind-be/internal/projectpath"
	scalargo "github.com/bdpiprava/scalar-go"
	"github.com/gin-gonic/gin"
)

func SetupRouter(srv *server.Server) {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	RegisterBookRoutes(srv, router.Group("/books"))

	specUrl := filepath.Join(projectpath.Root, "/docs/swagger.json")

	router.GET("/spec-url", func(c *gin.Context) {
		content, err := scalargo.NewV2(
			scalargo.WithSpecURL("/docs/swagger.json"),
			scalargo.WithMetaDataOpts(
				scalargo.WithTitle("Unwind"),
			),
			scalargo.WithServers(scalargo.Server{
				URL:         "http://localhost:8080",
				Description: "Default Server",
			}),
			scalargo.WithLayout(scalargo.LayoutClassic),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, content)
	})

	router.GET("/docs/swagger.json", func(c *gin.Context) {
		c.File(specUrl)
	})

	srv.Router = router
}

// router.GET("/reference", func(c *gin.Context) {
// 	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
// 		SpecURL: apiDir,
// 		CustomOptions: scalar.CustomOptions{
// 			PageTitle: "Unwind Api",
// 		},
// 	})
// 	if err != nil {
// 		c.String(http.StatusInternalServerError, "Failed to generate API reference: %v", err)
// 		return
// 	}
//
// 	c.Header("Content-Type", "text/html; charset=utf-8")
// 	c.String(http.StatusOK, htmlContent)
// })
