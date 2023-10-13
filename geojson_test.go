package gcfbackend

import (
	"fmt"
	"testing"
)

func TestGCHandlerFunc(t *testing.T) {
	data := GCHandlerFunc("mONGO", "GIS", "geogis")

	fmt.Printf("%+v", data)
}
