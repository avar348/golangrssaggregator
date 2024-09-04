package models

import (
	"time"

	"github.com/avar348/golangrssaggregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowtoFeedFollow(dbFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		UserID:    dbFeed.UserID,
		FeedID:    dbFeed.FeedID,
	}
}

func DatabaseFeedFollowstoFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {

	feedFollows := []FeedFollow{}
	for _, dbFeed := range dbFeedFollows {
		feedFollows = append(feedFollows, DatabaseFeedFollowtoFeedFollow(dbFeed))
	}
	return feedFollows
}
