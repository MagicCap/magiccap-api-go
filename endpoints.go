package main

import (
	"net/http"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/labstack/echo/v4"
)

func (p *Platform) checkUpdate(ctx echo.Context) error {
	c := new(Check)
	if err := ctx.Bind(c); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bad request", "error": err.Error()})
	}
	var r Release
	p.db.Where("type = ?", strings.ToLower(string(c.Channel))).Last(&r)

	latestVersion, err := semver.NewVersion(r.Version)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "something went wrong"})
	}
	requestVersion, err := semver.NewVersion(c.Version)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "not a valid version"})
	}

	checkResp := CheckResponse{
		Update: false,
	}

	if latestVersion.GreaterThan(requestVersion) {
		checkResp.Update = true
		checkResp.Release = r
	}

	return ctx.JSON(http.StatusOK, checkResp)
}
