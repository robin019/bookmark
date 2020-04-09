package route

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/robin019/bookmark/src/controller"
	"github.com/robin019/bookmark/src/middleware"
	"github.com/robin019/bookmark/src/utils/config"
)

var addr = fmt.Sprintf("%v:%v", config.Get("server.host"), config.Get("server.port"))

// Run maps the routing path and keeps listening for request
func Run() {
	app := iris.Default()

	// CORS
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CorsMiddleware)

	user := app.Party("user")
	{
		user.Get("/", controller.GetUsers)
		user.Post("/", controller.CreateUser)
	}

	app.Run(iris.Addr(addr))
}
