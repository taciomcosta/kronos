Feature: Create a job
    Scenario: Job is visualized after creation
        Given I create a new job
        When I list the existing jobs
        Then the new job should be listed


