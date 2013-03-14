Feature: Delete

    Scenario: one node connected to another
        Given a graph
        And an edge from node 0 to 1
        When delete the edge from node 0 to 1
        Then nodes 0 and 1 are not adjacent

    Scenario: one node not connected to another
        Given a graph
        And an edge from node 0 to 1
        When delete the edge from node 1 to 2
        Then nodes 0 and 2 are not adjacent
        Then nodes 1 and 2 are not adjacent
        And nodes 0 and 1 are adjacent

    Scenario: delete from empty graph
        Given a graph
        When delete the edge from node 0 to 1
