package fhirongo

import (
	"encoding/json"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson"
)

type ContainedResources []interface{}

type DomainResource struct {
	Resource          `bson:",inline"`
	Text              *Narrative         `bson:"text,omitempty" json:"text,omitempty"`
	Contained         ContainedResources `bson:"contained,omitempty" json:"contained,omitempty"`
	Extension         []Extension        `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension        `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}

// Convert contained resources from map[string]interfac{} to specific types.
// Custom marshalling methods on those types will then hide internal fields
// like @context and reference__id.
// func (x *ContainedResources) SetBSON(raw bson.Raw) (err error) {

// 	// alias type to avoid infinite loop when calling Unmarshal
// 	type containedResources ContainedResources
// 	x2 := (*containedResources)(x)
// 	if err = raw.Unmarshal(x2); err == nil {
// 		if x != nil {
// 			for i := range *x {
// 				(*x)[i], err = BSONMapToResource((*x)[i].(bson.M), true)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// Custom marshaller to add the resourceType property, as required by the specification
func (resource *DocumentReference) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "DocumentReference"
	// Dereferencing the pointer to avoid infinite recursion.
	// Passing in plain old x (a pointer to DocumentReference), would cause this same
	// MarshallJSON function to be called again
	return json.Marshal(*resource)
}

func (x *DocumentReference) GetBSON() (interface{}, error) {
	x.ResourceType = "DocumentReference"
	// See comment in MarshallJSON to see why we dereference
	return *x, nil
}

// The "documentReference" sub-type is needed to avoid infinite recursion in UnmarshalJSON
type documentReference DocumentReference

// Custom unmarshaller to properly unmarshal embedded resources (represented as interface{})
// func (x *DocumentReference) UnmarshalJSON(data []byte) (err error) {
// 	x2 := documentReference{}
// 	if err = json.Unmarshal(data, &x2); err == nil {
// 		if x2.Contained != nil {
// 			for i := range x2.Contained {
// 				x2.Contained[i], err = MapToResource(x2.Contained[i], true)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 		}
// 		*x = DocumentReference(x2)
// 		return x.checkResourceType()
// 	}
// 	return
// }

type DocumentReferenceContextComponent struct {
	BackboneElement   `bson:",inline"`
	Encounter         []Reference                                `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Event             []CodeableConcept                          `bson:"event,omitempty" json:"event,omitempty"`
	Period            *Period                                    `bson:"period,omitempty" json:"period,omitempty"`
	FacilityType      *CodeableConcept                           `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	PracticeSetting   *CodeableConcept                           `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	SourcePatientInfo *Reference                                 `bson:"sourcePatientInfo,omitempty" json:"sourcePatientInfo,omitempty"`
	Related           []DocumentReferenceContextRelatedComponent `bson:"related,omitempty" json:"related,omitempty"`
}

type DocumentReferenceContextRelatedComponent struct {
	BackboneElement `bson:",inline"`
	Identifier      *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Ref             *Reference  `bson:"ref,omitempty" json:"ref,omitempty"`
}
