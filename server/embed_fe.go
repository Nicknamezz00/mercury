/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package server

import (
	"embed"
	"github.com/Nicknamezz00/mercury/common/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/fs"
	"net/http"
)

//go:embed dist
var embeddedFiles embed.FS

func getFileSystem(path string) http.FileSystem {
	f, err := fs.Sub(embeddedFiles, path)
	if err != nil {
		panic(err)
	}
	return http.FS(f)
}

// embedFrontend use echo static middleware to serve "dist" folder
// refer: https://github.com/labstack/echo/blob/master/middleware/static.go
func embedFrontend(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper:    defaultAPIRequestSkipper,
		HTML5:      true,
		Filesystem: getFileSystem("dist"),
	}))

	assertsGroup := e.Group("assets")
	assertsGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderCacheControl, "max-age=31536000, immutable")
			return next(c)
		}
	})
	assertsGroup.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper:    defaultAPIRequestSkipper,
		HTML5:      true,
		Filesystem: getFileSystem("dist/assets"),
	}))
}

func defaultAPIRequestSkipper(c echo.Context) bool {
	path := c.Request().URL.Path
	return util.HasPrefixes(path, "/api", "/mercury.api.v2")
}
