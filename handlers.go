package main
import(
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/DavidMWeaver4/rssProject/internal/database"
)
//logins a user
func handlerLogin(s *state, cmd command) error{
	if len(cmd.Args) == 0{
		return fmt.Errorf("Username is required")
	}
	user, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil{
		return fmt.Errorf("User not found %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil{
		return err
	}
	fmt.Println("Current username has been set")
	return nil
}
//registers a new user
func handlerRegister(s *state, name command)error{
	if len(name.Args) == 0{
		return fmt.Errorf("Name is required")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: 			uuid.New(),
		CreatedAt: 		time.Now().UTC(),
		UpdatedAt: 		time.Now().UTC(),
	 	Name: 			name.Args[0],
	})
	if err != nil{
		return fmt.Errorf("Couldn't create user : %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil{
		return fmt.Errorf("Coudln't set current user: %w", err)
	}
	fmt.Println("User was created")
	fmt.Println(user)
	return nil
}
//resets the database
func handlerReset(s *state, cmd command) error{
	err := s.db.ResetUsers(context.Background())
	if err != nil{
		return err
	}
	fmt.Println("Database reset.")
	return nil
}
//fetches the full list of user Names and shows current login
func handlerUsers(s *state, cmd command) error{
	users, err := s.db.GetUsers(context.Background())
	if err != nil{
		return err
	}
	for _, user := range users{
		if user == s.cfg.CurrentUserName{
			fmt.Printf("* %s (current)\n", user)
			continue
		}
		fmt.Printf("* %s\n", user)
	}
	return nil
}

func handlerAgg(s *state, cmd command) error{
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil{
		return err
	}
	fmt.Printf("%+v\n",feed)
	return nil
}
