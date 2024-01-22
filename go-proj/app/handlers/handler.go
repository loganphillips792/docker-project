package handlers

import (
	"context"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/loganphillips792/kubernetes-project/components"
	"github.com/loganphillips792/kubernetes-project/services"
	"log/slog"
	"net/http"
)

// func GetCount

// func PostCount

type CountService interface {
	Increment(ctx context.Context)
	// Get(ctx context.Context)
}

type DefaultHandler struct {
	Log           *slog.Logger
	CountServices CountService
}

func NewHandler(log *slog.Logger, s CountService) *DefaultHandler {
	return &DefaultHandler{
		Log:           log,
		CountServices: s,
	}
}

func (h *DefaultHandler) GetCount(c echo.Context) error {
	h.Log.Info(
		"incoming request",
		"method", "GET",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)

	count := 5
	return renderView(c, 200, components.Page(count, count))

}

func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	h.Log.Info(
		"incoming request",
		"method", "POST",
		"time_taken_ms", 158,
		"path", "/hello/world?q=search",
		"status", 200,
		"user_agent", "Googlebot/2.1 (+http://www.google.com/bot.html)",
	)
}

type ViewProps struct {
	Counts services.Counts
}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	components.Page(props.Counts.Global, props.Counts.Global).Render(r.Context(), w)
}

func renderView(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}
