package db

var leaderboardKey = "leaderboard"

type Leaderboard struct {
	Count int
	Users []*User
}

func (db *Database) GetLeaderboard() (*Leaderboard, error) {
	scores := db.Client.ZRangeWithScores(leaderboardKey, 0, -1)
	if scores == nil {
		return nil, ErrNil
	}

	count := len(scores.Val())
	users := make([]*User, count)
	for idx, member := range scores.Val() {
		users[count-1-idx] = &User{
			Username: member.Member.(string),
			Points:   int(member.Score),
			Rank:     count - idx,
		}
	}

	leaderboard := &Leaderboard{
		Count: count,
		Users: users,
	}
	return leaderboard, nil
}
