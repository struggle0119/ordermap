package ordermap

// element include key and value, reduce for loop
type Element struct {
	key, value interface{}
}

// build element
func newElement(key, value interface{}) *Element {
	return &Element{
		key:   key,
		value: value,
	}
}
