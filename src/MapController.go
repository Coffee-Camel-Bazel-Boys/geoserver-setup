package mapController

import (
	"strings"
	"net/http"
	"time"
	"encoding/json"
	_ "fmt"
)

type GeoRequest struct {
	service string;
	version string;
	outputFormat string;
	typeName string;
	srsName string;
	bbox []string;

}

type GeoResponse struct {
	GeoType string `json:"type"`;
	Features []Feature `json:"features"`;
	TotalFeatures int `json:"totalFeatures"`;
	NumberMatched int `json:"numverMatched"`;
	NumberReturned int `json:"numberReturned"`;
	TimeStamp time.Time `json:"timeStamp"`;
	Crs string `json:"crs"`;
}

func HandleRequest(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		version := request.URL.Query().Get("version");
		if (version != "1.1.0") {
			writer.WriteHeader(http.StatusNotImplemented);
			writer.Write([]byte("501 - we only support 1.1.0 here :)"));
			return;
		}
		geoRequest := initializeGeoRequest(request)
		writer.Write(geoData(geoRequest));
	default:
		writer.WriteHeader(http.StatusNotImplemented)
	    writer.Write([]byte("501 - Nope, try a get"))
	}

}

func initializeGeoRequest(request *http.Request) GeoRequest {
	return GeoRequest {
		service: request.URL.Query().Get("service"),
		version: request.URL.Query().Get("version"),
		outputFormat: request.URL.Query().Get("outputFormat"),
		typeName: request.URL.Query().Get("typename"),
		srsName: request.URL.Query().Get("srsName"),
		bbox: strings.Split(request.URL.Query().Get("bbox"), ",")};
}

func geoData(geoRequest GeoRequest) []byte {

	var features []Feature;
	features = getSection(geoRequest);

	response := GeoResponse{
		GeoType: "FeatureCollection",
		Features: features,
		TotalFeatures: len(features),
		NumberMatched: len(features),
		NumberReturned: len(features),
		TimeStamp: time.Now(),
		Crs: ""};

	responseJson, error := json.Marshal(response);
	
	if (error != nil) {
		// TODO!
		panic(error)
		return []byte("TODO");
	}
	
	return responseJson;
}

func getSection(geoRequest GeoRequest) []Feature {
	if(geoRequest.typeName == "parcel") {
		return GetParcelData(geoRequest.bbox);
	}
	// FIXME, 404
	return nil;
}