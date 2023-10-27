package gcfbackend

import (
	"context"
	"fmt"
	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/whatsauth/watoken"
	"testing"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	//privateKey, publicKey := watoken.GenerateKey()
	//fmt.Println(privateKey)
	//fmt.Println(publicKey)
	hasil, err := watoken.Encode("rofiasoy", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "cihuy",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGO", "GIS", context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "cihuy",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGO", "GIS", context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	conn := GetConnectionMongo("MONGO", "GIS")
	pass, _ := pasproj.HashPass("rofiganteng")
	data := RegisterStruct{
		Username: "rofiasoy",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}
