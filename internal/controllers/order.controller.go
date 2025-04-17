package controllers

import (
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/services"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService services.OrderServiceInterface
}

func NewOrderController(orderService services.OrderServiceInterface) *OrderController {
	return &OrderController{orderService: orderService}
}

func (r *OrderController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("limit"))
	if pageSize <= 0 {
		pageSize = 10
	}

	customer, total, err := r.orderService.GetAll(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Status: true,
		Data:   customer,
		Error:  "",
		Meta:   map[string]interface{}{"total": total, "page": page, "size": pageSize},
	})
	return
}

func (r *OrderController) GetDetail(c *gin.Context) {
	id := c.Param("id")
	order, err := r.orderService.GetDetail(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Status: true,
		Data:   order,
		Error:  "",
	})
	return
}

func (r *OrderController) Create(c *gin.Context) {
	var payload models.OrderCreateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	err := r.orderService.Create(c.Request.Context(), &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Status: true,
		Data:   nil,
		Error:  "",
	})
	return
}

func (r *OrderController) Update(c *gin.Context) {
	id := c.Param("id")
	var payload models.OrderUpdateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}
	payload.ID = id

	err := r.orderService.Update(c.Request.Context(), &payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status: true,
		Data:   nil,
		Error:  "",
	})
	return
}

func (r *OrderController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  "ID is required",
		})
		return
	}

	err := r.orderService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status: true,
		Data:   nil,
		Error:  "",
	})
	return
}
