package urbanapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RData struct {
	Data []struct {
		Meaning string `json:"meaning"`
		Example string `json:"example"`
	}`json:"data"`
}
func Run(message string) (string,string){
	endpoint := "https://unofficialurbandictionaryapi.com/api/search?term="
	var d RData
	out, err := connect(endpoint,message); if err != nil {
		log.Println(err)
	}
	d1,err := decode(out,d); if err!= nil {
		log.Println(err)
	}
	if len(d1.Data) == 0 {
		return "No defintion found","no example found"
	}
	  return d1.Data[0].Meaning, d1.Data[0].Example
}

func decode(data []byte, dStruct RData) (*RData,error){
	err := json.Unmarshal(data, &dStruct); if err != nil {
		log.Fatalln("Unable to unmarshal data", err)
	}
	return &dStruct, nil
}

func connect(endpoint string, term string) ([]byte,error){
	var client http.Client
	search := endpoint + term
	req, err := http.NewRequest("GET", search, nil); if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req); if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body); if err != nil {
		log.Fatalln(err)
		return nil,err
	}
	return body,nil
}
