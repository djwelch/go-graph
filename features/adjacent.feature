Feature: Adjacent

    Scenario: one node connected to another
        Given a graph
        And an edge from node 0 to 1
        Then nodes 0 and 1 are adjacent
        And nodes 1 and 0 are not adjacent

    Scenario: two nodes connected to each other
        Given a graph
        And an edge from node 0 to 1
        And an edge from node 1 to 0
        Then nodes 0 and 1 are adjacent
        And nodes 1 and 0 are adjacent

    Scenario: one node connected two others
        Given a graph
        And an edge from node 0 to 1
        And an edge from node 0 to 2
        Then nodes 0 and 1 are adjacent
        And nodes 0 and 2 are adjacent
        And nodes 1 and 2 are not adjacent
