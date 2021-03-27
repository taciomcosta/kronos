Feature: Notifier
    Scenario: Notifier is listed after creation
        Given I provide valid data for notifier creation
        When I create a new notifier
        And I list the existing notifiers
        Then the new notifier is listed
