package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xcodeme21/go-test-project/models"
	"gorm.io/gorm"
)

type ProductsController struct {
	DB    *gorm.DB
	DBTwo *gorm.DB
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

func (c *ProductsController) UpdateProduct(ctx *gin.Context) {
	var source []models.SourceProduct
	var destination models.DestinationProduct
	err := c.DB.Find(&source).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": source, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	for _, v := range source {
		updateQuery := c.DB.Model(&destination).Where("id = ?", v.ID)
		updateQuery.Update("product_name", v.ProductName)
		updateQuery.Update("qty", v.Qty)
		updateQuery.Update("selling_price", v.SellingPrice)
		updateQuery.Update("promo_price", v.PromoPrice)
		updateQuery.Update("created_at", v.CreatedAt)
		updateQuery.Update("updated_at", v.UpdatedAt)
	}

	var newDestination []models.DestinationProduct
	newErr := c.DB.Find(&source).Error
	if newErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": newDestination, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": newDestination, "error_message": nil, "status": ctx.Writer.Status()})
}
