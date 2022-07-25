package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Populate and Lookup
//
//  Add elements to the maps that you've declared in the
//  first exercise, and try them by looking up for the keys.
//
//  Either use the `make()` or `map literals`.
//
//  After completing the exercise, remove the data and check
//  that your program still works.
//
//
//  1. Phone numbers by last name
//     --------------------------
//     bowen  202-555-0179
//     dulin  03.37.77.63.06
//     greco  03489940240
//
//     Print the dulin's phone number.
//
//  2. Product availability by Product ID
//     ----------------
//     617841573 true
//     879401371 false
//     576872813 true
//
//     Is Product ID 879401371 available?
//
//
//  3. Multiple phone numbers by last name
//     ------------------------------------------------------
//     bowen  [202-555-0179]
//     dulin  [03.37.77.63.06 03.37.70.50.05 02.20.40.10.04]
//     greco  [03489940240 03489900120]
//
//     What is Greco's second phone number?
//
//
//  4. Shopping basket by Customer ID
//     -------------------------------
//     100 [617841573:4 576872813:2]
//     101 [576872813:5 657473833:20]
//     102 []
//
//     How many of 576872813 the customer 101 is going to buy?
//                (Product ID)  (Customer ID)
//
//
// EXPECTED OUTPUT
//
//   1. Run the solution to see the output
//   2. Here is the output with empty maps:
//
//      dulin's phone number: N/A
//      Product ID #879401371 is not available
//      greco's 2nd phone number: N/A
//      Customer #101 is going to buy 5 from Product ID #576872813.
//
// ---------------------------------------------------------

func main() {
	var (
		phones         map[string]string
		multiplePhones map[string][]string
		basket         map[int]map[int]int
		products       map[int]string
	)

	// #1
	phones = make(map[string]string)
	phones["bowen"] = "202-555-0179"
	phones["dulin"] = "03.37.77.63.06"
	phones["greco"] = "03489940240"

	// 3
	multiplePhones = make(map[string][]string)
	multiplePhones["bowen"] = []string{"202-555-0179"}
	multiplePhones["dulin"] = []string{"03.37.77.63.06", "03.37.70.50.05", "02.20.40.10.04"}
	multiplePhones["greco"] = []string{"03489940240", "03489900120"}

	// #3
	products = make(map[int]string)
	products[617841573] = "true"
	products[879401371] = "false"
	products[576872813] = "true"

	// #4
	basket = make(map[int]map[int]int)
	basket[100] = map[int]int{617841573: 4, 576872813: 2}
	basket[101] = map[int]int{576872813: 5, 657473833: 20}
	basket[102] = map[int]int{}

	fmt.Printf("dulin's phone number: %s\n", phones["dulin"])
	fmt.Printf("Product ID #879401371 is %s\n", products[879401371])
	fmt.Printf("greco's 2nd phone number: %s\n", multiplePhones["greco"][1])
	fmt.Printf("Customer #101 is going to buy %d from Product ID #%d.\n", basket[101][576872813], 576872813)
}
