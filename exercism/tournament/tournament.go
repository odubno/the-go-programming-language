package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Scores tracks the outcome of each game
type Scores struct {
	Name           string
	MP, W, D, L, P int
}

// tallyScores updates scores for all games
func tallyScores(r io.Reader, scores map[string]Scores) error {
	// bufio reader helps with textual input
	scanner := bufio.NewScanner(r)

	// split text by lines
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		text := scanner.Text()

		// skip empty lines and comments
		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}

		row := strings.Split(text, ";")
		if len(row) < 3 {
			return fmt.Errorf("invalid data")
		}
		teamOne, teamTwo, outcome := row[0], row[1], row[2]

		// Access the struct for each team
		// If a struct for a team doesn't already exist, it's created
		scoreTeamOne := scores[teamOne]
		scoreTeamTwo := scores[teamTwo]

		// Update Name for each team
		scoreTeamOne.Name = teamOne
		scoreTeamTwo.Name = teamTwo

		// Increment the struct values of Matches Played
		scoreTeamOne.MP++
		scoreTeamTwo.MP++

		// If outcome is a win, then teamOne won -> 3 points
		// If outcome is a loss, then teamOne lost -> 0 points
		// If outcome is a draw, then both teams get 1 point each
		if outcome == "win" {
			scoreTeamOne.W++
			scoreTeamTwo.L++
			scoreTeamOne.P += 3
		} else if outcome == "loss" {
			scoreTeamOne.L++
			scoreTeamTwo.W++
			scoreTeamTwo.P += 3
		} else if outcome == "draw" {
			scoreTeamOne.D++
			scoreTeamTwo.D++
			scoreTeamOne.P++
			scoreTeamTwo.P++
		} else {
			return fmt.Errorf("the outcome is invalid")
		}

		// Update results for each team
		scores[teamOne] = scoreTeamOne
		scores[teamTwo] = scoreTeamTwo
	}
	return nil
}

// Tally keeps track of the score
func Tally(r io.Reader, w io.Writer) error {

	// Create a map that maps string to Scores struct
	scores := map[string]Scores{}

	err := tallyScores(r, scores)
	if err != nil {
		return err
	}
	results := sortScores(scores)
	w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n"))
	for _, outcome := range results {
		result := fmt.Sprintf("%-30v |  %v |  %v |  %v |  %v |  %v\n", outcome.Name, outcome.MP, outcome.W, outcome.D, outcome.L, outcome.P)
		w.Write([]byte(result))
	}
	return nil
}

// sortScores sorts scores by Points (P), while keeping original order if tied
func sortScores(scores map[string]Scores) []Scores {

	// Create a slice for all team results
	results := []Scores{}
	for _, v := range scores {
		// A slice helps preserve the original order
		results = append(results, v)
	}

	// Sort by points (P), decreasing order, while keeping original order if tied
	sort.SliceStable(results, func(i, j int) bool {
		// If "draw", then order by Name i.e. alphabetically
		if results[i].P == results[j].P {
			return results[i].Name < results[j].Name
		}
		// order in decreasing order
		return results[i].P > results[j].P
	})

	return results
}
