package maps

// A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.
// Therefore, you should never initialize an empty map variable:
// var m map[string]string    like this
// instead:
// var dictionary = map[string]string{}
//
//// OR
//
// var dictionary = make(map[string]string)
// Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime panic.
type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")

)

type DictionaryErr string

// DictionaryErr implements the error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
// A map value is a pointer to a runtime.hmap structure.
// So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	if err != nil {
		return ErrWordDoesNotExist
	}

	d[word] = newDefinition

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
