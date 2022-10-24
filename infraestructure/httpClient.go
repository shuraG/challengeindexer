package infraestructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"example/indexer/domain"
)

type HttpClient struct {
	host   string
	user   string
	pass   string
	client http.Client
}

func NewHttpClient(host, user, pass string, timeout int) HttpClient {
	return HttpClient{
		host:   host,
		user:   user,
		pass:   pass,
		client: http.Client{Timeout: 5 * time.Second},
	}
}

func (h HttpClient) SaveBulk(emails []domain.Email, index string) bool {

	request := emailBulkRequest{
		Index:   index,
		Records: emails,
	}

	json_data, _ := json.Marshal(request)

	req, _ := http.NewRequest(http.MethodPost, h.host+saveBulkEndpoint, bytes.NewBuffer(json_data))
	req.SetBasicAuth(h.user, h.pass)
	res, err := h.client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(responseData))

	return true

}

func (h HttpClient) FindByMatch(index, keyword string, start, limit int) EmailZincResponse {

	client := http.Client{Timeout: 5 * time.Second}

	request := `{
			"search_type": "matchphrase",
			"query": {
				"term": "%s"
			},
			"sort_fields": ["-@timestamp"],
			"from": %d,
			"max_results": %d
		}`
	request = fmt.Sprintf(request, keyword, start, limit)

	req, _ := http.NewRequest(http.MethodPost, h.host+fmt.Sprintf(findByMatch, index), strings.NewReader(request))
	req.SetBasicAuth(h.user, h.pass)
	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err)
	}
	var emailZincResponse EmailZincResponse
	json.Unmarshal(responseData, &emailZincResponse)

	return emailZincResponse
}
