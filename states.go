package main
import(
	"github.com/DavidMWeaver4/rssProject/internal/config"
	"github.com/DavidMWeaver4/rssProject/internal/database"
)

type state struct{
	cfg *config.Config
	db *database.Queries
}
