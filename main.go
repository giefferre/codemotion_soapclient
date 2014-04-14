package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	// instancing a request
	request, err := NewSoapRequest()
	if err != nil {
		fmt.Println(err)
	}

	// set the request address, type and address
	request.SetRequest("http://www.webservicex.net/geoipservice.asmx", "GetGeoIp", "8.8.8.8")

	// run the request
	byteBody, err := request.Do()
	if err != nil {
		fmt.Println(err)
	}

	// get the response
	var response SoapGetGeoIpResponse

	// unmarshal
	err = xml.Unmarshal(byteBody, &response)
	if err != nil {
		fmt.Println(err)
	}

	// get the result object
	getGeoIpResult := response.Body.GetGeoIPResponse.GetGeoIPResponseResult

	// printing the output
	if getGeoIpResult.ReturnCode == "1" {
		fmt.Println(
			fmt.Sprintf("Results for IP %s:\n", getGeoIpResult.IP),
			fmt.Sprintf("\t- Country:\t %s \n", getGeoIpResult.CountryName),
			fmt.Sprintf("\t- Country Code:\t %s \n", getGeoIpResult.CountryCode),
		)
	}

}
