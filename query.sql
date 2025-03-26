-- name: CreateUser :one
INSERT INTO users (
  username, password
) VALUES (
  ?, ?
)
RETURNING id;

-- name: GetUser :one
SELECT * FROM users
WHERE username = ? LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUsers :many
SELECT username FROM users;

-- name: CreateSession :exec
INSERT INTO sessions (
  user_id, token, expires_at
) VALUES (
  ?, ?, ?
);

-- name: GetSession :one
SELECT * FROM sessions
WHERE token = ? LIMIT 1;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE token = ?;

-- name: CreateListing :exec
INSERT INTO listings (
  title, starting_price, description, owner_id, image_url, created_at, category_id
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

-- name: EditListing :exec
UPDATE listings 
SET 
    title=?,            
    starting_price=?,  
    description=?,      
    image_url=?,        
    category_id=?,
    created_at=?      
WHERE 
    id=?;

-- name: GetListings :many
SELECT * FROM listings WHERE active = TRUE;

-- name: GetListing :one
SELECT * FROM listings WHERE id = ? LIMIT 1;

-- name: GetListingsByOwnerId :many
SELECT * FROM listings
WHERE owner_id = ?;

-- name: CloseListing :exec
UPDATE listings
SET active = FALSE
WHERE id = ?;

-- name: GetUserWatchlist :many
SELECT listings.* FROM listings
INNER JOIN watchlist ON listings.id = watchlist.listing_id
WHERE watchlist.user_id = ?;

-- name: AddToWatchlist :exec
INSERT INTO watchlist (
  user_id, listing_id
) VALUES (
  ?, ?
);

-- name: RemoveFromWatchlist :exec
DELETE FROM watchlist
WHERE user_id = ? AND listing_id = ?;

-- name: IsInWatchlist :one
SELECT * FROM watchlist
WHERE user_id = ? AND listing_id = ? LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories;

-- name: GetCategoryById :one
SELECT * FROM categories
WHERE id = ? LIMIT 1;

-- name: GetListingsByCategoryId :many
SELECT * FROM listings
WHERE category_id = ?;

-- name: GetBidsByListingId :many
SELECT * FROM bids
WHERE listing_id = ?;

-- name: CreateBid :exec
INSERT INTO bids (
  amount, user_id, listing_id, owner_name
) VALUES (
  ?, ?, ?, ?
);

-- name: GetMaxBid :one
SELECT * FROM bids
WHERE listing_id = ?
ORDER BY amount DESC
LIMIT 1;

-- name: GetCommentsByListingId :many
SELECT * FROM comments
WHERE listing_id = ?
ORDER BY id DESC;

-- name: CreateComment :exec
INSERT INTO comments (
  comment, listing_id, owner_id, owner_name
) VALUES ( 
?, ?, ?, ? 
);