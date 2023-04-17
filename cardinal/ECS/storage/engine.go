package storage

import (
	"github.com/argus-labs/cardinal/ECS/component"
	"github.com/argus-labs/cardinal/ECS/entity"
	"github.com/argus-labs/cardinal/ECS/filter"
)

type ComponentStorage interface {
	PushComponent(component component.IComponentType, index ArchetypeIndex) error
	Component(archetypeIndex ArchetypeIndex, componentIndex ComponentIndex) ([]byte, error)
	SetComponent(ArchetypeIndex, ComponentIndex, []byte) error
	MoveComponent(ArchetypeIndex, ComponentIndex, ArchetypeIndex) error
	SwapRemove(archetypeIndex ArchetypeIndex, componentIndex ComponentIndex) ([]byte, error)
	Contains(archetypeIndex ArchetypeIndex, componentIndex ComponentIndex) (bool, error)
}

type ComponentStorageManager interface {
	GetComponentStorage(cid component.TypeID) ComponentStorage
	GetComponentIndexStorage(cid component.TypeID) ComponentIndexStorage
}

type ComponentIndexStorage interface {
	ComponentIndex(ArchetypeIndex) (ComponentIndex, bool, error)
	SetIndex(ArchetypeIndex, ComponentIndex) error
	IncrementIndex(ArchetypeIndex) error
	DecrementIndex(ArchetypeIndex) error
}

type EntityLocationStorage interface {
	ContainsEntity(entity.ID) (bool, error)
	Remove(entity.ID) error
	Insert(entity.ID, ArchetypeIndex, ComponentIndex) error
	Set(entity.ID, *Location) error
	Location(entity.ID) (*Location, error)
	ArchetypeIndex(id entity.ID) ArchetypeIndex
	ComponentIndexForEntity(id entity.ID) ComponentIndex
	Len() (int, error)
}

type ArchetypeComponentIndex interface {
	Push(layout *Layout)
	SearchFrom(filter filter.LayoutFilter, start int) *ArchetypeIterator
	Search(layoutFilter filter.LayoutFilter) *ArchetypeIterator
}

type ArchetypeAccessor interface {
	PushArchetype(index ArchetypeIndex, layout *Layout)
	Archetype(index ArchetypeIndex) ArchetypeStorage
	Count() int
}

type ArchetypeStorage interface {
	Layout() *Layout
	Entities() []entity.Entity
	SwapRemove(entityIndex int) entity.Entity
	LayoutMatches(components []component.IComponentType) bool
	PushEntity(entity entity.Entity)
	Count() int
}

type EntryStorage interface {
	SetEntry(entity.ID, *Entry) error
	GetEntry(entity.ID) (*Entry, error)
	SetEntity(entity.ID, Entity)
	SetLocation(entity.ID, Location)
}

type EntityManager interface {
	Destroy(Entity)
	NewEntity() (Entity, error)
}
