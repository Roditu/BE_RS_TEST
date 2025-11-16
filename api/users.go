package api

import (
	"net/http"
	"time"

	db "github.com/Roditu/BE_RS_TEST/db/sqlc"
	"github.com/Roditu/BE_RS_TEST/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorRespone(err))
		return
	}

	_, err := server.store.GetUserByUsername(ctx, req.Username)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "user with this name already exists"})
		return
	}

	hash, err := util.HashPassword(req.Password)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorRespone(err))
        return
    }
	
	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hash,
	}

	user, err := server.store.CreateUser(ctx, arg)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, errorRespone(err))
        return
    }

    user, err := server.store.GetUserByUsername(ctx, req.Username)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, errorRespone(err))
        return
    }

    if err := util.CheckPasswordHash(req.Password, user.Password); err != nil {
        ctx.JSON(http.StatusUnauthorized, errorRespone(err))
        return
    }

    token, err := server.tokenMaker.CreateToken(int32(user.UserID), 24*time.Hour)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, errorRespone(err))
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "access_token": token,
        "user":         user,
    })
}

func (server *Server) getUserByUserId(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(int64)

	user, err := server.store.GetUserByUserId(ctx, userID)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}