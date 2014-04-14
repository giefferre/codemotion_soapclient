package main

import (
	"encoding/xml"
	"errors"
	"github.com/franela/goreq"
	"io/ioutil"
	"time"
)

// REQUEST
type SoapRequest struct {
	address           string
	XMLName           xml.Name `xml:"soapenv:Envelope"`
	SoapenvNamespace1 string   `xml:"xmlns:soapenv,attr"`
	SoapenvNamespace2 string   `xml:"xmlns:web,attr"`
	Header            RequestHeader
	Body              RequestBody
}

type RequestHeader struct {
	XMLName xml.Name `xml:"soapenv:Header"`
}

type RequestBody struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Content RequestContent
}

type RequestContent interface {
}

type RequestContentGetGeoIp struct {
	XMLName   xml.Name `xml:"web:GetGeoIP"`
	IpAddress string   `xml:"web:IPAddress"`
}

func NewSoapRequest() (*SoapRequest, error) {
	request := new(SoapRequest)
	request.SoapenvNamespace1 = "http://schemas.xmlsoap.org/soap/envelope/"
	request.SoapenvNamespace2 = "http://www.webservicex.net/"

	return request, nil
}

func (this *SoapRequest) SetRequest(serviceAddress string, contentType string, content string) error {

	this.address = serviceAddress

	switch contentType {
	case "GetGeoIp":
		requestContent := new(RequestContentGetGeoIp)
		requestContent.IpAddress = content
		this.Body.Content = requestContent
		break
	default:
		return errors.New("Unrecognized Request Type: " + contentType)
	}

	return nil
}

func (this *SoapRequest) Do() ([]byte, error) {
	formattedXml, err := xml.Marshal(this)
	if err != nil {
		return nil, err
	}

	httpResponse, err := goreq.Request{
		Method:      "POST",
		Uri:         this.address,
		ContentType: "text/xml;charset=UTF8",
		Body:        formattedXml,
		Timeout:     30 * time.Second,
	}.Do()
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != 200 {
		return nil, errors.New("Unable to retrieve status")
	}

	byteBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	return byteBody, nil
}
