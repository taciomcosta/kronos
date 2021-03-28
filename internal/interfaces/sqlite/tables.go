package sqlite

type jobRow struct {
	name    string
	command string
	tick    string
	status  bool
}

type notifierRow struct {
	name  string
	ntype int
}

type slackRow struct {
	authToken  string
	channelIds string
}
