-- name: InsertNewEntry :exec
INSERT INTO demo_table (
    special
) VALUES (
    @special::MyEnum[]
);

-- name: GetAll :many
SELECT * FROM demo_table ORDER BY id;