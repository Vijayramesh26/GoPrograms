package getapistatus

import (
	"fmt"
	"goprograms/common"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

type ApiDetails struct {
	APIName    string
	APIURL     string
	StatusCode int
	StatusMsg  string
}
type ApiStatusResponse struct {
	ApiDetailsArr []ApiDetails `json:"apiDetailsArr"`
	IsApiDown     string       `json:"isApiDown"`
	Status        string       `json:"status"`
	ErrMsg        string       `json:"errMsg"`
}

/*
Purpose: This method serves as the main entry point for checking service statuses and generating a summary.
Parameter: None
Response:

    On Success:
        This function returns successfully with a ServiceResponse containing service details and status information.
    On Error:
        This function returns an error if any exception occurs during execution.

Author: VIJAY
Date: 30-APRIL-2024
*/

func APIServiceStatusMain() {
	log.Println("ServiceStatusMain (+)")
	// Initialize email record
	lRespRec, lErr := ReadnCheckSatus()
	if lErr != nil {
		// Handle error and log email summary
		lRespRec.ErrMsg = "GASASSM-001" + lErr.Error()
		lRespRec.Status = common.ErrorCode
		log.Println("GASASSM-001" + lErr.Error())
	} else {
		lRespRec.Status = common.SuccessCode
		if lRespRec.IsApiDown == "Y" {
			log.Println("All API is Down", lRespRec.ApiDetailsArr)
		} else {
			log.Println("All API is UP", lRespRec.ApiDetailsArr)
		}
	}
	log.Println("ServiceStatusMain (-)")
}

/*
Purpose: This method reads service details from a TOML configuration file and checks the status of each service.
Parameter: None
Response:

    On Success:
        This function returns successfully with a ServiceResponse containing service details and status information.
    On Error:
        This function returns an error if any exception occurs during execution.

Author: VIJAY
Date: 30-APRIL-2024
*/

func ReadnCheckSatus() (lApiRec ApiStatusResponse, lErr error) {
	log.Println("ReadnCheckSatus (+)")
	// Define a struct to hold the TOML service data
	var lApiDataRec struct {
		apiDetailsArr []ApiDetails
	}

	// Load TOML data from file
	_, lErr = toml.DecodeFile("./toml/serviceConfig.toml", &lApiDataRec)
	if lErr != nil {
		// Handle error and return
		log.Println("ReadnCheckSatus", common.ErrorCode, " GASRCS-001 "+lErr.Error())
		return lApiRec, fmt.Errorf(" GASRCS-001 -->" + lErr.Error())
	} else {
		// Iterate through each service and check status
		for _, lServerRec := range lApiDataRec.apiDetailsArr {
			StatusMsg, StatusCode, lErr := CheckServiceStatus(lServerRec.APIURL)
			if lErr != nil {
				// Handle error and return
				log.Println("ReadnCheckSatus", common.ErrorCode, " GASRCS-002 "+lErr.Error())
				return lApiRec, fmt.Errorf(" GASRCS-002 -->" + lErr.Error())
			} else {
				lServerRec.StatusCode = StatusCode
				lServerRec.StatusMsg = StatusMsg
				lApiRec.ApiDetailsArr = append(lApiRec.ApiDetailsArr, lServerRec)

				//send only when there is API down
				if StatusCode != http.StatusOK {
					lApiRec.IsApiDown = "Y"
				}
			}
		}
	}
	log.Println("ReadnCheckSatus (-)")
	return lApiRec, nil
}

/*
Purpose: This method checks the status of a given URL by making an HTTP GET request.
Parameter: pUrl - The URL to check the status of
Response:

    On Success:
        This function returns the status message, status code, and any potential error encountered during the HTTP request.
    On Error:
        This function returns an error if any exception occurs during execution.

Author: VIJAY
Date: 30-APRIL-2024
*/

func CheckServiceStatus(pUrl string) (lStatusMsg string, lStatusCode int, lErr error) {
	log.Println("CheckServiceStatus (+)")
	// Make HTTP GET request
	lResponse, lErr := http.Get(pUrl)
	if lErr != nil {
		// Handle error
		lStatusMsg = "Failed to reach URL"
		lStatusCode = 0
		//this is to avoid the upper program discontinue checking of next urls.
		lErr = nil
		log.Println("GASCSS-001 -->", lErr.Error())
	} else {
		// Close response body and check HTTP status code
		defer lResponse.Body.Close()
		if lResponse.StatusCode == http.StatusOK {
			lStatusMsg = "Service is up"
			lStatusCode = lResponse.StatusCode
		} else {
			lStatusMsg = "Service is down"
			lStatusCode = lResponse.StatusCode
		}
	}
	log.Println("CheckServiceStatus (-)")
	return
}
