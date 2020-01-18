//字典树
package meta

type TrieNode struct {
	Val    byte
	IsWord bool

	Children []*TrieNode
}

func NewTrieNode(val byte) *TrieNode {
	return &TrieNode{
		Val:      val,
		IsWord:   true,
		Children: nil,
	}
}

func (root *TrieNode) TrieBuild(strs []string) *TrieNode {
	if root == nil {
		root = NewTrieNode(nil)
	}

	if len(strs) == 0 {
		return root
	}

	for str := range strs {
		root.TrieBuildByString(str)
	}

	return root
}

func (root *TrieNode) TrieBuildByString(str string) {
	if len(str) == 0 {
		return
	}

	var (
		cur *TrieNode
		pre *TrieNode
	)
	pre = root
	for _, b := range str {
		cur = pre.findInChildren(b)

		if cur == nil {
			pre.Children = append(root.Children, NewTrieNode(b))
			pre.IsWord = false
			continue
		}

		pre = cur
	}

}

func (root *TrieNode) findInChildren(b byte) *TrieNode {
	if len(root.Children) == 0 {
		return nil
	}

	for _, child := range root.Children {
		if child.Val == b {
			return child
		}
	}

	return nil
}
