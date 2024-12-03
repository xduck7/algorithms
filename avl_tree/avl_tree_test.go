package avltree

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AVL Tree", func() {
	var root *Node

	BeforeEach(func() {
		root = nil
	})

	It("should create a tree and calculate its height", func() {
		root = Insert(root, 10)
		root = Insert(root, 20)
		root = Insert(root, 30)
		Expect(GetHeightOfTree(root)).To(Equal(2))
	})

	It("should balance the tree after insertion", func() {
		root = Insert(root, 30)
		root = Insert(root, 20)
		root = Insert(root, 10)
		Expect(root.Value).To(Equal(20))
		Expect(root.Left.Value).To(Equal(10))
		Expect(root.Right.Value).To(Equal(30))
		Expect(GetHeightOfTree(root)).To(Equal(2))
	})

	It("should handle multiple insertions and maintain balance", func() {
		values := []int{50, 20, 60, 10, 25, 70}
		for _, v := range values {
			root = Insert(root, v)
			fmt.Printf("After inserting %d:\n", v)
			fmt.Print("In-order traversal: ")
			PrintInOrder(root)
			fmt.Println()
			Expect(ValidateBalance(root)).To(BeTrue())
		}
		Expect(GetHeightOfTree(root)).To(Equal(3))
		Expect(root.Value).To(Equal(50))
		Expect(root.Left.Value).To(Equal(20))
		Expect(root.Right.Value).To(Equal(60))
	})

	It("should return height of an empty tree as zero", func() {
		Expect(GetHeightOfTree(root)).To(Equal(0))
	})

	It("should maintain balance for sequential insertions", func() {
		for i := 1; i <= 10; i++ { // 10 узлов
			root = Insert(root, i)
			Expect(ValidateBalance(root)).To(BeTrue())
		}
		Expect(GetHeightOfTree(root)).To(Equal(4))
	})
})

func TestTestingPkg(t *testing.T) {
	t.Parallel()

	RegisterFailHandler(Fail)
	RunSpecs(t, "testing testutil package")
}
