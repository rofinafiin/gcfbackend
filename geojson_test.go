package gcfbackend

import (
	"fmt"
	"testing"
)

func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("mongodb+srv://rofinafiin:aXz4RdVqUVIQcqa1@rofinafiinsdata.9fyvx4r.mongodb.net", "GIS", "geogis")

	fmt.Printf("%+v", data)
}
