// core/handlers/books/book.go
package books

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/adeyemialameen04/unwind-be/core/server"
	"github.com/adeyemialameen04/unwind-be/internal/db/repository"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

type Handler struct {
	srv *server.Server
}

func NewBookHandler(srv *server.Server) *Handler {
	return &Handler{srv: srv}
}

// GetBooks godoc
// @Summary      Get all books
// @Description  Get a list of all books
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {array}   repository.Book
// @Router       /books [get]
func (book *Handler) GetBooks(c *gin.Context) {
	ctx := context.Background()
	tx, err := book.srv.DB.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}
	defer tx.Rollback(ctx)

	repo := repository.New(tx)
	books, err := repo.FindAllBooks(ctx)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	if books != nil {
		c.JSON(http.StatusOK, books)
	}
}

// CreateBook godoc
// @Summary      Create a new book
// @Description  Insert a new book into the database
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      repository.InsertBookParams  true  "Book data"
// @Success      201   {object}  repository.Book             "Book created successfully"
// @Failure      400   {object}  map[string]string            "Invalid request data"
// @Failure      500   {object}  map[string]string            "Failed to start transaction or insert book"
// @Router       /books [post]

func (h *Handler) CreateBook(c *gin.Context) {
	g := galidator.New().CustomMessages(galidator.Messages{
		"required": "$field is required",
	})
	customizer := g.Validator(repository.InsertBookParams{})

	ctx := context.Background()

	var req repository.InsertBookParams

	// Bind the incoming JSON to the InsertBookParams struct
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotAcceptable, gin.H{"detail": customizer.DecryptErrors(err), "err": err.Error()})
		return
	}

	// Start a transaction
	tx, err := h.srv.DB.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}
	defer tx.Rollback(ctx)

	// Insert book using the repository
	repo := repository.New(tx)
	book, err := repo.InsertBook(ctx, req)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert book"})
		return
	}

	// Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "data": book})
}
