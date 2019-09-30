package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (p *Platform) addRelease(ctx echo.Context) error {
	rp := new(ReleasePublish)
	if err := ctx.Bind(rp); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
	}
}
