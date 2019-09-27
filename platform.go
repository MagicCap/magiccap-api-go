package main

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// Platform represents the API platform
type Platform struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}
