package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hatrungt1n/go-api-sample/configs"
	"github.com/hatrungt1n/go-api-sample/models"
	util "github.com/hatrungt1n/go-api-sample/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

// @Summary Get a list of users
// @Description get all users
// @Produce  json
// @Success 200 {array} util.Responses
// @Router /users [get]
func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			util.APIErrorResponse(c, http.StatusInternalServerError, http.MethodGet, err.Error())
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				util.APIErrorResponse(c, http.StatusInternalServerError, http.MethodGet, err.Error())
				return
			}

			users = append(users, singleUser)
		}

		util.APIResponse(c, "success", http.StatusOK, http.MethodGet, map[string]interface{}{"users": users})
	}
}

// @Summary Create a user
// @Description create a user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Success 201 {object} util.Responses
// @Router /user [post]
func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()
		fmt.Printf("%v", *userCollection)

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			util.APIErrorResponse(c, http.StatusBadRequest, http.MethodPost, err.Error())
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			util.APIErrorResponse(c, http.StatusBadRequest, http.MethodPost, validationErr.Error())
			return
		}

		newUser := models.User{
			// Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Age:      user.Age,
			Location: user.Location,
		}

		_, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			util.APIErrorResponse(c, http.StatusInternalServerError, http.MethodPost, err.Error())
			return
		}

		util.APIResponse(c, "success", http.StatusCreated, http.MethodPost, map[string]interface{}{"user": newUser})
	}
}
