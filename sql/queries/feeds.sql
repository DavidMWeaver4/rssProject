-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, last_fetched_at, name, url, user_id)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = $1;

-- name: GetFeedsForUser :many
SELECT * FROM feeds
WHERE user_id = $1;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: DeleteUserFeed :exec
DELETE FROM feeds
WHERE user_id = $1;

-- name: GetFeedsFromURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetUserWhoMadeFeed :one
SELECT user_id FROM feeds
WHERE url = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1,
updated_at = $2
WHERE id = $3;
