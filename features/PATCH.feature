Feature: As a QA,I want to verify if PATCH requests are working as expected for https://reqres.in/
  Background: Set up the url
    Given I set the url

Scenario: Verify if PATCH request is working as expected
  When I send PATCH request with endpoint "/api/users/2" and with updated data "data.name"
  And status code should be 200
  Then data in response should be updated "data.name" as well