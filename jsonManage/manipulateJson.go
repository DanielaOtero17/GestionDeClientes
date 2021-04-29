package jsonManage

import (
	"encoding/json"
	"fmt"

	"github.com/danielaotero17/modelo"
)

func DecodejsonProduct(datos []byte) []modelo.Product {
	fmt.Println("\nDECODIFICANDO PRODUCTOS ... ")

	var arreglo []modelo.Product

	errorsito := json.Unmarshal(datos, &arreglo)

	if errorsito != nil {
		fmt.Printf("Error decodificando: %v\n", errorsito)
	}

	return arreglo
}

func DecodejsonTransactions(datos []byte) []modelo.Transaction {
	fmt.Println("\nDECODIFICANDO TRANSACCIONES ... ")

	var arreglo []modelo.Transaction

	errorsito := json.Unmarshal(datos, &arreglo)

	if errorsito != nil {
		fmt.Printf("Error decodificando: %v\n", errorsito)
	}

	return arreglo
}

func DecodejsonBuyers(datos []byte) []modelo.Buyer {
	fmt.Println("\nDECODIFICANDO COMPRADORES ... ")

	// Se define la variable que alojará los valores decodificado
	var arreglo []modelo.Buyer

	// Y ahora decodificamos pasando el apuntador, utilizando unmarshal
	errorsito := json.Unmarshal(datos, &arreglo)

	if errorsito != nil {
		fmt.Printf("Error decodificando: %v\n", errorsito)
	}

	return arreglo
}
