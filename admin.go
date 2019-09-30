package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/digitalocean/godo"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v6"
	"golang.org/x/oauth2"
)

func (p *Platform) getCDNUrl(bucketName string) (string, error) {
	tokenSource := &TokenSource{
		AccessToken: p.settings.DigitalOcean.Token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

}

func (p *Platform) addRelease(ctx echo.Context) error {
	rp := new(ReleasePublish)
	if err := ctx.Bind(rp); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bad request"})
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
	rp.Release.PackageURL = ""

	p.db.Create(rp)

	n, err := p.storage.PutObject(p.settings.Spaces.BucketName, reader, reader.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

	return ctx.JSON(http.StatusOK, map[string]string{"message": "OK"})
}
