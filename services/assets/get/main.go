package main

import (
	"encoding/json"
	"net/http"
	"omnidoc/db"
	"omnidoc/lib"
	"omnidoc/models"
	"omnidoc/types"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Validate Index Request
	a := request.QueryStringParameters["a"]
	u := request.QueryStringParameters["u"]
	t := request.QueryStringParameters["t"]

	if a == "" && u == "" {
		return lib.APIResponse(http.StatusBadRequest, "Provide appID or userID")
	}

	var appID, userID int
	var err error

	if a != "" {
		appID, err = strconv.Atoi(a)
		if err != nil {
			return lib.APIResponse(http.StatusBadRequest, err.Error())
		}
	}

	if u != "" {
		userID, err = strconv.Atoi(u)
		if err != nil {
			return lib.APIResponse(http.StatusBadRequest, err.Error())
		}
	}

	// Open Database Connection
	pgConn := db.PGConn{}
	db2, err := pgConn.GetConnection()
	defer db2.Close()
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	// Create filter for database query
	filter := models.Asset{}
	if appID != 0 {
		filter.AppID = appID
	}

	if userID != 0 {
		filter.UserID = userID
	}

	if t != "" {
		filter.Type = t
	}

	var assets []models.Asset

	db2.Where(filter).Find(&assets)
	if db2.Error != nil {
		return lib.APIResponse(http.StatusNotFound, db2.Error.Error())
	}

	// Create response
	resps := make([]types.GetResponse, len(assets))
	for i, asset := range assets {
		psURL, err := lib.GetS3PresignedURL(asset.FileName)
		if err != nil {
			return lib.APIResponse(http.StatusInternalServerError, err.Error())
		}

		resps[i] = types.GetResponse{Asset: asset, SignedURL: psURL}
	}

	res, err := json.Marshal(resps)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	return lib.APIResponse(http.StatusOK, string(res))
}

func main() {
	lambda.Start(handler)
}
