package gcfbackend

import (
	"context"
	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertUserdata(MongoConn *mongo.Database, username, password string) (InsertedID interface{}) {
	req := new(RegisterStruct)
	req.Username = username
	req.Password = password
	return pasproj.InsertOneDoc(MongoConn, "user", req)
}

func UpdateNameGeo(Mongoenv, dbname string, ctx context.Context, val LonLatProperties) (UpdateID interface{}) {
	conn := GetConnectionMongo(Mongoenv, dbname)
	filter := bson.D{{"volume", val.Volume}}
	update := bson.D{{"$set", bson.D{
		{"name", val.Name},
	}}}
	res, err := conn.Collection("lonlatpost").UpdateOne(ctx, filter, update)
	if err != nil {
		return "Gagal Update"
	}
	return res
}

func DeleteDataGeo(Mongoenv, dbname string, ctx context.Context, val LonLatProperties) (DeletedId interface{}) {
	conn := GetConnectionMongo(Mongoenv, dbname)
	filter := bson.D{{"volume", val.Volume}}
	res, err := conn.Collection("lonlatpost").DeleteOne(ctx, filter)
	if err != nil {
		return "Gagal Delete"
	}
	return res
}

func IsExist(Tokenstr, PublicKey string) bool {
	id := watoken.DecodeGetId(PublicKey, Tokenstr)
	if id == "" {
		return false
	}
	return true
}

func GetCoordinateNear(MongoConn *mongo.Database, colname string, coordinate []float64) (result []GeoJson, err error) {
	filter := bson.M{"geometry.coordinates": bson.M{
		"$near": bson.M{
			"$geometry": bson.M{
				"type":        "LineString",
				"coordinates": coordinate,
			},
		},
	}}
	curr, err := MongoConn.Collection(colname).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer curr.Close(context.Background())
	err = curr.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
