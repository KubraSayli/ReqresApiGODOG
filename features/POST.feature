Feature: As a QA,I want to verify if POST requests are working as expected for https://reqres.in/

  Background: Set up the url
    Given I set the url

  Scenario: Verify if POST request is working as expected for creating user
    When I send post request "/api/users" with "testData/CreateUserBody.json"
    And status code should be 201
    Then I verify if data from POST body matches with data from response
      | name | job |


  Scenario: Verify if POST request is working as expected for login
    When I send post request "/api/login" with "testData/Login.json"
    And status code should be 200
    Then I verify if data from response
      | token | QpwL5tke4Pnpja7X4 |

  @negative
  Scenario: Verify if POST request with /api/register contains error message "Missing password"
    When I send post request "/api/register" with "testData/Register.json"
    And status code should be 400
    Then I verify if error message contains "Missing password"
