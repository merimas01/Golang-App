package controllers

import (
	"net/http"
	"strconv"

	"Golang-App/interfaces"

	"github.com/gin-gonic/gin"
)

type BaseController[T any, Tinsert any, Tupdate any, Tsearch any] struct {
	Service interfaces.ICRUDService[T, Tinsert, Tupdate, Tsearch]
}

func (ctrl *BaseController[T, Tinsert, Tupdate, Tsearch]) Create(c *gin.Context) {
	var input Tinsert

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := ctrl.Service.Create(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity)
}

func (ctrl *BaseController[T, Tinsert, Tupdate, Tsearch]) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input Tupdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := ctrl.Service.Update(&input, uint(id))
	if err != nil {
		//c.JSON(http.StatusNotFound, gin.H{"error": "Update unsuccessful"})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entity)
}

func (ctrl *BaseController[T, Tinsert, Tupdate, Tsearch]) GetAll(c *gin.Context) {
	var search Tsearch

	// Bind query parameters to Tsearch struct (page, pageSize, filters)
	if err := c.ShouldBindQuery(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	// Pass by value, not pointer
	result, err := ctrl.Service.GetAll(&search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ctrl *BaseController[T, Tinsert, Tupdate, TSearch]) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	entity, err := ctrl.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}

	c.JSON(http.StatusOK, entity)
}

func (ctrl *BaseController[T, Tinsert, Tupdate, Tsearch]) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Entity deleted successfully"})
}
