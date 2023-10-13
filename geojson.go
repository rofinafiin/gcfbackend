package gcfbackend

import "encoding/json"

func GCHandlerFunc(Mongostring, dbname, colname string) []byte {
	koneksyen := GetConnectionMongo(Mongostring, dbname)
	datageo := GetAllGeoData(koneksyen, colname)

	jsoncihuy, _ := json.Marshal(datageo)

	return jsoncihuy
}
