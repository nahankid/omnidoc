package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"omnidoc/db"
	"omnidoc/lib"
	"omnidoc/models"
	"omnidoc/types"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/validator.v9"
)

var db2 *gorm.DB

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Validate Create Request
	req, err := validateRequest(request)
	if err != nil {
		return lib.APIResponse(http.StatusBadRequest, err.Error())
	}

	// Open Database Connection
	pgConn := db.PGConn{}
	db2, err = pgConn.GetConnection()
	defer db2.Close()
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	res, err := createAsset(req, request)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	return lib.APIResponse(http.StatusCreated, res)
}

func validateRequest(request events.APIGatewayProxyRequest) (types.CreateRequest, error) {
	var req types.CreateRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		log.Println("validateRequest", err.Error())
		return req, err
	}

	v := validator.New()
	validateErr := v.Struct(req)

	if validateErr != nil {
		log.Println("validateRequest", err.Error())
		return req, err
	}

	log.Println("validateRequest", req)

	return req, nil
}

func createAsset(req types.CreateRequest, request events.APIGatewayProxyRequest) (string, error) {
	db2.AutoMigrate(&models.Asset{})

	// Get PutPresignedURL
	key := fmt.Sprintf("/%d/%d/%s", req.AppID, req.UserID, req.FileName)
	psURL, err := lib.GetS3PresignedURL(key, lib.PutObjectRequest)
	if err != nil {
		log.Println("createAsset", err.Error())
		return "", err
	}

	// Create the asset
	asset := models.Asset{
		UserID:   req.UserID,
		AppID:    req.AppID,
		FileName: key,
		Type:     req.Type,
		Attrs:    postgres.Jsonb{RawMessage: json.RawMessage(req.Attrs)},
		APIKey:   request.RequestContext.Identity.APIKey,
	}

	db2.Create(&asset)
	if db2.Error != nil {
		log.Println("createAsset", err.Error())
		return "", db2.Error
	}

	res, err := json.Marshal(psURL)
	if err != nil {
		log.Println("createAsset", err.Error())
		return "", err
	}

	log.Println("createAsset", string(res))

	return string(res), nil
}

func main() {
	lambda.Start(handler)
}
