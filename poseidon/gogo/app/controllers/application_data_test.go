package controllers

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"testing"

	"github.com/poseidon/app/concerns/kodo"
)

func Test_Application_Data(t *testing.T) {
	filename := path.Clean("../../tmp/data/birds.json")

	kodoclient := kodo.New(Config.Qiniu.Kodo)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	images := strings.Split(string(data), ",")
	for _, image := range images {
		image = strings.Replace(image, ":", "", -1)
		// image = strings.Replace(image, ".", "", -1)

		urlobj, err := url.Parse(image)
		if err != nil {
			continue
		}

		imgurl := urlobj.Query().Get("imgurl")
		if imgurl == "" {
			continue
		}
		println("Scrabe image from ", imgurl)

		req, err := http.NewRequest("GET", imgurl, nil)
		if err != nil {
			continue
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			continue
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			println("Read", imgurl, ":", err.Error())

			continue
		}
		resp.Body.Close()

		name := imgurl[strings.LastIndex(imgurl, "/")+1:]

		err = kodoclient.Put(name, data)
		if err != nil {
			println("Upload(", name, ", ", imgurl, "):", err.Error())
		}
	}
}
