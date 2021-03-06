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

type assignmentRow struct {
	jobName      string
	notifierName string
	onErrorOnly  bool
}

var tablesStmts []string = []string{
	`CREATE TABLE IF NOT EXISTS job(
		name TEXT PRIMARY KEY,
		command TEXT,
		tick TEXT,
		status BOOLEAN
	)`,
	`CREATE TABLE IF NOT EXISTS execution(
		job_name TEXT,
		date DATE,
		STATUS TEXT,
		mem_usage INTEGER,
		cpu_time INTEGER
	)`,
	`CREATE TABLE IF NOT EXISTS notifier(
		name TEXT PRIMARY KEY,
		type INTEGER
	)`,
	`CREATE TABLE IF NOT EXISTS slack(
		auth_token TEXT,
		channel_ids TEXT,
		notifier_name TEXT REFERENCES notifier(name)
	)`,
	`CREATE TABLE IF NOT EXISTS assignment(
		job_name TEXT REFERENCES job(name),
		notifier_name TEXT REFERENCES notifier(name),
		on_error_only BOOLEAN,
		PRIMARY KEY (job_name, notifier_name)

	)`,
}
