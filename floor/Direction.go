package floor

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	return [...]string{"North", "South", "East", "West"}[d]
}