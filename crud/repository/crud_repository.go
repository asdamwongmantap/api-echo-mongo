package repository

import (
	"context"
	"fmt"
	"github.com/asdamwongmantap/api-echo-mongo/crud"
	"github.com/asdamwongmantap/api-echo-mongo/crud/model"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (cr CrudRepository) InsertData(ctx context.Context, req model.DataProductRequest) (err error) {

	dataReq := bson.M{
		"product_name": req.ProductName,
	}

	query, err := cr.mongoDB.Collection("product").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		productID := oid.Hex()
		dataUpdateProductID := bson.M{"_id": oid}
		dataObjectID := bson.M{"$set": bson.M{"product_id": productID}}
		_, err := cr.mongoDB.Collection("product").UpdateOne(ctx, dataUpdateProductID, dataObjectID)
		if err != nil {
			log.Println("error")
		}
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID ", err))
		log.Println("error")
	}

	return err
}

func (cr CrudRepository) UpdateData(ctx context.Context, req model.DataProductRequest) (err error) {

	dataUpdateProductID := bson.M{"product_id": req.ProductID}
	dataObjectID := bson.M{"$set": bson.M{
		"product_name": req.ProductName,
	}}
	_, err = cr.mongoDB.Collection("product").UpdateOne(ctx, dataUpdateProductID, dataObjectID)
	if err != nil {
		log.Println("error")
	}

	return err
}
