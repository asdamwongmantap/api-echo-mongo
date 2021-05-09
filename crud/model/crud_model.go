package model

type (
	GetDataResponse struct {
		Data []DataProduct `json:"product"`
	}
	DataProduct struct {
		ProductID   string `json:"product_id,omitempty" bson:"product_id"`
		ProductName string `json:"product_name" bson:"product_name"`
	}
	DataProductRequest struct {
		ProductID   string `json:"product_id,omitempty" bson:"product_id"`
		ProductName string `json:"product_name" bson:"product_name"`
	}
)
