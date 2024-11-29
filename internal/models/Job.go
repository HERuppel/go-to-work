package models

import "time"

type JobType string

const (
	FullTime   JobType = "FULL_TIME"
	PartTime   JobType = "PART_TIME"
	Contract   JobType = "CONTRACT"
	Internship JobType = "INTERNSHIP"
)

type Job struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Type        JobType   `json:"type"`
	Location    string    `json:"location,omitempty"`
	SalaryRange string    `json:"salary_range,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsActive    bool      `json:"is_active"`
	RecruiterId uint64    `json:"recruiter_id"`
}
