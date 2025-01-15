package types

// Job represents a unit of work in the ETL pipeline
type Job struct {
	ID     int
	Result []Result
}

type Result struct {
	 ResultID int
	 CvssScores string
	 JobID int
}
