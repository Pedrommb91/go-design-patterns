package composite

import (
	"testing"
)

func TestFileOperation(t *testing.T) {
	file := &File{Name: "TestFile"}
	expected := "TestFile"
	if got := file.Operation(); got != expected {
		t.Errorf("File.Operation() = %v, want %v", got, expected)
	}
}

func TestDirectoryOperation(t *testing.T) {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	dir := &Directory{Name: "TestDir"}
	dir.Add(file1)
	dir.Add(file2)

	expected := "TestDir: [File1 File2]"
	if got := dir.Operation(); got != expected {
		t.Errorf("Directory.Operation() = %v, want %v", got, expected)
	}
}

func TestDirectoryAddRemove(t *testing.T) {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	dir := &Directory{Name: "TestDir"}
	dir.Add(file1)
	dir.Add(file2)
	dir.Remove(file1)

	expected := "TestDir: [File2]"
	if got := dir.Operation(); got != expected {
		t.Errorf("Directory.Operation() after remove = %v, want %v", got, expected)
	}
}
func TestCompositeWorkflow(t *testing.T) {
	file1 := &File{Name: "File 1"}
	file2 := &File{Name: "File 2"}
	file3 := &File{Name: "File 3"}

	dir1 := &Directory{Name: "Directory 1"}
	dir1.Add(file1)
	dir1.Add(file2)

	dir2 := &Directory{Name: "Directory 2"}
	dir2.Add(dir1)
	dir2.Add(file3)

	if got := file1.Operation(); got != "File 1" {
		t.Errorf("File.Operation() = %v, want %v", got, "File 1")
	}

	if got := dir1.Operation(); got != "Directory 1: [File 1 File 2]" {
		t.Errorf("Directory.Operation() = %v, want %v", got, "Directory 1: [File 1 File 2]")
	}

	if got := dir2.Operation(); got != "Directory 2: [Directory 1: [File 1 File 2] File 3]" {
		t.Errorf("Directory.Operation() = %v, want %v", got, "Directory 2: [Directory 1: [File 1 File 2] File 3]")
	}
}
