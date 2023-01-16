package health

import (
	"github.com/labstack/echo"
	"net/http"
)

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "OK-Email api status")
}
