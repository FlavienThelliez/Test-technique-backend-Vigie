package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Orders struct which contains
// an array of orders
type Orders struct {
	Orders []Order `json:"orders"`
}

// Order struct which contains all 5 informations about an order.
type Order struct {
	ID           string  `json:"id"`
	Marketplace  string  `json:"marketplace"`
	Country      string  `json:"country"`
	Amount_cents float32 `json:"amount_cents"`
	Created_at   string  `json:"created_at"`
}

// Une fonction récupérant tout les "marketplace" et nous retournant un tableau de ses derniers
// Nécessaire pour déterminer les "marketplace" distincts et plus tard le revenu total par marketplace
// Attention aux "marketplace" vide !
/*
func getMarkets(orders Orders) []string{

}
*/

// Calcule le revenu total et par markets
// Prend en argument les orders et les markets
/*
func revenu(orders Orders, markets){

}
*/

// Recense les commandes suspectes. Sous-fonction appelé par la fonction "revenu"
// func suspiciousOrder(){}

func main() {
	// Open our jsonFile passed as an argument (or use "orders.json")
	var jsonFile *os.File
	var err error
	if len(os.Args) > 1 {
		jsonFile, err = os.Open(os.Args[1])
	} else {
		jsonFile, err = os.Open("orders.json")
	}

	// if os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened orders.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened File as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// we initialize our Orders array
	var orders Orders

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'orders' which we defined above
	json.Unmarshal(byteValue, &orders)

	// we iterate through every oder within our orders array and print out the full order as just an example
	/*
		for i := 0; i < len(orders.Orders); i++ {
			fmt.Print("Order id: " + orders.Orders[i].ID)
			fmt.Print("; Markerplace: " + orders.Orders[i].Marketplace)
			fmt.Print("; Country: " + orders.Orders[i].Country)

			// Stocké dans une variable car on va vouloir récupérer se montant plus tard.
			var tmp = orders.Orders[i].Amount_cents / 100
			fmt.Printf("; Amount_€: %.2f", tmp)
			fmt.Println("; Created_at: " + orders.Orders[i].Created_at)
		}
	*/

	//On récupére nos marketplaces
	//var markets := getMarkets(orders)

}
