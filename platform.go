package main

import (
	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go/v6"
	"go.uber.org/zap"
)

// Platform represents the API platform
type Platform struct {
	logger   *zap.SugaredLogger
	db       *gorm.DB
	storage  *minio.Client
	settings *Settings
}
