package routes

import (
	getcollection "Samudai/Collection"
	database "Samudai/databases"
	model "Samudai/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "Payment")

	postId := c.Param("postId")
	var post model.Payment

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	edited := bson.M{"sender": post.Sender,
		"chain_id":         post.Chain_id,
		"initiated_at":     post.Initiated_at,
		"status":           post.Status,
		"created_by":       post.Created_by,
		"payment_type":     post.Payment_type,
		"updated_at":       time.Now(),
		"completed_at":     post.Completed_at,
		"transaction_hash": post.Transaction_hash,
		"task_id":          post.Task_id,
		"currency":         post.Currency,
		"amount":           post.Amount,
		"dao":              post.Dao,
		"receiver":         post.Receiver}

	result, err := postCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}
