package frontend

import (
	"context"
	"embed"
	"html/template"
	"io"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/usememos/memos/internal/util"
	"github.com/usememos/memos/server/profile"
	"github.com/usememos/memos/store"
)

//go:embed dist
var embeddedFiles embed.FS

type FrontendService struct {
	Profile *profile.Profile
	Store   *store.Store
}

func NewFrontendService(profile *profile.Profile, store *store.Store) *FrontendService {
	return &FrontendService{
		Profile: profile,
		Store:   store,
	}
}

var baseURL = ""

func indexHander(c echo.Context) error {
	// open index.html
	file, err := embeddedFiles.Open("dist/index.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer file.Close()
	// render it.
	b, err := io.ReadAll(file)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().WriteHeader(http.StatusOK)
	tmpl, err := template.New("index.html").Parse(string(b))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	err = tmpl.Execute(c.Response().Writer, map[string]any{
		"baseurl": baseURL,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (f *FrontendService) Serve(_ context.Context, e *echo.Echo) {
	skipper := func(c echo.Context) bool {
		pathTmp := c.Path()
		return pathTmp == "/" || pathTmp == "/index.html" || util.HasPrefixes(pathTmp, "/api", "/memos.api.v1")
	}

	e.GET("/", indexHander)
	e.GET("/index.html", indexHander)
	// save baseUrl from profile.
	baseURL = f.Profile.BaseURL

	// Use echo static middleware to serve the built dist folder.
	// Reference: https://github.com/labstack/echo/blob/master/middleware/static.go
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5:      false,
		Filesystem: getFileSystem("dist"),
		Skipper:    skipper,
	}), func(skipper middleware.Skipper) echo.MiddlewareFunc {
		return func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) (err error) {
				// skip
				if skipper(c) {
					return next(c)
				}
				// skip assets
				if util.HasPrefixes(c.Path(), "/assets") {
					return next(c)
				}
				// otherwise (NotFound), serve index.html
				return indexHander(c)
			}
		}
	}(skipper))
	// Use echo gzip middleware to compress the response.
	// Reference: https://echo.labstack.com/docs/middleware/gzip
	e.Group("assets").Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}), func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderCacheControl, "max-age=31536000, immutable")
			return next(c)
		}
	}, middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: getFileSystem("dist/assets"),
	}))
}

func getFileSystem(path string) http.FileSystem {
	fs, err := fs.Sub(embeddedFiles, path)
	if err != nil {
		panic(err)
	}
	return http.FS(fs)
}
