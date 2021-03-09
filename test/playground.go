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

// 
func PostorderTraversalV2(root *TreeNode) {
	if root == nil {
		return
	}

	var current *TreeNode
	var preNode *TreeNode
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	
	for len(stack) != nil {
		current = stack[len(stack) -1]
		if (current.Left == nil && current.Right == nil) || (preNode != nil && (preNode == current.Left || preNode == current.Right )) {
			// do something to currentNode
			stack = stack[:len(stack) -1]
			preNode = current
		}else {
			// 因为是后序遍历，所以要注意顺序 先入栈右孩子 再入栈左孩子
			if current.Right != nil {
				stack = append(stack, current.Right)
			}

			if current.Left != nil {
				stack = append(stack, current.Left)
			}

		}
	}
}

// 修改前序遍历
func PostorderTraversalV3(root *TreeNode) {

	stack := make([]*TreeNode, 0)
	list := make([]*TreeNode, 0)

	for root!= nil || len(stack) != 0 {
		if root!=nil {
			stack = append(stack, root)
			list = append(stack, root)
			root = root.Right
		}else {
			temp := stack[len(stack) -1]
			stack = stack[:len(stack) -1]
			root = temp.Left
		}
	}

	for i:= len(list) -1 ; i >= 0 ; i -- {
		// do something to list[i]
	}

}