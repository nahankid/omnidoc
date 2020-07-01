# OmniDoc

![Go](https://github.com/nahankid/omnidoc/workflows/Go/badge.svg)
![GitHub](https://img.shields.io/github/license/nahankid/omnidoc)
[![Maintainability](https://api.codeclimate.com/v1/badges/add8791ba98cf2cc2a5e/maintainability)](https://codeclimate.com/github/nahankid/omnidoc/maintainability)
![GitHub repo size](https://img.shields.io/github/repo-size/nahankid/omnidoc)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/nahankid/omnidoc)

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
	"obj_type": "app",
	"obj_id": 402,
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
| obj_type  | string   | Object Type for which the document is being stored - app or user  | 
| obj_id    | int      | Object ID for which the document is being stored.          | 
| type      | string   | Type of the document being stored. Valid types:<br />Loan Agreement<br />CAF<br />SOC<br />NOC<br />Welcome Letter<br />Foreclosure Letter<br />SOA<br />RPS<br />Delivery Order<br />Insurance Form<br />Vehicle Registration Certificate<br />Passport<br />PAN<br />Aadhaar<br />Driving License<br />Voter ID Card<br />NREGA Job Card<br />Utility Bill<br />Bank Statement<br />Rent Agreement<br />Experian Bureau Report<br />CIBIL Bureau Report<br />CRIF Bureau Report<br />Equifax Bureau Report<br />CKYC XML<br />OKYC XML<br />Aadhaar Front<br />Aadhaar Back<br />Driving License Front<br />Driving License Back<br />Voter ID Card Front<br />Voter ID Card Back<br />Passport Front<br />Passport Back<br />Photo |
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
- GET /?o=app&id=402
- GET /?o=app&id=402&t=Loan Agreement
```

### Parameters

| Name         | Type     | Description                                                 |
| ------------ | ---------| ----------------------------------------------------------- | 
| o            | string   | Object Type for which the documents are to be retrieved     | 
| id           | int      | Object ID for which the documents are to be retrieved       | 
| t            | string   | Type of document to be retrieved                            | 


### Response

```
[
[
    {
        "asset": {
            "obj_type": "app",
            "obj_id": 121,
            "filename": "/app/121/loanagreement.pdf",
            "type": "Loan Agreement",
            "attrs": {
                "key1": "value1"
            },
            "created_at": "2020-05-23T19:12:19.308369Z",
            "updated_at": "2020-05-23T19:12:19.308369Z"
        },
        "signed_url": {
            "url": "<SIGNED_URL",
            "expires_at": "2020-05-24T04:42:31.6585508Z"
        }
    },
  
]
```


## Contributing to OmniDoc

Fork, fix, then send me a pull request.


## License

MIT
