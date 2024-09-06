CREATE TABLE todos (
    id        INTEGER PRIMARY KEY,
    title     VARCHAR(100) NOT NULL,
    details   TEXT,
    completed BOOLEAN      NOT NULL DEFAULT FALSE
);
