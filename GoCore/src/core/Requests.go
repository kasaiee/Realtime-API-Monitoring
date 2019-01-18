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

}