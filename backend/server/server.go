package server

// TODO: pick a better name for the package

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/database"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	todoGroup := router.Group("/todo")

	todoGroup.GET("", HandleGETall)
	todoGroup.GET("/:id", HandleGET)
	todoGroup.PUT("/:id", HandlePUT)
	todoGroup.DELETE("/:id", HandleDELETE)
	todoGroup.POST("", HandlePOST)

	return router
}

func HandleGETall(ctx *gin.Context) {
	// TODO: might deserve to be a part of HandleGET
	a := database.HandleQueryAll()
	output := ""
	for _, b := range a {
		output += b.Body + "\n"
	}
	ctx.String(http.StatusOK, output)
}

func HandleGET(ctx *gin.Context) {
	// TODO: users would want more query types i bet
	queryId := ctx.Param("id")
	ctx.String(http.StatusOK, database.HandleQuery(queryId).Body) // TODO: return different status on failure
	//TODO: maybe don't just return the body
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
	ctx.String(http.StatusOK, "Something happened") // TODO: handle response
}
