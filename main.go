// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Let's first read the `config.json` file
	content, err := ioutil.ReadFile(os.Getenv("JSON_FILE"))
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Let's be freiendly
	fmt.Fprintln(w, "¡Hola!")

	// Let's print the unmarshalled data!
	// m is a map[string]interface.
	// loop over keys and values in the map.
	for k, v := range payload {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
