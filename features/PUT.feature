Feature: As a QA,I want to verify if PUT requests are working as expected for https://reqres.in/
  Background: Set up the url
    Given I set the url

  Scenario Outline: Verify if PUT request is working as expected
    When I send PUT request with endpoint "<endpoint>" and with updated data "<data1>" and "<data2>"
    And status code should be 200
    Then data in response should be updated as well for PUT request "<data1>" , "<data2>"
    Examples:
      | endpoint | data1 | data2 |
      | /api/users/2 | data.name | data.job |