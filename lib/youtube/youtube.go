package youtube

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	apiKey = os.Getenv("YOUTUBE_API_KEY")
)

func fetchData(path, params string) {
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/youtube/v3/%s?%s&key=%s", path, params, apiKey))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("get:\n", string(body))
}

// GetChannelData ...
func GetChannelData() /*(model.Channel, err)*/ {
	fetchData("channels", "part=snippet%2CbrandingSettings&id=UCqB90BBr6eNRaJl-kl30Xxw")

}
