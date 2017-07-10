package db

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
