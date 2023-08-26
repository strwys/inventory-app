package handler

import (
	"net/http"

	"github.com/cecepsprd/inventory-app/model"
	"github.com/cecepsprd/inventory-app/service"
	"github.com/cecepsprd/inventory-app/utils/logger"
	"github.com/labstack/echo/v4"
)

type inventoryHandler struct {
	inventoryService service.InventoryService
}

func NewInventoryHandler(e *echo.Echo, inventoryService service.InventoryService) {
	handler := inventoryHandler{
		inventoryService: inventoryService,
	}

	e.POST("/api/inventory/in", handler.InventoryIn)
	e.POST("/api/inventory/out", handler.InventoryOut)
	e.GET("/api/inventory/report", handler.StockReport)
}

func (h *inventoryHandler) InventoryIn(c echo.Context) error {
	var req model.InventoryInRequest

	if err := c.Bind(&req); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.inventoryService.InventoryIn(req); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (h *inventoryHandler) InventoryOut(c echo.Context) error {
	var req model.InventoryOutRequest

	if err := c.Bind(&req); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := h.inventoryService.InventoryOut(req); err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func (h *inventoryHandler) StockReport(c echo.Context) error {

	reports, err := h.inventoryService.StockReport()
	if err != nil {
		logger.Log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    reports,
	})
}
