package main

import (
	controllers "mrkresnofatih/gobookapi/controllers"

	gin "github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	controllers.AddControllers(app)
	app.Run(":3000")

}
