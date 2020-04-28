# OmniDoc

![Go](https://github.com/nahankid/shortie/workflows/Go/badge.svg)
![GitHub](https://img.shields.io/github/license/nahankid/sho.rt)
[![Maintainability](https://api.codeclimate.com/v1/badges/eb06a36a6fcda7abc6f2/maintainability)](https://codeclimate.com/github/nahankid/sho.rt/maintainability)
![GitHub repo size](https://img.shields.io/github/repo-size/nahankid/sho.rt)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/nahankid/sho.rt)


OmniDoc is a Document Management System, written in Go and automatically deployed on AWS, using:

- AWS API Gateway
- AWS Lambda
- AWS S3
- RDS (PostgreSQL)
- AWS Serverless Application Model (SAM)


## Quick start

## **Add a document to DMS**

```POST /``` 

### Parameters

| Name      | Type     | Description                                              |
| --------- | ---------| -------------------------------------------------------- | 
| app_id    | int      | Application ID for which the document is being stored.   | 
| user_id   | int      | User ID for which the document is being stored.          | 
| type      | string   | Type of the document being stored.                       | 
| filename  | string   | Filename of the document to be stored.                   | 
| attrs     | JSON     | Any metadata to be stored with the document.             | 


### Response

| Name         | Type     | Description                                              |
| ------------ | ---------| -------------------------------------------------------- | 
| url          | string   | Presigned URL for uploading the file                     | 
| expires_at   | string   | Time at which the the presigned URL expires.             | 




## **Get documents from DMS**

```GET /``` 

### Parameters

| Name         | Type     | Description                                                 |
| ------------ | ---------| ----------------------------------------------------------- | 
| u            | string   | User ID for which the documents are to be retrieved         | 
| a            | string   | Application ID for which the documents are to be retrieved  | 
| t            | string   | Type of document to be retrieved                            | 

### Response

```[
    {
        "asset": {
            "ID": 5,
            "CreatedAt": "2020-04-28T03:15:12.526861Z",
            "UpdatedAt": "2020-04-28T03:15:12.526861Z",
            "DeletedAt": null,
            "AppID": 4002,
            "UserID": 401,
            "FileName": "/4002/401/Hidden Characteristics of Unicorns.pdf",
            "Type": "Loan Agreement",
            "Attrs": {
                "key1": "value1"
            }
        },
        "signed_url": {
            "url": ""
        }
    }
]```


## Contributing to Shortie

Fork, fix, then send me a pull request.

## License

MIT
