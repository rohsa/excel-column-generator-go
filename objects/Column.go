package objects

// Column is the struct used as a request and response body for the "/column" handler
type Column struct {
	StartColumn string 		`json:"StartColumn"`
	Count int32				`json:"Count"`
	ColumnNames []string	`json:"ColumnNames"`
}
