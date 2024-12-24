package urbanapi
import (
	"log"
	"net/http"
	"encoding/json"
	"io"
	"fmt"
)

type RData struct {
	Data []struct {
		Meaning string `json:"meaning"`
		Example string `json:"example"`
	}
}
func Run(message string) string{
	endpoint := "https://unofficialurbandictionaryapi.com/api/search?term="
	log.Println("fetching definition")
	var d RData
	out, err := connect(endpoint,message); if err != nil {
		log.Println(err)
	}
	d1,err := decode(out,d); if err!= nil {
		log.Println(err)
	}
  return fmt.Sprintf("Meaning %s ",d1.Data[0].Meaning)
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
