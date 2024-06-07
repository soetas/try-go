package collections

type List interface {
	Get(pos int) string
	Set(pos int, val string)
	Insert(pos int, val string)
	Remove(val string)
	RemoveAt(pos int)
	Count(val string)
	Size() int
	IsEmpty() bool
}
