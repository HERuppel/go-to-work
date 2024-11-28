package models

import "time"

type StatusType string

const (
	Pending  StatusType = "PENDING"
	Approved StatusType = "APPROVED"
	Rejected StatusType = "REJECTED"
)

type Application struct {
	ID          uint       `json:"id"`
	JobId       uint       `json:"job_id"`
	CandidateId uint       `json:"candidate_id"`
	AppliedAt   time.Time  `json:"applied_at"`
	Status      StatusType `json:"status"`
}
