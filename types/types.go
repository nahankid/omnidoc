package types

import (
	"encoding/json"
	"fmt"
	"omnidoc/lib"
	"omnidoc/models"
	"strings"
)

// CreateRequest struct
type CreateRequest struct {
	ObjectType string          `json:"obj_type"`
	ObjectID   int64           `json:"obj_id"`
	FileName   string          `json:"filename"`
	Type       string          `json:"type"`
	Attrs      json.RawMessage `json:"attrs"`
}

// GetResponse struct
type GetResponse struct {
	Asset     models.Asset       `json:"asset"`
	SignedURL lib.S3PresignedURL `json:"signed_url"`
}

// DocumentCodes constant
var DocumentCodes = map[string]int{
	"loan agreement":                      101,
	"caf":                                 102,
	"soc":                                 103,
	"noc":                                 104,
	"welcome letter":                      105,
	"foreclosure letter":                  106,
	"soa":                                 107,
	"rps":                                 108,
	"delivery order":                      109,
	"insurance form":                      110,
	"vehicle registration certificate":    111,
	"undertaking docket":                  112,
	"fi report":                           113,
	"fcu report":                          114,
	"property":                            115,
	"nach":                                116,
	"pdc":                                 117,
	"drd":                                 118,
	"approval email":                      119,
	"others":                              120,
	"application docket":                  121,
	"application docket (rectified)":      122,
	"existing lan documents":              123,
	"sanction letter":                     124,
	"application form":                    125, 
	"passport":                            201,
	"pan":                                 202,
	"aadhaar":                             203,
	"driving license":                     204,
	"voter id card":                       205,
	"nrega job card":                      206,
	"utility bill":                        207,
	"bank statement":                      208,
	"rent agreement":                      209,
	"experian bureau report":              210,
	"cibil bureau report":                 211,
	"crif bureau report":                  212,
	"equifax bureau report":               213,
	"ckyc xml":                            214,
	"okyc xml":                            215,
	"aadhaar front":                       216,
	"aadhaar back":                        217,
	"driving license front":               218,
	"driving license back":                219,
	"voter id card front":                 220,
	"voter id card back":                  221,
	"passport front":                      222,
	"passport back":                       223,
	"photo":                               224,
	"address proof":                       225,
	"constitution of entity":              226,
	"commercial cibil bureau report":      227,
	"pd questionnaire":                    228,
	"student data":                        229,
	"ca certificate":                      230,
	"audit report":                        231,
	"valuation report":                    232,
	"legal report":                        233,
	"financial statement analysis report": 234,
	"bank statement analysis report":      235,
	"posidex output":                      236,
	"debtors and creditors list":          237,
	"cma data/projections":                238,
	"udyam certificate":                   239,
	"ownership proof":                     240,
	"kyc docket":                          241,
	"gst":								   242	  
}

// Valid function validates request
func (cr CreateRequest) Valid() error {
	dc := DocumentCodes[strings.ToLower(cr.Type)]

	if dc == 0 {
		return fmt.Errorf("type is missing or invalid")
	}

	if cr.ObjectID == 0 {
		return fmt.Errorf("obj_id is required")
	}

	// It's an app document
	if dc > 100 && dc < 200 && cr.ObjectType != "app" {
		return fmt.Errorf("%s requires obj_type 'app'", cr.Type)
	}

	// It's a borrower document
	if dc > 200 && cr.ObjectType != "user" {
		return fmt.Errorf("%s requires obj_typee 'user'", cr.Type)
	}

	return nil
}
