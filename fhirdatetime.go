package fhirongo

import (
	"time"
)

type Precision string
type FHIRDateTime struct {
	Time      time.Time
	Precision Precision
}
