package flyweight

// TeamFactory ...
type TeamFactory interface {
	GetTeam(teamID int) *Team
	GetNumberOfObjects() int
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

// NewTeamFactory ...
func NewTeamFactory() TeamFactory {
	return &teamFlyweightFactory{
		createdTeams: make(map[int]*Team),
	}
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}

	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team

	return t.createdTeams[teamID]
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}

func getTeamFactory(team int) Team {
	switch team {
	case TeamB:
		return Team{
			ID:   2,
			Name: "team_b",
		}
	default:
		return Team{
			ID:   1,
			Name: "team_a",
		}
	}
}
