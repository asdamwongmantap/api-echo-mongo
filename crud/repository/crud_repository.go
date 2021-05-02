package repository

import (
	"context"
	"github.com/asdamwongmantap/api-echo-mongo/crud"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type CrudRepository struct {
	mongoDB *mongo.Database
}

func NewCrudRepository(mongo *mongo.Database) crud.CrudRepositoryI {
	return &CrudRepository{
		mongoDB: mongo,
	}
}

func (cr CrudRepository) GetAllData(ctx context.Context) (crudResp model.GetDataResponse, err error) {

	query, err := cr.mongoDB.Collection("product").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return model.GetDataResponse{}, err
	}
	defer query.Close(ctx)

	listDataProduct := make([]model.DataProduct, 0)
	for query.Next(ctx) {
		var row model.DataProduct
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listDataProduct = append(listDataProduct, row)
	}

	crudResp = model.GetDataResponse{Data: listDataProduct}

	return crudResp, err
}
