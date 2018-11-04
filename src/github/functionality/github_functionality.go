package functionality

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/msalcantara/godog-jenkins/src/github/page"
)

func GetStatusRequestGitHub(url string) (int, error) {
	response, err := http.Get(url)
	if err != nil {
		return 500, err
	}
	return response.StatusCode, nil
}

func GetRepositories(url string) ([]page.Repository, error) {
	body, err := request(url)
	if err != nil {
		return nil, err
	}
	var repositories []page.Repository
	if err := json.Unmarshal(body, &repositories); err != nil {
		return nil, err
	}
	return repositories, nil
}

func request(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
