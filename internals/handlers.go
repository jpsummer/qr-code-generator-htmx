package internals

import (
	"context"
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/jpsummer/qr-code-generator-htmx/views"
	"net/http"
	"time"
)

const appTimeout = time.Second * 10

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}

func (app *Config) indexPageHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		defer cancel()

		render(ctx, http.StatusOK, views.Index())
	}
}
