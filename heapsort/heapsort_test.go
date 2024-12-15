package heapsort_test

import (
	"heapsort/heapsort"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Heapsort", func() {
	It("должен корректно сортировать массив", func() {
		arr := []int{4, 10, 3, 5, 1}
		expected := []int{1, 3, 4, 5, 10}
		heapsort.Heapsort(arr)
		Expect(arr).To(Equal(expected))
	})

	It("должен обрабатывать пустой массив", func() {
		arr := []int{}
		expected := []int{}
		heapsort.Heapsort(arr)
		Expect(arr).To(Equal(expected))
	})

	It("должен обрабатывать массив из одного элемента", func() {
		arr := []int{42}
		expected := []int{42}
		heapsort.Heapsort(arr)
		Expect(arr).To(Equal(expected))
	})

	It("должен корректно сортировать уже отсортированный массив", func() {
		arr := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		heapsort.Heapsort(arr)
		Expect(arr).To(Equal(expected))
	})

	It("должен корректно сортировать массив с одинаковыми элементами", func() {
		arr := []int{5, 5, 5, 5, 5}
		expected := []int{5, 5, 5, 5, 5}
		heapsort.Heapsort(arr)
		Expect(arr).To(Equal(expected))
	})
})

func TestTestingPkg(t *testing.T) {
	t.Parallel()

	RegisterFailHandler(Fail)
	RunSpecs(t, "testing testutil package")
}
