package tree_test

import (
	"testing"

	"github.com/DuncanScu/golux/tree"
	"github.com/stretchr/testify/assert"
)

func TestNewPathTree(t *testing.T) {
	sut := tree.NewPathTree()
	assert.NotNil(t, sut.Root)
}
