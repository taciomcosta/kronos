package sqlite

var findAllJobsSQL = "SELECT * FROM job"

var findOneJobSQL = "SELECT * FROM job WHERE name=?"

var describeJobSQL = `SELECT
   j.name AS name,
   j.command as command,
   j.tick AS tick,
   exec.last_execution,
   j.status,
   exec.executions_succeeded,
   exec.executions_failed,
   exec.average_cpu,
   exec.average_mem
FROM job j
LEFT JOIN (
   SELECT 
   	MAX(e.job_name) AS job_name,
   	MAX(e.date) AS last_execution,
   	MAX(e.date) AS last_execution,
   	COUNT(CASE e.status WHEN 'Succeeded' THEN 1 ELSE null END) AS executions_succeeded,
   	COUNT(CASE e.status WHEN 'Failed' THEN 1 ELSE null END) AS executions_failed,
   	AVG(e.cpu_time) AS average_cpu,
   	AVG(e.mem_usage) AS average_mem
   FROM execution e
   WHERE e.job_name=?
   GROUP BY e.job_name
) AS exec
ON j.name = exec.job_name
WHERE j.name=?`

var insertJobSQL = "INSERT INTO job VALUES(?, ?, ?, ?)"

var deleteJobSQL = "DELETE FROM job where name=?"

var insertExecutionSQL = "INSERT INTO execution VALUES(?, ?, ?, ?, ?)"

var findiAllExecutions = "SELECT * FROM execution ORDER BY date DESC LIMIT ? OFFSET ?"

var findJobExecutions = "SELECT * FROM execution WHERE job_name = ? ORDER BY date DESC LIMIT ? OFFSET ?"

var updateJobSQL = "UPDATE job SET status=? WHERE name=?"
