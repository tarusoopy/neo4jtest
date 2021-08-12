package main

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

func AddNode(driver neo4j.Driver) (string, error) {
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (company:会社) SET company.name = $name RETURN company.name + ', from node ' + id(company)",
			map[string]interface{}{"name": "VarioSecure"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}
