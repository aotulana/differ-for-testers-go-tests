package tests

import (
	"encoding/json"
	"io/ioutil"
)

type (
	leftSideResponse struct {
		Left string
	}

	rightSideResponse struct {
		Right string
	}
)

func (suite *TestSuite) TestSideEndpointShouldReturnExactValueInRequest() {
	expectedStatusCode := 200

	for _, t := range []struct {
		testName string
		side     string
		value    string
	}{
		{"leftSideShouldReturnLeftSideValueOnly", "left", "abcd"},
		{"rightSideShouldReturnRightSideValueOnly", "right", "dcba"},
	} {
		suite.Run(t.testName, func() {
			id, err := suite.GenerateRandomID()
			response, err := suite.SetSideValue(id, t.side, t.value)

			if err != nil {
				panic(err)
			}

			responseBody, err := ioutil.ReadAll(response.Body)
			suite.NoError(err)

			switch t.side {
			case "left":
				actualResponseBody := leftSideResponse{}
				expectedResponseBody := leftSideResponse{t.value}
				responseDecodingErr := json.Unmarshal(responseBody, &actualResponseBody)

				suite.NoError(responseDecodingErr)
				suite.Equal(expectedStatusCode, response.StatusCode)
				suite.Equal(expectedResponseBody, actualResponseBody, string(responseBody))
				break
			case "right":
				actualResponseBody := rightSideResponse{}
				expectedResponseBody := rightSideResponse{t.value}
				responseDecodingErr := json.Unmarshal(responseBody, &actualResponseBody)

				suite.NoError(responseDecodingErr)
				suite.Equal(expectedStatusCode, response.StatusCode)
				suite.Equal(expectedResponseBody, actualResponseBody, string(responseBody))
				break
			}
			defer response.Body.Close()
		})
	}
}
