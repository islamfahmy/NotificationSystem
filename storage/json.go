package storage

import (
	"encoding/json"
	"io/ioutil"
	"log"
	subj "main/subjects"
	"os"
)

type JsonSaveAgent struct{}

func CreateJsonSaveAgent() *JsonSaveAgent {
	return &JsonSaveAgent{}
}
func checkorCreateFile(fileName string) error {
	if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
		csvFile, err := os.Create(fileName + ".json")
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
			return err
		}
		csvFile.Close()
	}
	return nil
}
func (*JsonSaveAgent) Save(obj map[int]subj.Observer, fileName string) error { // need to change it to specific type
	err := checkorCreateFile(fileName)
	if err != nil {
		return err
	}
	s := []interface{}{}
	for _, v := range obj {
		var temp interface{}
		encodedjson, err := json.Marshal(v.Export())
		if err != nil {
			log.Fatalf("failed to marshal: %s", err)
		}
		if err := json.Unmarshal(encodedjson, &temp); err != nil {
			log.Fatalf("failed to unmarshal: %s", err)

		}
		s = append(s, temp)

	}
	b, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		log.Fatal("final marshal error: ", err)
	}

	ioutil.WriteFile(fileName+".json", b, 0644)

	return nil
}
