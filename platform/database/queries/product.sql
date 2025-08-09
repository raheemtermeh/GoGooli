-- platform/database/queries/product.sql

-- name: GetAllProducts :many
SELECT * FROM products ORDER BY id;