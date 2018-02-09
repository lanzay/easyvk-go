package easyvk

import (
	"fmt"
	"encoding/json"
)

//https://vk.com/dev/groups.get
type Group struct {
	vk *VK
}

type GroupGetParams struct {
	UserId   int
	Extended bool
	Count    int
	Offset   int
	Fields   string
}

type GroupGetIdResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

type (
	GroupGetResponse struct {
		Count int        `json:"count"`
		Items []GroupObj `json:"items"`
	}
	GroupObj struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		IsClosed   int    `json:"is_closed"`
		Type       string `json:"type"`
		IsAdmin    int    `json:"is_admin"`
		IsMember   int    `json:"is_member"`
		Photo50    string `json:"photo_50"`
		Photo100   string `json:"photo_100"`
		Photo200   string `json:"photo_200"`
	}
)

//groups.get Возвращает список сообществ указанного пользователя.
//https://vk.com/dev/groups.get
func (o *Group) Get(p GroupGetParams) (GroupGetResponse, error) {
	
	method := "groups.get"
	
	fields := p.Fields
	if p.Fields == "" {
		fields = "city, country, place, description, wiki_page, members_count, counters, start_date, finish_date, can_post, can_see_all_posts, activity, status, contacts, links, fixed_post, verified, site, can_create_topic"
	}
	
	count := "1000" //MAX
	if p.Count != 0 {
		count = fmt.Sprint(p.Count)
	}
	
	params := map[string]string{
		"user_id":  fmt.Sprint(p.UserId),
		"extended": boolConverter(p.Extended),
		"count":    count,
		"offset":   fmt.Sprint(p.Offset),
		"fields":   fields,
	}
	
	resp, err := o.vk.Request(method, params)
	if err != nil {
		return GroupGetResponse{}, err
	}
	//fmt.Println(string(resp))
	
	var res GroupGetResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return GroupGetResponse{}, err
	}
	return res, nil
}

func (o *Group) GetAll(p GroupGetParams) (GroupGetResponse, error) {
	
	fields := p.Fields
	if p.Fields == "" {
		fields = "nickname, domain, sex, bdate, city, country, timezone, photo_50, photo_100, photo_200_orig, has_mobile, contacts, education, online, relation, last_seen, status, can_write_private_message, can_see_all_posts, can_post, universities"
	}
	
	offset := 0
	all := GroupGetResponse{}
	for {
		params := GroupGetParams{
			UserId:   p.UserId,
			Extended: p.Extended,
			Count:    1000,
			Offset:   offset,
			Fields:   fields,
		}
		items, err := o.Get(params)
		if err != nil {
			return GroupGetResponse{}, err
		}
		all.Items = append(all.Items, items.Items...)
		if (len(items.Items) + offset) >= items.Count {
			break
		}
		offset += len(items.Items)
	}
	
	return all, nil
	
}

//===========================================================================================
//groups.getMembers Возвращает список участников сообщества.
//https://vk.com/dev/groups.getMembers

type GroupGetMembersParams struct {
	GroupId int
	Sort    string
	Count   int
	Offset  int
	Fields  string
}

type (
	GroupGetMembersResponse struct {
		Count int                `json:"count"`
		Items []UsersGetResponse `json:"items"`
	}
)

//groups.getMembers Возвращает список участников сообщества.
//https://vk.com/dev/groups.getMembers
func (o *Group) GetMembers(p GroupGetMembersParams) (GroupGetMembersResponse, error) {
	
	method := "groups.getMembers"
	
	fields := p.Fields
	if p.Fields == "" {
		fields = "sex, bdate, city, country, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, online_mobile, lists, domain, has_mobile, contacts, connections, site, education, universities, schools, can_post, can_see_all_posts, can_see_audio, can_write_private_message, status, last_seen, common_count, relation, relatives"
	}
	
	count := "1000" //MAX
	if p.Count != 0 {
		count = fmt.Sprint(p.Count)
	}
	
	params := map[string]string{
		"group_id": fmt.Sprint(p.GroupId),
		"sort":     fmt.Sprint(p.Sort),
		"count":    count,
		"offset":   fmt.Sprint(p.Offset),
		"fields":   fields,
	}
	
	resp, err := o.vk.Request(method, params)
	if err != nil {
		return GroupGetMembersResponse{}, err
	}
	//fmt.Println(string(resp))
	
	var res GroupGetMembersResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return GroupGetMembersResponse{}, err
	}
	return res, nil
}

func (o *Group) GetMembersAll(p GroupGetMembersParams) (GroupGetMembersResponse, error) {
	
	fields := p.Fields
	if p.Fields == "" {
		fields = "nickname, domain, sex, bdate, city, country, timezone, photo_50, photo_100, photo_200_orig, has_mobile, contacts, education, online, relation, last_seen, status, can_write_private_message, can_see_all_posts, can_post, universities"
	}
	
	offset := 0
	all := GroupGetMembersResponse{}
	for {
		params := GroupGetMembersParams{
			GroupId: p.GroupId,
			Sort:    p.Sort,
			Count:   1000,
			Offset:  offset,
			Fields:  fields,
		}
		items, err := o.GetMembers(params)
		if err != nil {
			return GroupGetMembersResponse{}, err
		}
		all.Items = append(all.Items, items.Items...)
		if (len(items.Items) + offset) >= items.Count {
			break
		}
		offset += len(items.Items)
	}
	
	return all, nil
	
}
