CREATE TABLE blood_pressure (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    systolic INT NOT NULL,
    diastolic INT NOT NULL,
    pulse INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_blood_pressure_user_id
    ON blood_pressure(user_id);

CREATE INDEX idx_blood_pressure_user_created
    ON blood_pressure(user_id, created_at DESC);