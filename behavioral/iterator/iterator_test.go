package iterator

import (
	"testing"
)

func TestBookShelfIterator(t *testing.T) {
	books := []Book{
		{Title: "Design Patterns", Author: "Erich Gamma"},
		{Title: "Clean Code", Author: "Robert C. Martin"},
		{Title: "Refactoring", Author: "Martin Fowler"},
	}

	bookShelf := &BookShelf{Books: books}
	iterator := bookShelf.CreateIterator()

	for i := 0; iterator.HasNext(); i++ {
		book := iterator.Next()
		if book != books[i] {
			t.Errorf("Expected book %v, got %v", books[i], book)
		}
	}

	if iterator.HasNext() {
		t.Errorf("Expected false, got true")
	}

	if next := iterator.Next(); next != (Book{}) {
		t.Errorf("Expected empty sturct, got %v", next)
	}
}
