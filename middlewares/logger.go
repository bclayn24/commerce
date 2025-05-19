package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[${time_custom}] ${method} ${uri} ${status} ${bytes_out}\n",
		CustomTimeFormat: "01/Jan/2006 15:04:05",
		Output:           os.Stdout,
	})
}
