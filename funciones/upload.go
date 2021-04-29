package funciones

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/danielaotero17/jsonManage"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

func GetDgraphClient() (*dgo.Dgraph, CancelFunc) {

	direccion := "localhost:9080"
	conn, err := grpc.Dial(direccion, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	dc := api.NewDgraphClient(conn)
	dgraphClient := dgo.NewDgraphClient(dc)

	if err != nil {
		log.Fatalf("While trying to login %v", err.Error())
	}

	return dgraphClient, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func ConectarBd(data []byte) {

	dg, cancel := GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	mu := &api.Mutation{
		SetJson: data,
	}

	res, err2 := txn.Mutate(ctx, mu)
	if err2 != nil {
		log.Fatal("Hola", err2)
	}

	err3 := txn.Commit(context.Background())
	if err3 != nil {
		log.Fatal(err3)
	}

	log.Println(string(res.Json))

}

func Upload_data(data int64) {

	date := strconv.FormatInt(data, 16)

	data_compradores := ReadBuyers(date)
	dataBuyers := jsonManage.DecodejsonBuyers(data_compradores)
	dataProductos := ReadProducts(date)
	dataTransacciones := ReadTransactions(date)

	bytesBuyers, err1 := json.Marshal(dataBuyers)
	bytesProducts, err2 := json.Marshal(dataProductos)
	bytesTransaction, err3 := json.Marshal(dataTransacciones)

	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	if err3 != nil {
		fmt.Println(err3)
	}

	//colocar subrutinas de go al principio del llamado de un método, me permite crear hilos.
	go ConectarBd(bytesBuyers)
	ConectarBd(bytesProducts)
	ConectarBd(bytesTransaction)
}
