package main

type file struct {
	name string
}

func (f *file) search(keyword string) {
	//TODO: output a String "Searching for keyword "test" in file "file"
}

func (f *file) getName() string {
	return f.name
}
