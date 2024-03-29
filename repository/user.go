package repository

import (
	"context"
	"net/http"
	"strconv"

	"jwt-project/database"

	"jwt-project/database/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (sR mongoRepository) GetResults(c *gin.Context, ctx context.Context) *mongo.Cursor {
	matchStage, groupStage, projectStage := sR.getStages(c)
	result, _ := database.Collection(database.Connect(), model.TABLE).Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage,
	})

	return result
}

func (mongoRepository) getStages(c *gin.Context) (primitive.D, primitive.D, primitive.D) {
	recordPerPage, errorConvertionRecord := strconv.Atoi(c.Query("recordPerPage"))
	if errorConvertionRecord != nil || recordPerPage < 1 {
		recordPerPage = 10
	}

	page, errorConvertionPage := strconv.Atoi(c.Query("page"))
	if errorConvertionPage != nil || page < 1 {
		page = 1
	}

	startIndex, errorConvertionStartIndex := strconv.Atoi(c.Query("startIndex"))
	if errorConvertionStartIndex != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Provide a valid integer start number"})
	}

	matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}

	groupStage := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: bson.D{{Key: "_id", Value: "null"}}},
		{Key: "total_count", Value: bson.D{{Key: "$sum", Value: 1}}},
		{Key: "data", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}}}}}

	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "total_count", Value: 1},
			{Key: "user_items", Value: bson.D{{Key: "$slice", Value: []interface{}{"$data", startIndex, recordPerPage}}}}}}}

	return matchStage, groupStage, projectStage
}
