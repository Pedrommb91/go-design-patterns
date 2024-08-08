package memento

// EditorState represents the state of the text editor content.
type EditorState struct {
	content string
}

// Editor is the originator that creates a memento containing a snapshot of its current state.
type Editor struct {
	content string
}

// SetContent sets the content of the editor.
func (e *Editor) SetContent(content string) {
	e.content = content
}

// GetContent returns the content of the editor.
func (e *Editor) GetContent() string {
	return e.content
}

// CreateState creates a memento with the current content of the editor.
func (e *Editor) CreateState() *EditorState {
	return &EditorState{content: e.content}
}

// Restore restores the editor's content from a memento object.
func (e *Editor) Restore(state *EditorState) {
	e.content = state.content
}

// History maintains the history of editor states.
type History struct {
	states []*EditorState
}

// Push adds a state to the history.
func (h *History) Push(state *EditorState) {
	h.states = append(h.states, state)
}

// Pop retrieves the most recent state from the history.
func (h *History) Pop() *EditorState {
	if len(h.states) == 0 {
		return nil
	}
	lastIndex := len(h.states) - 1
	state := h.states[lastIndex]
	h.states = h.states[:lastIndex] // Remove the last state
	return state
}
