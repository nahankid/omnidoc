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
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jinzhu/gorm"
)

var db2 *gorm.DB

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	filter, err := validateRequest(request)
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

	c := make(chan int)
	go recordRequest(request, filter, c)

	assets, err := getAssets(filter)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	log.Println("Recorded Request", <-c)

	return lib.APIResponse(http.StatusOK, assets)
}

func validateRequest(request events.APIGatewayProxyRequest) (models.Asset, error) {
	// Validate Index Request
	obj := request.QueryStringParameters["o"]
	id := request.QueryStringParameters["id"]
	t := request.QueryStringParameters["t"]

	filter := models.Asset{}

	if obj == "" && id == "" && !(obj == "app" || obj == "user") {
		log.Printf("Missing or invalid obj_type %s and obj_id %s", obj, id)
		return filter, fmt.Errorf("Missing or invalid obj_type %s and obj_id %s", obj, id)
	}

	oid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("validateRequest", err.Error())
		return filter, err
	}

	// Create filter for database query
	filter.ObjectType = obj
	filter.ObjectID = oid

	if t != "" {
		filter.Type = t
	}

	log.Println("Validated Request", filter)
	return filter, nil
}

func recordRequest(request events.APIGatewayProxyRequest, filter models.Asset, c chan int) {
	// Record visit
	db2.AutoMigrate(&models.Visit{})

	visit := models.Visit{
		ObjectType: filter.ObjectType,
		ObjectID:   filter.ObjectID,
		IP:         request.RequestContext.Identity.SourceIP,
		UserAgent:  request.RequestContext.Identity.UserAgent,
		APIKey:     request.RequestContext.Identity.APIKey,
		VisitedAt:  time.Now(),
	}

	log.Println("Recording visit", visit)

	db2.Create(visit)
	if db2.Error != nil {
		log.Println("recordRequest", db2.Error.Error())
	}

	c <- 1
}

func getAssets(filter models.Asset) (string, error) {
	var assets []models.Asset

	db2.Where(filter).Select("DISTINCT ON (file_name) *").Order("file_name, updated_at DESC").Find(&assets)
	// db2.Where(filter).Find(&assets)
	if db2.Error != nil {
		log.Println("getAssets", db2.Error.Error())
		return "", db2.Error
	}

	// Create response
	resps := make([]types.GetResponse, len(assets))
	for i, asset := range assets {
		psURL, err := lib.GetS3PresignedURL(asset.FileName, lib.GetObjectRequest)
		if err != nil {
			log.Println("getAssets", err.Error())
			return "", err
		}

		resps[i] = types.GetResponse{Asset: asset, SignedURL: psURL}
	}

	res, err := json.Marshal(resps)
	return string(res), err
}

func main() {
	lambda.Start(handler)
}
