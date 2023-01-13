package BinaryTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data = []int{1, 2, 3, 4, 5}

func TestTree_Insert(t *testing.T) {

	var tree Node
	for i := 0; i < len(data); i++ {
		tree.Insert(&tree, data[i])
	}
	assert.NotEmpty(t, tree, "Tree should not be empty")
}

func TestTree_Find(t *testing.T) {
	var data = []int{7, 8, 1, 3, 9}
	var tree Node
	for i := 0; i < len(data); i++ {
		tree.Insert(&tree, data[i])
	}
	_, success := tree.Find(&tree, 4)
	assert.False(t, success)
	tree.Insert(&tree, 4)
	_, success = tree.Find(&tree, 4)
	assert.True(t, success)
}
