Feature: Org Account E2E sanity
  Scenario Outline: E2E sanity
   Given I created account <accountId> and <status>
   When I get account <accountId>
   Then Validate account with id <accountId> and status <status>
   Then Delete the account <accountId> and <version> with error "false"

   Examples:
   | accountId                            | status    | version |
   | ad27e265-9605-4b4b-a0e5-3003ea9cc4dc | pending   |   0     |
   | d528cfe8-546a-496c-b086-41c170c56d30 | confirmed |   0     |
   | 79fc881b-9b8f-4e83-bd2d-c518f1a87f48 | failed    |   0     |

Scenario Outline: Validate error message for account creation
   Given Create account with <key> and <value> with validataion
   Then Get error message <error>

   Examples:
   | key        | value    |   error  |
   | accountId  | 123      |  id in body must be of type uuid |
   | status     | Pending  |  status in body should be one of [pending confirmed failed] |
   | orgId      | 123      | organisation_id in body must be of type uuid |
   | country    | G3       | country in body should match |
   | country    | GBR      | country in body should match |

Scenario Outline: Validate error message for account creation
   Given I created account <value> and <status>
   When Create account with <key> and <value> with validataion
   Then Get error message <error>
   Then Delete the account <value> and <version> with error "false"

   Examples:
   | key        | version | value       |  status | error  |
   | accountId  | 0       |d528cfe8-546a-496c-b086-41c170c56d30      |  pending |  Account cannot be created as it violates a duplicate constraint |
   
Scenario Outline: Get Account with error
   Given I get error account <accountId>
   When Get error message <error>

   Examples:
   | accountId                            | error    |
   | ad27e265-9605-4b4b-a0e5-3003ea9cc9dc | record ad27e265-9605-4b4b-a0e5-3003ea9cc9dc does not exist   |
   | ad27e265-9605 | id is not a valid uuid   |

   Scenario Outline: Delete account with 404
      Given Delete the account <accountId> and <version> with error "true"
      When Get error message <error>

      Examples:
      | accountId                            | version  | error    |
      | ad27e265-9605-4b4b-a0e5-3003ea9cc9dc | 0        |  Unknown error, status code: 404   |
   
   Scenario Outline: Delete account with invalid version
      Given I created account <accountId> and <status>
      When Delete the account <accountId> and <version> with error "true"
      Then Get error message <error>
      Then Delete the account <accountId> and 0 with error "false"

      Examples:
      | accountId                            | status  | version    | error             |
      | ad27e265-9605-4b4b-a0e5-3003ea9cc9ac | pending | 345        | invalid version   |
