// Package models provides data models and structures used across the application.
package models

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type PermissionType string

const (
	ReadPermission  PermissionType = "read"
	WritePermission PermissionType = "write"
)
