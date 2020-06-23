package array

type DynamicArray struct {
	size     int
	elements []int
}

func NewDynamicArray(size int) *DynamicArray {
	return &DynamicArray{size: size}
}

func (array *DynamicArray) Clear() {
	array.size = 0
}

func (array *DynamicArray) Size() int {
	return array.size
}

func (array *DynamicArray) IsEmpty() bool {
	return array.size == 0
}

func (array *DynamicArray) Contains(element int) bool {
	for i := 0; i < array.size; i++ {
		if array.elements[i] == element {
			return true
		}
	}
	return false
}

func (array *DynamicArray) Get(index int) int {
	if index < 0 || index > array.size {
		return -1
	}
	return array.elements[index]
}

func (array *DynamicArray) Set(index, element int) int {
	if index < 0 || index > array.size {
		return -1
	}
	oldElement := array.elements[index]
	array.elements[index] = element
	return oldElement
}

func (array *DynamicArray) IndexOf(element int) int {
	for i := 0; i < array.size; i++ {
		if array.elements[i] == element {
			return i
		}
	}
	return -1
}

func (array *DynamicArray) AddElement(element int) {

}

func (array *DynamicArray) AddElementByIndex(index, element int) int {
	return 0
}

func (array *DynamicArray) Remove(index int) int {
	return 0
}
