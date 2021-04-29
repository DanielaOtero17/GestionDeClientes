package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/danielaotero17/funciones"
	"github.com/danielaotero17/jsonManage"
	"github.com/danielaotero17/modelo"
)

/*Esta funcion retorna un arreglo byte que contiene la información de los productos recomendados.
La recomendación se hacen a partir de los productos que han comprado otros compradores que han usado la misma ip.
Esto es porque si dos personas utilizan la misma ip, entonces se encuentran relacionadas en cuanto al lugar.
(ciudad, país, zona geográfica, barrio, casa). Teniendo en cuenta que el lugar en el que estén dos personas,
influye en sus gustos y necesidades, puede que los productos que uno de los dos compra, le interese al otro.
Por eso, a mi parecer, tomar en cuenta la ip sería una buena idea de "recomendación"
*/
func recomendaciones(id string) []byte {

	cadena, compradoresRelacionados := OtrosCompradoresMismaIp(id)

	print(cadena)

	var arregloCompradores []modelo.Buyer

	err2 := json.Unmarshal(compradoresRelacionados, &arregloCompradores)

	if err2 != nil {
		log.Println("Ha ocurrido un error al hacer unmarshall a los compradores relacionados")
	}

	var productosRecomendados []modelo.Product
	var idProductos []string

	//var idAuxes []string
	for i := 0; i < len(arregloCompradores); i++ {

		response := ComprasBuyerId(arregloCompradores[i].Id)

		transacciones := jsonManage.DecodejsonTransactions([]byte(JsonTransacciones(response)))

		for j := 0; j < len(transacciones); j++ {

			for p := 0; p < len(transacciones[j].ProductIds); p++ {

				//idAuxes = append(idAuxes, transacciones[j].ProductIds[p])

				if existeString(transacciones[j].ProductIds[p], idProductos) == false {

					idProductos = append(idProductos, transacciones[j].ProductIds[p])
				}
			}
		}
	}

	productosRecomendados = queryProducts(idProductos)

	recomendadosJson, err := json.Marshal(productosRecomendados)
	if err != nil {
		log.Println("Ha ocurrido un error al hacer marshall a los productos recomendados")
	}

	cadenaRecomendados := "{" + string(34) + "ProductosRecomendados" + string(34) + ": " + string(recomendadosJson) + "}"

	return []byte(cadenaRecomendados)

}

func existeProducto(producto modelo.Product, arreglo []modelo.Product) bool {

	for i := 0; i < len(arreglo); i++ {
		if arreglo[i].Id == producto.Id {
			return true
		}
	}
	return false
}

func existeString(dato string, arreglo []string) bool {

	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] == dato {
			return true
		}
	}
	return false

}

func foundTransactionByBuyer(id string) []modelo.Transaction {

	dg, cancel := funciones.GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(
		`
		{
			HistorialCompras(func:has(buyerId))@filter(eq(buyerId, %s)){
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
		println("found transaction by Ip. Ha ocurrido un error al consultar la transaccion en la base de datos: ", err)
	}

	transaccion := jsonManage.DecodejsonTransactions([]byte(JsonTransacciones(res)))

	return transaccion
}
