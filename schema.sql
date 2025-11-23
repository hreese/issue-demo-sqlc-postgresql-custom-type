CREATE TYPE MyEnum AS ENUM (
    'one',
    'two',
    'three',
    'four'
    );

CREATE TABLE IF NOT EXISTS demo_table (
    id SERIAL PRIMARY KEY,
    special MyEnum[]
);
