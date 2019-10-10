package main

import "github.com/jinzhu/gorm"

// ReleaseType is a Enum representing the release channel
type ReleaseType string

const (
	// Alpha is an alpha level release
	Alpha ReleaseType = "alpha"
	// Beta is an beta level release
	Beta ReleaseType = "beta"
	// Stable is an stable level release
	Stable ReleaseType = "stable"
)

type (
	// Admin represents an admin user with access key
	Admin struct {
		Username  string `json:"username" gorm:"unique;not null"`
		AccessKey string `json:"access_key" gorm:"unique;not null"`
	}

	// Release represents a specific released version of the software
	Release struct {
		gorm.Model
		Version  string      `gorm:"unique;not null" json:"version"`
		Info     string      `json:"description"`
		Type     ReleaseType `gorm:"not null" json:"channel"`
		FileName string      `json:"file_name"`
		Hash     string      `json:"hash"`
	}

	// ReleasePublish represents a release to publish including base64 encoded bytes for the binary
	ReleasePublish struct {
		Data    string  `json:"data"`
		Release Release `json:"release"`
	}

	// Check represents a request to check the current version
	Check struct {
		Version string      `json:"current_version"`
		Channel ReleaseType `json:"channel"`
	}

	// CheckResponse contains the reply to a Check request
	CheckResponse struct {
		Update  bool    `json:"update_available"`
		Release Release `json:"release"`
	}

	// ReleasesRequest represents a request to get a list of releases, leave channel blank to request all channels
	ReleasesRequest struct {
		Channel ReleaseType `json:"channel"`
		Limit   int         `json:"limit"`
	}

	// ReleasesResponse represents a response to a ReleasesRequest
	ReleasesResponse struct {
		Releases []Release `json:"releases"`
	}

	// Settings represents the settings file
	Settings struct {
		DB struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			User     string `json:"user"`
			DB       string `json:"db"`
			Password string `json:"password"`
		} `json:"db"`
		Spaces struct {
			Endpoint        string `json:"endpoint"`
			AccessKeyID     string `json:"access_key_id"`
			SecretAccessKey string `json:"secret_access_key"`
			BucketName      string `json:"bucket_name"`
			CDNEndpoint     string `json:"cdn_endpoint"`
		} `json:"spaces"`
		Port string `json:"listen_port"`
	}
)
