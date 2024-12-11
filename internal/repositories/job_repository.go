package repositories

import (
	"context"
	"go-to-work/internal/models"

	"github.com/jackc/pgx/v5"
)

type JobRepositoryInterface interface {
	Create(ctx context.Context, tx pgx.Tx, job models.Job) (uint64, error)
}

type JobRepository struct{}

func NewJobRepository() JobRepositoryInterface {
	return &JobRepository{}
}

func (jobRepository *JobRepository) Create(ctx context.Context, tx pgx.Tx, job models.Job) (uint64, error) {
	query := `
		INSERT INTO jobs (title, description, type, location, salary_range, is_active, recruiter_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;
	`

	var id uint64

	err := tx.QueryRow(
		ctx,
		query,
		job.Title,
		job.Description,
		job.Type,
		job.Location,
		job.SalaryRange,
		job.IsActive,
		job.RecruiterId,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
