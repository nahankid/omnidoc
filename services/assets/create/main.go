package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"omnidoc/db"
	"omnidoc/lib"
	"omnidoc/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/validator.v9"
)

type createRequest struct {
	UserID   int             `json:"user_id"`
	AppID    int             `json:"app_id"`
	Type     string          `json:"type"`
	FileName string          `json:"filename"`
	Attrs    json.RawMessage `json:"attrs"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Validate Login Request
	var req createRequest
	err := json.Unmarshal([]byte(request.Body), &req)
	if err != nil {
		return lib.APIResponse(http.StatusBadRequest, err.Error())
	}

	v := validator.New()
	validateErr := v.Struct(req)

	if validateErr != nil {
		return lib.APIResponse(http.StatusBadRequest, err.Error())
	}

	// Open Database Connection
	pgConn := db.PGConn{}
	db2, err := pgConn.GetConnection()
	defer db2.Close()
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	db2.AutoMigrate(&models.Asset{})

	// Get PutPresignedURL
	key := fmt.Sprintf("/%d/%d/%s", req.UserID, req.AppID, req.FileName)
	psURL, err := lib.PutS3PresignedURL(key)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	// Create the asset
	asset := models.Asset{UserID: req.UserID, AppID: req.AppID, Type: req.Type, Attrs: postgres.Jsonb{RawMessage: json.RawMessage(req.Attrs)}}
	db2.Create(&asset)

	res, err := json.Marshal(psURL)
	if err != nil {
		return lib.APIResponse(http.StatusInternalServerError, err.Error())
	}

	return lib.APIResponse(http.StatusCreated, string(res))
}

func main() {
	lambda.Start(handler)
}
