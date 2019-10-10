package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v6"
)

func (p *Platform) addRelease(ctx echo.Context) error {
	rp := new(ReleasePublish)
	if err := ctx.Bind(rp); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bad request", "error": err.Error()})
	}

	data, err := base64.StdEncoding.DecodeString(rp.Data)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid base64"})
	}

	hasher := sha256.New()
	hasher.Write(data)

	rp.Release.Hash = fmt.Sprintf("%x", hasher.Sum(nil))
	reader := bytes.NewReader(data)

	u := uuid.New()
	rp.Release.FileName = u.String()

	_, err = p.storage.PutObject(p.settings.Spaces.BucketName, u.String(), reader, reader.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "error", "error": err.Error()})
	}
	result := p.db.Create(&rp.Release)
	if result.Error != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "error", "error": result.Error.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
