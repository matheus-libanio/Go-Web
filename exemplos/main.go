package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

func main() {

	marshallingAndUnmarshalling()
	encondingAndDecoding()
	roteador()
}

func marshallingAndUnmarshalling() {
	type product struct {
		Name      string
		Price     int
		Published bool
	}

	p := product{
		Name:      "MacBook Pro",
		Price:     1500,
		Published: true,
	}
	// Marshalling - Serializando
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	//Unmarshalling - Desserializando

	newJsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`

	var r product

	if err := json.Unmarshal([]byte(newJsonData), &r); err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
}

func encondingAndDecoding() {
	//Encoding

	// It is necessary to create our type Encode
	// for this the NewEncoder function is called
	// this receives a streaming as a parameter
	// we use one of the standard streams offered by the OS Stdout pkg
	// stdout generates a stream to a file that is printed to the console.

	myEncoder := json.NewEncoder(os.Stdout)

	// prepare the information you want to send in json format to the streaming
	type MyData struct {
		ProductID string
		Price     float64
	}
	data := MyData{
		ProductID: "XASD",
		Price:     25.50,
	}

	// the Encode method is invoked.
	// internally this method makes a kind of marshall and writes it to the stream

	myEncoder.Encode(data)

	//DECODING
	const jsonStream = `
    {"ProductID": "AXW123", "Price": 25.50}
    {"ProductID": "NLBR17", "Price": 357.58}
    {"ProductID": "KNUB82", "Price": 150}
    `
	// It is necessary to create our type Decode, for this the NewDecoder function is called
	// this receives a streaming as a parameter
	// A jsonStream variable is created and the NewReader method of the strings pkg is used
	// NewReader generates a streaming for the text string it receives.

	myStreaming := strings.NewReader(jsonStream)
	myDecoder := json.NewDecoder(myStreaming)

	// streaming behaves so that each line in the jsonStrem text is streamed separately
	// we go through all the data transmitted in the streaming until the end of the text is reached
	for {
		// the variable on which the data is going to be written is created
		var data MyData
		// the decode method is invoked
		// Decode is responsible for reading the data transmitted by the streaming and transforming it from json to our interface
		if err := myDecoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// The received data is printed
		log.Println("Data:", data.ProductID, data.Price)
	}

}

func roteador() {

	// criando uma rota com  a biblioteca chi e o server
	router := chi.NewRouter()

	router.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		// set code 200
		w.WriteHeader(200)

		// set boody
		w.Write([]byte(`{"message":"Hello World!}`))
	})

	//run
	http.ListenAndServe(":8080", router)
}
