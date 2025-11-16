package api

import (
	"database/sql"
	"net/http"

	db "github.com/Roditu/BE_RS_TEST/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPersonRequest struct {
	Name    string `json:"name" binding:"required"`
	Ambition string `json:"ambition"`
}

func (server *Server) createPerson(ctx *gin.Context) {
	var req createPersonRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorRespone(err))
		return
	}

	_, err := server.store.GetPersonByName(ctx, req.Name)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "person with this name already exists"})
		return
	}
	
	arg := db.CreatePersonParams{
		Name: req.Name,
		Ambition: sql.NullString{
			String: req.Ambition,
			Valid:  req.Ambition != "",
		},
	}

	person, err := server.store.CreatePerson(ctx, arg)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	ctx.JSON(http.StatusOK, person)
}

type getPersonRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type personResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Ambition string `json:"ambition,omitempty"` // omit if empty
}

func (server *Server) getPerson(ctx *gin.Context) {
	var req getPersonRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorRespone(err))
		return
	}

	person, err := server.store.GetPerson(ctx, req.ID)

	resp := personResponse{
		ID:   person.ID,
		Name: person.Name,
	}

	if person.Ambition.Valid {
		resp.Ambition = person.Ambition.String
	}

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorRespone(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}