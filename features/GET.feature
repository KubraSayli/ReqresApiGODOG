Feature: As a QA,I want to verify if GET requests are working as expected for https://reqres.in/
  Background: Set up the url
    Given I set the url

  Scenario: Verify if GET request is working as expected for list of users
    When I send GET request to the endpoint with "/api/users?page=2"
    And status code should be 200
    Then the number of data should be 6 for each page
    And the 11 th users name is "George"


  Scenario Outline: Verify if GET request is working as expected for one user
    When I send GET request to the endpoint with "/api/users/2"
    And status code should be 200
    Then I verify response is as expected for related filter "<testData1>" ,"<testData2>" , "<testData3>" , "<testData4>"
    Examples:
      | testData1 | testData2  | testData3       | testData4      |
      | data.id   | data.email | data.first_name | data.last_name |

  @this
Scenario: Verify /api/users/23 endpoint returns null
  When I send GET request to the endpoint with "/api/users/23"
  And status code should be 404
  Then response should be null



