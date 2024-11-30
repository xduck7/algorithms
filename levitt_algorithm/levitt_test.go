package main

import (
	l "levitt_algorithm/levitt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Levitt Algorithm", func() {

	var (
		graph  *l.Graph
		levitt *l.Levitt
	)

	BeforeEach(func() {
		graph = l.NewGraph([]int{1, 2, 3, 4, 5})
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 3, 2)
		graph.AddEdge(1, 3, 4)
		graph.AddEdge(3, 4, 1)
		graph.AddEdge(4, 5, 3)

		levitt = l.NewLevitt(graph, []int{1})
	})

	It("should calculate the correct shortest paths from source vertex", func() {
		levitt.Run()

		distances := levitt.GetDistances()

		Expect(distances[1]).To(Equal(0))
		Expect(distances[2]).To(Equal(1))
		Expect(distances[3]).To(Equal(3))
		Expect(distances[4]).To(Equal(4))
		Expect(distances[5]).To(Equal(7))
	})

	It("should handle graphs with cycles correctly", func() {
		graph := l.NewGraph([]int{1, 2, 3})
		graph.AddEdge(1, 2, 2)
		graph.AddEdge(2, 3, 1)
		graph.AddEdge(3, 1, 1)

		levitt := l.NewLevitt(graph, []int{1})
		levitt.Run()

		distances := levitt.GetDistances()

		Expect(distances[1]).To(Equal(0))
		Expect(distances[2]).To(Equal(2))
		Expect(distances[3]).To(Equal(3))
	})

	It("should handle multiple source vertices", func() {
		graph := l.NewGraph([]int{1, 2, 3, 4})
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 3, 2)
		graph.AddEdge(1, 3, 4)
		graph.AddEdge(3, 4, 1)

		levitt := l.NewLevitt(graph, []int{1, 3})
		levitt.Run()

		distances := levitt.GetDistances()

		Expect(distances[1]).To(Equal(0))
		Expect(distances[2]).To(Equal(1))
		Expect(distances[3]).To(Equal(0))
		Expect(distances[4]).To(Equal(1))
	})

	It("should handle graphs with equal weight edges", func() {
		graph := l.NewGraph([]int{1, 2, 3})
		graph.AddEdge(1, 2, 2)
		graph.AddEdge(2, 3, 2)
		graph.AddEdge(1, 3, 2)

		levitt := l.NewLevitt(graph, []int{1})
		levitt.Run()

		distances := levitt.GetDistances()

		Expect(distances[1]).To(Equal(0))
		Expect(distances[2]).To(Equal(2))
		Expect(distances[3]).To(Equal(2))
	})
})

func TestTestingPkg(t *testing.T) {
	t.Parallel()

	RegisterFailHandler(Fail)
	RunSpecs(t, "testing testutil package")
}
