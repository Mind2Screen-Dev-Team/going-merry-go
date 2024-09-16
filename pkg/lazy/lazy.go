package lazy

import "sync"

type Supplier[T any] func() (T, error)

// Loader is a generic Lazy loader.
type Loader[T any] struct {
	onc      sync.Once
	set      bool
	val      T
	err      error
	supplier Supplier[T]
}

// New creates a new Loader.
func New[T any](supplier Supplier[T]) Loader[T] {
	return Loader[T]{
		onc:      sync.Once{},
		supplier: supplier,
	}
}

// Value loads if the value is not loaded yet,
// and returns the value.
func (c *Loader[T]) Value() T {
	c.onc.Do(func() {
		val, err := c.supplier()
		c.val = val
		c.err = err
		c.set = true
		// release the supplier, so the GC can collect it.
		c.supplier = nil
	})
	return c.val
}

func (c *Loader[T]) Error() error {
	if !c.Loaded() {
		return nil
	}
	return c.err
}

// Loaded returns true if the value is loaded.
func (c *Loader[T]) Loaded() bool { return c.set }
