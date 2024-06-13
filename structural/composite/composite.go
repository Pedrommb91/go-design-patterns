package composite

// Component is the interface for all objects in the composition
type Component interface {
	Operation() string
}

// File represents a file in the filesystem
type File struct {
	Name string
}

func (f *File) Operation() string {
	return f.Name
}

// Directory represents a directory that may contain files or other directories
type Directory struct {
	Name     string
	children []Component
}

func (d *Directory) Add(child Component) {
	d.children = append(d.children, child)
}

func (d *Directory) Remove(child Component) {
	for i, comp := range d.children {
		if comp == child {
			d.children = append(d.children[:i], d.children[i+1:]...)
			break
		}
	}
}

func (d *Directory) Operation() string {
	result := d.Name + ": ["
	for _, child := range d.children {
		result += child.Operation() + " "
	}
	result = result[:len(result)-1] + "]"
	return result
}
