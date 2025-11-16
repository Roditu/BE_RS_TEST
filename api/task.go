package api

import (
	"net/http"
	"strconv"

	db "github.com/Roditu/BE_RS_TEST/db/sqlc"

	"github.com/gin-gonic/gin"
)

type addTaskRequest struct {
	Todo  string `json:"todo" binding:"required"`
	Exp 	int64  `json:"exp" binding:"required"`
}

func (server *Server) AddTask(ctx *gin.Context) {
	var req addTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorRespone(err))
		return
	}

	userID := ctx.MustGet("user_id").(int64)

	task, err := server.store.CreateTask(ctx, db.CreateTaskParams{
		Todo:   req.Todo,
		Exp:  req.Exp,
		UserID: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (server *Server) FinishTask(ctx *gin.Context) {
	taskIDParam := ctx.Param("id")
	taskID64, err := strconv.ParseInt(taskIDParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorRespone(err))
		return
	}

	userID := ctx.MustGet("user_id").(int64)

	task, err := server.store.GetTaskByUserId(ctx, taskID64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorRespone(err))
		return
	}

	if task.UserID != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Mark task as done
	task, err = server.store.UpdateTaskStatus(ctx, db.UpdateTaskStatusParams{
		TaskID 	: taskID64,
		Status	: "COMPLETE",
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	// Increase user points
	_, err = server.store.AddUserExp(ctx, db.AddUserExpParams{
		UserID	: userID,
		Exp			: task.Exp,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":     "task completed",
		"earnedPoint": task.Exp,
	})
}

func (server *Server) ListTasks(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(int64)
	tasks, err := server.store.ListTasksByUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorRespone(err))
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}