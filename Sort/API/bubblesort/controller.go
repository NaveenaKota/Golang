package bubblesort

import (
	     "net/http"
		 "encoding/json"
		 "io/ioutil"
		 "fmt"
		 )

func bubblesort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the input JSON
	var input []int
	err = json.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Sort the array
	BubbleSort(input)

	// Encode and send the sorted array as the response
	json.NewEncoder(w).Encode(input)
}

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(array)
}

