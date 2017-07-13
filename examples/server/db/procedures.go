package db

// NOTE: The process of creating Access Types: first create your basic access methods e.g. read, write, etc.
// Then create function types matching the function signatures of the access methods
// Then create fabric methods for those function types, and convert the access methods to those function types
// Now you have all your original access methods but in a fabricated format.

func myRead(e *ElementNode) (*ElementNode, error) {
	return e, nil
}

// MyReadFunc is the type conversion of MyRead to an Access Type
var MyReadFunc = ElementRead(myRead)

func myOtherRead(e *ElementNode) (*ElementNode, error) {
	return e, nil
}

// MyOtherReadFunc is the type conversion of MyOtherRead to an Access Type
var MyOtherReadFunc = ElementRead(myOtherRead)
