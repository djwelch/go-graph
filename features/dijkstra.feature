Feature: Dijkstra

    Scenario: Test
        Given a graph
        And an edge from node 0 to 1
        And an edge from node 1 to 2
        And an edge from node 2 to 4
        And an edge from node 0 to 3
        And an edge from node 3 to 4
        When I run the dijkstra algorithm from node 0 to node 4
        Then the path is 0,3,4

