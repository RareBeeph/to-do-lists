package server

// TODO: pick a better name for the package

import (
	"io"
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
		HandlePOST(ctx)
	})

	server.Run() // TODO: this probably isn't ideal
}

func HandleGET(ctx *gin.Context) {
	// TODO: users would want more query types i bet
	queryId := ctx.Param("id")
	ctx.String(http.StatusOK, database.HandleQuery(queryId).CreatedAt.String()) // TODO: return different status on failure
}

func HandlePUT(ctx *gin.Context) {
	queryID := ctx.Param("id")
	body, _ := (io.ReadAll(ctx.Request.Body))
	response, _ := database.HandleUpdate(queryID, (string)(body)) // TODO: actually use the error return for logic
	ctx.String(http.StatusOK, response)
}

func HandleDELETE(ctx *gin.Context) {
	queryId := ctx.Param("id")
	response := database.HandleDelete(queryId)
	ctx.String(http.StatusOK, response) // TODO: return different status on failure
}

func HandlePOST(ctx *gin.Context) {
	entry := database.TodoEntry{}
	body, _ := io.ReadAll(ctx.Request.Body)
	entry.Body = (string)(body)
	database.HandleCreate(entry)
	// TODO: handle response
}
