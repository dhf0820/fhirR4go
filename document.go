package fhirR4go

import (
	"encoding/json"
	"fmt"

	//"time"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/pretty"
	//"github.com/tidwall/pretty"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetDocumentReference will return a result which has an array of document references
func (c *Connection) FindDocumentReferences(qry string) (*DocumentResults, error) {
	fmt.Printf("\n%sDocumentReference%s\n", c.BaseURL, qry)
	body, err := c.Query(fmt.Sprintf("DocumentReference/%s", qry))
	if err != nil {
		return nil, err
	}
	data := DocumentResults{}

	b := body
	//bodyStr := pretty.Pretty(b[:])
	//fmt.Printf("\n\n\n@@@ RAW DocumentReference: %s\n\n\n", bodyStr)

	//json.NewDecoder(body).Decode(&data)
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Printf("GetDocumentReference err: %s\n", err.Error())
		return nil, err
	}
	return &data, nil
}

// GetDocumentReference will return a result which has an array of document references
func (c *Connection) GetPatientDocumentReferences(pid string) (*DocumentResults, error) {
	qry := fmt.Sprintf("?patient=%s", pid)
	fmt.Printf("\n%sDocumentReference%s\n", c.BaseURL, qry)
	body, err := c.Query(fmt.Sprintf("DocumentReference/%s", qry))
	if err != nil {
		return nil, err
	}
	data := DocumentResults{}

	b := body
	//bodyStr := pretty.Pretty(b[:])
	//fmt.Printf("\n\n\n@@@ RAW DocumentReference: %s\n\n\n", bodyStr)

	//json.NewDecoder(body).Decode(&data)
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Printf("GetDocumentReference err: %s\n", err.Error())
		return nil, err
	}
	return &data, nil
}

//FindDocumentReference will return one document
func (c *Connection) GetDocumentReference(pid string) (*DocumentResults, error) {
	qry := fmt.Sprintf("?patient=%s", pid)
	fmt.Printf("%sDocumentReference%s\n", c.BaseURL, qry)
	body, err := c.Query(fmt.Sprintf("DocumentReference%s", qry))
	if err != nil {
		log.Errorf("FhirDocumentReference cerner returned error: %s", err.Error())
		return nil, err
	}
	fmt.Printf("\n\n\n@@@ RAW DocumentReference: %s\n\n\n", pretty.Pretty(body))
	data := &DocumentResults{}
	if err := json.Unmarshal(body, data); err != nil {
		log.Errorf("FhirDocumentReference Unmarshal error: %s", err.Error())
		return nil, err
	}
	fmt.Printf("FindDocumentReference:50 returning all %s\n", spew.Sdump(data))
	return data, nil
}

