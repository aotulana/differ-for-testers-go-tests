package tests

import (
	"bytes"
	"crypto/rand"
	"differ-for-testers-go-tests/integration/config"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type (
	// TestSuite defines the tests to run as part of the integration tests
	// It also defines some variables to be used across the test suite
	TestSuite struct {
		suite.Suite
		Host string
		Port int
	}
)

const basePath = "/diffassign/v1/diff"

// TestRunSuite will be run by the 'go test' command, so within it, we
// can run our suite using the Run(*testing.T, TestingSuite) function.
func TestRunSuite(t *testing.T) {
	testSuite := new(TestSuite)
	suite.Run(t, testSuite)
}

// SetupSuite runs before the test suite to setup the required variables
func (suite *TestSuite) SetupSuite() {
	config.Load()
	suite.Host = config.Conf.Host
	suite.Port = config.Conf.Port
}

// GenerateRandomID returns random positive integer value
func (suite *TestSuite) GenerateRandomID() (id int, err error) {
	min := 1
	max := 50
	i, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return
	}
	id = int(i.Int64()) + min
	return
}

// SetSideValue calls the side endpoint
func (suite *TestSuite) SetSideValue(id int, side, value string) (resp *http.Response, err error) {
	requestBodyJSON, _ := json.Marshal(value)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s:%d%s/%d/%s", suite.Host, suite.Port, basePath, id, side), bytes.NewBuffer(requestBodyJSON))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	resp = response

	return
}
