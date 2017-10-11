package logviewer

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewViewer(t *testing.T) {
	viewer := NewViewer("Test Viewer")
	assert.True(t, viewer != nil, "Viewer shouldn't be nil")
}
