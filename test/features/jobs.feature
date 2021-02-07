Feature: Create a job
    Scenario: Job is visualized after creation
        Given I provide valid data for job creation
        When I create a new job
        And I list the existing jobs
        Then the new job should be listed

    Scenario: Invalid job creation
        Given I provide invalid data for job creation
        When I create a new job
        Then an error message is shown
