-- name: CreateObra :one
INSERT INTO obra (titulo, descripcion, artista, precio, vendida)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetObraById :one
SELECT * 
FROM obra 
WHERE id = $1;

-- name: ListObras :many
SELECT *            
FROM obra
ORDER BY id DESC;

-- name: UpdateObra :exec
UPDATE obra
SET titulo = $2, descripcion = $3, artista = $4, precio = $5, vendida = $6
WHERE id = $1;

-- name: DeleteObra :exec
DELETE FROM obra
WHERE id = $1;

-- name: ListAvailableObras :many
SELECT * 
FROM obra
WHERE vendida = FALSE
ORDER BY titulo DESC;