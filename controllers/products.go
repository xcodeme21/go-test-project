package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-test-project/models"
	"gorm.io/gorm"
)

type ProductsController struct {
	DB *gorm.DB
}

func (c *ProductsController) ListSourceProduct(ctx *gin.Context) {
	var products []models.SourceProduct
	err := c.DB.Find(&products).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": products, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products, "error_message": nil, "status": ctx.Writer.Status()})
}

func (c *ProductsController) ListDestinationProduct(ctx *gin.Context) {
	var products []models.DestinationProduct
	err := c.DB.Find(&products).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": products, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products, "error_message": nil, "status": ctx.Writer.Status()})
}