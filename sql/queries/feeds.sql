-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id) 
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: GetFeeds :many
select * from feeds;

-- name: GetNextFeedToFetch :many
select * from feeds 
order by last_fetch asc nulls first
limit $1;


-- name: MarkFeedFetched :one
update feeds set last_fetch = NOW(),
updated_at = NOW()
where id = $1
returning *;