package ReqresApiGODOG

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/tidwall/gjson"
	"github/ReqresApiGODOG/utilities"
	"gopkg.in/resty.v1"
	"strconv"
	"strings"
)

var url string
var contentType string
var requestBody string
var responseBody string
var statusCode int
var updatedRequestBody string

func iSetTheUrl() {
	url = utilities.JsonReader("Configuration.json", "createUser.Url")
	contentType = utilities.JsonReader("Configuration.json", "createUser.contentType")
}

func iSendPostRequestWith(pathParameter, requestBodyPath string) (err error) {
	//set the request body by parsing in json file in test data
	requestBody = utilities.JsonReader(requestBodyPath, "data")

	//Creating post request:
	client := resty.New()
	Resp, _ := client.R().
		SetHeader("Content-Type", contentType).
		SetBody([]byte(requestBody)).
		Post(url + pathParameter)
	responseBody = string(Resp.Body())
	statusCode = Resp.StatusCode()
	return err
}

func statusCodeShouldBe(expectedStatusCode int) (err error) {
	//Assert status code by using assert methods from utilities
	utilities.AssertEqualInt(expectedStatusCode, statusCode)
	return err
}

func iVerifyIfDataFromPOSTBodyMatchesWithDataFromResponse(jsonPaths *godog.Table) (err error) {
	//For expected data: Read json file from test data package with the help of Json Reader
	//For actual data:Read response body from post request method with the help of Json Reader
	//Note that, we gat json path from feature file, this way we can modify the data easily from feature file
	for j := 0; j < len(jsonPaths.Rows[0].Cells); j++ {
		expectedData := gjson.Get(requestBody, jsonPaths.Rows[0].Cells[j].Value).String()
		actualData := gjson.Get(responseBody, jsonPaths.Rows[0].Cells[j].Value).String()
		utilities.AssertEqual(expectedData, actualData)
	}
	return err
}

func iVerifyIfDataFromResponse(dataTable *godog.Table) (err error) {
	expectedData := dataTable.Rows[0].Cells[1].Value
	actualData := gjson.Get(responseBody, dataTable.Rows[0].Cells[0].Value).String()
	utilities.AssertEqual(expectedData, actualData)
	return err

}

func iVerifyIfErrorMessageContains(errorMessage string) (err error) {
	utilities.AssertContain(responseBody, errorMessage)
	return err
}

//=====GET REQUESTS=========
func iSendGETRequestToTheEndpointWith(pathParameter string) (err error) {
	client := resty.New()
	Resp, _ := client.R().
		SetHeader("Content-Type", contentType).
		Get(url + pathParameter)
	responseBody = string(Resp.Body())
	statusCode = Resp.StatusCode()
	return err
}

func theNumberOfDataShouldBeForEachPage(expectedNumber int) (err error) {
	actualNumber := strings.Count(responseBody, "id")
	utilities.AssertEqualInt(expectedNumber, actualNumber)
	return err
}

func theThUsersNameIs(arg1 int, name string) (err error) {
	actualName := gjson.Get(responseBody, "data.4.first_name").String()
	utilities.AssertEqual(name, actualName)
	return err
}

func iVerifyResponseIsAsExpectedForRelatedFilter(testData1, testData2, testData3, testData4 string) (err error) {
	expectedData1, _ := strconv.Atoi(utilities.JsonReader("testData/OneUser.json", testData1))
	actualData1, _ := strconv.Atoi(gjson.Get(responseBody, testData1).String())
	utilities.AssertEqualInt(expectedData1, actualData1)
	expectedData2 := utilities.JsonReader("testData/OneUser.json", testData2)
	actualData2 := gjson.Get(responseBody, testData2).String()
	utilities.AssertEqual(expectedData2, actualData2)
	expectedData3 := utilities.JsonReader("testData/OneUser.json", testData3)
	actualData3 := gjson.Get(responseBody, testData3).String()
	utilities.AssertEqual(expectedData3, actualData3)
	expectedData4 := utilities.JsonReader("testData/OneUser.json", testData4)
	actualData4 := gjson.Get(responseBody, testData4).String()
	utilities.AssertEqual(expectedData4, actualData4)
	return err
}

func responseShouldBeNull() (err error) {
	utilities.AssertEqual(responseBody, "{}")
	return err
}

