package main

import (
	"time"

	"github.com/chkev05/Go-Projects/project1/internal/database"
	"github.com/google/uuid"
)

// user model for the API response
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

// Converts a database.User to a User model for API response
func databaseUserToUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		APIKey:    u.ApiKey,
	}
}

// feed model for the API response
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"api_key"`
	UserID    uuid.UUID `json:"user_id"`
}

// converts a database.Feed to a Feed model for API response
func databaseFeedToFeed(f database.Feed) Feed {
	return Feed{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Url:       f.Url,
		UserID:    f.UserID,
	}
}

// converts a slice of database.Feed to a slice of Feed models for API response
func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := []Feed{}
	for _, f := range feeds {
		result = append(result, databaseFeedToFeed(f))
	}
	return result
}

// FeedFollow model for the API response
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

// converts a database.FeedFollow to a FeedFollow model for API response
func databaseFeedFollowToFeedFollow(ff database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        ff.ID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserID:    ff.UserID,
		FeedID:    ff.FeedID,
	}
}

// converts a slice of database.FeedFollow to a slice of FeedFollow models for API response
func databaseFeedFollowsToFeedFollows(feeds []database.FeedFollow) []FeedFollow {
	result := []FeedFollow{}
	// var result []FeedFollow
	for _, f := range feeds {
		result = append(result, databaseFeedFollowToFeedFollow(f))
	}
	return result
}

// Post model for the API response
type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

// converts a database.Post to a Post model for API response
func databasePostToPost(p database.Post) Post {
	var description *string
	if p.Description.Valid {
		description = &p.Description.String
	}
	return Post{
		ID:          p.ID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Title:       p.Title,
		Description: description,
		PublishedAt: p.PublishedAt,
		Url:         p.Url,
		FeedID:      p.FeedID,
	}
}

// converts a slice of database.Post to a slice of Post models for API response
func databasePostsToPosts(posts []database.Post) []Post {
	result := []Post{}
	for _, p := range posts {
		result = append(result, databasePostToPost(p))
	}
	return result
}
