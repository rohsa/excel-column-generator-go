# excel-column-generator-go
This project generates the excel column names using a javascript `POST` request. It displays a web UI to the user taking three inputs:
1. Start column name (It is optional and defaults to `A`)
2. Row count (Dislays the given number of rows)
3. Column Count (Displays the given number of columns)

The above input data is used to create an ajax POST request to get the column names. The backend logic is implemented using the GoLang/Go Programming. The column names received as a repsponse is displayed in tablular format on the web UI.

**PS:** This project uses the go package `excelize` library. The library size being huge it could not be vendored. Please make sure to fetch the module from the path: `github.com/360EntSecGroup-Skylar/excelize`

# Running the project
1. Please make sure the project follows the correct path order (`github.com/rohsa/excel-column-generator-go`)
2. Please make sure you're not using/blocking the port `8080` on your localhost
2. `cd` to the project directory and run `go build main.go`
3. Open your web browser and hit the url `www.localhost:8080`
