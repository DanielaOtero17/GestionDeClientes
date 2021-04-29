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

func OtrosCompradoresMismaIp(id string) ([]byte, []byte) {

	log.Println("Entrando a buscar compradores misma ip")

	jsonResponseTransactions := ComprasBuyerId(id)

	responseOrganizada := JsonTransacciones(jsonResponseTransactions)

	transEncontradas := jsonManage.DecodejsonTransactions([]byte(responseOrganizada))

	var ips []string
	var buyersCoincidencias []modelo.Buyer

	for i := 0; i < len(transEncontradas); i++ {

		ips = append(ips, transEncontradas[i].Ip)
	}

	for i := 0; i < len(ips); i++ {

		//entra a consultar la transaccion con la ip relacionaa, y la guarda en el objeto transaccion
		transaccion := foundTransactionsByIp(ips[i])

		id_OtroComprador := transaccion.BuyerId

		//entra a consultar el comprador que ha comprado con esa ip, teniendo en cuenta su id
		comprador := foundBuyerById(id_OtroComprador)

		/*
			 verifica que el id del comprador encontrado, sea diferente del comprador que se busca, para evitar
			que el comprador que se busca en el endpoint se incluya entre los que se buscan.
			Dentro del if se agrega el comprador que tenga la misma ip, se hace mediante un filtrado,
			esto es para evitar que la información del comprador se repita, ya que puede estar en diferentes
			transacciones con la misma ip.
		*/
		if existeComprador(comprador, buyersCoincidencias) == false && comprador.Id != id {

			buyersCoincidencias = append(buyersCoincidencias, comprador)
		}
	}

	// Este es el formato json de los compradores que han tenido la misma ip.
	compradoresJsonIp, err := json.Marshal(buyersCoincidencias)

	if err != nil {
		fmt.Println("Ha ocurrido un error al hacer marshal de los compradores con ip similar: ", err)
	}

	// cadena es el formato personalizado que le doy, para que se muestre organizadamente.
	cadena := "{" + string(34) + "otrosCompradores" + string(34) + ":"
	cadena += string(compradoresJsonIp)
	cadena += "}"

	if err != nil {
		log.Println("Ha ocurrido un error al hacer marshal el formatoJson de compradores ip relacionados", err)
	}

	return []byte(cadena), compradoresJsonIp
}

func foundBuyerById(id string) modelo.Buyer {

	dg, cancel := funciones.GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(
		`
		{
			Comprador(func:has(id))@filter(eq(id, %s)){
				id
				name
				age
			}
	} 
	`, id)

	res, err := txn.Query(ctx, q)

	if err != nil {
		log.Println("found buyer. Ha ocurrido un error al consultar el comprador en la base de datos: ", err)
	}
	comprador := jsonManage.DecodejsonBuyers([]byte(JsonBuyers(res)))

	return comprador[0]
}

func existeComprador(comprador modelo.Buyer, arreglo []modelo.Buyer) bool {

	for i := 0; i < len(arreglo); i++ {
		if arreglo[i].Id == comprador.Id {
			return true
		}
	}
	return false
}

func foundTransactionsByIp(ip string) modelo.Transaction {

	dg, cancel := funciones.GetDgraphClient()
	defer cancel()

	ctx := context.Background()

	txn := dg.NewTxn()
	defer txn.Discard(ctx)

	q := fmt.Sprintf(
		`
		{
			HistorialCompras(func:has(ip))@filter(eq(ip, %s)){
				id_trans
				buyerId
				ip
				device
				productIds
			}
	} 
	`, ip)

	res, err := txn.Query(ctx, q)

	if err != nil {
		println("found transaction by Ip. Ha ocurrido un error al consultar la transaccion en la base de datos: ", err)
	}

	transaccion := jsonManage.DecodejsonTransactions([]byte(JsonTransacciones(res)))

	return transaccion[0]

}
