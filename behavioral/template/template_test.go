package template

import "testing"

func TestChessGame(t *testing.T) {
	chess := &Chess{MaxTurns: 1}

	chess.Start()
	if chess.turn != 1 {
		t.Errorf("Expected turn to be 1 after start, got %d", chess.turn)
	}

	chess.TakeTurn()
	if chess.turn != 2 {
		t.Errorf("Expected turn to be 2 after one take turn, got %d", chess.turn)
	}

	if chess.HasEnded() != true {
		t.Errorf("Expected game to be ended after one turn, got %v", chess.HasEnded())
	}

	chess.End()
	// Note: Since End() does not change state, we do not test it here.
	// In a real-world scenario, we might check for side effects if any.
}

func TestPlayOneGame(t *testing.T) {
	chess := &Chess{MaxTurns: 1}
	PlayOneGame(chess)
	if chess.turn != 2 {
		t.Errorf("Expected turn to be 2 after playing one game, got %d", chess.turn)
	}
	if chess.HasEnded() != true {
		t.Errorf("Expected game to be ended after playing one game, got %v", chess.HasEnded())
	}
}
