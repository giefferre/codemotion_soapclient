package main

import (
	"encoding/xml"
)

// RESPONSE
// generic
type SoapGenericResponse struct {
	XMLName       xml.Name `xml:"Envelope"`
	XSINamespace  string   `xml:"xmlns:xsi,attr"`
	XSDNamespace  string   `xml:"xmlns:xsd,attr"`
	SoapNamespace string   `xml:"xmlns:soap,attr"`
}

// GetGeoIpResponse
type SoapGetGeoIpResponse struct {
	SoapGenericResponse
	Body SoapGetGeoIpBodyResponse
}

type SoapGetGeoIpBodyResponse struct {
	XMLName          xml.Name `xml:"Body"`
	GetGeoIPResponse SoapGetGeoIpResponseBody
}

type SoapGetGeoIpResponseBody struct {
	XMLName                xml.Name `xml:"GetGeoIPResponse"`
	Namespace              string   `xml:"xmlns,attr"`
	GetGeoIPResponseResult SoapGetGeoIpResponseResult
}

type SoapGetGeoIpResponseResult struct {
	XMLName           xml.Name `xml:"GetGeoIPResult"`
	ReturnCode        string   `xml:"ReturnCode"`
	IP                string   `xml:"IP"`
	ReturnCodeDetails string   `xml:"ReturnCodeDetails"`
	CountryName       string   `xml:"CountryName"`
	CountryCode       string   `xml:"CountryCode"`
}
