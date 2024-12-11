package models

import (
	"errors"
	"strings"
	"time"
)

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

func (job *Job) Validate() error {
	if job.Title == "" {
		return errors.New("TITLE_REQUIRED")
	}
	if job.Description == "" {
		return errors.New("DESCRIPTION_REQUIRED")
	}
	if !validateJobType(job.Type) {
		return errors.New("INVALID JOB TYPE, MUST BE: FULL_TIME, PART_TIME, CONTRACT, or INTERNSHIP")
	}
	if job.RecruiterId == 0 {
		return errors.New("RECRUITER_REQUIRED")
	}

	return nil
}

var validJobTypes = []JobType{
	FullTime,
	PartTime,
	Contract,
	Internship,
}

func validateJobType(jobType JobType) bool {
	for _, validType := range validJobTypes {
		if strings.ToUpper(string(jobType)) == string(validType) {
			return true
		}
	}
	return false
}
