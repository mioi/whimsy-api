package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAllPlants(t *testing.T) {
	req := httptest.NewRequest("GET", "/plants", nil)
	w := httptest.NewRecorder()

	AllPlants(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if _, ok := response["plants"]; !ok {
		t.Error("response missing 'plants' field")
	}

	if _, ok := response["count"]; !ok {
		t.Error("response missing 'count' field")
	}
}

func TestAllAnimals(t *testing.T) {
	req := httptest.NewRequest("GET", "/animals", nil)
	w := httptest.NewRecorder()

	AllAnimals(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if _, ok := response["animals"]; !ok {
		t.Error("response missing 'animals' field")
	}

	if _, ok := response["count"]; !ok {
		t.Error("response missing 'count' field")
	}
}

func TestAllColors(t *testing.T) {
	req := httptest.NewRequest("GET", "/colors", nil)
	w := httptest.NewRecorder()

	AllColors(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if _, ok := response["colors"]; !ok {
		t.Error("response missing 'colors' field")
	}

	if _, ok := response["count"]; !ok {
		t.Error("response missing 'count' field")
	}
}

func TestAllNames(t *testing.T) {
	req := httptest.NewRequest("GET", "/names", nil)
	w := httptest.NewRecorder()

	AllNames(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expectedFields := []string{"plants", "animals", "colors", "totals"}
	for _, field := range expectedFields {
		if _, ok := response[field]; !ok {
			t.Errorf("response missing '%s' field", field)
		}
	}
}

func TestRandomPlants(t *testing.T) {
	req := httptest.NewRequest("GET", "/plants/random?count=5", nil)
	w := httptest.NewRecorder()

	RandomPlants(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if _, ok := response["plants"]; !ok {
		t.Error("response missing 'plants' field")
	}

	if count, ok := response["count"].(float64); !ok || count != 5 {
		t.Errorf("expected count 5, got %v", response["count"])
	}
}

func TestRandomNames(t *testing.T) {
	req := httptest.NewRequest("GET", "/names/random?count=3&parts=2", nil)
	w := httptest.NewRecorder()

	RandomNames(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if _, ok := response["names"]; !ok {
		t.Error("response missing 'names' field")
	}

	if count, ok := response["count"].(float64); !ok || count != 3 {
		t.Errorf("expected count 3, got %v", response["count"])
	}

	if parts, ok := response["parts"].(float64); !ok || parts != 2 {
		t.Errorf("expected parts 2, got %v", response["parts"])
	}
}

func TestHealth(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	Health(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if status, ok := response["status"]; !ok || status != "ok" {
		t.Errorf("expected status 'ok', got %v", response["status"])
	}
}

func TestGetCountParam(t *testing.T) {
	tests := []struct {
		url        string
		expected   int
		defaultVal int
	}{
		{"/?count=5", 5, 1},
		{"/?count=0", 1, 1},
		{"/?count=invalid", 10, 10},
		{"/", 3, 3},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		result := getCountParam(req, test.defaultVal)
		if result != test.expected {
			t.Errorf("for URL %s, expected %d, got %d", test.url, test.expected, result)
		}
	}
}

func TestGetPartsParam(t *testing.T) {
	tests := []struct {
		url        string
		expected   int
		defaultVal int
	}{
		{"/?parts=1", 1, 2},
		{"/?parts=3", 3, 2},
		{"/?parts=4", 2, 2},
		{"/?parts=0", 2, 2},
		{"/", 2, 2},
	}

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.url, nil)
		result := getPartsParam(req, test.defaultVal)
		if result != test.expected {
			t.Errorf("for URL %s, expected %d, got %d", test.url, test.expected, result)
		}
	}
}

func TestGetRandomItems(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}

	result := getRandomItems(items, 3)
	if len(result) != 3 {
		t.Errorf("expected 3 items, got %d", len(result))
	}

	result = getRandomItems(items, 10)
	if len(result) != len(items) {
		t.Errorf("expected %d items when count exceeds available, got %d", len(items), len(result))
	}

	seen := make(map[string]bool)
	result = getRandomItems(items, 5)
	for _, item := range result {
		if seen[item] {
			t.Error("found duplicate item in result")
		}
		seen[item] = true
	}
}

func TestHandleRandomItems(t *testing.T) {
	mockGetAll := func() []string {
		return []string{"item1", "item2", "item3"}
	}

	mockGetRandom := func() (string, error) {
		return "random_item", nil
	}

	t.Run("count=1 uses random function", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?count=1", nil)
		w := httptest.NewRecorder()

		handleRandomItems(w, req, "plants", mockGetAll, mockGetRandom)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		items, ok := response["plants"].([]interface{})
		if !ok {
			t.Error("response missing or invalid 'plants' field")
		}

		if len(items) != 1 {
			t.Errorf("expected 1 item, got %d", len(items))
		}

		if items[0] != "random_item" {
			t.Errorf("expected 'random_item', got %v", items[0])
		}
	})

	t.Run("count>1 uses getRandomItems", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?count=2", nil)
		w := httptest.NewRecorder()

		handleRandomItems(w, req, "animals", mockGetAll, mockGetRandom)

		if w.Code != http.StatusOK {
			t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		items, ok := response["animals"].([]interface{})
		if !ok {
			t.Error("response missing or invalid 'animals' field")
		}

		if len(items) != 2 {
			t.Errorf("expected 2 items, got %d", len(items))
		}
	})

	t.Run("unrecognized category returns 501", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/?count=1", nil)
		w := httptest.NewRecorder()

		handleRandomItems(w, req, "unknown_category", mockGetAll, mockGetRandom)

		if w.Code != http.StatusNotImplemented {
			t.Errorf("expected status %d, got %d", http.StatusNotImplemented, w.Code)
		}

		var response map[string]string
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}

		if error, ok := response["error"]; !ok || error != "not implemented yet" {
			t.Errorf("expected error 'not implemented yet', got %v", response["error"])
		}
	})
}
