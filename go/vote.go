package main

// Vote Model
type Vote struct {
	// ID          int
	// UserID      int
	CandidateID int
	VoteCount   int
	// Keyword     string
}

func getVoteCountByCandidateID(candidateID int) (count int) {
	row := db.QueryRow("SELECT vote_count AS count FROM votes WHERE candidate_id = ?", candidateID)
	row.Scan(&count)
	return
}

func getUserVotedCount(userID int) (count int) {
	row := db.QueryRow("SELECT voted AS count FROM users WHERE id =  ?", userID)
	row.Scan(&count)
	return
}

func createVote(voteCount int, userID int, candidateID int, keyword string) {
	db.Exec("update votes set vote_count = vote_count + ? where candidate_id = ?", voteCount, candidateID)
	db.Exec("UPDATE users set voted = voted + ? where id = ?", voteCount, userID)
	db.Exec("UPDATE keyword set count = count + ? where candidate_id = ? and keyword = ?", voteCount, candidateID, keyword)
}

func getVoiceOfSupporter(candidateIDs []int) (voices []string) {
	rows, err := db.Query(`
    SELECT keyword
    FROM votes
    WHERE candidate_id IN (?)
    GROUP BY keyword
    ORDER BY count DESC
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
