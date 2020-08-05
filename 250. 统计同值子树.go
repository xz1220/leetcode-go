/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func countUnivalSubtrees(root *TreeNode) int {

    var count int
    var dfs func(*TreeNode)bool

    dfs=func(root *TreeNode)bool{
        if root==nil{
            return true
        }

        left:=dfs(root.Left)
        right:=dfs(root.Right)

        var flag bool=false
        if left && right{
            if root.Left==nil && root.Right==nil{
            flag=true
            count++
            }else if root.Left==nil && root.Right!=nil{
                if root.Val==root.Right.Val{
                    flag=true
                    count++
                }
            }else if root.Left!=nil && root.Right==nil{
                if root.Val==root.Left.Val{
                    flag=true
                    count++
                }
            }else{
                if root.Val==root.Left.Val && root.Val==root.Right.Val{
                    flag=true
                    count++
                }
            }
        }
        

        return  flag
    }

    dfs(root)

    return count
}