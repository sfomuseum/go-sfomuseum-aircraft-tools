package tools

import (
	"github.com/sfomuseum/go-sfomuseum-aircraft/sfomuseum"
	// sfomuseum_props "github.com/sfomuseum/go-sfomuseum-geojson/properties/sfomuseum"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/properties/whosonfirst"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/utils"
)

func SFOMuseumAircraftFromFeature(f geojson.Feature) (*sfomuseum.Aircraft, error) {

	wof_id := whosonfirst.Id(f)
	name := whosonfirst.Name(f)

	sfom_id := utils.Int64Property(f.Bytes(), []string{"properties.sfomuseum:aircraft_id"}, -1)

	concordances, err := whosonfirst.Concordances(f)

	if err != nil {
		return nil, err
	}

	a := &sfomuseum.Aircraft{
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

	return a, nil
}
