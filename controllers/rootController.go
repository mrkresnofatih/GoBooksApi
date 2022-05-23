package controllers

import (
	gin "github.com/gin-gonic/gin"
)

func AddControllers(engine *gin.Engine) {
	routing := engine.Group("/api/v1")
	AddBookEndpoints(routing)
}
