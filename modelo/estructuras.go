package modelo

// NOTA: Los nombres de las variables, deben empezar con mayúscula. Y dentro del Json, deben ser en minúscula.
type Product struct {
	Id    string `json:"id_p,omitempty"`
	Name  string `json:"name_p,omitempty"`
	Price int    `json:"price,omitempty"`
}

type Buyer struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

type Transaction struct {
	Id         string   `json:"id_trans,omitempty"`
	BuyerId    string   `json:"buyerId,omitempty"`
	Ip         string   `json:"ip,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIds []string `json:"productIds,omitempty"`
}
