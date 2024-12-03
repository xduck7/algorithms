package avltree

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

func NewNode(value int) *Node {
	return &Node{
		Value:  value,
		Height: 1,
	}
}

func GetHeight(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func UpdateHeight(node *Node) {
	if node != nil {
		node.Height = max(GetHeight(node.Left), GetHeight(node.Right)) + 1
	}
}

func BalanceFactor(node *Node) int {
	if node == nil {
		return 0
	}
	return GetHeight(node.Left) - GetHeight(node.Right)
}

func RotateRight(y *Node) *Node {
	x := y.Left
	T := x.Right
	x.Right = y
	y.Left = T

	UpdateHeight(y)
	UpdateHeight(x)
	return x
}

func RotateLeft(x *Node) *Node {
	y := x.Right
	T := y.Left
	y.Left = x
	x.Right = T

	UpdateHeight(x)
	UpdateHeight(y)
	return y
}

func Balance(node *Node) *Node {
	UpdateHeight(node)
	if BalanceFactor(node) > 1 {
		if BalanceFactor(node.Left) < 0 {
			node.Left = RotateLeft(node.Left)
		}
		return RotateRight(node)
	}
	if BalanceFactor(node) < -1 {
		if BalanceFactor(node.Right) > 0 {
			node.Right = RotateRight(node.Right)
		}
		return RotateLeft(node)
	}
	return node
}

func Insert(node *Node, value int) *Node {
	if node == nil {
		return NewNode(value)
	}
	if value < node.Value {
		node.Left = Insert(node.Left, value)
	} else if value > node.Value {
		node.Right = Insert(node.Right, value)
	} else {
		return node
	}
	return Balance(node)
}

func GetHeightOfTree(root *Node) int {
	return GetHeight(root)
}

func PrintInOrder(node *Node) {
	if node == nil {
		return
	}
	PrintInOrder(node.Left)
	print(node.Value, " ")
	PrintInOrder(node.Right)
}

func ValidateBalance(node *Node) bool {
	if node == nil {
		return true
	}
	balanceFactor := BalanceFactor(node)
	if balanceFactor < -1 || balanceFactor > 1 {
		println("Unbalanced at node", node.Value, "with balance factor", balanceFactor)
		return false
	}
	return ValidateBalance(node.Left) && ValidateBalance(node.Right)
}
