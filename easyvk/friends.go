package easyvk

import (
	"fmt"
	"encoding/json"
)

//https://vk.com/dev/friends.get
type Friends struct {
	vk *VK
}

type (
	FriendsGetIdResponse struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	}
	FriendsGetResponse struct {
		Count int `json:"count"`
		Items []struct {
			ID                     int          `json:"id"`
			FirstName              string       `json:"first_name"`
			LastName               string       `json:"last_name"`
			Sex                    int          `json:"sex"`
			Nickname               string       `json:"nickname,omitempty"`
			MaidenName             string       `json:"maiden_name"`
			Domain                 string       `json:"domain"`
			ScreenName             string       `json:"screen_name"`
			Bdate                  string       `json:"bdate"`
			City                   City         `json:"city"`
			Country                Country      `json:"country"`
			Photo50                string       `json:"photo_50"`
			Photo100               string       `json:"photo_100"`
			Photo200               string       `json:"photo_200"`
			PhotoMax               string       `json:"photo_max"`
			Photo200Orig           string       `json:"photo_200_orig"`
			Photo400Orig           string       `json:"photo_400_orig"`
			PhotoMaxOrig           string       `json:"photo_max_orig"`
			PhotoID                string       `json:"photo_id"`
			HasPhoto               int          `json:"has_photo"`
			HasMobile              int          `json:"has_mobile"`
			IsFriend               int          `json:"is_friend"`
			FriendStatus           int          `json:"friend_status"`
			Online                 int          `json:"online"`
			WallComments           int          `json:"wall_comments"`
			CanPost                int          `json:"can_post"`
			CanSeeAllPosts         int          `json:"can_see_all_posts"`
			CanSeeAudio            int          `json:"can_see_audio"`
			CanWritePrivateMessage int          `json:"can_write_private_message"`
			CanSendFriendRequest   int          `json:"can_send_friend_request"`
			MobilePhone            string       `json:"mobile_phone"`
			HomePhone              string       `json:"home_phone"`
			Twitter                string       `json:"twitter"`
			Instagram              string       `json:"instagram"`
			Facebook               string       `json:"facebook"`
			Facebook_name          string       `json:"facebook_name"`
			Skype                  string       `json:"skype"`
			Site                   string       `json:"site"`
			Status                 string       `json:"status"`
			LastSeen               PlatformInfo `json:"last_seen"`
		} `json:"items"`
	}
)

type FriendsGetParams struct {
	UserId int
	Order  string
	Count  int
	Offset int
	Fields string
}

func (o *Friends) GetId(p FriendsGetParams) (FriendsGetIdResponse, error) {
	
	params := map[string]string{
		"user_id": fmt.Sprint(p.UserId),
		"order":   p.Order,
		"count":   fmt.Sprint(p.Count),
		"offset":  fmt.Sprint(p.Offset),
	}
	
	resp, err := o.vk.Request("friends.get", params)
	if err != nil {
		return FriendsGetIdResponse{}, err
	}
	
	//fmt.Println(string(resp))
	
	var res FriendsGetIdResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return FriendsGetIdResponse{}, err
	}
	
	return res, nil
	
}

//https://vk.com/dev/friends.get
func (o *Friends) Get(p FriendsGetParams) (FriendsGetResponse, error) {
	
	method := "friends.get"
	
	fields := p.Fields
	if p.Fields == "" {
		fields = "nickname, domain, sex, bdate, city, country, timezone, photo_50, photo_100, photo_200_orig, has_mobile, contacts, education, online, relation, last_seen, status, can_write_private_message, can_see_all_posts, can_post, universities"
	}
	
	count := "1000" //MAX
	if p.Count != 0 {
		count = fmt.Sprint(p.Count)
	}
	
	params := map[string]string{
		"user_id": fmt.Sprint(p.UserId),
		"order":   p.Order,
		"count":   count,
		"offset":  fmt.Sprint(p.Offset),
		"fields":  fields,
	}
	
	resp, err := o.vk.Request(method, params)
	if err != nil {
		return FriendsGetResponse{}, err
	}
	
	//fmt.Println(string(resp))
	
	var res FriendsGetResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return FriendsGetResponse{}, err
	}
	
	return res, nil
	
}

func (o *Friends) GetAll(p FriendsGetParams) (FriendsGetResponse, error) {
	
	fields := p.Fields
	if p.Fields == "" {
		fields = "nickname, domain, sex, bdate, city, country, timezone, photo_50, photo_100, photo_200_orig, has_mobile, contacts, education, online, relation, last_seen, status, can_write_private_message, can_see_all_posts, can_post, universities"
	}
	
	offset := 0
	all := FriendsGetResponse{}
	for {
		params := FriendsGetParams{
			UserId: p.UserId,
			Order:  p.Order,
			Count:  1000,
			Offset: offset,
			Fields: fields,
		}
		items, err := o.Get(params)
		if err != nil {
			return FriendsGetResponse{}, err
		}
		all.Items = append(all.Items, items.Items...)
		if (len(items.Items) + offset) >= items.Count {
			break
		}
		offset += len(items.Items)
	}
	
	return all, nil
}

func (o *Friends) GetAllById(id int) (FriendsGetResponse, error) {
	
	offset := 0
	all := FriendsGetResponse{}
	for {
		params := FriendsGetParams{
			UserId: id,
			Count:  1000,
			Offset: offset,
		}
		items, err := o.Get(params)
		if err != nil {
			return FriendsGetResponse{}, err
		}
		all.Items = append(all.Items, items.Items...)
		if (len(items.Items) + offset) >= items.Count {
			break
		}
		offset += len(items.Items)
	}
	
	return all, nil
}