package api

import (
	"reflect"
	"testing"
)

func TestAsyncGet(t *testing.T) {

	responseAPI, err := AsyncGet()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedType := []*HttpResponse{}
	if reflect.TypeOf(responseAPI) != reflect.TypeOf(expectedType) {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(responseAPI) != 6 {
		t.Fatalf("Expected no error, got %v", err)
	}

}
