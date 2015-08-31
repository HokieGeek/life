package cgol

import (
	// "fmt"
	"bytes"
	"strconv"
)

type PondStatus int

const (
	Active PondStatus = 1
	Stable PondStatus = 2
	Dead   PondStatus = 3
)

func (t PondStatus) String() string {
	s := ""

	if t&Active == Active {
		s += "Active"
	} else if t&Stable == Stable {
		s += "Stable"
	} else if t&Dead == Dead {
		s += "Dead"
	}

	return s
}

type NeighborsSelector int

const (
	NEIGHBORS_ORTHOGONAL NeighborsSelector = 1
	NEIGHBORS_OBLIQUE    NeighborsSelector = 2
	NEIGHBORS_ALL        NeighborsSelector = 3
)

func (t NeighborsSelector) String() string {
	s := ""

	if t&NEIGHBORS_ORTHOGONAL == NEIGHBORS_ORTHOGONAL {
		s += "NEIGHBORS_ORTHOGONAL"
	} else if t&NEIGHBORS_OBLIQUE == NEIGHBORS_OBLIQUE {
		s += "NEIGHBORS_OBLIQUE"
	} else if t&NEIGHBORS_ALL == NEIGHBORS_ALL {
		s += "NEIGHBORS_ALL"
	}

	return s
}

type Pond struct {
	gameboard         *Gameboard
	NumLiving         int
	Status            PondStatus
	neighborsSelector NeighborsSelector
	living            map[int]map[int]GameboardLocation
}

func (t *Pond) GetNeighbors(organism GameboardLocation) []GameboardLocation {
	switch {
	case t.neighborsSelector == NEIGHBORS_ORTHOGONAL:
		return t.gameboard.GetOrthogonalNeighbors(organism)
	case t.neighborsSelector == NEIGHBORS_OBLIQUE:
		return t.gameboard.GetObliqueNeighbors(organism)
	case t.neighborsSelector == NEIGHBORS_ALL:
		return t.gameboard.GetAllNeighbors(organism)
	}

	return make([]GameboardLocation, 0)
}

func (t *Pond) isOrganismAlive(organism GameboardLocation) bool {
	return (t.GetOrganismValue(organism) >= 0)
}

func (t *Pond) GetNumLiving() int {
	return len(t.living)
}

func (t *Pond) GetOrganismValue(organism GameboardLocation) int {
	// fmt.Printf("\tgetNeighborCount(%s)\n", organism.String())
	return t.gameboard.GetValue(organism)
}

func (t *Pond) setOrganismValue(organism GameboardLocation, numNeighbors int) {
	// fmt.Printf("\tsetNeighborCount(%s, %d)\n", organism.String(), numNeighbors)
	originalNumNeighbors := t.GetOrganismValue(organism)

	// Write the value to the gameboard
	t.gameboard.SetValue(organism, numNeighbors)

	// Update the living count if organism changed living state
	if originalNumNeighbors < 0 && numNeighbors >= 0 {
		// TODO: add to 'living'
		t.NumLiving++
	} else if originalNumNeighbors >= 0 && numNeighbors < 0 {
		t.NumLiving--
		// TODO: remove from 'living'
	}
}

func (t *Pond) calculateNeighborCount(organism GameboardLocation) (int, []GameboardLocation) {
	numNeighbors := 0
	neighbors := t.GetNeighbors(organism)
	for _, neighbor := range neighbors {
		if t.isOrganismAlive(neighbor) {
			numNeighbors++
		}
	}
	return numNeighbors, neighbors
}

func (t *Pond) init(initialLiving []GameboardLocation) {
	// Initialize the first organisms and set their neighbor counts
	t.living = make(map[int]map[int]GameboardLocation)
	for _, organism := range initialLiving {
		// TODO: this logic needs to move into its own place function with channel accessors
		_, keyExists := t.living[organism.Y]
		if !keyExists {
			t.living[organism.Y] = make(map[int]GameboardLocation)
		}
		t.living[organism.Y][organism.X] = organism
		t.setOrganismValue(organism, 0)
	}
}

func (t *Pond) String() string {
	var buf bytes.Buffer
	buf.WriteString("Neighbor selection: ")
	buf.WriteString(t.neighborsSelector.String())
	buf.WriteString("\nLiving organisms: ")
	buf.WriteString(strconv.Itoa(t.NumLiving))
	buf.WriteString("\tStatus: ")
	buf.WriteString(t.Status.String())
	buf.WriteString("\n")
	buf.WriteString(t.gameboard.String())

	return buf.String()
}

func NewPond(rows int, cols int, neighbors NeighborsSelector) *Pond {
	p := new(Pond)

	// Create values
	p.NumLiving = 0
	p.Status = Active

	// Add the given values
	p.gameboard = NewGameboard(GameboardDims{Height: rows, Width: cols})
	p.neighborsSelector = neighbors

	return p
}