package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
)

func CreateServer() {
	server := gin.Default()

	server.GET("/:user/:id/", func(ctx *gin.Context) {
		HandleGET(ctx)
	})

	server.PUT("/:user/:id/", func(ctx *gin.Context) {
		HandlePUT(ctx)
	})

	server.DELETE("/:user/:id/", func(ctx *gin.Context) {
		HandleDELETE(ctx)
	})

	server.POST("/:user/", func(ctx *gin.Context) {
		HandlePOST(ctx.Request)
	})

	server.Run()
}

func HandleGET(ctx *gin.Context) {
	// TODO: determine query from req
	queryId := ctx.Param("id")
	ctx.String(http.StatusOK, database.HandleQuery(queryId).CreatedAt.String())
}

func HandlePUT(ctx *gin.Context) {
	// TODO: determine query from req
	database.HandleUpdate()
}

func HandleDELETE(ctx *gin.Context) {
	// TODO: determine query from ctx
	database.HandleDelete()
}

func HandlePOST(req *http.Request) {
	var entry database.TodoEntry
	// TODO: determine entry contents from req
	database.HandleCreate(entry)
}
