package watchlist

import (
	"context"
	"errors"

	"github.com/adeyemialameen04/unwind-be/core/server"
	"github.com/adeyemialameen04/unwind-be/internal/db/repository"
	"github.com/adeyemialameen04/unwind-be/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
)

type Handler struct {
	srv *server.Server
}

func NewWatchListHandler(srv *server.Server) *Handler {
	if srv == nil {
		panic("server cannot be nil")
	}
	return &Handler{srv: srv}
}

// GetWatchList godoc
//
//	@Summary		Get User WatchList
//	@Description	Retrieves a user watch list.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		AccessTokenBearer
//	@Success		200		{object}	server.SuccessResponse{data=[]repository.WatchList}
//	@Failure		401		{object}	server.UnauthorizedResponse			"Unauthorized"
//	@Failure		404		{object}	server.NotFoundResponse				"Profile not found"
//	@Failure		500		{object}	server.InternalServerErrorResponse	"Internal server error"
//	@Router			/user/watch-list [get]
func (h *Handler) GetWatchList(c *gin.Context) {
	ctx := context.Background()
	tx, err := h.srv.DB.Begin(ctx)
	if err != nil {
		server.SendInternalServerError(c, err, server.WithMessage("Error beginning transaction"))
		return
	}
	defer tx.Rollback(ctx)

	userId, err := domain.GetUserIDFromContext(c)
	if err != nil {
		return
	}

	repo := repository.New(tx)
	watchList, err := repo.GetUserWatchList(ctx, userId)
	if err != nil {
		server.SendInternalServerError(c, err, server.WithMessage("Failed to retrieve watch list"))
		return
	}

	if err := tx.Commit(ctx); err != nil {
		server.SendInternalServerError(c, err, server.WithMessage("Failed to commit transaction"))
		return
	}

	if watchList == nil {
		server.SendNotFound(c, errors.New("You don't have anything in your watch list"))
		return
	}

	server.SendSuccess(c, watchList, server.WithMessage("WatchList retrieved successfully"))
}

// AddToList godoc
//
//	@Summary		Adds an anime to a user watch list
//	@Description	Retrieves a user watch list.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		AccessTokenBearer
//
// @Param			Anime	body		repository.AddToListParams	true	"Anime Data"
//
//	@Success		200		{object}	server.SuccessResponse{data=repository.WatchList}
//	@Failure		401		{object}	server.UnauthorizedResponse			"Unauthorized"
//	@Failure		404		{object}	server.NotFoundResponse				""
//	@Failure		500		{object}	server.InternalServerErrorResponse	"Internal server error"
//	@Router			/user/watch-list [post]
func (h *Handler) AddToList(c *gin.Context) {
	var (
		ctx = context.Background()
		req repository.AddToListParams
	)

	if err := h.validateAddtoListRequest(c, &req); err != nil {
		return
	}

	userId, err := domain.GetUserIDFromContext(c)
	if err != nil {
		return
	}

	tx, err := h.srv.DB.Begin(ctx)
	if err != nil {
		server.SendInternalServerError(c, err, server.WithMessage("Error beginning transaction"))
		return
	}
	defer tx.Rollback(ctx)

	req.UserID = userId
	repo := repository.New(tx)
	watchList, err := repo.AddToList(ctx, req)
	if err != nil {
		server.SendInternalServerError(c, err, server.WithMessage("Fail to Add to watch list"))
		return
	}

	if err := tx.Commit(ctx); err != nil {
		server.SendInternalServerError(c, err, server.WithMessage("Failed to commit transaction"))
		return
	}

	server.SendSuccess(c, watchList, server.WithMessage("Added to watchList successfully"))
}

func (h *Handler) validateAddtoListRequest(c *gin.Context, req *repository.AddToListParams) error {
	g := galidator.New().CustomMessages(galidator.Messages{
		"required": "$field is required",
		"min":      `$field can't be less than {min}`,
	})
	customizer := g.Validator(repository.AddToListParams{})

	if err := c.ShouldBindJSON(req); err != nil {
		server.SendValidationError(c, customizer.DecryptErrors(err))
		return err
	}
	return nil
}
