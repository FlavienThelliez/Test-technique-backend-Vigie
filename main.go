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

func main() {
	/*name := "Go Developers"
	fmt.Println("Azure for", name)*/
	// Open our jsonFile
	jsonFile, err := os.Open("orders.json")
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

	// we iterate through every oder within our orders array and
	// print out the full order as just an example
	for i := 0; i < len(orders.Orders); i++ {
		fmt.Print("Order id: " + orders.Orders[i].ID)
		fmt.Print("; Markerplace: " + orders.Orders[i].Marketplace)
		fmt.Print("; Country: " + orders.Orders[i].Country)

		// Code invalide pour traiter le montant (problèmes de type)
		/*var tmp = strconv.Itoa(orders.Orders[i].Amount_cents)
		fmt.Print("; Amount_cent: " + tmp)*/

		// Stocké dans une variable car on va vouloir récupérer se montant plus tard.
		var tmp = orders.Orders[i].Amount_cents / 100
		fmt.Printf("; Amount_€: %.2f", tmp)
		fmt.Println("; Created_at: " + orders.Orders[i].Created_at)
	}
}
