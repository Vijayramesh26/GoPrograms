package api

import (
	"database/sql"
	"fmt"
	"log"
)

// This method is used to fetch data for this purpose from this table.
func SelectRecordsMethod(pParameterName string) (lVariable string, lErr error) {
	// Log the start of the SelectRecordsMethod function
	log.Println("SelectRecordsMethod (+)")

	// Connect to the database using a connection key
	lDb, lErr := DBConnection("key")
	if lErr != nil {
		// Log an error message if the database connection fails
		log.Println("ASRM:001", lErr.Error())
		// Return the local variable and an error with a specific code and the error message
		return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-001) " + lErr.Error())
	} else {
		// Ensure the database connection is closed when the function exits
		defer lDb.Close()

		// Prepare the SQL statement to retrieve data from the table (replace with actual query)
		lCoreString := `// Enter Select Query`
		lStmt, lErr := lDb.Prepare(lCoreString)
		if lErr != nil {
			// Log an error message if the statement preparation fails
			log.Println("ASRM:002", lErr.Error())
			// Return the local variable and an error with a specific code and the error message
			return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-002) " + lErr.Error())
		} else {
			// Ensure the prepared statement is closed when the function exits
			defer lStmt.Close()

			// Execute the prepared statement with the provided parameter
			lRows, lErr := lStmt.Query(pParameterName)
			if lErr != nil {
				// Log an error message if the query execution fails
				log.Println("ASRM:003", lErr.Error())
				// Return the local variable and an error with a specific code and the error message
				return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-003) " + lErr.Error())
			} else {
				// Ensure the result set is closed when the function exits
				defer lRows.Close()

				// Process the result set
				for lRows.Next() {
					// Scan the row and store the result in the local variable
					lErr := lRows.Scan(&lVariable)
					if lErr != nil {
						// Log an error message if scanning the row fails
						log.Println("ASRM:004", lErr.Error())
						// Return the local variable and an error with a specific code and the error message
						return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-004) " + lErr.Error())
					} else {
						// Additional logic for processing the result can be added here
					}
				}
			}
		}
	}

	// Log the end of the SelectRecordsMethod function
	log.Println("SelectRecordsMethod (-)")

	// Return the local variable and no error
	return lVariable, nil
}

// This method is used to insert or update records based on the flag.
func InsertUpdateMethod(pParameterName, pFlag string) error {
	// Log the start of the InsertUpdateMethod function
	log.Println("InsertUpdateMethod (+)")

	// Establish a connection to the database using a connection key
	lDb, lErr := DBConnection("key")
	if lErr != nil {
		// Log an error message if the database connection fails
		log.Println("AIUM-001 ", lErr.Error())
		// Return an error with a specific code and the error message
		return fmt.Errorf("InsertUpdateMethod - (AIUM-001) " + lErr.Error())
	}
	// Ensure the database connection is closed when the function exits
	defer lDb.Close()

	var lCorestring string
	var lExecResult sql.Result

	// Prepare the SQL statement based on the flag (INSERT or UPDATE)
	switch {
	case pFlag == common.INSERT:
		// Define the SQL insert query string (replace with actual query)
		lCorestring = `Enter Insert Query`
		// Execute the SQL insert query with the provided parameter
		lExecResult, lErr = lDb.Exec(lCorestring, pParameterName)
	case pFlag == common.UPDATE:
		// Define the SQL update query string (replace with actual query)
		lCorestring = `Enter Update Query`
		// Execute the SQL update query with the provided parameter
		lExecResult, lErr = lDb.Exec(lCorestring, pParameterName)
	}

	// Check if there was an error executing the query
	if lErr != nil {
		// Log an error message if the query execution fails
		log.Println("AIUM-002 ", lErr.Error())
		// Return an error with a specific code and the error message
		return fmt.Errorf("InsertUpdateMethod - (AIUM-002) " + lErr.Error())
	} else {
		// Check the number of rows affected by the insert or update query
		rowsAffected, _ := lExecResult.RowsAffected()
		if lErr != nil {
			// Log an error message if fetching the affected rows count fails
			log.Println("AIUM-003 ", lErr.Error())
		} else {
			// Log the number of rows affected and a success message
			log.Printf("InsertUpdateMethod Rows affected: %d\n", rowsAffected)
			log.Println("Record Inserted or Updated successfully")
		}
	}

	// Log the end of the InsertUpdateMethod function
	log.Println("InsertUpdateMethod (-)")
	return nil
}

// This method is used to insert records into the database.
func InsertRecords(pParameterName string) error {
	// Log the start of the InsertRecords function
	log.Println("InsertRecords (+)")

	// Establish a connection to the database using a connection key
	lDb, lErr := DBConnection("key")
	if lErr != nil {
		// Log an error message if the database connection fails
		log.Println("AIR-001 ", lErr.Error())
		// Return an error with a specific code and the error message
		return fmt.Errorf("InsertRecords - (AIR-001) " + lErr.Error())
	} else {
		// Ensure the database connection is closed when the function exits
		defer lDb.Close()

		// Define the SQL insert query string (replace with actual query)
		lSqlString := `//Enter Insert Query`

		// Execute the SQL insert query with the provided parameter
		lExecResult, lErr := lDb.Exec(lSqlString, pParameterName)
		if lErr != nil {
			// Log an error message if the query execution fails
			log.Println("AIR-002 ", lErr.Error())
			// Return an error with a specific code and the error message
			return fmt.Errorf("InsertRecords - (AIR-002) " + lErr.Error())
		} else {
			// Check the number of rows affected by the insert query
			rowsAffected, lErr := lExecResult.RowsAffected()
			if lErr != nil {
				// Log an error message if fetching the affected rows count fails
				log.Println("AIR-003 ", lErr.Error())
			} else {
				// Log the number of rows affected and a success message
				log.Printf("InsertRecords Rows affected: %d\n", rowsAffected)
				log.Println("Record Inserted successfully")
			}
		}
	}

	// Log the end of the InsertRecords function
	log.Println("InsertRecords (-)")
	return nil
}
