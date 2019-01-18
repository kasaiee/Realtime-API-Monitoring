import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGetRequestWithoutToken(url string) int {
	// send get request and receive response bytes
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()


	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// cast response bytes to json
	responseCodeInstance := ResponseCode{}
	err = json.Unmarshal(bodyBytes, &responseCodeInstance)
	
	if err != nil {
		fmt.Println(err)
		return 700
	}

	return responseCodeInstance.ResponseCode

}