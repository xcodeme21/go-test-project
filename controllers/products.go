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
	err := c.DB.Order("id asc").Find(&products).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": products, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products, "error_message": nil, "status": ctx.Writer.Status()})
}

func (c *ProductsController) ListDestinationProduct(ctx *gin.Context) {
	var products []models.DestinationProduct
	err := c.DB.Order("id asc").Find(&products).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": products, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": products, "error_message": nil, "status": ctx.Writer.Status()})
}

func (c *ProductsController) UpdateDestinationProduct(ctx *gin.Context) {
	var source []models.SourceProduct
	var destination []models.DestinationProduct
	err := c.DB.Find(&source).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"data": source, "error_message": err.Error(), "status": ctx.Writer.Status()})
		return
	}
	for _, v := range source {
		updateData := map[string]interface{}{"product_name": v.ProductName, "qty": v.Qty, "selling_price": v.SellingPrice, "promo_price": v.PromoPrice, "created_at": v.CreatedAt, "updated_at": v.UpdatedAt}
		err2 := c.DBTwo.Model(&destination).Where("id = ?", v.ID).Where("product_name = ?", v.ProductName).Updates(updateData).Error
		if err2 != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"data": nil, "error_message": "Update failed", "status": ctx.Writer.Status()})
			return
		}
	}

	var newDestination []models.DestinationProduct
	c.DBTwo.Order("id asc").Find(&newDestination)

	ctx.JSON(http.StatusOK, gin.H{"data": newDestination, "error_message": nil, "status": ctx.Writer.Status()})
}
