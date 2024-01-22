package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(e *echo.Echo, handler *DefaultHandler) {
	e.GET("/", handler.GetCount)
}

// e.GET("/hello", func(c echo.Context) error {
// 	component := components.Hello("John")
// 	return component.Render(context.Background(), c.Response().Writer)
// })

// e.GET("/", func(c echo.Context) error {
// 	component := components.Page(5, 5)
// 	return component.Render(context.Background(), c.Response().Writer)
// })

// e.POST("/", func(c echo.Context) error {
// 	component := components.Page(6, 5)
// 	return component.Render(context.Background(), c.Response().Writer)
// })
