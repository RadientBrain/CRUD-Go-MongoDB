package routes

import (
	getcollection "Samudai/Collection"
	database "Samudai/databases"
	model "Samudai/model"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// POST endpoint
func CreatePost(c *gin.Context) {
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "payment")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.Payment)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	postPayload := model.Payment{
		// ID:               primitive.NewObjectID(),
		Sender:           post.Sender,
		Payment_id:       uuid.New(),
		Chain_id:         post.Chain_id,
		Initiated_at:     post.Initiated_at,
		Status:           post.Status,
		Created_by:       post.Created_by,
		Payment_type:     post.Payment_type,
		Updated_at:       time.Now(),
		Created_at:       time.Now(),
		Completed_at:     post.Completed_at,
		Transaction_hash: post.Transaction_hash,
		Task_id:          post.Task_id,
		Currency:         post.Currency,
		Amount:           post.Amount,
		Dao:              post.Dao,
		Receiver:         post.Receiver,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}
