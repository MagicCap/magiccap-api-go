package main

import "github.com/labstack/echo/v4"

func (p *Platform) checkKey(key string, c echo.Context) (bool, error) {
	var user Admin
	result := p.db.Where(&Admin{AccessKey: key}).First(&user)

	if result.Error != nil {
		return false, nil
	}
	return true, nil
}
