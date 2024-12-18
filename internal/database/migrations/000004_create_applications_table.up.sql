BEGIN;

CREATE TYPE status_type AS ENUM ('PENDING', 'APPROVED', 'REJECTED');

CREATE TABLE applications(
  id SERIAL PRIMARY KEY,
  job_id INT NOT NULL,
  candidate_id INT NOT NULL,
  applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  status status_type DEFAULT 'PENDING',

  CONSTRAINT applications_job_id_fk FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT applications_candidate_id_fk FOREIGN KEY (candidate_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;