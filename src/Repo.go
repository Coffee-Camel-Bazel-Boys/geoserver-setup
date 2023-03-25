package mapController

import (
	"database/sql"
	"strconv"
	"fmt"
	"encoding/json"
	
	_ "github.com/lib/pq"
)

type Feature struct {
	Type string `json:"type"`;
	Id string `json:"id"`;
	Geometry Geometry `json:"geometry"`;
	Geometry_name string `json:"geometry_name"`;
	Properties map[string]interface{} `json:"properties"`;
}

type Geometry struct {
	Type string `json:"type"`;
	Coordinates  [][][][]float64 `json:"coordinates"` 
}

// FIXME
const parcelQuery = `SELECT ogc_fid, 
						ST_AsGeoJSON(wkb_geometry) AS geometry,
						b.boundary
					FROM regina_parcel, 
						(select ST_MakeEnvelope($1,$2,$3,$4)::geography AS boundary) as b
					WHERE ST_DWITHIN(wkb_geometry, boundary, 0);`;

const crimeMapQuery = `SELECT 
						ogc_fid, 
						ST_AsGeoJSON(wkb_geometry) AS geometry,
						b.boundary,
						row_to_json(crime_statistics) AS data,
						community.name
					FROM community
						LEFT JOIN crime_statistics ON community.name = crime_statistics.community,
							(select ST_MakeEnvelope($1,$2,$3,$4)::geography AS boundary) as b
						WHERE ST_DWITHIN(wkb_geometry, boundary, 0)`;

func GetParcelData(bbox []string) []Feature {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
	HOST, PORT, USER, POSTGRES_PASSWORD, DATABASE);
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err);
	}

	// FIXME, cleanup
	b0, err := strconv.ParseFloat(bbox[0], 64); 
	b1, err := strconv.ParseFloat(bbox[1], 64); 
	b2, err := strconv.ParseFloat(bbox[2], 64); 
	b3, err := strconv.ParseFloat(bbox[3], 64);

	rows, err := db.Query(parcelQuery, b0, b1, b2, b3);

	if (err != nil) {
	  panic(err);
	}
	defer rows.Close();

	var parcels []Feature;
	for rows.Next() {
		var ogc_fid string;
		var geometry []byte;
		var boundary string;
		err := rows.Scan(&ogc_fid, &geometry, &boundary);
		if (err != nil) {
			// FIXME
			panic(err);
		}
		parcels = append(parcels, getFeature(ogc_fid, geometry, nil));
	}

	err = rows.Err()
	if err != nil {
	  panic(err)
	}

	defer db.Close();
	err = db.Ping();
	if err != nil {
	  panic(err);
	}

	return parcels;
}


func GetCrimeData(bbox []string) []Feature {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
	HOST, PORT, USER, POSTGRES_PASSWORD, DATABASE);
	
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err);
	}

	// FIXME, cleanup
	b0, err := strconv.ParseFloat(bbox[0], 64); 
	b1, err := strconv.ParseFloat(bbox[1], 64); 
	b2, err := strconv.ParseFloat(bbox[2], 64); 
	b3, err := strconv.ParseFloat(bbox[3], 64);

	rows, err := db.Query(crimeMapQuery, b0, b1, b2, b3);

	if (err != nil) {
	  panic(err);
	}
	defer rows.Close();

	var crimeFeatures []Feature;
	for rows.Next() {
		var ogc_fid string;
		var geometry []byte;
		var boundary string;
		var data []byte;
		var communityName string;

		err := rows.Scan(&ogc_fid, &geometry, &boundary, &data, &communityName);
		if (err != nil) {
			// FIXME
			panic(err);
		}

		crimeFeatures = append(crimeFeatures, getFeature(ogc_fid, geometry, getDataValue(data)));
	}

	err = rows.Err()
	if err != nil {
	  panic(err)
	}

	db.Close();

	return crimeFeatures;
}

func getFeature(ogc_fid string, geometry []byte, data map[string]interface{}) Feature {
	var typedGeometry Geometry;
	json.Unmarshal(geometry, &typedGeometry);
	return Feature{
		Type: "Feature",
		Id: "feature." + ogc_fid,
		Geometry: typedGeometry,
		Geometry_name: "the_geom",
		Properties: data};
}

func getDataValue(value []byte) map[string]interface{} {
	var parsedData map[string]interface{};
	json.Unmarshal(value, &parsedData);
	return parsedData;
}


