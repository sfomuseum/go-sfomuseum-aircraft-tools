package main

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/sfomuseum/go-sfomuseum-aircraft-tools/template"
	"github.com/sfomuseum/go-sfomuseum-aircraft/sfomuseum"
	"github.com/sfomuseum/go-sfomuseum-geojson/feature"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
	"github.com/whosonfirst/go-whosonfirst-index"
	"io"
	"log"
	"os"
	"sync"
)

func main() {

	data := flag.String("data", "/usr/local/data/sfomuseum-data-aircraft", "...")

	flag.Parse()

	lookup := make([]sfomuseum.Aircraft, 0)
	mu := new(sync.RWMutex)

	cb := func(fh io.Reader, ctx context.Context, args ...interface{}) error {

		f, err := feature.LoadFeatureFromReader(fh)

		if err != nil {
			return err
		}

		wof_id := whosonfirst.Id(f)
		name := whosonfirst.Name(f)

		sfom_id := utils.Int64Property(f.Bytes(), []string{"properties.sfomuseum:aircraft_id"}, -1)

		concordances, err := whosonfirst.Concordances(f)

		if err != nil {
			return err
		}

		a := sfomuseum.Aircraft{
			WOFID:       wof_id,
			SFOMuseumID: int(sfom_id),
			Name:        name,
		}

		code, ok := concordances["icao:designator"]

		if ok {
			a.ICAODesignator = code
		}

		id, ok := concordances["wd:id"]

		if ok {
			a.WikidataID = id
		}

		mu.Lock()
		lookup = append(lookup, a)
		mu.Unlock()

		return nil
	}

	idx, err := index.NewIndexer("repo", cb)

	if err != nil {
		log.Fatal(err)
	}

	err = idx.IndexPath(*data)

	if err != nil {
		log.Fatal(err)
	}

	enc, err := json.Marshal(lookup)

	if err != nil {
		log.Fatal(err)
	}

	vars := template.AircraftDataVars{
		Package: "sfomuseum",
		Data:    string(enc),
	}

	err = template.RenderAircraftData(os.Stdout, &vars)

	if err != nil {
		log.Fatal(err)
	}
}
