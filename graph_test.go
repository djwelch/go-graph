package graph

import (
    . "github.com/djwelch/go-gherkin"
    . "github.com/djwelch/go-matchers"
    "strings"
    "testing"
)

var g Interface
var gpath []interface{}

type mockEdge struct {
    distance int
}

func (e *mockEdge) Distance() int {
    return e.distance
}

func GivenAGraph(w *World) {
    g = New()
    AssertThat(w, g, Not(Equals(nil)))
}

func GivenAnEdge(w *World, n0 string, n1 string) {
    g.Add(n0, n1)
}

func GivenAnEdgeDist(w *World, n0 string, n1 string, d int) {
    g.AddWithEdge(n0, n1, &mockEdge{d})
}

func ThenNodesAreAdjacent(w *World, n0 string, n1 string) {
    result := g.Adjacent(n0, n1)
    AssertThat(w, result, IsTrue)
}

func ThenNodesAreNotAdjacent(w *World, n0 string, n1 string) {
    result := g.Adjacent(n0, n1)
    AssertThat(w, result, IsFalse)
}

func WhenDeleteAnEdge(w *World, n0 string, n1 string) {
    g.Delete(n0, n1)
}

func ThenNodeIsANeighbor(w *World, n0 string, n1 string) {
    _, ok := g.Neighbors(n0)[n1]
    AssertThat(w, ok, IsTrue)
}

func ThenNodeIsNotANeighbor(w *World, n0 string, n1 string) {
    _, ok := g.Neighbors(n0)[n1]
    AssertThat(w, ok, IsFalse)
}

func WhenIRunDijkstra(w *World, n0 string, n1 string) {
    path := Dijkstra(g, n0, n1)
    gpath = make([]interface{}, len(path))
    for i, v := range path {
        gpath[i] = interface{}(v)
    }
}

func WhenIRunDijkstraFails(w *World, n0 string, n1 string) {
    panicked := false
    func() {
        defer func() {
            if r := recover(); r != nil {
                panicked = true
            }
        }()
        Dijkstra(g, n0, n1)
    }()
    AssertThat(w, panicked, IsTrue)
}

func ThenThePathIs(w *World, p string) {
    generic := []interface{} { }
    if p != "empty" {
        splits := strings.Split(p, ",")
        generic = make([]interface{}, len(splits))
        for i, v := range splits {
            generic[i] = interface{}(v)
        }
    }
    AssertThat(w, gpath, HasExactly(generic...))
}

func Test(t *testing.T) {
    Given("a graph", GivenAGraph)
    Given("an edge from node ([^ ]*) to ([^ ]*)$", GivenAnEdge)
    Given("an edge from node ([^ ]*) to ([^ ]*) with distance ([^ ]*)", 
        GivenAnEdgeDist)
    When("delete the edge from node ([^ ]*) to ([^ ]*)", WhenDeleteAnEdge)
    Then("nodes ([^ ]*) and ([^ ]*) are adjacent", ThenNodesAreAdjacent)
    Then("nodes ([^ ]*) and ([^ ]*) are not adjacent", ThenNodesAreNotAdjacent)
    Then("node ([^ ]*) has neighbor node (.*)", ThenNodeIsANeighbor)
    Then("node ([^ ]*) does not have neighbor node ([^ ]*)", ThenNodeIsNotANeighbor)
    When("I run the dijkstra algorithm from node ([^ ]*) to node ([^ ]*)$",
        WhenIRunDijkstra)
    When("I run the dijkstra algorithm from node ([^ ]*) to node ([^ ]*) it fails",
        WhenIRunDijkstraFails)
    Then("the path is (.*)", ThenThePathIs)
    TearDown(func() { g = nil })
    Run(t)
}

