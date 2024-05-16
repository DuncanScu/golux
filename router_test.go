package golux_test

import (
	"testing"

	"github.com/DuncanScu/golux"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	sut := golux.NewRouter()
	assert.NotNil(t, sut.PathTree)
}
