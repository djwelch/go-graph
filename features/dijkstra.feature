Feature: Dijkstra

    Scenario: Finds the optimal path
        Given a graph
        And an edge from node 0 to 1 with distance 2
        And an edge from node 1 to 2 with distance 2
        And an edge from node 2 to 4 with distance 1
        And an edge from node 0 to 3 with distance 3
        And an edge from node 3 to 4 with distance 1
        When I run the dijkstra algorithm from node 0 to node 4
        Then the path is 0,3,4

    Scenario: Finds the optimal path
        Given a graph
        And an edge from node 0 to 1 with distance 2
        When I run the dijkstra algorithm from node 0 to node 1
        Then the path is 0,1

    Scenario: Finds the optimal path
        Given a graph
        And an edge from node 0 to 1 with distance 3
        And an edge from node 0 to 2 with distance 1
        And an edge from node 2 to 1 with distance 1
        When I run the dijkstra algorithm from node 0 to node 1
        Then the path is 0,2,1

    Scenario: Doesn't fail with empty graph
        Given a graph
        When I run the dijkstra algorithm from node 0 to node 1
        Then the path is empty

    Scenario: Doesn't fail trying to go to non-existant destination
        Given a graph
        And an edge from node 0 to 1 with distance 3
        When I run the dijkstra algorithm from node 0 to node 2
        Then the path is empty

    Scenario: Graphs are directional
        Given a graph
        And an edge from node 0 to 1 with distance 3
        When I run the dijkstra algorithm from node 1 to node 0
        Then the path is empty

    Scenario: Fails when an edge doesn't implement EdgeDistance
        Given a graph
        And an edge from node 0 to 1
        When I run the dijkstra algorithm from node 0 to node 1 it fails

