package main
import(
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/DavidMWeaver4/rssProject/internal/database"
)

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
