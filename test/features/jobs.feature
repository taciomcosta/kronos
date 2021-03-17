Feature: Job
    Scenario: Job is listed after creation
        Given I provide valid data for job creation
        When I create a new job
        And I list the existing jobs
        Then the new job is listed

    Scenario: Invalid job creation
        Given I provide invalid data for job creation
        When I create a new job
        Then an error message is shown

    Scenario: Job is not visualized after deletion
        Given I provide valid data for job creation
        And I create a new job
        When I delete the new job
        And I list the existing jobs
        Then the new job is not listed

    Scenario: Job is described after creation
        Given I provide valid data for job creation
        When I create a new job
        And I describe the new job
        Then the new job is detailed
