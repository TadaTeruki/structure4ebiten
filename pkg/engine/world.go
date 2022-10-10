package engine

type World struct {
	Width  int
	Height int
	DEM    [][]float64
}

func NewWorld(width, height int) *World {
	world := new(World)
	world.Width = width
	world.Height = height
	world.DEM = make([][]float64, world.Height)
	for y := 0; y < world.Height; y++ {
		world.DEM[y] = make([]float64, world.Width)
		for x := 0; x < world.Width; x++ {
			world.DEM[y][x] = 0
		}
	}
	return world
}
