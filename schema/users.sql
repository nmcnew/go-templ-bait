CREATE TABLE user (
    id TEXT PRIMARY KEY,
    name TEXT,
    password TEXT NULL
);

CREATE TABLE contact_method (
    id TEXT PRIMARY KEY,
    user_id TEXT,
    method_type_id INT,
    method TEXT,
    FOREIGN KEY(user_id) REFERENCES user(id),
    FOREIGN KEY(type_id) REFERENCES contact_method_type(id)
);

