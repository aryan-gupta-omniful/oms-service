package controllers

// import (
// 	"context"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"

// 	"oms-service/models"
// )

// type OrderController struct {
// 	collection *mongo.Collection
// }

// func NewOrderController(db *mongo.Database) *OrderController {
// 	return &OrderController{
// 		collection: db.Collection("orders"),
// 	}
// }

// // GetAllOrders godoc
// // @Summary Get all orders
// // @Description Get all orders from the database
// // @Tags orders
// // @Accept json
// // @Produce json
// // @Success 200 {array} models.Order
// // @Router /orders [get]
// func (oc *OrderController) GetAllOrders(c *gin.Context) {
// 	var orders []models.Order

// 	cursor, err := oc.collection.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching orders"})
// 		return
// 	}
// 	defer cursor.Close(context.Background())

// 	if err := cursor.All(context.Background(), &orders); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding orders"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, orders)
// }

// // GetOrderByID godoc
// // @Summary Get order by ID
// // @Description Get order details by order ID
// // @Tags orders
// // @Accept json
// // @Produce json
// // @Param id path string true "Order ID"
// // @Success 200 {object} models.Order
// // @Router /orders/{id} [get]
// func (oc *OrderController) GetOrderByID(c *gin.Context) {
// 	id := c.Param("id")
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	var order models.Order
// 	err = oc.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&order)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching order"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, order)
// }
