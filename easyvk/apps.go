package easyvk

import (
	"fmt"
	"encoding/json"
	"log"
)

//https://vk.com/dev/apps
type Apps struct {
	vk *VK
}

type AppsGetParams struct {
	AppId         int
	AppIds        []int
	Platform      string
	Extended      bool
	ReturnFriends bool
	Fields        string
}

type (
	AppsGetResponse struct {
		Count int       `json:"count"`
		Items []AppsObj `json:"items"`
	}
	AppsObj struct {
		Id              int       `json:"id"`
		Title           string    `json:"title"`
		Type            string    `json:"type"`
		Section         string    `json:"section"`
		AuthorUrl       string    `json:"author_url"`
		AuthorGroup     int       `json:"author_group"`
		MembersCount    int       `json:"members_count"`
		//PublishedDate   time.Time `json:"published_date"`
		GenreId         int       `json:"genre_id"`
		Genre           string    `json:"genre"`
		LeaderboardType int       `json:"leaderboard_type"`
		IsInCatalog     int       `json:"is_in_catalog"`
	}
)

func (o *Apps) GetById(id int) AppsObj {
	p := AppsGetParams{
		AppId: id,
	}
	apps, err := o.Get(p)
	if err != nil {
		log.Panicln("[ERR] Apps.get", err)
	}
	return apps.Items[0]
}

//https://vk.com/dev/apps.get
//Возвращает данные о запрошенном приложении.
func (o *Apps) Get(p AppsGetParams) (AppsGetResponse, error) {
	
	method := "apps.get"
	
	fields := p.Fields
	if p.Fields == "" {
		//fields = "sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives, counters,screen_name,timezone"
		fields = "city, country, place, description, wiki_page, members_count, counters, start_date, finish_date, can_post, can_see_all_posts, activity, status, contacts, links, fixed_post, verified, site, can_create_topic"
	}
	if p.Platform == "" {
		p.Platform = "web"
	}
	
	params := map[string]string{
		"app_id": fmt.Sprint(p.AppId),
		//"app_ids":        fmt.Sprint(strings.Join(p.AppIds)),
		"platform":       p.Platform,
		"extended":       boolConverter(p.Extended),
		"return_friends": boolConverter(p.ReturnFriends),
		"fields":         fields,
	}
	
	resp, err := o.vk.Request(method, params)
	if err != nil {
		return AppsGetResponse{}, err
	}
	//fmt.Println(string(resp))
	
	var res AppsGetResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return AppsGetResponse{}, err
	}
	return res, nil
}
