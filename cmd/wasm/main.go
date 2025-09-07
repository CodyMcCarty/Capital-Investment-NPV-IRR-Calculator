package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func DoubleLogic(input float64) float64 {
	return input * 2
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the input number from form data.
	inputStr := r.FormValue("number")
	input, err := strconv.ParseFloat(inputStr, 64)
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Apply the doubling logic.
	result := DoubleLogic(input)

	// Return JSON response.
	response := map[string]float64{"doubled": result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/double", doubleHandler)
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
