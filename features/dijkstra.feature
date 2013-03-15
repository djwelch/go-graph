Feature: Dijkstra

    Scenario: Test
        Given a graph
        And an edge from node 0 to 1 with distance 2
        And an edge from node 1 to 2 with distance 2
        And an edge from node 2 to 4 with distance 1
        And an edge from node 0 to 3 with distance 3
        And an edge from node 3 to 4 with distance 1
        When I run the dijkstra algorithm from node 0 to node 4
        Then the path is 0,3,4

