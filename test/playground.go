package test


type TreeNode struct {	
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func frontOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0)

	for root!=nil || len(stack) != 0 {
		if root != nil {
			// Do something
			stack = append(stack, root)
			root = root.Left
		}else {
			root = stack[len(stack) - 1]
			stack = stack[:len(stack)-1]
			root  = root.Right
		}
	}
} 

func middleOrderV1(root *TreeNode) {
	stack := make([]*TreeNode, 0)

	for root != nil || len(root) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		}else {
			root = stack[len(stack) - 1]
			stack = stack[:len(stack) -1]
			// Do something
			root = root.Right
		}
	}
}

func middleOrderV2(root *TreeNode) {
	stack := make([]*TreeNode, 0)

	for root != nil || len(root) != 0 {
		for root!= nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) -1]
		stack = stack[:len(stack) -1]
		root = root.Right
	}
}

func afterOrderV1(root *TreeNode) {
	type StateTreeNode struct {
		node *TreeNode
		isFirst *TreeNode
	}

	tempStateNode := &StateTreeNode{}
	pre := root
	StateTreeNodeStack := make([]*StateTreeNode, 0)
	
	for pre != nil || len(StateTreeNodeStack) != 0 {
		if pre != nil {
			// 第一次访问节点
			StateTreeNodeStack = append(StateTreeNodeStack, &StateTreeNode{
				node: pre,
				isFirst: true
			})
			pre = pre.Left
		}else {
			// 出栈
			tempStateNode = StateTreeNodeStack[len(StateTreeNodeStack) -1]
			// 第一次出栈，第二次访问节点
			if tempStateNode.isFirst {
				tempStateNode.isFirst = false
				pre = tempStateNode.node.Right
			}else {
				// Do something
				StateTreeNodeStack = StateTreeNodeStack[:len(StateTreeNodeStack) -1]
				pre = nil // 第三次访问节点，应该置为nil
			}

		}
	}
}