package main

import "errors"

// Candidate Model
type Candidate struct {
	ID             int
	Name           string
	PoliticalParty string
	Sex            string
}

// CandidateElectionResult type
type CandidateElectionResult struct {
	ID             int
	Name           string
	PoliticalParty string
	Sex            string
	VoteCount      int
}

// PartyElectionResult type
type PartyElectionResult struct {
	PoliticalParty string
	VoteCount      int
}

func getAllCandidate() (candidates []Candidate) {
	return memcandidates
}

func getCandidate(candidateID int) (c Candidate, err error) {
	for _, cand := range memcandidates {
		if cand.ID == candidateID {
			return cand, nil
		}
	}
	return c, errors.New("cand not found")
}

func getCandidateByName(name string) (c Candidate, err error) {
	for _, cand := range memcandidates {
		if cand.Name == name {
			return cand, nil
		}
	}
	return c, errors.New("cand not found")
}

func getAllPartyName() (partyNames []string) {
	return allPartyName
}

func getCandidatesByPoliticalParty(party string) (candidates []Candidate) {
	for _, cand := range memcandidates {
		if cand.Name == party {
			candidates = append(candidates, cand)
		}
	}
	return candidates
}

func getElectionResult() (result []CandidateElectionResult) {
	rows, err := db.Query(`
		SELECT c.id, c.name, c.political_party, c.sex, IFNULL(v.count, 0)
		FROM candidates AS c
		LEFT OUTER JOIN
	  	(SELECT candidate_id, COUNT(*) AS count
	  	FROM votes
	  	GROUP BY candidate_id) AS v
		ON c.id = v.candidate_id
		ORDER BY v.count DESC`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		r := CandidateElectionResult{}
		err = rows.Scan(&r.ID, &r.Name, &r.PoliticalParty, &r.Sex, &r.VoteCount)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, r)
	}
	return
}
