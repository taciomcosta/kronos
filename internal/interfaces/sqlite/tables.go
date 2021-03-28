package sqlite

type notifierRow struct {
	name  string
	ntype int
}

type slackRow struct {
	authToken  string
	channelIds string
}
