package cardinal

import (
	"pkg.world.dev/world-engine/cardinal/ecs"
	"pkg.world.dev/world-engine/cardinal/ecs/entity"
)

// Search allowed for the querying of entities within a World.
type Search struct {
	impl *ecs.Search
}

// NewSearch creates a new Search.
func (w *World) NewSearch(filter ecs.Filterable) (*Search, error) {
	q, err := w.implWorld.NewSearch(filter)
	if err != nil {
		return nil, err
	}
	return &Search{q}, nil
}

// SearchCallBackFn represents a function that can operate on a single EntityID, and returns whether the next EntityID
// should be processed.
type SearchCallBackFn func(EntityID) bool

// Each executes the given callback function on every EntityID that matches this search. If any call to callback returns
// falls, no more entities will be processed.
func (q *Search) Each(w *World, callback SearchCallBackFn) {
	q.impl.Each(w.implWorld, func(eid entity.ID) bool {
		return callback(eid)
	})
}

// Count returns the number of entities that match this search.
func (q *Search) Count(w *World) (int, error) {
	return q.impl.Count(w.implWorld)
}

// First returns the first entity that matches this search.
func (q *Search) First(w *World) (id EntityID, err error) {
	return q.impl.First(w.implWorld)
}