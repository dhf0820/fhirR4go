package fhirR4go

import (
	"time"
)

type Precision string
type FHIRDateTime struct {
	Time      time.Time
	Precision Precision
}
