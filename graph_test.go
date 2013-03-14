package graph

import (
    . "github.com/djwelch/go-gherkin"
    . "github.com/tychofreeman/go-matchers"
    "testing"
)

var g Interface

func GivenAGraph(w *World) {
    g = New()
    AssertThat(w, g, Not(Equals(nil)))
}

func GivenAnEdge(w *World, n0 string, n1 string) {
    g.Add(n0, n1)
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

func Test(t *testing.T) {
    Given("a graph", GivenAGraph)
    Given("an edge from node (.*) to (.*)", GivenAnEdge)
    When("delete the edge from node (.*) to (.*)", WhenDeleteAnEdge)
    Then("nodes (.*) and (.*) are adjacent", ThenNodesAreAdjacent)
    Then("nodes (.*) and (.*) are not adjacent", ThenNodesAreNotAdjacent)
    Then("node (.*) has neighbor node (.*)", ThenNodeIsANeighbor)
    Then("node (.*) does not have neighbor node (.*)", ThenNodeIsNotANeighbor)
    TearDown(func() { g = nil })
    Run()
}

