package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	Price       float64           `bson:"price"`
	Stock       int               `bson:"stock"`
	Image       string            `bson:"image"`
	CreatedAt   time.Time         `bson:"createdAt"`
	V           int               `bson:"__v"`
}

type Customer struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
	Phone string `bson:"phone"`
	Address string`bson:"address"`
}

type OrderProduct struct {
	Product  primitive.ObjectID `bson:"product"`
	Quantity int               `bson:"quantity"`
	Price    float64           `bson:"price"`
}

type Order struct {
	ID        primitive.ObjectID `bson:"_id"`
	Customer  Customer          `bson:"customer"`
	Products  []OrderProduct    `bson:"products"`
	Total     float64           `bson:"total"`
	Status    string            `bson:"status"`
	CreatedAt time.Time         `bson:"createdAt"`
	V         int               `bson:"__v"`
}

type Orders struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    ProductID string            `bson:"product_id"`
    Name      string            `bson:"name"`
    Email     string            `bson:"email"`
    Quantity  int               `bson:"quantity"`
}

// Constantes para el estado de las órdenes
const (
	OrderStatusPending   = "pendiente"
	OrderStatusPaid      = "pagado"
	OrderStatusDelivered = "enviado"
	OrderStatusCancelled = "cancelado"
)