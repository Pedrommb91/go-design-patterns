package memento

import (
	"testing"

	"gotest.tools/assert"
)

func TestMementoPattern(t *testing.T) {
	editor := Editor{}
	history := History{}

	editor.SetContent("First State")
	history.Push(editor.CreateState())

	editor.SetContent("Second State")
	history.Push(editor.CreateState())

	editor.SetContent("Third State")

	assert.Equal(t, "Third State", editor.GetContent(), "Expecting Third State")

	// Restore to the second state
	secondState := history.Pop()
	if secondState != nil {
		editor.Restore(secondState)
	}
	assert.Equal(t, "Second State", editor.GetContent(), "Expecting Second State")

	// Restore to the first state
	firstState := history.Pop()
	if firstState != nil {
		editor.Restore(firstState)
	}
	assert.Equal(t, "First State", editor.GetContent(), "Expecting First State")
}

func TestMementoPattern_EmptyPop(t *testing.T) {
	editor := Editor{}
	history := History{}

	// Attempt to restore from an empty history (should not change the editor's content)
	editor.SetContent("Initial State")
	restoredState := history.Pop()
	assert.Assert(t, restoredState == nil, "Expecting nil")
}
