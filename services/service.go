package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/danielaotero17/funciones"
	"github.com/danielaotero17/jsonManage"
	"github.com/danielaotero17/modelo"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/go-chi/chi/v5"
)

func GetUrlParam(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	w.Write([]byte(id))
}

func CodeServer(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("welcome"))

}

func ImportWithoutDate(w http.ResponseWriter, r *http.Request) {
	fechaActual := time.Now().Unix()
	funciones.Upload_data(fechaActual)
	w.Write([]byte("{message: 'se han agregado los datos con la fecha actual'}"))
}

func Import(w http.ResponseWriter, r *http.Request) {

	date := chi.URLParam(r, "date")

	s, err := strconv.ParseInt(date, 10, 64)

	if err != nil {
		log.Print("ha ocurrido un error al parsear el string", err)
	}

	funciones.Upload_data(s)
	w.Write([]byte("{message: 'se han agregado los datos con la fecha dada'}"))

}

func Buyers(w http.ResponseWriter, r *http.Request) {
	dg, cancel := funciones.GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := `
	{
	  listadoCompradores(func: has(id)) {
		id
		name
		age
	  }
	} 
	`
	res, err := txn.Query(ctx, q)

	if err != nil {
		println("Ha ocurrido un error al consultar los compradores de la base de datos: ", err)
	} else {
		w.Write(res.Json)
	}

}

func BuyersById(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	dg, cancel := funciones.GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(
		`
	{
	  listadoCompradores(func: has(id)) @filter(eq(id, %s)){
		id
		name
		age
	  }
	} 
	`, id)

	res, err := txn.Query(ctx, q)

	if err != nil {
		println("service.ByersById - Ha ocurrido un error al consultar el comprador en la base de datos: ", err)

	} else {

		// Cadena, es la variable que guarda el formato de res.json estándar (vacío)
		cadena := "{" + string(34) + "listadoCompradores" + string(34) + ":[]}"

		//Si res.json no está vacío, entonces procede a buscar la información relacionada con el comprador suministrado.
		if string(res.Json) != (cadena) {

			cadenabyte, respu := OtrosCompradoresMismaIp(id)

			respuesta := "{" + string(34) + "InformacionComprador" + string(34) + ":["
			respuesta += string(productsByBuyer(id)) + ","
			respuesta += string(cadenabyte) + ","
			respuesta += string(recomendaciones(id))
			respuesta += "]}"

			log.Println(cadenabyte, respu)

			w.Write([]byte(respuesta))

		} else {
			w.Write([]byte("{message: 'No se ha encontrado ningún comprador con el id suministrado'}"))

		}
	}

}

func ComprasBuyerId(id string) *api.Response {

	dg, cancel := funciones.GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(
		`
		{
			HistorialCompras(func:has(id_trans))@filter(eq(buyerId, %s)){
				id_trans
				buyerId
				ip
				device
				productIds
			}
	} 
	`, id)

	res, err := txn.Query(ctx, q)

	if err != nil {
		log.Println("service. ComprasBuyerId - Ha ocurrido un error al consultar el comprador en la base de datos: ", err)
	}

	return res

}

func productsByBuyer(idBuyer string) []byte {

	productsByte := ComprasBuyerId(idBuyer)

	transEncontradas := jsonManage.DecodejsonTransactions([]byte(JsonTransacciones(productsByte)))

	type InstanciaCompra struct {
		Transaccion        modelo.Transaction `json: "trans_compra,omitempty"`
		ProductosComprados []modelo.Product   `json: "prods_compra,omitempty"`
	}

	var instanciasCompras []InstanciaCompra

	for i := 0; i < len(transEncontradas); i++ {

		var instanciaInterna InstanciaCompra

		idProductos := transEncontradas[i].ProductIds

		instanciaInterna.Transaccion = transEncontradas[i]

		auxi := queryProducts(idProductos)

		for j := 0; j < len(auxi); j++ {
			instanciaInterna.ProductosComprados = append(instanciaInterna.ProductosComprados, auxi[j])
		}

		instanciasCompras = append(instanciasCompras, instanciaInterna)
	}

	dataInstancias, err := json.Marshal(instanciasCompras)

	formatoInit := "{" + string(34) + "historialCompras" + string(34) + ":"
	formatoInit += string(dataInstancias)
	formatoInit += "}"

	if err != nil {
		fmt.Println("Error al hacer marshal a la data de instancias: ", err)
	}

	log.Println(dataInstancias)

	return []byte(formatoInit)
}

func queryProducts(idProductos []string) []modelo.Product {

	var productosEncontrados []modelo.Product

	for i := 0; i < len(idProductos); i++ {

		dg, cancel := funciones.GetDgraphClient()
		defer cancel()

		ctx := context.Background()

		txn := dg.NewTxn()
		defer txn.Discard(ctx)

		q := fmt.Sprintf(

			`
			{
				Producto(func: has(id_p))@filter(eq(id_p, %s)){
				  id_p
				  name_p
				  price
				}
			  } 
		`, idProductos[i])

		res, err := txn.Query(ctx, q)

		if err != nil {
			println("Ha ocurrido un error al consultar el producto en la base de datos: ", err)
		} else {

			stri := "{" + string(34) + "Producto" + string(34) + ":"

			//hago split al res.json, para ponerlo en el formato que deso
			split1 := strings.Split(string(res.Json), stri)

			//hago un trim para quitar el sufijo "}"
			datosProds := strings.TrimSuffix(split1[1], "}")

			foundProducts := jsonManage.DecodejsonProduct([]byte(datosProds))

			productosEncontrados = append(productosEncontrados, foundProducts[0])
		}
	}

	return productosEncontrados
}

func JsonTransacciones(res *api.Response) string {

	stri := "{" + string(34) + "HistorialCompras" + string(34) + ":"

	split1 := strings.Split(string(res.Json), stri)

	datosTrans := strings.TrimSuffix(split1[1], "}")

	return datosTrans

}

func JsonBuyers(res *api.Response) string {

	stri := "{" + string(34) + "Comprador" + string(34) + ":"

	split1 := strings.Split(string(res.Json), stri)

	datosComp := strings.TrimSuffix(split1[1], "}")

	return datosComp

}
