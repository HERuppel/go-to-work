package usecases

import (
	"context"
	"go-to-work/internal/database"
	"go-to-work/internal/models"
	"go-to-work/internal/repositories"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type JobUseCase struct {
	pool          *pgxpool.Pool
	jobRepository repositories.JobRepositoryInterface
}

func NewJobUseCase(pool *pgxpool.Pool, jobRepository repositories.JobRepositoryInterface) *JobUseCase {
	return &JobUseCase{
		pool:          pool,
		jobRepository: jobRepository,
	}
}

func (jobUseCase *JobUseCase) Create(ctx context.Context, job models.Job) (uint64, error) {
	id, err := database.WithTransaction(jobUseCase.pool, ctx, func(tx pgx.Tx) (uint64, error) {
		if err := job.Validate(); err != nil {
			return 0, err
		}

		id, err := jobUseCase.jobRepository.Create(ctx, tx, job)
		if err != nil {
			return 0, err
		}

		return id, nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
