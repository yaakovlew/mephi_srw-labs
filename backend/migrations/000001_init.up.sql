CREATE TABLE users_done (
    user_id bigint not null,
    internal_lab_id bigint not null,
    external_lab_id bigint not null,
    is_done boolean not null,
    percentage int not null default 0,
    token varchar(500) default '',
    step int default 0,
    variance json,
    CONSTRAINT check_percentage_range CHECK (percentage >= 0 AND percentage <= 100),
    CONSTRAINT constraint_line_unique UNIQUE (user_id, internal_lab_id, external_lab_id)
);

CREATE TABLE bank_variance (
    id bigserial not null Primary key,
    variance json not null
);

CREATE TABLE bank_variance_1a (
    id bigserial not null Primary key,
    variance json not null
);

CREATE TABLE bank_variance_1b (
    id bigserial not null Primary key,
    variance json not null
);