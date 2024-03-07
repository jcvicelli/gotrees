package tree

type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func NewTreeNode[T any](value T) *Node[T] {
	return &Node[T]{Val: value}
}

func NewTree[T any]() *Node[string] {
	nodeA := NewTreeNode("A")
	nodeB := NewTreeNode("B")
	nodeC := NewTreeNode("C")
	nodeD := NewTreeNode("D")
	nodeE := NewTreeNode("E")
	nodeF := NewTreeNode("F")
	nodeG := NewTreeNode("G")

	nodeA.Left = nodeB
	nodeA.Right = nodeC
	nodeB.Left = nodeD
	nodeB.Right = nodeE
	nodeC.Right = nodeF
	nodeF.Right = nodeG

	return nodeA
}

func NewTreeNumber[T any]() *Node[int] {
	nodeA := NewTreeNode(3)
	nodeB := NewTreeNode(11)
	nodeC := NewTreeNode(4)
	nodeD := NewTreeNode(4)
	nodeE := NewTreeNode(-2)
	nodeF := NewTreeNode(1)

	nodeA.Left = nodeB
	nodeA.Right = nodeC
	nodeB.Left = nodeD
	nodeB.Right = nodeE
	nodeC.Right = nodeF

	return nodeA
}

func DepthFirstValues(root *Node) []string {

	result := []string{root.Val}

	if root.Left == nil && root.Right == nil {
		return result
	}

	//stack!
	stack := []*Node{root.Right, root.Left}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		result = append(result, item.Val)
		//remove from stack, add leafs
		stack = stack[:len(stack)-1]
		if item.Right != nil {
			stack = append(stack, item.Right)
		}
		if item.Left != nil {
			stack = append(stack, item.Left)
		}

	}

	return result
}

func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func DepthFirstValuesRecursive(root *Node) []string {
	var result []string

	// Helper function using recursion
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root)
	return result
}

func BreadthFirstValues(root *Node) []string {

	var result = []string{}

	if root == nil {
		return result
	}

	//queue!
	queue := []*Node{root}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		result = append(result, item.Val)

		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}

	}

	return result
}
