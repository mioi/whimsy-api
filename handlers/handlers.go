package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/mioi/whimsy"
)

func JSONMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func Root(w http.ResponseWriter, r *http.Request) {
	categories := whimsy.Categories()
	categoryNames := make([]string, len(categories))
	totalWords := 0

	for i, cat := range categories {
		categoryNames[i] = cat.Name
		totalWords += len(cat.Words)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":      "ok",
		"name":        "whimsy-api",
		"version":     "1.0.0",
		"categories":  categoryNames,
		"total_words": totalWords,
	})
}

func AllPlants(w http.ResponseWriter, r *http.Request) {
	plants := whimsy.Plants()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"plants": plants,
		"count":  len(plants),
	})
}

func AllAnimals(w http.ResponseWriter, r *http.Request) {
	animals := whimsy.Animals()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"animals": animals,
		"count":   len(animals),
	})
}

func AllColors(w http.ResponseWriter, r *http.Request) {
	colors := whimsy.Colors()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"colors": colors,
		"count":  len(colors),
	})
}

func AllNames(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"plants":  whimsy.Plants(),
		"animals": whimsy.Animals(),
		"colors":  whimsy.Colors(),
		"totals": map[string]int{
			"plants":  len(whimsy.Plants()),
			"animals": len(whimsy.Animals()),
			"colors":  len(whimsy.Colors()),
		},
	}

	json.NewEncoder(w).Encode(response)
}

func RandomPlants(w http.ResponseWriter, r *http.Request) {
	handleRandomItems(w, r, "plants", whimsy.Plants, whimsy.RandomPlant)
}

func RandomAnimals(w http.ResponseWriter, r *http.Request) {
	handleRandomItems(w, r, "animals", whimsy.Animals, whimsy.RandomAnimal)
}

func RandomColors(w http.ResponseWriter, r *http.Request) {
	handleRandomItems(w, r, "colors", whimsy.Colors, whimsy.RandomColor)
}

func RandomNames(w http.ResponseWriter, r *http.Request) {
	count := getCountParam(r, 10)
	parts := getPartsParam(r, 2) // default 2 parts per name

	names := make([]string, count)
	for i := 0; i < count; i++ {
		name, _ := whimsy.RandomName(parts)
		names[i] = name
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"names": names,
		"count": len(names),
		"parts": parts,
	})
}

func Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func handleRandomItems(w http.ResponseWriter, r *http.Request, category string, getAllItems func() []string, getRandomItem func() (string, error)) {
	categories := whimsy.Categories()
	validCategory := false
	for _, validCat := range categories {
		if validCat.Name == category {
			validCategory = true
			break
		}
	}

	if !validCategory {
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "not implemented yet",
		})
		return
	}

	count := getCountParam(r, 1)

	var items []string
	if count == 1 {
		item, _ := getRandomItem()
		items = []string{item}
	} else {
		items = getRandomItems(getAllItems(), count)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		category: items,
		"count":  len(items),
	})
}

func getCountParam(r *http.Request, defaultVal int) int {
	if c := r.URL.Query().Get("count"); c != "" {
		if parsed, err := strconv.Atoi(c); err == nil && parsed > 0 {
			return parsed
		}
	}
	return defaultVal
}

func getPartsParam(r *http.Request, defaultVal int) int {
	if p := r.URL.Query().Get("parts"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed >= 1 && parsed <= 3 {
			return parsed
		}
	}
	return defaultVal
}

func getRandomItems(items []string, count int) []string {
	if count >= len(items) {
		return items
	}

	selected := make([]string, count)
	used := make(map[int]bool)

	for i := 0; i < count; i++ {
		var idx int
		for {
			idx = rand.Intn(len(items))
			if !used[idx] {
				break
			}
		}
		selected[i] = items[idx]
		used[idx] = true
	}

	return selected
}
