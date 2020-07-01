package processor

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/rohsa/excel-column-generator-go/objects"
)

// GetColumnNames generates a list of excel column names by using the "columnObject" data
// It takes the starting name as the value of startColumn,
// if it's empty then it defaults to first column name, i.e. "A"
func GetColumnNames(columnObject objects.Column)([]string, error) {
	count := int(columnObject.Count)
	startIndex, err := excelize.ColumnNameToNumber(columnObject.StartColumn)
	if err != nil {
		return nil, fmt.Errorf("invalid start column: %s", err.Error())
	}

	var columnNames []string
	for i := startIndex; i < (startIndex + count); i++ {
		columnName, err := excelize.ColumnNumberToName(i)
		if err != nil {
			return nil, err
		}
		columnNames = append(columnNames, columnName)
	}
	return columnNames, nil
}
