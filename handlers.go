package main

import (
	"context"
	"fmt"
	"time"

	"github.com/DavidMWeaver4/rssProject/internal/database"
	"github.com/google/uuid"
)

// logins a user
func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Username is required")
	}
	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("User not found %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Println("Current username has been set")
	return nil
}

// registers a new user
func handlerRegister(s *state, name command) error {
	if len(name.Args) == 0 {
		return fmt.Errorf("Name is required")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name.Args[0],
	})
	if err != nil {
		return fmt.Errorf("Couldn't create user : %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Coudln't set current user: %w", err)
	}
	fmt.Println("User was created")
	fmt.Println(user)
	return nil
}

// resets the database
func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Database reset.")
	return nil
}

// fetches the full list of user Names and shows current login
func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user)
			continue
		}
		fmt.Printf("* %s\n", user)
	}
	return nil
}

// fetches a xml post from a website
func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

// add a feed to the feeds table
func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Missing Arguments")
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

// gets feeds information
func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for i := range feeds {
		fmt.Printf("Name: %s\n", feeds[i].Name)
		fmt.Printf("URL: %s\n", feeds[i].Url)
		user, err := s.db.GetUserWhoMadeFeed(context.Background(), feeds[i].Url)
		if err != nil {
			return err
		}
		name, err := s.db.GetUserByID(context.Background(), user)
		if err != nil {
			return err
		}
		fmt.Printf("Created By: %+v\n", name)
	}
	return nil
}

// follows a feed
func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Missing Arguments")
	}
	feeds, err := s.db.GetFeedsFromURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	feedFollows, err := s.db.CreateFeedFollows(context.Background(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feeds.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Name: %s\n", feedFollows.UserID)
	fmt.Printf("URL: %s\n", feedFollows.FeedID)
	return nil
}

// lists all feeds followed
func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, ff := range feedFollows {
		fmt.Printf("%s\n", ff.FeedName)
	}
	return nil
}

// takes a feed url and unfollows it for current users
func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Missing URL")
	}
	feed, err := s.db.GetFeedsFromURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}
	err = s.db.DeleteFeedFollowsByUserAndFeed(context.Background(), database.DeleteFeedFollowsByUserAndFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println("Unfollowed!")
	return nil
}

// middleware function for Currently Logged In User
func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error,
) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
