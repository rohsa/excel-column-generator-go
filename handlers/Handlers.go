package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/rohsa/excel-column-generator-go/objects"
	"github.com/rohsa/excel-column-generator-go/processor"
	"github.com/rohsa/excel-column-generator-go/rendoring"
	"net/http"
	"path/filepath"
)

var viewTemplate = filepath.Join("view")

// ViewHandler is the request handler for the home page
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	rendoring.RenderTemplate(w, viewTemplate)
}

// GetColumnHandler is the request handler for the api "/column"
// POST:
// Takes the request body of type "objects.Column"
// Processes the user data to calculate the excel column names from the given
// startColumn, rowCount and columnCount
func GetColumnHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		var columnObject objects.Column
		err := decoder.Decode(&columnObject)
		if err != nil {
			http.Error(w, "Error parsing request", http.StatusInternalServerError)
		}
		columnNames, err := processor.GetColumnNames(columnObject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		}
		columnObject.ColumnNames = columnNames
		responseBytes, err := json.Marshal(columnObject)
		if err != nil {
			// Return internal server error
			http.Error(w, "Empty date cannot be processed", http.StatusInternalServerError)
			return
		}
		_, err = fmt.Fprintf(w, "%s\n", responseBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
