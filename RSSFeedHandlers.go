package main

import (
	//"github.com/DavidMWeaver4/rssProject/internal/config"
	//"github.com/DavidMWeaver4/rssProject/internal/database"
	"context"
	"io"
	"encoding/xml"
	"net/http"
	"html"
)

func fetchFeed(ctx context.Context, feedURL string)(*RSSFeed, error){
	//get request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil{
		return nil, err
	}
	//setup client
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	//read body
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	//unmarshal
	var rssF RSSFeed
	err = xml.Unmarshal(body, &rssF)
	if err != nil{
		return nil, err
	}
	rssF.escapedText()
	return &rssF, nil
}
func(rss *RSSFeed) escapedText(){
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	for i:= range rss.Channel.Item{
		rss.Channel.Item[i].Title = html.UnescapeString(rss.Channel.Item[i].Title)
		rss.Channel.Item[i].Description = html.UnescapeString(rss.Channel.Item[i].Description)
	}
	return
}
