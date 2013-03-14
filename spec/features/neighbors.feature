Feature: Neighbors

    Scenario: one node connected to another
        Given a graph
        And an edge from node 0 to 1
        Then node 0 has neighbor node 1
        And node 1 does not have neighbor node 0

    Scenario: empty graph
        Given a graph
        Then node 0 does not have neighbor node 1
