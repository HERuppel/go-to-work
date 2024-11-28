ALTER TABLE applications DROP CONSTRAINT applications_job_id_fk;
ALTER TABLE applications DROP CONSTRAINT applications_candidate_id_fk;
DROP TABLE IF EXISTS applications;
DROP TYPE status_type;