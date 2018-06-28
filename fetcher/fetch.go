package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/unicode"
)

func Fetch (url string) ([]byte, error) {

	//resp, err := http.Get(url)


	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil ,err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return nil,
				fmt.Errorf("Wrong status code: %d",
				resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	urfReader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(urfReader)


}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e , _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
