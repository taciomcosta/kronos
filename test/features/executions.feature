Feature: Execution
    Scenario: Execution is listed 
        Given that I create a job
        When the job finishes 1 execution
        And I list all job execution history
        Then 1 execution should is listed
