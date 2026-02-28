package main
import(
	"context"
	"fmt"
	"time"
	"log"
	"database/sql"
	"github.com/DavidMWeaver4/rssProject/internal/database"
)
// fetches a xml post from a website
func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1{
		return fmt.Errorf("Missing arguments")
	}
	timing_between, err := time.ParseDuration(cmd.Args[0])
	if err != nil{
		return err
	}
	fmt.Printf("Collecting feeds every %+v\n", timing_between)
	ticker := time.NewTicker(timing_between)
	for ; ; <- ticker.C{
		scrapeFeeds(s)
	}
}

//scrapes feeds from a last fetched timer from oldest to newest
func scrapeFeeds(s *state){
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil{
		log.Printf("Could not get next feed to fetch %v", err)
		return
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt:		sql.NullTime{
			Time:	time.Now().UTC(),
			Valid:	true,
		},
		UpdatedAt:			time.Now().UTC(),
		ID:					next.ID,
	})
	if err != nil{
		log.Printf("Could not mark feed as fetched %v", err)
		return
	}
	feed, err := fetchFeed(context.Background(),next.Url)
	if err != nil{
		log.Printf("Could not fetch next feed %v", err)
		return
	}
	for _, item := range feed.Channel.Item{
		fmt.Printf("%s\n", item.Title)
	}
	return
}
