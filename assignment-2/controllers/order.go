package controllers

import (
	"assignment-2/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	db *gorm.DB
}

type ReqOrder struct {
	OrderedAt    time.Time `json:"orderedAt"`
	CustomerName string    `json:"customerName"`
	Items        []ReqItem `json:"items"`
}

type ReqItem struct {
	LineItemId  uint   `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

func NewOrderController(db *gorm.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}

func (o *OrderController) GetAllOrder(ctx *gin.Context) {
	var orders []models.Order

	err := o.db.Preload("Items").Find(&orders).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"data":    orders,
	})
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var newReqOrder ReqOrder

	err := ctx.ShouldBindJSON(&newReqOrder)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	newOrder := models.Order{
		CustomerName: newReqOrder.CustomerName,
		OrderedAt:    newReqOrder.OrderedAt,
	}

	err = o.db.Create(&newOrder).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	for i, item := range newReqOrder.Items {
		newItem := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     newOrder.ID,
		}
		err = o.db.Create(&newItem).Error
		if err != nil {
			badRequestResponse(ctx, err.Error())
			return
		}
		itemRef := &newReqOrder.Items[i]
		itemRef.LineItemId = newItem.ID
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"data":    newReqOrder,
	})
}

func (o *OrderController) GetOneOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderID")
	var order models.Order

	err := o.db.Preload("Items").First(&order, orderId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			notFoundResponse(ctx, "Order data not found")
			return
		}
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"data":    order,
	})
}

func (o *OrderController) UpdateOrder(ctx *gin.Context) {
	var newReqOrder ReqOrder

	err := ctx.ShouldBindJSON(&newReqOrder)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	orderId := ctx.Param("orderID")
	var order models.Order
	var result []string

	err = o.db.Preload("Items").First(&order, orderId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			notFoundResponse(ctx, "Order data not found")
			return
		}
		badRequestResponse(ctx, err.Error())
		return
	}

	newOrder := models.Order{
		CustomerName: newReqOrder.CustomerName,
		OrderedAt:    newReqOrder.OrderedAt,
	}

	err = o.db.Model(&order).Updates(newOrder).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	} else {
		result = append(result, fmt.Sprintf("Order ID %d: updated", order.ID))
	}

	for _, item := range newReqOrder.Items {
		newItem := models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderID:     order.ID,
		}
		var itemModel models.Item
		err = o.db.First(&itemModel, item.LineItemId).Error
		if err != nil {
			result = append(result, fmt.Sprintf("Item ID %d: Error: %s", item.LineItemId, err.Error()))
			err = nil
			continue
		}

		err = o.db.Model(&itemModel).Updates(newItem).Error
		if err != nil {
			result = append(result, fmt.Sprintf("Item ID %d: Error: %s", item.LineItemId, err.Error()))
			err = nil
			continue
		}

		result = append(result, fmt.Sprintf("Item ID %d: updated", itemModel.ID))
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

func (o *OrderController) DeleteOrder(ctx *gin.Context) {
	orderId := ctx.Param("orderID")
	var order models.Order
	var items []models.Item

	err := o.db.First(&order, orderId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			notFoundResponse(ctx, "Order data not found")
			return
		}
		badRequestResponse(ctx, err.Error())
		return
	}

	err = o.db.Delete(&order).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			notFoundResponse(ctx, err.Error())
			return
		}
		badRequestResponse(ctx, err.Error())
		return
	} else {
		o.db.Where("order_id = ?", order.ID).Delete(&items)
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("Order ID %d has been deleted", order.ID),
	})
}
