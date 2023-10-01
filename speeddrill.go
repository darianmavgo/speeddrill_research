package main

type knowledge struct {
	id        int
	fact      string
	factmatch string
	points    int
	addDate   interface{}
}
type userJournal struct {
	id        string
	userId    int
	event     string
	eventDate interface{}
}

type gameConfig struct {
	id                    int
	tps                   int
	speedUpAfter          int
	slowDownAfter         int
	sessionCapMinutes     int
	titleImage            string
	fallRateStart         int
	wrongAnimation        string
	rightAnimation        string
	winEvent              string
	loseEvent             string
	restartEvent          string
	pointsRight           int
	pointsWrong           int
	winAnimation          string
	loseAnimation         string
	winSound              string
	loseSound             string
	rightSound            string
	wrongSound            string
	gettingCloseSound     string
	gettingCloseAnimation string
}
type user struct {
	userId    int
	username  string
	topScore  int
	worstFact string
}
