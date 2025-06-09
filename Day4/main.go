package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/subtract", subtractHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	http.HandleFunc("/divide", divideHandler)

	fmt.Println("Calculator API running on :8080")
	http.ListenAndServe(":8080", nil)
}

func getOperands(r *http.Request) (float64, float64, error) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid value for 'a'")
	}

	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid value for 'b'")
	}

	return a, b, nil
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a + b
	fmt.Fprintf(w, "Result: %.2f", result)
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a - b
	fmt.Fprintf(w, "Result: %.2f", result)
}

func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := a * b
	fmt.Fprintf(w, "Result: %.2f", result)
}

func divideHandler(w http.ResponseWriter, r *http.Request) {
	a, b, err := getOperands(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if b == 0 {
		http.Error(w, "Division by zero is not allowed", http.StatusBadRequest)
		return
	}
	result := a / b
	fmt.Fprintf(w, "Result: %.2f", result)
}
