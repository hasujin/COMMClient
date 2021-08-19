package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	marine "github.com/hasujin/COMMClient/marineV21"
)

func main() {

	if len(os.Args) < 2 {
		marine.Help()
		return
	}

	mode := *flag.String("mode", getEnv("MODE", "CLI"), "Defualt Mode is CLI") // CLI / DOCKING /...
	tokenChannel := *flag.String("tokenchannel", getEnv("TOKENCHANNEL", "starpoly"), "Deafult Channel is starpoly")
	tokenCC := *flag.String("tokencc", getEnv("TOKENCC", "cat"), "Deafult chaincode is cat")

	marine.SetMode(mode)
	//	myArgs := []string{"./marine", "7", "127.0.0.1:50051", "Admin", "starpoly", "did", `{"Args":["queryHistoryByKey","1|BbsCofe7K9xupRHU4ycA6c"]}`}
	myArgs := os.Args
	result, err := marine.ExecuteMarine(myArgs, "Admin", tokenChannel, tokenCC)

	if mode != "CLI" {
		fmt.Printf("\nstatus: %v\n", result.Status)
		fmt.Printf("payload: %v\n", string(result.Payload))
	}

	if err != nil {
		fmt.Printf("%v\n", err)
	}

}

func getEnv(key, defaultValue string) string {
	if value, existed := os.LookupEnv(key); existed {
		return value
	}
	return defaultValue
}

func getIntEnv(key string, defaultValue int) int {
	if value, existed := os.LookupEnv(key); existed {
		ret, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return defaultValue
		} else {
			return int(ret)
		}
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value, existed := os.LookupEnv(key); existed {
		ret, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		} else {
			return ret
		}
	}
	return defaultValue
}
