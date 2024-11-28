DROP TRIGGER IF EXISTS trigger_job_updated_at ON jobs;
ALTER TABLE jobs DROP CONSTRAINT jobs_recruiter_id_fk;
DROP TABLE IF EXISTS jobs;
DROP TYPE job_type;