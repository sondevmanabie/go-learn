package file_reader

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type FileReader interface {
	Read(res interface{}) (interface{}, error)
	Write(data interface{}) error
}

type EntityReader struct {
	Entity string
	Path   string
}

func (r *EntityReader) Read(res interface{}) (interface{}, error) {
	if r.Entity != "student" && r.Entity != "book" {
		return nil, errors.New("not match any entity")
	}
	data, err := ioutil.ReadFile(r.Path + "/" + r.Entity + ".json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if err := json.Unmarshal(data, res); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func (r *EntityReader) Write(data interface{}) error {
	if r.Entity != "student" && r.Entity != "book" {
		return errors.New("not match any entity")
	}
	convertedData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = ioutil.WriteFile(r.Path+"/"+r.Entity+".json", convertedData, 0644)
	return err
}
