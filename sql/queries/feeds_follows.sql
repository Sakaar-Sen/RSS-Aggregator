-- name: CreateFeedFollow :one
insert into feeds_follows (id, created_at, updated_at, feed_id, user_id)
values ($1, $2, $3, $4, $5)
returning *;

-- name: GetFeedFollows :many
select * from feeds_follows where user_id = $1;


-- name: DeleteFeedFollow :exec
delete from feeds_follows where id = $1 AND user_id = $2;