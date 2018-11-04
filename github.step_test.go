package steps

import (
	"fmt"
	"strconv"

	"github.com/DATA-DOG/godog"
	"github.com/msalcantara/godog-jenkins/src/github/functionality"
)

var statusCode string

func queRealizeiUmRequestParaOEndPoint(url string) error {
	status, err := functionality.GetStatusRequestGitHub(url)
	if err != nil {
		return err
	}
	statusCode = strconv.Itoa(status)
	return nil
}

func oResponseDaAPIDeveSer(statusFeature string) error {
	if statusFeature != statusCode {
		return fmt.Errorf("Expected: %s, Got: %s", statusFeature, statusCode)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que realizei um request para o end point "([^"]*)"$`, queRealizeiUmRequestParaOEndPoint)
	s.Step(`^o response da API deve ser "([^"]*)"$`, oResponseDaAPIDeveSer)
}
