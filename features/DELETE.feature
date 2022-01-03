Feature: As a QA,I want to verify if DELETE requests are working as expected for https://reqres.in/
  Background: Set up the url
    Given I set the url

    Scenario:  Verify if DELETE request is working as expected
      When I send DELETE request with endpoint "/api/users/2"
      And status code should be 204
      Then data should be deleted