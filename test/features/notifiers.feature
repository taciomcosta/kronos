Feature: Notifier
    Scenario: Notifier is listed after creation
        Given I provide valid data for notifier creation
        When I create a new notifier
        And I list the existing notifiers
        Then the new notifier is listed

    Scenario: Invalid notifier creation
        Given I provide invalid data for notifier creation
        When I create a new notifier
        Then an error message is shown for notifier

    Scenario: Notifier is not visualized after deletion
        Given I provide valid data for notifier creation
        And I create a new notifier
        When I delete the new notifier
        And I list the existing notifiers
        Then the new notifier is not listed
