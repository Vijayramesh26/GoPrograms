package api

import (
	"Sample/ftdb"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

// ResponseStruct holds the API response structure
type ResponseStruct struct {
	ResponseArr []ResponseArr `json:"responseArr"` // Array of response data
	Status      string        `json:"status"`      // Status of the response
	ErrMsg      string        `json:"errMsg"`      // Error message if any
}

// ResponseArr is a sample response array structure
type ResponseArr struct {
	Field1 string `json:"field1"` // Example field in response
	Field2 int    `json:"field2"` // Another field in response
}

// RequestStruct holds the incoming request structure
type RequestStruct struct {
	Param1 string `json:"param1" validate:"required"` // First parameter, required
	Param2 int    `json:"param2" validate:"required"` // Second parameter, required
}

// SampleAPI handles the main API logic
func SampleAPI(lHttpWriter http.ResponseWriter, lHttpRequest *http.Request) {
	log.Println("SampleAPI (+)")

	// Set headers to allow CORS (Cross-Origin Resource Sharing)
	lHttpWriter.Header().Set("Access-Control-Allow-Origin", "*")
	lHttpWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	lHttpWriter.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	lHttpWriter.Header().Set("Access-Control-Allow-Headers", "USER, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Initialize response structure
	var lResponseRec ResponseStruct

	// Check if the request method is POST
	if lHttpRequest.Method == http.MethodPost {
		// 1. Collect: Collect and unmarshal the incoming request data
		var lRequestData RequestStruct
		lErr := CollectMethod(lHttpRequest, &lRequestData)
		if lErr != nil {
			log.Println("SampleAPI:001 - Error in CollectMethod:", lErr)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg = "Error in request collection: " + lErr.Error()
			goto Complete // Jump to the completion step if there's an error
		}

		// 2. Validate: Validate the collected data
		lErr = ValidateMethod(lRequestData)
		if lErr != nil {
			log.Println("SampleAPI:002 - Validation failed:", lErr)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg = "Validation failed: " + lErr.Error()
			goto Complete // Jump to the completion step if validation fails
		}

		// 3. Construct: Process and construct batch information from validated data
		lBatchId, lErr := ConstructMethod(lRequestData)
		if lErr != nil {
			log.Println("SampleAPI:003 - Construction error:", lErr)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg = "Data construction failed: " + lErr.Error()
			goto Complete // Jump to the completion step if construction fails
		}

		// 4. Communicate: Interact with other services or databases and get data
		lResponseArr, lErr := CommunicateMethod(lBatchId)
		if lErr != nil {
			log.Println("SampleAPI:004 - Communication error:", lErr)
			lResponseRec.Status = "Error"
			lResponseRec.ErrMsg = "Communication failure: " + lErr.Error()
			goto Complete // Jump to the completion step if communication fails
		}

		// Populate response structure with data
		lResponseRec.ResponseArr = lResponseArr
		lResponseRec.Status = "Success" // Set status to success if no errors
	}

Complete:
	// Finalize the response and send it back to the client
	CompleteMethod(lResponseRec, lHttpWriter)
	log.Println("SampleAPI (-)")
}

// 1. Collect: CollectMethod reads and unmarshals data from the HTTP request
func CollectMethod(lHttpRequest *http.Request, lRequestData *RequestStruct) error {
	log.Println("CollectMethod (+)")

	// Read the body of the request
	lBody, lErr := ioutil.ReadAll(lHttpRequest.Body)
	if lErr != nil {
		return errors.New("failed to read request body") // Return error if reading the body fails
	}

	// Unmarshal the request body into the RequestStruct
	if lErr = json.Unmarshal(lBody, &lRequestData); lErr != nil {
		return errors.New("failed to unmarshal request data") // Return error if unmarshaling fails
	}

	log.Println("CollectMethod (-)")
	return nil // Return nil if no errors
}

// 2. Validate: ValidateMethod validates the request data using struct tags
func ValidateMethod(lRequestData RequestStruct) error {
	log.Println("ValidateMethod (+)")

	// Initialize the validator
	lValidate := validator.New()

	// Validate the request data
	lErr := lValidate.Struct(lRequestData)
	if lErr != nil {
		return errors.New("request validation failed") // Return error if validation fails
	}

	log.Println("ValidateMethod (-)")
	return nil // Return nil if validation is successful
}

// 3. Construct: ConstructMethod processes data and constructs batch information
func ConstructMethod(lRequestData RequestStruct) (string, error) {
	log.Println("ConstructMethod (+)")

	// Example logic to create a batch ID based on request data
	lBatchId := fmt.Sprintf("BATCH-%s-%d", lRequestData.Param1, lRequestData.Param2)

	// Example logic for batch ID suffix based on Param2 value
	if lRequestData.Param2 > 10 {
		lBatchId += "-HIGH"
	} else {
		lBatchId += "-LOW"
	}

	log.Println("ConstructMethod (-)")
	return lBatchId, nil // Return batch ID and no error
}

// 4. Communicate: CommunicateMethod interacts with other services to retrieve data
// CommunicateMethod calls SelectRecordsMethod to retrieve data from the database
// and processes the result before returning it.
func CommunicateMethod(pParameterName string) (lResponseArr []ResponseArr, lErr error) {
	// Log the start of the CommunicateMethod function
	log.Println("CommunicateMethod (+)")
	lData := "asd"
	// Call SelectRecordsMethod to fetch data based on the parameter
	// lData, lErr := SelectRecordsMethod(pParameterName)
	if lErr != nil {
		// Log and return an error if SelectRecordsMethod fails
		log.Println("CommunicateMethod: Error in SelectRecordsMethod", lErr.Error())
		return lResponseArr, fmt.Errorf("CommunicateMethod - Error: %s", lErr.Error())
	}

	// Simulate fetching data, in a real scenario this could be a database query or an API call
	lResponseArr = []ResponseArr{
		{Field1: lData, Field2: 1}, // Example response item 1
		{Field1: lData, Field2: 2}, // Example response item 2
	}

	// Log the successful completion of the CommunicateMethod function
	log.Println("CommunicateMethod (-)")

	// Return the fetched data and no error
	return lResponseArr, nil
}

// 5. Complete: CompleteMethod sends the response back to the frontend
func CompleteMethod(lResponseRec ResponseStruct, lHttpWriter http.ResponseWriter) {
	log.Println("CompleteMethod (+)")

	// Marshal the response struct into JSON
	lData, lErr := json.Marshal(lResponseRec)
	if lErr != nil {
		// If marshalling fails, return an error response
		http.Error(lHttpWriter, "Error marshaling response data: "+lErr.Error(), http.StatusInternalServerError)
		return
	}

	// Send a successful response with status 200
	lHttpWriter.WriteHeader(http.StatusOK)
	lHttpWriter.Write(lData)

	log.Println("CompleteMethod (-)")
}

// SelectRecordsMethod retrieves records from the database based on a given parameter.
// It connects to the database, prepares a query, and returns the result.
func SelectRecordsMethod(pParameterName string) (lVariable string, lErr error) {
	// Log the start of the SelectRecordsMethod function
	log.Println("SelectRecordsMethod (+)")

	// Step 1: Establish a database connection
	// Replace with your actual database connection details
	lDb, lErr := ftdb.LocalDbConnect("main")
	if lErr != nil {
		// Log and return an error if the database connection fails
		log.Println("ASRM:001", lErr.Error())
		return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-001) " + lErr.Error())
	}
	defer lDb.Close() // Ensure the database connection is closed after the function finishes

	// Step 2: Prepare the SQL query
	// Replace this with your actual query for selecting records from the database
	lCoreString := `SELECT column_name FROM your_table WHERE parameter_column = ?`
	lStmt, lErr := lDb.Prepare(lCoreString)
	if lErr != nil {
		// Log and return an error if preparing the SQL statement fails
		log.Println("ASRM:002", lErr.Error())
		return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-002) " + lErr.Error())
	}
	defer lStmt.Close() // Ensure the statement is closed when done

	// Step 3: Execute the query with the provided parameter
	lRows, lErr := lStmt.Query(pParameterName)
	if lErr != nil {
		// Log and return an error if executing the query fails
		log.Println("ASRM:003", lErr.Error())
		return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-003) " + lErr.Error())
	}
	defer lRows.Close() // Ensure the result set is closed after iteration

	// Step 4: Process the rows returned by the query
	// Iterate over the rows and scan the data into the local variable
	for lRows.Next() {
		if lErr := lRows.Scan(&lVariable); lErr != nil {
			// Log and return an error if scanning the row fails
			log.Println("ASRM:004", lErr.Error())
			return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-004) " + lErr.Error())
		}
	}

	// Log the end of the SelectRecordsMethod function
	log.Println("SelectRecordsMethod (-)")

	// Return the retrieved variable and no error
	return lVariable, nil
}
