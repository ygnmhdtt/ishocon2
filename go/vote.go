package main

import (
	"fmt"
	"strings"
)

// Vote Model
type Vote struct {
	ID          int
	UserID      int
	CandidateID int
	Keyword     string
}

func getVoteCountByCandidateID(candidateID int) (count int) {
	row := db.QueryRow("SELECT COUNT(*) AS count FROM votes WHERE candidate_id = ?", candidateID)
	row.Scan(&count)
	return
}

func getUserVotedCount(userID int) (count int) {
	row := db.QueryRow("SELECT COUNT(*) AS count FROM votes WHERE user_id =  ?", userID)
	row.Scan(&count)
	return
}

func createVote(voteCount int, userID int, candidateID int, keyword string) {
	var args []string
	for i := 1; i <= voteCount; i++ {
		args = append(args, fmt.Sprintf("(%d, %d, '%s')", userID, candidateID, keyword))
	}
	sql := fmt.Sprintf("INSERT INTO votes (user_id, candidate_id, keyword) VALUES %s", strings.Join(args, ","))
	db.Exec(sql)
}

func getVoiceOfSupporter(candidateIDs []int) (voices []string) {
	rows, err := db.Query(`
    SELECT keyword
    FROM votes
    WHERE candidate_id IN (?)
    GROUP BY keyword
    ORDER BY COUNT(*) DESC
    LIMIT 10`)
	if err != nil {
		return nil
	}

	defer rows.Close()
	for rows.Next() {
		var keyword string
		err = rows.Scan(&keyword)
		if err != nil {
			panic(err.Error())
		}
		voices = append(voices, keyword)
	}
	return
}
