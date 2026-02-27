package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/DavidMWeaver4/rssProject/internal/config"
	"github.com/DavidMWeaver4/rssProject/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	//check paramenter length
	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments")
		os.Exit(1)
	}
	//read config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	//get database
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	//get commands
	commandName := os.Args[1]
	arguments := os.Args[2:]
	dbQueries := database.New(db)
	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}
	//handle commands
	cmds := commands{
		Handlers: make(map[string]func(*state, command) error),
	}
	cmd := command{
		Name: commandName,
		Args: arguments,
	}

	//command registers
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
