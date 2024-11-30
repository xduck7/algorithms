package main

import (
	l "levitt_algorithm/levitt"
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Levitt Algorithm", func() {

	var (
		graph  *l.Graph
		levitt *l.Levitt
	)

	ginkgo.BeforeEach(func() {
		graph = l.NewGraph([]int{1, 2, 3, 4, 5})
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 3, 2)
		graph.AddEdge(1, 3, 4)
		graph.AddEdge(3, 4, 1)
		graph.AddEdge(4, 5, 3)

		levitt = l.NewLevitt(graph, []int{1})
	})

	ginkgo.It("should calculate the correct shortest paths from source vertex", func() {
		levitt.Run()

		distances := levitt.GetDistances()

		gomega.Expect(distances[1]).To(gomega.Equal(0))
		gomega.Expect(distances[2]).To(gomega.Equal(1))
		gomega.Expect(distances[3]).To(gomega.Equal(3))
		gomega.Expect(distances[4]).To(gomega.Equal(4))
		gomega.Expect(distances[5]).To(gomega.Equal(7))
	})

	ginkgo.It("should handle graphs with cycles correctly", func() {
		graph := l.NewGraph([]int{1, 2, 3})
		graph.AddEdge(1, 2, 2)
		graph.AddEdge(2, 3, 1)
		graph.AddEdge(3, 1, 1)

		levitt := l.NewLevitt(graph, []int{1})
		levitt.Run()

		distances := levitt.GetDistances()

		gomega.Expect(distances[1]).To(gomega.Equal(0))
		gomega.Expect(distances[2]).To(gomega.Equal(2))
		gomega.Expect(distances[3]).To(gomega.Equal(3))
	})

	ginkgo.It("should handle multiple source vertices", func() {
		graph := l.NewGraph([]int{1, 2, 3, 4})
		graph.AddEdge(1, 2, 1)
		graph.AddEdge(2, 3, 2)
		graph.AddEdge(1, 3, 4)
		graph.AddEdge(3, 4, 1)

		levitt := l.NewLevitt(graph, []int{1, 3})
		levitt.Run()

		distances := levitt.GetDistances()

		gomega.Expect(distances[1]).To(gomega.Equal(0))
		gomega.Expect(distances[2]).To(gomega.Equal(1))
		gomega.Expect(distances[3]).To(gomega.Equal(0))
		gomega.Expect(distances[4]).To(gomega.Equal(1))
	})

	ginkgo.It("should handle graphs with equal weight edges", func() {
		graph := l.NewGraph([]int{1, 2, 3})
		graph.AddEdge(1, 2, 2)
		graph.AddEdge(2, 3, 2)
		graph.AddEdge(1, 3, 2)

		levitt := l.NewLevitt(graph, []int{1})
		levitt.Run()

		distances := levitt.GetDistances()

		gomega.Expect(distances[1]).To(gomega.Equal(0))
		gomega.Expect(distances[2]).To(gomega.Equal(2))
		gomega.Expect(distances[3]).To(gomega.Equal(2))
	})
})

func TestTestingPkg(t *testing.T) {
	t.Parallel()

	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "testing testutil package")
}
