# excel-column-generator-go
This project generates the excel column names using a javascript `POST` request. It displays a web UI to the user taking three inputs:
1. Start column name (It is optional and defaults to `A`)
2. Row count (Dislays the given number of rows)
3. Column Count (Displays the given number of columns)

This project uses the go package 'excelize' library. The library size being huge it could not be vendored. Please make sure to fetch the module from the path: `github.com/360EntSecGroup-Skylar/excelize`
