package pagination

import "context"

type Page[T any] struct {
	Items      []T
	NextCursor string
}

type FetchPage[T any] func(ctx context.Context, cursor string) (Page[T], error)

type Iterator[T any] struct {
	fetch     FetchPage[T]
	cursor    string
	buffer    []T
	exhausted bool
}

func NewIterator[T any](startCursor string, fetch FetchPage[T]) *Iterator[T] {
	return &Iterator[T]{fetch: fetch, cursor: startCursor}
}

func (it *Iterator[T]) Next(ctx context.Context) (T, bool, error) {
	var zero T
	if len(it.buffer) == 0 && !it.exhausted {
		pg, err := it.fetch(ctx, it.cursor)
		if err != nil {
			return zero, false, err
		}
		it.buffer = pg.Items
		it.cursor = pg.NextCursor
		if len(it.buffer) == 0 && it.cursor == "" {
			it.exhausted = true
		}
	}
	if len(it.buffer) == 0 {
		return zero, false, nil
	}
	item := it.buffer[0]
	it.buffer = it.buffer[1:]
	return item, true, nil
}
