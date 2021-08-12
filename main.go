package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	uri := "neo4j://localhost:7687"
	user := "neo4j"
	passwd := "hn3437"

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(user, passwd, ""))
	if err != nil {
		return
	}
	defer driver.Close()

	// Search Graph
	Search()

	// Add node
	trunsctionResult, err := AddNode(driver)

	// Add arrow

	// Delete node

	// Delete arrow

	// Change node

	// Change arrow

	fmt.Println(trunsctionResult, err)
}
