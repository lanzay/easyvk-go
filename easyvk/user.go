package easyvk

import (
	"fmt"
	"encoding/json"
)

//https://vk.com/dev/users.get
type User struct {
	vk *VK
}

type (
	UsersGetResponse struct {
		ID                     int          `json:"id"`
		URL                    string       `json:"url,omitempty"`
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
		CropPhoto              CropPhoto    `json:"crop_photo"`
		Verified               int          `json:"verified"`
		FollowersCount         int          `json:"followers_count"`
		Blacklisted            int          `json:"blacklisted"`
		BlacklistedByMe        int          `json:"blacklisted_by_me"`
		IsFavorite             int          `json:"is_favorite"`
		IsHiddenFromFeed       int          `json:"is_hidden_from_feed"`
		CommonCount            int          `json:"common_count"`
		Career                 []Career     `json:"career"`
		Military               []Military   `json:"military"`
		University             int          `json:"university"`
		UniversityName         string       `json:"university_name"`
		Faculty                int          `json:"faculty"`
		FacultyName            string       `json:"faculty_name"`
		Graduation             int          `json:"graduation"`
		HomeTown               string       `json:"home_town"`
		Relation               int          `json:"relation"`
		Personal struct {
			Religion   string `json:"religion"`
			InspiredBy string `json:"inspired_by"`
			PeopleMain int    `json:"people_main"`
			LifeMain   int    `json:"life_main"`
			Smoking    int    `json:"smoking"`
			Alcohol    int    `json:"alcohol"`
		} `json:"personal"`
		Interests    string       `json:"interests"`
		Music        string       `json:"music"`
		Activities   string       `json:"activities"`
		Movies       string       `json:"movies"`
		Tv           string       `json:"tv"`
		Books        string       `json:"books"`
		Games        string       `json:"games"`
		Universities []University `json:"universities"`
		Schools      []School     `json:"schools"`
		About        string       `json:"about"`
		Relatives    []Relative   `json:"relatives"`
		Quotes       string       `json:"quotes"`
		Deactivated  string       `json:"deactivated"`
	}
	Country struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	
	City struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	
	CropPhoto struct {
		Photo struct {
			ID       int    `json:"id"`
			AlbumID  int    `json:"album_id"`
			OwnerID  int    `json:"owner_id"`
			Photo75  string `json:"photo_75"`
			Photo130 string `json:"photo_130"`
			Photo604 string `json:"photo_604"`
			Width    int    `json:"width"`
			Height   int    `json:"height"`
			Text     string `json:"text"`
			Date     int    `json:"date"`
			PostID   int    `json:"post_id"`
		} `json:"photo"`
		Crop struct {
			X  float64 `json:"x"`
			Y  float64 `json:"y"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"crop"`
		Rect struct {
			X  float64 `json:"x"`
			Y  float64 `json:"y"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"rect"`
	}
	
	Military struct {
		//информация о военной службе пользователя. Объект, содержащий следующие поля:
		Unit       string `json:"unit"`       //(string) — номер части;
		Unit_id    int    `json:"unit_id"`    //(integer) — идентификатор части в базе данных;
		Country_id int    `json:"country_id"` //(integer) — идентификатор страны, в которой находится часть;
		From       int    `json:"from"`       //(integer) — год начала службы;
		Until      int    `json:"until"`      //(integer) — год окончания службы.
	}
	
	Career struct {
		Group_id   int    `json:"group_id"`   //integer) — идентификатор сообщества (если доступно, иначе company);
		Company    string `json:"company"`    // (string) — название компании (если доступно, иначе group_id);
		Country_id int    `json:"country_id"` // (integer) — идентификатор страны;
		City_id    int    `json:"city_id"`    // (integer) — идентификатор города (если доступно, иначе city_name);
		City_name  string `json:"city_name"`  // (string) — название города (если доступно, иначе city_id);
		From       int    `json:"from"`       // (integer) — год начала работы;
		Until      int    `json:"until"`      // (integer) — год окончания работы;
		Position   string `json:"position"`   // (string) — должность.
	}
	Contacts struct {
		Mobile_phone string `json:"mobile_phone,omitempty"`
		Home_phone   string `json:"home_phone,omitempty"`
	}
	
	// GeoPlace contains geographical information like City, Country
	GeoPlace struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	// PlatformInfo contains information about time and platform
	PlatformInfo struct {
		Time     EpochTime `json:"time"`
		Platform int       `json:"platform"`
		
		/*
		platform (integer) — тип платформы, через которую был осуществлён последний вход. Возможные значения:
			1 — мобильная версия;
			2 — приложение для iPhone;
			3 — приложение для iPad;
			4 — приложение для Android;
			5 — приложение для Windows Phone;
			6 — приложение для Windows 10;
			7 — полная версия сайта.
		*/
	}
	// University contains information about the university
	University struct {
		ID              int    `json:"id"`
		Country         int    `json:"country"`
		City            int    `json:"city"`
		Name            string `json:"name"`
		Faculty         int    `json:"faculty"`
		FacultyName     string `json:"faculty_name"`
		Chair           int    `json:"chair"`
		ChairName       string `json:"chair_name"`
		Graduation      int    `json:"graduation"`
		EducationForm   string `json:"education_form"`
		EducationStatus string `json:"education_status"`
	}
	// School contains information about schools
	School struct {
		//ID         int    `json:"id"`
		ID         string `json:"id"`
		Country    int    `json:"country"`
		City       int    `json:"city"`
		Name       string `json:"name"`
		YearFrom   int    `json:"year_from"`
		YearTo     int    `json:"year_to"`
		Class      string `json:"class"`
		TypeStr    string `json:"type_str,omitempty"`
		Speciality string `json:"speciality,omitempty"`
	}
	// Relative contains information about relative to the user
	Relative struct {
		ID   int    `json:"id"`   // negative id describes non-existing users (possibly prepared id if they will register)
		Type string `json:"type"` // like `parent`, `grandparent`, `sibling`
		Name string `json:"name,omitempty"`
	}
)

type UsersGetParams struct {
	UserId   int
	UserIds  []int
	Fields   string
	NameCase string
}

func (o *User) GetById(id int) ([]UsersGetResponse, error) {
	
	p := UsersGetParams{
		UserId: id,
	}
	return o.Get(p)
}

func (o *User) Get(p UsersGetParams) ([]UsersGetResponse, error) {
	
	fields := p.Fields
	if p.Fields == "" {
		fields = "about, activities, bdate, blacklisted, blacklisted_by_me, books, can_post, can_see_all_posts, can_see_audio, can_send_friend_request, can_write_private_message, career, city, common_count, connections, contacts, counters, country, crop_photo, domain, education, exports, first_name_{case}, followers_count, friend_status, games, has_mobile, has_photo, home_town, interests, is_favorite, is_friend, is_hidden_from_feed, last_name_{case}, last_seen, lists, maiden_name, military, movies , music , nickname, occupation, online, personal, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_id, photo_max, photo_max_orig, quotes, relatives, relation, schools, screen_name, sex, site, status, timezone, trending, tv, universities, verified"
	}
	
	user_ids := fmt.Sprint(p.UserId)
	for _, v := range p.UserIds {
		user_ids += "," + fmt.Sprint(v)
	}
	
	params := map[string]string{
		"user_ids":  user_ids,
		"fields":    fields,
		"name_case": p.NameCase,
	}
	
	resp, err := o.vk.Request("users.get", params)
	if err != nil {
		return []UsersGetResponse{}, err
	}
	
	//fmt.Println(string(resp))
	
	var res []UsersGetResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return []UsersGetResponse{}, err
	}
	
	return res, nil
	
}

/*
func (obj *UserObject) setURL() string {
	
	//TODO в сам объект не пишется
	obj.URL	= "https://vk.com/id" + fmt.Sprint(obj.ID)
	return obj.URL
}
*/

type UserGetSubscriptionsParams struct {
	UserId   int
	Extended bool
	Count    int
	Offset   int
	Fields   string
}

type UserGetSubscriptionsIdResponse struct {
	Users struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"users"`
	Groups struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"groups"`
}

//GetSubscriptions на что подписан пользователь
//https://vk.com/dev/users.getSubscriptions
func (o *User) GetSubscriptionsId(p UserGetSubscriptionsParams) (UserGetSubscriptionsIdResponse, error) {
	
	method := "users.getSubscriptions"
	
	//fields := p.Fields
	/*
	var fields string
	if p.Fields == "" {
		fields = "photo_id, verified, sex, bdate, city, country, home_town, has_photo, photo_50, photo_100, photo_200_orig, photo_200, photo_400_orig, photo_max, photo_max_orig, online, domain, has_mobile, contacts, site, education, universities, schools, status, last_seen, followers_count, common_count, occupation, nickname, relatives, relation, personal, connections, exports, wall_comments, activities, interests, music, movies, tv, books, games, about, quotes, can_post, can_see_all_posts, can_see_audio, can_write_private_message, can_send_friend_request, is_favorite, is_hidden_from_feed, timezone, screen_name, maiden_name, crop_photo, is_friend, friend_status, career, military, blacklisted, blacklisted_by_me"
	}*/
	
	params := map[string]string{
		"user_id":  fmt.Sprint(p.UserId),
		"extended": boolConverter(p.Extended),
		"count":    fmt.Sprint(p.Count),
		"offset":   fmt.Sprint(p.Offset),
		//"fields":   fields,
	}
	
	resp, err := o.vk.Request(method, params)
	if err != nil {
		return UserGetSubscriptionsIdResponse{}, err
	}
	//fmt.Println(string(resp))
	
	var res UserGetSubscriptionsIdResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return UserGetSubscriptionsIdResponse{}, err
	}
	return res, nil
}

type UserGetFollowersParams struct {
	UserId int
	Count  int
	Offset int
	Fields string
}

type UserGetFollowersIdResponse struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

//users.getFollowers кто подписан на пользователя
//https://vk.com/dev/users.getFollowers
func (o *User) GetFollowersId(p UserGetFollowersParams) (UserGetFollowersIdResponse, error) {
	
	method := "users.getFollowers"
	
	count := fmt.Sprint(p.Count)
	if p.Count == 0 {
		count = "1000" //MAX
	}
	params := map[string]string{
		"user_id": fmt.Sprint(p.UserId),
		"count":   count,
		"offset":  fmt.Sprint(p.Offset),
		//"fields":   fields,
	}
	
	resp, err := o.vk.Request(method, params)
	if err != nil {
		return UserGetFollowersIdResponse{}, err
	}
	//fmt.Println(string(resp))
	
	var res UserGetFollowersIdResponse
	if err = json.Unmarshal(resp, &res); err != nil {
		return UserGetFollowersIdResponse{}, err
	}
	return res, nil
}
