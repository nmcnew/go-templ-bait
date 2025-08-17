CREATE TABLE oauth2_service (
    id INT PRIMARY KEY,
    service_name TEXT,
    client_id TEXT,
    client_secret TEXT,
    config_uri TEXT,
    jwks_uri TEXT NULL
);

CREATE TABLE user_auth (
    id TEXT PRIMARY KEY,
    user_id TEXT,
    auth_type_id TEXT, -- For Example: oauth2, saml, oidc, etc.
    auth_service TEXT,
    access_token TEXT,
    refresh_token TEXT,
    FOREIGN KEY(user_id) REFERENCES user(id)
);
