package fhirR4go

import (
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"

	"fmt"
	//"os"
	"testing"
	//"time"
	"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
	//"github.com/davecgh/go-spew/spew"
	//log "github.com/sirupsen/logrus"
)

//const pid = "4342009"
//const baseurla = "https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/"

func TestGetDocReferences(t *testing.T) {
	fmt.Printf("Test run a FHIR query\n")
	c := New(baseurl)
	Convey("Run a query", t, func() {

		fmt.Printf("GetRefReport\n")
		//url := fmt.Sprintf("%sDiagnosticReport?patient=12724066",baseurla)
		// data, err := c.GetDiagnosticReports("?patient=12724066")
		// So(err, ShouldBeNil)
		// So(data, ShouldNotBeNil)
		results, err := c.FindDocumentReferences("?patient=12724066")
		So(err, ShouldBeNil)
		So(results, ShouldNotBeNil)
		So(len(results.Entry), ShouldBeGreaterThan, 1)
		//fmt.Printf("Results: %s\n", spew.Sdump(results))

	})
}

func TestGetReferenceByID(t *testing.T) {
	fmt.Printf("Test run a FHIR query\n")
	c := New(baseurl)
	Convey("Run a query", t, func() {

		fmt.Printf("GetRefReferenceByID\n")
		//url := fmt.Sprintf("%sDiagnosticReport?patient=12724066",baseurla)
		// data, err := c.GetDiagnosticReports("?patient=12724066")
		// So(err, ShouldBeNil)
		// So(data, ShouldNotBeNil)
		rpt, err := c.GetDocumentReference("/197466727")
		So(err, ShouldBeNil)
		So(rpt, ShouldNotBeNil)
		fmt.Printf("rpt_id: %s\n", spew.Sdump(rpt))

	})
}

//https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/DiagnosticReport?patient=12714066
//https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/DiagnosticReport?patient=12724066
