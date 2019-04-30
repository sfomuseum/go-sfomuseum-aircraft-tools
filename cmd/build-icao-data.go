package main

import (
	"context"
	"flag"
	"github.com/sfomuseum/go-sfomuseum-aircraft-tools/template"	
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	// curl 'https://www4.icao.int/doc8643/External/AircraftTypes' -H 'Connection: keep-alive' --data ''

	source := flag.String("source", "https://www4.icao.int/doc8643/External/AircraftTypes", "...")

	flag.Parse()

	rsp, err := 
	
	defer rsp.Body.Close()
	
	body, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		log.Fatal(err)
	}
	
	vars := template.AircraftDataVars{
		Package: "icao",
		Data:    string(body),
	}

	err = template.RenderAircraftData(os.Stdout, &vars)

	if err != nil {
		log.Fatal(err)
	}
}
