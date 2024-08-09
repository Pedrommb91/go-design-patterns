package template

// Game defines the skeleton of the algorithm to play a game.
// It is a template method defining the order of the steps to execute the game.
type Game interface {
	Start()
	TakeTurn()
	End()
	HasEnded() bool
}

// PlayOneGame is the template method.
func PlayOneGame(g Game) {
	g.Start()
	for !g.HasEnded() {
		g.TakeTurn()
	}
	g.End()
}

// Chess implements the Game interface and provides the specific details
// for playing a game of chess.
type Chess struct {
	turn     int
	MaxTurns int
}

func (c *Chess) Start() {
	// Initialize game of chess
	c.turn = 1
	println("Start a game of chess")
}

func (c *Chess) TakeTurn() {
	// Process a turn for the player
	println("Turn", c.turn, "taken in chess game")
	c.turn++
}

func (c *Chess) End() {
	// End the game of chess
	println("Game of chess ended")
}

func (c *Chess) HasEnded() bool {
	// Determine if the game of chess has ended
	return c.turn > c.MaxTurns
}