//===========PATCH Requests=====
func iSendPATCHRequestWithEndpointAndWithUpdatedData(pathParameter, data string) (err error) {
	//update data in post body
	updatedRequestBody = utilities.JsonSetter("testData/CreateUserBody.json", data, "Joe")
	client := resty.New()
	Resp, _ := client.R().
		SetHeader("Content-Type", contentType).
		SetBody([]byte(updatedRequestBody)).
		Patch(url + pathParameter)
	responseBody = string(Resp.Body())
	statusCode = Resp.StatusCode()
	return err
}

func dataInResponseShouldBeUpdatedAsWell(pathOfData string) (err error) {
	expectedData := gjson.Get(updatedRequestBody, pathOfData).String()
	actualData := gjson.Get(responseBody, pathOfData).String()
	utilities.AssertEqual(expectedData, actualData)
	return err
}

//====DELETE Request=====

func iSendDELETERequestWithEndpoint(pathParameter string)  {
	client := resty.New()
	Resp, _ := client.R().
		SetHeader("Content-Type", contentType).
		Delete(url+pathParameter)
	responseBody = string(Resp.Body())
	statusCode = Resp.StatusCode()
}

func dataShouldBeDeleted() (err error) {
	utilities.AssertEqualInt(len(responseBody), 0)
	return err
}

//======PUT Requests=====
func iSendPUTRequestWithEndpointAndWithUpdatedDataAnd(pathParameter, data, data2 string) (err error) {
	//update data in post body
	updatedRequestBody = utilities.JsonSetter("testData/CreateUserBody.json", data, "Joe")
	updatedRequestBody = utilities.JsonSetter("testData/CreateUserBody.json", data2, "QA")
fmt.Println(updatedRequestBody)
	client := resty.New()
	Resp, _ := client.R().
		SetHeader("Content-Type", contentType).
		SetBody([]byte(updatedRequestBody)).
		Patch(url + pathParameter)
	responseBody = string(Resp.Body())
	statusCode = Resp.StatusCode()
	return err
}

func dataInResponseShouldBeUpdatedAsWellForPUTRequest(pathOfData, pathOfData2 string) (err error) {
	expectedData := gjson.Get(updatedRequestBody, pathOfData).String()
	actualData := gjson.Get(responseBody, pathOfData).String()
	expectedData2 := gjson.Get(updatedRequestBody, pathOfData2).String()
	actualData2 := gjson.Get(responseBody, pathOfData2).String()
	utilities.AssertEqual(expectedData, actualData)
	utilities.AssertEqual(expectedData2, actualData2)
	return err
}




func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I set the url$`, iSetTheUrl)
	ctx.Step(`^I send post request "([^"]*)" with "([^"]*)"$`, iSendPostRequestWith)
	ctx.Step(`^status code should be (\d+)$`, statusCodeShouldBe)
	ctx.Step(`^I verify if data from POST body matches with data from response$`, iVerifyIfDataFromPOSTBodyMatchesWithDataFromResponse)
	ctx.Step(`^I verify if data from response$`, iVerifyIfDataFromResponse)
	ctx.Step(`^I verify if error message contains "([^"]*)"$`, iVerifyIfErrorMessageContains)
	ctx.Step(`^I send GET request to the endpoint with "([^"]*)"$`, iSendGETRequestToTheEndpointWith)
	ctx.Step(`^the number of data should be (\d+) for each page$`, theNumberOfDataShouldBeForEachPage)
	ctx.Step(`^the (\d+) th users name is "([^"]*)"$`, theThUsersNameIs)
	ctx.Step(`^I verify response is as expected for related filter "([^"]*)" ,"([^"]*)" , "([^"]*)" , "([^"]*)"$`, iVerifyResponseIsAsExpectedForRelatedFilter)
	ctx.Step(`^response should be null$`, responseShouldBeNull)
	ctx.Step(`^I send PATCH request with endpoint "([^"]*)" and with updated data "([^"]*)"$`, iSendPATCHRequestWithEndpointAndWithUpdatedData)
	ctx.Step(`^data in response should be updated "([^"]*)" as well$`, dataInResponseShouldBeUpdatedAsWell)
	ctx.Step(`^I send DELETE request with endpoint "([^"]*)"$`, iSendDELETERequestWithEndpoint)
	ctx.Step(`^data should be deleted$`, dataShouldBeDeleted)
	ctx.Step(`^I send PUT request with endpoint "([^"]*)" and with updated data "([^"]*)" and "([^"]*)"$`, iSendPUTRequestWithEndpointAndWithUpdatedDataAnd)
	ctx.Step(`^data in response should be updated as well for PUT request "([^"]*)" , "([^"]*)"$`, dataInResponseShouldBeUpdatedAsWellForPUTRequest)

}
