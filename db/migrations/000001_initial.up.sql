CREATE TABLE apps (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    slug varchar(255) NOT NULL
);

CREATE TABLE app_fields (
    id serial PRIMARY KEY,
    app_id integer NOT NULL,
    field varchar(255) NOT NULL,
    CONSTRAINT fk_app FOREIGN KEY (app_id) REFERENCES apps (id)
);

CREATE TABLE integrations (
    id serial PRIMARY KEY,
    name text NOT NULL,
    app_id integer NOT NULL,
    CONSTRAINT fk_app FOREIGN KEY (app_id) REFERENCES apps (id)
);

CREATE TABLE integration_app_field_values (
    id serial PRIMARY KEY,
    integration_id integer NOT NULL,
    app_field_id integer NOT NULL,
    value text NOT NULL,
    CONSTRAINT fk_integration FOREIGN KEY (integration_id) REFERENCES integrations (id),
    CONSTRAINT fk_app_field FOREIGN KEY (app_field_id) REFERENCES app_fields (id)
);

