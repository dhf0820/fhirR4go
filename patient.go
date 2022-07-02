package fhirR4go

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	//"github.com/tidwall/pretty"
	//"github.com/davecgh/go-spew/spew"

	log "github.com/sirupsen/logrus"
)

// GetPatient will return patient information for a patient with id pid
func (c *Connection) GetPatient(pid string) (*Patient, error) {
	if pid == "" {
		msg := "fhir GetPatient param can not be blank"
		log.Errorf(msg)
		return nil, errors.New(msg)
	}
	//log.Infof("FHIR GetPatient url: %s/Patient/%v", c.BaseURL, pid)
	startTime := time.Now()
	qry := fmt.Sprintf("Patient/%v", pid)
	bytes, err := c.Query(qry)
	log.Infof("Query took %s", time.Since(startTime))
	if err != nil {
		msg := fmt.Sprintf("c.Query failed with [%s] err: %s\n", qry, err.Error())
		log.Error(msg)
		return nil, fmt.Errorf(msg)
	}
	//b := *body
	//fmt.Printf("\n\n\n@@@ Patient:22 RAW Patient: %s\n\n\n", pretty.Pretty(b))
	log.Infof("Length of bytes response: %d\n", len(bytes))
	data := Patient{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		fmt.Printf("UnMarshal GetPatient:25\n")
		return nil, err
	}
	//log.Infof("Query returning: %s", spew.Sdump(data))
	return &data, nil
}

func (c *Connection) FindFhirPatient(qry string) (*PatientResult, error) {
	//fmt.Printf("QRY: %s\n", qry)
	//fmt.Printf("With v: Patient%v\n", qry)
	//fmt.Printf("Patient%s\n", qry)
	//fmt.Printf("FHIR FindPatient url: %sPatient?%s\n", c.BaseURL, qry)
	query := fmt.Sprintf("Patient?%s", qry)
	bytes, err := c.Query(query)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ Patient 15 RAW Patient: %s\n\n\n", pretty.Pretty(b))
	data := PatientResult{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Connection) FindFhirPatients(qry string) (*PatientResult, error) {
	//fmt.Printf("QRY: %s\n", qry)
	//fmt.Printf("With v: Patient%v\n", qry)
	//fmt.Printf("Patient%s\n", qry)
	//fmt.Printf("FHIR FindPatient url: %sPatient?%s\n", c.BaseURL, qry)
	query := fmt.Sprintf("Patient%s", qry) // The query has the correect seperator(/, ?)
	bytes, err := c.Query(query)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("\n\n\n@@@ Patient 15 RAW Patient: %s\n\n\n", pretty.Pretty(b))
	data := PatientResult{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Connection) NextFhirPatients(url string) (*PatientResult, error) {
	//fmt.Printf("Next retrieving : %s\n", url)
	bytes, err := c.GetFhir(url)
	if err != nil {
		msg := fmt.Sprintf("NextPatient returned error: %s", err.Error())
		log.Errorf("%s", msg)
		return nil, err
	}

	data := PatientResult{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// Patient is a FHIR patient
type Patient struct {
	CacheID         primitive.ObjectID `json:"-" bson:"_id"`
	SessionId       string             `json:"-" bson:"sessiopn_id"`
	ResourceType    string             `json:"resourceType" bson:"resource_type"`
	ID              string             `json:"id" bson:"id"`
	Meta            MetaData           `json:"meta" bson:"meta"`
	Text            TextData           `json:"text" bson:"text"`
	Identifier      []Identifier       `json:"identifier" bson:"identifier"`
	Active          bool               `json:"active" bson:"active"`
	BirthDate       string             `json:"birthDate" bson:"birth_date"`
	Gender          string             `json:"gender" bson:"gender"`
	DeceasedBoolean bool               `json:"deceasedBoolean" bson:"deceased"`
	CareProvider    []Person           `json:"careProvider" bson:"care_provider"`
	Name            []HumanName        `json:"name" bson:"name"`
	Address         []Address          `json:"address" bson:"address"`
	Telecom         []Telecom          `json:"telecom" bson:"telecom"`
	MaritalStatus   Concept            `json:"maritalStatus" bson:"marital_status"`
	Communication   []Communication    `json:"communication" bson:"communication"`
	Extension       []Extension        `json:"extension" bson:"extension"`
	LastAccess      time.Time          `json:"-" bson:"last_access"`
}

// type PatientBundle struct {
// 	SearchResult
// 	Entry []struct {
// 		FullURL  string  `json:"fullUrl" bson:"full_url"`
// 		Resource Patient `json:"resource"`
// 	} `json:"entry"`
// }

type PatientBundle struct {
	ResourceType string `bson:"resource_type" json:"resourceType"`
	Id           string `bson:"id,omitempty" json:"id,omitempty"`
	Meta         *Meta  `bson:"meta,omitempty" json:"meta,omitempty"`
	//ImplicitRules *string       `bson:"implicit_rules,omitempty" json:"implicitRules,omitempty"`
	//Language      *string       `bson:"language,omitempty" json:"language,omitempty"`
	//Identifier    *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type Code `bson:"type" json:"type"` // document | message | transaton | transaction-response | batch | batch_response | history | searchset | collection
	//Timestamp     *string       `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	Total int                  `bson:"total,omitempty" json:"total,omitempty"`
	Link  []BundleLink         `bson:"link,omitempty" json:"link,omitempty"`
	Entry []PatientBundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	//Signature     *Signature    `bson:"signature,omitempty" json:"signature,omitempty"`
	SearchResult
}

type PatientBundleEntry struct {
	//Link     []BundleLink         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl  string               `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource *Patient             `bson:"resource,omitempty" json:"resource,omitempty"` //patient
	Search   *BundleEntrySearch   `bson:"search,omitempty" json:"search,omitempty"`
	Request  *BundleEntryRequest  `bson:"request,omitempty" json:"request,omitempty"`
	Response *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}
