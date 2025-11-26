package routes

import (
	"github.com/minseoi/gorizon/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// 식재료 관련 라우트
	e.GET("/ingredients", handlers.GetIngredients)          // 모든 식재료 조회
	e.POST("/ingredients", handlers.CreateIngredient)       // 식재료 추가
	e.PUT("/ingredients/:id", handlers.UpdateIngredient)    // 식재료 이름 변경
	e.DELETE("/ingredients/:id", handlers.DeleteIngredient) // 식재료 제거
}
