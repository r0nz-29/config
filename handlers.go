package main

import "fmt"

func handle(args []string) {
	switch args[0] {
	case GET:
		handleGet(args[1])
	case SET:
		handleSet(args[1:])
	case VIEW:
		handleView()
	default:
		return
	}
}

func handleView() {
	jsonMap := *getJsonFromFile()
	for key, val := range jsonMap {
		println("--------------------")
		fmt.Print(key + ": ")
		fmt.Println(val)
		println("--------------------")
	}
}

func handleGet(key string) {
	jsonMap := *getJsonFromFile()

	val, exists := jsonMap[key]

	if !exists {
		fmt.Println("No record found for key: " + key)
		return
	}

	fmt.Println(val)
}

func handleSet(kv []string) {
	key := kv[0]
	val := kv[1]

	jsonMap := *getJsonFromFile()
	jsonMap[key] = val

	persist(&jsonMap)
	println("saved")
}
