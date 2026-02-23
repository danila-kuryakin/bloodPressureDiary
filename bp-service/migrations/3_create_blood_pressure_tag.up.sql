CREATE TABLE blood_pressure_tag (
    pressure_id BIGINT NOT NULL REFERENCES blood_pressure(id) ON DELETE CASCADE,
    tag_id BIGINT NOT NULL REFERENCES user_tags(id),
    PRIMARY KEY (pressure_id, tag_id)
);

CREATE INDEX idx_bp_tag_pressure
    ON blood_pressure_tag(pressure_id);

CREATE INDEX idx_bp_tag_tag
    ON blood_pressure_tag(tag_id);