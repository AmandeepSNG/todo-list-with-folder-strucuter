package userRepository

import (
	"context"
	"time"
	helperFunctions "todolist-app/common/helper"
	constants "todolist-app/common/shared-constants"
	enums "todolist-app/common/types"
	"todolist-app/config/db"
	"todolist-app/modals"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserList() enums.Response {
	client, err := db.GetMongoDbClient()
	var userList []modals.User

	if err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database((string(constants.DATABASE_NAME))).Collection(string(constants.USER_SCHEMA_NAME))

	cur, err := collection.Find(context.TODO(), bson.D{
		primitive.E{},
	})

	if err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}

	for cur.Next(context.TODO()) {
		var user modals.User
		err := cur.Decode(&user)
		if err != nil {
			return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
		}

		userList = append(userList, user)
	}
	return helperFunctions.Response(userList, constants.SUCCESS, constants.USER_FETCHED_SUCCESSFULLY, nil)
}

func GetUserById(userId string) enums.Response {
	var userDetails *modals.User
	client, err := db.GetMongoDbClient()
	if err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}
	defer client.Disconnect(context.Background())
	whereCondition := bson.D{{"userId", userId}}
	collection := client.Database(string(constants.DATABASE_NAME)).Collection(string(constants.USER_SCHEMA_NAME))
	// coll.FindOne(context.TODO(), filter).Decode(&result)
	err = collection.FindOne(context.TODO(), whereCondition).Decode(&userDetails)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return helperFunctions.Response(nil, constants.NOT_FOUND, constants.USER_NOT_FOUND, err)
		}
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}
	return helperFunctions.Response(userDetails, constants.SUCCESS, constants.USER_FETCHED_SUCCESSFULLY, nil)
}

func CreateUser(user modals.User) enums.Response {
	client, err := db.GetMongoDbClient()

	if err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database(string(constants.DATABASE_NAME)).Collection(string(constants.USER_SCHEMA_NAME))

	_, insertionErr := collection.InsertOne(context.TODO(), user)

	if insertionErr != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}

	return helperFunctions.Response(user, constants.CREATED, constants.USER_CREATED_SUCCESSFULLY, nil)
}

func UpdateUser(userId string, user modals.User) enums.Response {
	client, err := db.GetMongoDbClient()
	if err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}
	defer client.Disconnect(context.Background())
	whereCondition := bson.D{{"userId", userId}}

	updateQuery := bson.D{
		{"$set", bson.D{
			{"firstName", user.FirstName},
			{"lastName", user.LastName},
			{"email", user.Email},
			{"mobileNumber", user.MobileNumber},
			{"address", user.Address},
			{"updatedAt", time.Now()},
		}},
	}

	collection := client.Database(string(constants.DATABASE_NAME)).Collection(string(constants.USER_SCHEMA_NAME))

	_, updationErr := collection.UpdateOne(context.TODO(), whereCondition, updateQuery)

	if updationErr != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}

	return helperFunctions.Response(user, constants.SUCCESS, constants.USER_UPDATED_SUCCESSFULLY, nil)
}

func DeleteUser(userId string) enums.Response {
	client, err := db.GetMongoDbClient()

	if err != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}
	defer client.Disconnect(context.Background())
	collection := client.Database(string(constants.DATABASE_NAME)).Collection(string(constants.USER_SCHEMA_NAME))
	whereCondition := bson.D{{"userId", userId}}
	_, deletionErr := collection.DeleteOne(context.TODO(), whereCondition)

	if deletionErr != nil {
		return helperFunctions.Response(nil, constants.SERVER_ERROR, constants.INTERNAL_SERVER_ERROR, err)
	}
	return helperFunctions.Response(nil, constants.SUCCESS, constants.USER_DELETED_SUCCESSFULLY, nil)
}