// Process the next page of DocRefs
func (c *Connection) NextFhirDocRefs(url string) (*DocumentResults, error) {
	//fmt.Printf("Next retrieving : %s\n", url)
	bytes, err := c.GetFhir(url)
	if err != nil {
		msg := fmt.Sprintf("NextPatient returned error: %s", err.Error())
		log.Errorf("%s", msg)
		return nil, err
	}

	data := DocumentResults{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

type DocumentResponse struct {
	//Bundle Bundle
	SearchResult
	Entry []struct {
		FullURL  string            `json:"fullUrl"`
		Document DocumentReference `json:"resource"`
	} `json:"entry"`
}
type DocumentResults struct {
	//Bundle Bundle
	SearchResult
	Entry []struct {
		FullURL  string            `json:"fullUrl"`
		Document DocumentReference `json:"resource"`
	}
}

// DocumentReference is a FHIR document
type ReferenceResults struct {
	Bundle
	Entry []struct {
		FullURL string            `json:"fullUrl"`
		DocRef  DocumentReference `json:"resource"`
	} `json:"entry"`
}

// DocumentReference is a single FHIR DocumentReference.
// Use DocumentReferences for a bundle.
type DocumentReference struct {
	DomainResource   `bson:",inline"`
	MasterIdentifier Identifier                            `bson:"masterIdentifier,omitempty" json:"masterIdentifier,omitempty"` //Version Specific
	Identifier       []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Text             *TextData                             `bson:"text" json:"text"`
	Subject          Reference                             `json:"subject" bson:"subject"`
	Type             CodeableConcept                       `json:"type" bson:"type"`
	Class            CodeableConcept                       `json:"class" bson:"class"`
	Author           []Reference                           `json:"author" bson:"author"`
	Custodian        Reference                             `json:"custodian" bson:"custodian"`
	Authenticator    Reference                             `json:"authenticator" bsn:"authenticator"`
	Created          *FHIRDateTime                         `json:"created" bson:"created"`
	Indexed          *FHIRDateTime                         `json:"indexed" bson:"indexed"`
	Status           Code                                  `json:"status" bson:"status"`
	DocStatus        Code                                  `json:"docStatus" bson:"doc_status"` //
	RelatesTo        []DocumentReferenceRelatesToComponent `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	Description      string                                `json:"description" bson:"description"`
	SecurityLabel    []CodeableConcept                     `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`

	Content []DocumentReferenceContentComponent `bson:"content,omitempty" json:"content,omitempty"`
	Context DocumentReferenceContextComponent   `bson:"context,omitempty" json:"context,omitempty"`
}

type DocumentReferenceRelatesToComponent struct {
	BackboneElement `bson:",inline"`
	Code            string     `bson:"code,omitempty" json:"code,omitempty"`
	Target          *Reference `bson:"target,omitempty" json:"target,omitempty"`
}

type DocumentReferenceContentComponent struct {
	//BackboneElement `bson:",inline"`
	Attachment *Attachment `bson:"attachment,omitempty" json:"attachment,omitempty"`
	Format     *Coding     `bson:"format,omitempty" json:"format,omitempty"`
}

// type DocumentReference struct {
// 	CacheID          primitive.ObjectID                  `bson:"cache_id" json:"cacheId"`
// 	SessionID        string                              `bson:"session_id" json:"sessionId"`
// 	ResourceType     string                              `bson:"resource_type" json:"resourceType"`
// 	ID               string                              `bson:"id"jbson:"id"`
// 	Meta             MetaData                            `bson:"meta" json:"meta"`
// 	Text             TextData                            `bson:"text" json:"text"`
// 	Identifier       []Identifier                        `json:"identifier" bson:"identifier"`
// 	Status           string                              `bson:"status" json:"status"`
// 	DocStatus        CodeableConcept                     `bson:"doc_status" json:"docSatus"`
// 	Type             Concept                             `bson:"type" json:"type"`
// 	Category         CodeableConcept                     `bson:"category" json:"category"`
// 	Subject          Person                              `bson:"subject" json:"subject"`
// 	EfectiveDateTime time.Time                           `bson:"effective_date_time" json:"date"`
// 	Author           Reference                           `bson:"author" json:"author"`
// 	Authenticator    Reference                           `bson:"authenticator" json:"authenticator"`
// 	Custodian        Reference                           `bson:"custodian" json:"custodian"`
// 	RelatesTo        DocumentReferenceRelatesToComponent `bson:"relates_to" json:"relatesTo"`
// 	Description      string                              `bson:"description" json:"description"`
// 	SecurityLabel    CodeableConcept                     `bson:"security_label" json:"securityLabel"`

// 	Content []struct {
// 		Attachment struct {
// 			ContentType string    `json:"contentType" bson:"content_type"`
// 			URL         string    `json:"url" bson:"url"`
// 			Title       string    `json:"title" bson:"title"`
// 			Creation    time.Time `json:"creation" bson:"creation"`
// 		} `json:"attachment" bson:"attachment"`
// 	} `json:"content" bson:"content"`
// 	Context EncounterContext `bson:"context" json:"context"`
// }

// "coding": [
// 	{
// 	  "system": "http://terminology.hl7.org/CodeSystem/data-absent-reason",
// 	  "code": "unknown",
// 	  "display": "Unknown"
// 	}
// type DocumentBundle struct {
// 	SearchResult
// 	Entry []struct {
// 		FullURL  string `json:"fullUrl" bson:"full_url"`
// 		Resource struct {
// 			Document
// 		} `json:"resource"`
// 	} `json:"entry"`
// }

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
// type Bundle struct {
// 	ResourceType *string `bson:"resource_type" json:"resourceType"`
// 	Id           *string `bson:"id,omitempty" json:"id,omitempty"`
// 	Meta         *Meta   `bson:"meta,omitempty" json:"meta,omitempty"`
// 	//ImplicitRules *string       `bson:"implicit_rules,omitempty" json:"implicitRules,omitempty"`
// 	//Language      *string       `bson:"language,omitempty" json:"language,omitempty"`
// 	//Identifier    *Identifier   `bson:"identifier,omitempty" json:"identifier,omitempty"`
// 	Type Code `bson:"type" json:"type"` // document | message | transaton | transaction-response | batch | batch_response | history | searchset | collection
// 	//Timestamp     *string       `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
// 	Total *int          `bson:"total,omitempty" json:"total,omitempty"`
// 	Link  []BundleLink  `bson:"link,omitempty" json:"link,omitempty"`
// 	Entry []BundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
// 	//Signature     *Signature    `bson:"signature,omitempty" json:"signature,omitempty"`
// }
type DocumentBundle struct {
	ResourceType string                `bson:"resource_type" json:"resourceType"`
	Id           string                `bson:"id,omitempty" json:"id,omitempty"`
	Meta         *Meta                 `bson:"meta,omitempty" json:"meta,omitempty"`
	Type         Code                  `bson:"type" json:"type"` // document | message | transaton | transaction-response | batch | batch_response | history | searchset | collection
	Total        int                   `bson:"total,omitempty" json:"total,omitempty"`
	Link         []BundleLink          `bson:"link,omitempty" json:"link,omitempty"`
	Entry        []DocumentBundleEntry `bson:"entry,omitempty" json:"entry,omitempty"`
	//Signature     *Signature    `bson:"signature,omitempty" json:"signature,omitempty"`
}

type DocumentBundleEntry struct {
	// Id                string              `bson:"id,omitempty" json:"id,omitempty"`
	// Extension         []Extension          `bson:"extension,omitempty" json:"extension,omitempty"`
	// ModifierExtension []Extension          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Link     []BundleLink         `bson:"link,omitempty" json:"link,omitempty"`
	FullUrl  string               `bson:"fullUrl,omitempty" json:"fullUrl,omitempty"`
	Resource DocumentReference    `bson:"resource,omitempty" json:"resource,omitempty"` //Document, patient, encounter. Diagnosticreport
	Search   *BundleEntrySearch   `bson:"search,omitempty" json:"search,omitempty"`
	Request  *BundleEntryRequest  `bson:"request,omitempty" json:"request,omitempty"`
	Response *BundleEntryResponse `bson:"response,omitempty" json:"response,omitempty"`
}
