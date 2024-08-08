package iterator

// Iterator provides a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Collection defines a standard collection that can create an Iterator.
type Collection interface {
	CreateIterator() Iterator
}

// Book is a simple struct representing a book in a collection.
type Book struct {
	Title  string
	Author string
}

// BookShelf is a collection of books that can be iterated over.
type BookShelf struct {
	Books []Book
}

// CreateIterator returns an iterator for the bookshelf.
func (b *BookShelf) CreateIterator() Iterator {
	return &BookShelfIterator{
		bookShelf: b,
		current:   0,
	}
}

// BookShelfIterator is the concrete iterator for the BookShelf collection.
type BookShelfIterator struct {
	bookShelf *BookShelf
	current   int
}

// HasNext checks if there is a next item in the collection.
func (i *BookShelfIterator) HasNext() bool {
	return i.current < len(i.bookShelf.Books)
}

// Next returns the next item in the collection.
func (i *BookShelfIterator) Next() interface{} {
	if i.HasNext() {
		book := i.bookShelf.Books[i.current]
		i.current++
		return book
	}
	return nil
}
