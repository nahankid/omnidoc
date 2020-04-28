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

You can add a document to DMS in two steps:

1. Get an asset pre-signed upload URL
2. Upload a file to the pre-signed URL

### Get an asset pre-signed upload URL

```POST /``` 

### Request
```
{
	"app_id": 4002,
	"user_id": 402,
	"type": "Loan Agreement",
	"attrs": {
		"key1": "value1",
		"key2": "value2"
	},
	"filename": "Meeting Mania.xlsx"
}
```

### Parameters

| Name      | Type     | Description                                              |
| --------- | ---------| -------------------------------------------------------- | 
| app_id    | int      | Application ID for which the document is being stored.   | 
| user_id   | int      | User ID for which the document is being stored.          | 
| type      | string   | Type of the document being stored.                       | 
| filename  | string   | Filename of the document to be stored.                   | 
| attrs     | JSON     | Any metadata to be stored with the document.             | 


### Response

```
{
    "url": "<SIGNED_URL>",
    "expires_at": "2020-04-28T03:29:26.882234986Z"
}
```

| Name         | Type     | Description                                              |
| ------------ | ---------| -------------------------------------------------------- | 
| url          | string   | Presigned URL for uploading the file                     | 
| expires_at   | string   | Time at which the the presigned URL expires.             | 


### Upload a file to the pre-signed URL
``` PUT <SIGNED_URL> ```




## **Retrieve documents from DMS**

You can rertrieve all documents from DMS for:
1. Application
2. User
3. Application and Type
4. User and Type

```
- GET /?u=401
- GET /?a=4001 
- GET /?a=4001&u=401
- GET /?a=4001&t=Loan Agreement
```

### Parameters

| Name         | Type     | Description                                                 |
| ------------ | ---------| ----------------------------------------------------------- | 
| u            | string   | User ID for which the documents are to be retrieved         | 
| a            | string   | Application ID for which the documents are to be retrieved  | 
| t            | string   | Type of document to be retrieved                            | 


### Response

```
[
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
            "url": "<SIGNED_URL>",
            "expires_at": "2020-04-28T03:29:55.334828375Z"
        }
    }
]
```


## Contributing to OmniDoc

Fork, fix, then send me a pull request.


## License

MIT


## Bringing to the next level

Next, you can use the following resources to know more about beyond hello world samples and how others structure their Serverless applications:

* [AWS Serverless Application Repository](https://aws.amazon.com/serverless/serverlessrepo/)
