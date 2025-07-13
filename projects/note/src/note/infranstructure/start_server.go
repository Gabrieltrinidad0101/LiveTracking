package infranstructure

import (
	"github.com/labstack/echo"
)

func Server() {
	e := echo.New()
	Router(e)
	e.Start("0.0.0.0:8000")
}
