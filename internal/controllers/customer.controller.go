package controllers

import (
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/services"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerController struct {
	customerService services.CustomerServiceInterface
}

func NewCustomerController(customerService services.CustomerServiceInterface) *CustomerController {
	return &CustomerController{customerService: customerService}
}

func (r *CustomerController) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("limit"))
	if pageSize <= 0 {
		pageSize = 10
	}

	customer, total, err := r.customerService.GetAll(c.Request.Context(), page, pageSize)
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

func (r *CustomerController) GetDetail(c *gin.Context) {
	id := c.Param("id")
	customer, err := r.customerService.GetDetail(c.Request.Context(), id)
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
		Data:   customer,
		Error:  "",
	})
	return
}

func (r *CustomerController) Create(c *gin.Context) {
	var payload models.CustomerCreateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}

	err := r.customerService.Create(c.Request.Context(), &payload)
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

func (r *CustomerController) Update(c *gin.Context) {
	var payload models.CustomerUpdateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  err.Error(),
		})
		return
	}
	payload.ID = c.Param("id")

	if payload.ID == "" {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: false,
			Data:   nil,
			Error:  "ID is required",
		})
		return
	}

	err := r.customerService.Update(c.Request.Context(), &payload)
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

func (r *CustomerController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := r.customerService.Delete(c.Request.Context(), id)
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

func (r *CustomerController) Search(c *gin.Context) {
	search := c.Query("keyword")
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.Query("limit"))
	if pageSize <= 0 {
		pageSize = 10
	}

	customer, total, err := r.customerService.Search(c.Request.Context(), search, page, pageSize)
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
