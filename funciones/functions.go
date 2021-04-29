package funciones

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/danielaotero17/modelo"
)

// para que las funciones sean públicas, deben empezar con mayúsculas.

func readData(link string) []byte {
	// Haciendo la petición para leer los datos con las url

	respuesta, err := http.Get(link)

	if err != nil {
		fmt.Println("error obteniendo respuesta: ", err)
	}

	// se cierra el cuerpo al terminar
	defer respuesta.Body.Close()

	cuerpoRespuesta, err := ioutil.ReadAll(respuesta.Body)

	if err != nil {
		fmt.Println("error leyendo respuesta: ", err)
	}

	return cuerpoRespuesta
}

func ReadBuyers(date string) []byte {

	fmt.Println("\nLEYENDO COMPRADORES... ")
	linkbuyers := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers?date=" + date

	return readData(linkbuyers)
}

func importAll() {

}

func ReadProducts(date string) []modelo.Product {

	fmt.Println("\nLEYENDO PRODUCTOS ")

	linkproducts := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products?date=" + date

	datos := string(readData(linkproducts))

	// arreglo que contiene las lineas del archivo
	arreglo := strings.Split(datos, "\n")

	var productos []modelo.Product
	var datosNoAgregados []string

	for i := 0; i < len(arreglo); i++ {

		var arrayAux []string
		var nuevoProd modelo.Product

		if strings.Contains(arreglo[i], string(34)) {

			splitComillas := strings.Split(arreglo[i], string(34))

			parteId := strings.TrimRight(splitComillas[0], "'")
			partePrice := strings.TrimLeft(splitComillas[2], "'")

			price, err := strconv.Atoi(partePrice)

			if err != nil {
				print("No se ha podido leer el precio: ", err)
			}
			nuevoProd = modelo.Product{parteId, splitComillas[1], price}

		} else {

			if strings.Contains(arreglo[i], "'") {
				arrayAux = strings.Split(arreglo[i], "'")

				price, err := strconv.Atoi(arrayAux[2])

				if err != nil {

					fmt.Println("error leyendo el precio", err)
					datosNoAgregados = append(datosNoAgregados, arreglo[i])
				}
				nuevoProd = modelo.Product{arrayAux[0], arrayAux[1], price}
			}
		}
		productos = append(productos, nuevoProd)
	}

	return productos

}

func ReadTransactions(date string) []modelo.Transaction {

	fmt.Println("\nLEYENDO TRANSACCIONES ")

	linktransactions := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions?date=" + date

	datos := string(readData(linktransactions))

	// arreglo que contiene las lineas del archivo, este archivo tiene espacios nulls
	arreglo := strings.Split(datos, "\000\000")

	var trans []modelo.Transaction

	for i := 0; i < len(arreglo); i++ {

		item := strings.Split(arreglo[i], "\000")

		if len(item) == 5 {

			ids := strings.TrimLeft(item[4], "(")
			ids = strings.TrimRight(ids, ")")

			productsIds := strings.Split(ids, ",")

			laTrans := modelo.Transaction{item[0], item[1], item[2], item[3], productsIds}

			trans = append(trans, laTrans)
		}

	}
	return trans
}
