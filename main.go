package main
import(
"github.com/DavidMWeaver4/rssProject/internal/config"
"fmt"
"log"
)
func main(){
	cfg, err := config.Read()
	if err != nil{
		log.Fatal(err)
	}
	err = cfg.SetUser("David")
	if err != nil{
		log.Fatal(err)
	}
	cfg2, err :=config.Read()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(cfg)
	fmt.Println(cfg2)
}
