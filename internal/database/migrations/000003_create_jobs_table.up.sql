BEGIN;

CREATE TYPE job_type AS ENUM('FULL_TIME', 'PART_TIME', 'CONTRACT', 'INTERNSHIP');

CREATE TABLE jobs(
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  type job_type NOT NULL,
  location VARCHAR(100),
  salary_range VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  is_acitve BOOLEAN DEFAULT TRUE,
  recruiter_id INT NOT NULL,

  CONSTRAINT jobs_recruiter_id_fk FOREIGN KEY (recruiter_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);


CREATE  INDEX idx_jobs_title ON jobs (title);
CREATE  INDEX idx_jobs_location ON jobs (location);
CREATE  INDEX idx_jobs_recruiter_id ON jobs (recruiter_id);

CREATE TRIGGER trigger_job_updated_at
BEFORE UPDATE ON jobs
FOR EACH ROW
EXECUTE FUNCTION updated_at();

COMMIT;