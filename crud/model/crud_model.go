package model

type (
	GetDataResponse struct {
		Data []DataProduct `json:"product"`
	}
	DataProduct struct {
		Name string `json:"name" bson:"name"`
	}
)
