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
	// Release represents a specific released version of the software
	Release struct {
		gorm.Model
		Version    string      `gorm:"unique;not null" json:"version"`
		Info       string      `json:"description"`
		Type       ReleaseType `gorm:"not null" json:"channel"`
		PackageURL string      `json:"url"`
	}

	// Check represents a request to check the current version
	Check struct {
		Version string      `json:"current_version"`
		Channel ReleaseType `json:"channel"`
	}

	// CheckResponse contains the reply to a Check request
	CheckResponse struct {
		Update  bool    `json:"update_available`
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
		Port string `json:"listen_port"`
	}
)
