package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/lucasdsampaio/starwars/api"
	"github.com/lucasdsampaio/starwars/models"
)

func TestCreatePlanetRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	values := map[string]string{
		"name":    "Tatooine",
		"climate": "arid",
		"terrain": "desert",
	}

	jsonValue, _ := json.Marshal(values)

	url := fmt.Sprintf("%s/planet", ts.URL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))

	var bodyResponseAPI api.ResponseAPI
	if err := json.NewDecoder(resp.Body).Decode(&bodyResponseAPI); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}

func TestListPlanetsRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/planets", ts.URL))

	planets := struct {
		data []models.Planet
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&planets); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
func TestGetPlanetByNameRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/planet-name/Tatooine", ts.URL))

	planets := struct {
		data []models.Planet
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&planets); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
func TestGetPlanetByIDRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/planet-id/1", ts.URL))

	planets := struct {
		data models.Planet
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&planets); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}

func TestDeletePlanetRoute(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()

	client := &http.Client{}
	url := fmt.Sprintf("%s/planet/1", ts.URL)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	planets := struct {
		data bool
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&planets); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

}
