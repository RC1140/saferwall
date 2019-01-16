// Copyright 2018 Saferwall. All rights reserved.
// Use of this source code is governed by Apache v2 license
// license that can be found in the LICENSE file.

package route

import (
	"github.com/labstack/echo"
	"github.com/saferwall/saferwall/web/app/handler/file"
	"github.com/saferwall/saferwall/web/app/handler/user"
	"github.com/saferwall/saferwall/web/app/middleware"
)

// New create an echo insance
func New() *echo.Echo {

	// Create `echo` instance
	e := echo.New()

	// Setup middlwares
	middleware.Init(e)

	// handle /files endpoint.
	e.GET("/v1/files", file.GetFiles)
	e.POST("/v1/files", file.PostFiles)
	e.PUT("/v1/files", file.PutFiles)
	e.DELETE("/v1/files", file.DeleteFiles)

	// handle /files/:sha256 endpoint.
	e.GET("/v1/files/:sha256", file.GetFile)
	e.PUT("/v1/files/:sha256", file.PutFile)
	e.DELETE("/v1/files/:sha256", file.DeleteFile)

	// handle /users endpoint.
	e.GET("/v1/users", user.GetUsers)
	e.POST("/v1/users", user.PostUsers)
	e.PUT("/v1/users", user.PutUsers)
	e.DELETE("/v1/users", user.DeleteUsers)

	// handle /users/:username  endpoint.
	e.GET("/v1/users/:username", user.GetUser)
	e.POST("/v1/users/:username", user.PostUser)
	e.PUT("/v1/users/:username", user.PutUser)
	e.DELETE("/v1/users/:username", user.DeleteUser)

	return e
}
