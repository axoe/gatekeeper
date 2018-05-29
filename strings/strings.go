package main

import (
	"encoding/json"
	"fmt"
)

type GatekeeperResponse map[string]interface{}

func main() {
	resMap := GatekeeperResponse{
		"ARN":           "aaa",
		"CreatedDate":   12321,
		"Name":          "some name",
		"SecretString":  "a string",
		"VersionId":     "version id",
		"VersionStages": []string{"SOME_STAGE"},
	}

	res2Map := GatekeeperResponse{
		"ARN":           "aaa",
		"CreatedDate":   12321,
		"Name":          "some name",
		"SecretString":  map[string]string{"some json key": "the value"},
		"VersionId":     "version id",
		"VersionStages": []string{"SOME_STAGE"},
	}

	resBytes, err := json.Marshal(resMap)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	res2Bytes, err := json.Marshal(res2Map)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("Marshalled JSON resMap: %s\n", resBytes)
	fmt.Printf("Marshalled JSON res2Map: %s\n", res2Bytes)

	var res GatekeeperResponse
	err = json.Unmarshal(resBytes, &res)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	var res2 GatekeeperResponse
	err = json.Unmarshal(res2Bytes, &res2)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("Unmarshalled JSON res: %+v\n", res)
	fmt.Printf("Unmarshalled JSON res2: %+v\n", res2)

	switch val := res["SecretString"].(type) {
	case string:
		fmt.Printf("SecretString is a string: %s\n", val)

	case interface{}:
		fmt.Printf("SecretString is an interface: %v\n", val)
	}

	switch val := res2["SecretString"].(type) {
	case string:
		fmt.Printf("SecretString is a string: %s\n", val)

	case interface{}:
		fmt.Printf("SecretString is an interface{}: %v\n", val)
		val2, ok := val.(map[string]interface{})
		if !ok {
			fmt.Printf("Error: not a map[string]string\n")
			return
		}
		fmt.Printf("SecretString is now a map[string]interface{}: %v\n", val2)

		for key, val3 := range val2 {
			val4, ok := val3.(string)
			if !ok {
				fmt.Printf("Error: not a string\n")
				return
			}
			fmt.Printf("key: %s, : val4: %s\n", key, val4)
		}
	}
}
