// An example implementation the qlearning interfaces. Can be run
// with go run hangman.go.
//
// Word list provided by https://github.com/first20hours/google-10000-english
package services

import (
	"fmt"
)

const (
	startingLives = 6

	Lost   = 3
	Active = 1
	Won    = 2

	Buy  = 1
	Sell = -1
)

type StateData struct {
	State  []int64
	Signal int64
}

var (
	//Positions = []int64{112312, 1241232, 1342132} // hashmap ile her state'in içeriğindeki alım satım ve tutma sinyalleri unique olarak tutulacak.

	debug      bool = false
	progressAt int  = 1000
)

// Game represents the state of any given game of Hangman. It implements
// Agent, Rewarder, and State.
type Game struct {
	State     StateData
	Lives     int
	Attempted map[int64]bool
	Positions map[int64]int64
	debug     bool
}

// NewGame creates a new Hangman game for the given word. If debug
// is true, Game.Log messages will print to stdout.
func NewGame(states StateData, debug bool, possiblePositions map[int64]int64, lives int) *Game {
	game := &Game{debug: debug}
	game.New(states, possiblePositions, lives)

	return game
}

// New resets the current game to a new game for the given word.
func (game *Game) New(states StateData, possiblePositions map[int64]int64, lives int) {
	game.State = states
	game.Lives = lives
	game.Attempted = make(map[int64]bool, len(possiblePositions))
	game.Positions = possiblePositions
}

// Returns Lost, Active, or Won based on the game's current state.
func (game *Game) IsComplete() int {
	/*
		if game.Lives > 0 {
			return Active
		} else if game.Profit > 0 {
			return Won
		} else if game.Profit <= 0 {
			return Lost
		}
	*/
	return 0
}

// Choose applies a character attempt in the current game, returning
// true if char is present in Game.Word.
//
// Choose updates the s's state.
func (game *Game) Choose(position int64) bool {
	game.Attempted[position] = true

	hit := false

	for _, key := range game.State.State {
		if key == position {
			game.Lives -= 1
			hit = true
		}
	}

	if !hit {
		return false
	}

	return true
}

// Reward returns a score for a given StateAction. Reward is a
// member of the Rewarder interface. If the choice is found in
// the game's word, a positive score is returned. Otherwise, a static
// -1000 is returned.
func (game *Game) Reward(action *StateAction, signal int64) float64 {
	choice := action.Action.Int()

	if int64(signal) == choice {
		return 10
	} else {
		return -10
	}

}

// Next creates a new slice of Action instances. A possible
// action is created for each character that has not been attempted in
// in the game.
func (game *Game) Next() []Action {
	actions := make([]Action, 0, len(game.State.State))

	for check, _ := range game.Positions {
		attempted := game.Attempted[check]
		if !attempted {
			actions = append(actions, &Choice{Position: check})
		}
	}

	return actions
}

// Log is a wrapper of fmt.Printf. If Game.debug is true, Log will print
// to stdout.
func (game *Game) Log(msg string, args ...interface{}) {
	if game.debug {
		logMsg := fmt.Sprintf("[GAME %v] (%d moves, %d lives) %s\n", game.State, len(game.Attempted), game.Lives, msg)
		fmt.Printf(logMsg, args...)
	}
}

// String returns a consistent hash for the current game state to be
// used in a Agent.
func (game *Game) String() string {
	return fmt.Sprintf("%v", game.State.State)
}

// Choice implements Action for a character choice in a game
// of Hangman.
type Choice struct {
	Position int64
}

// String returns the character for the current action.
func (choice *Choice) Int() int64 {
	return choice.Position
}

// String returns the character for the current action.
func (choice *Choice) String() string {
	return fmt.Sprintf("%v", choice.Position)
}

// Apply updates the state of the game for a given character choice.
func (choice *Choice) Apply(state State) State {
	game := state.(*Game)
	game.Choose(choice.Position)

	return game
}

func Learn(allStates []StateData, agent *SimpleAgent, allPossibilities map[int64]int64, basePercentage float64, lives int) *SimpleAgent {
	var (
		wins     = 0
		lastWins = 0
		count    = 0
		profit   = .0

		// Our agent has a learning rate of 0.7 and discount of 1.0.
	)

	progress := func() {
		// Print our progress every 1000 rows.
		if count > 0 && count%progressAt == 0 {
			rate := float32(wins-lastWins) / float32(progressAt) * 100.0
			lastWins = wins
			fmt.Printf("%d games played: %d WINS %d LOSSES %.0f%% WIN RATE %.2f PROFIT\n", count, wins, count-wins, rate, profit)
		}
	}

	// Let's play 5 million games
	func(states []StateData) {
		for i := 0; i < lives; i++ {
			profit = .0
			for count = 0; count < len(states); count++ {
				signal := states[count].Signal
				// Get a new word and game for each iteration...
				game := NewGame(states[count], debug, allPossibilities, lives) //TODO: tüm olası pozisyonlar ve hash tablosu entegre edilecek
				game.Log("Game created")

				// While the game is still active, we'll continue to update
				// our agent and learn from its choices.
				// Pick the next move, which is going to be a letter choice.
				action := Next(agent, game)

				// Whatever that choice is, let's update our model for its
				// impact. If the character chosen is in the game's word,
				// then this action will be positive. Otherwise, it will be
				// negative.
				agent.Learn(action, game, signal)

				// Reward doesn't change state so we can check what the
				// reward would be for this action, and report how the
				// game changed.
				if game.Reward(action, signal) > 0.0 {
					game.Log("%s was selected", action.Action.Int())
				}
				game.Log("%s was incorrect", action.Action.Int())

				progress()
			}
			fmt.Println(profit)
		}
	}(allStates)

	progress()

	fmt.Printf("\nAgent performance: %d games played, %d WINS %d LOSSES %.0f%% WIN RATE %.2f PROFIT\n", count, wins, count-wins, float32(wins)/float32(count)*100.0, profit)

	return agent
}

func Predict(state StateData, agent *SimpleAgent, allPossibilities map[int64]int64, lives int) (prediction int64) {
	game := NewGame(state, false, allPossibilities, lives)
	action := Next(agent, game)

	return action.Action.Int()
}
