// Package easyvk provides you simple way
// to work with VK API.
package easyvk

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"golang.org/x/net/html"
	"net/http/cookiejar"
	"io"
	"errors"
	"time"
	//"bytes"
	"strings"
	"log"
	"github.com/influxdata/influxdb-client"
)

const (
	version = "5.71"
	apiURL  = "https://api.vk.com/method/"
	authURL = "https://oauth.vk.com/authorize?" +
		"client_id=%s" +
		"&scope=%s" +
		"&redirect_uri=https://oauth.vk.com/blank.html" +
		"&display=wap" +
		"&v=%s" +
		"&response_type=token"
	vkReqPerSec  = 3
	vkReqEveryMs = 340 * time.Millisecond //1s = 1000ms //350-Ок
)

var LogResp func(string, map[string]string, *json.RawMessage)

// VK defines a set of functions for
// working with VK API.
type VK struct {
	AccessToken string
	LastTimeReq time.Time
	Version     string
	Account     Account
	Board       Board
	Fave        Fave
	Likes       Likes
	Photos      Photos
	Status      Status
	Upload      Upload
	Wall        Wall
	User        User
	Friends     Friends
	Group       Group
	Apps        Apps
}

var logT *influxdb.UDPWriter

func init() {
	
	inflixServer := "localhost:8089"
	
	var err error
	logT, err = influxdb.NewUDPWriter(inflixServer)
	if err != nil {
		log.Println(err)
	}
}

func SetLogger(fn func(string, map[string]string, *json.RawMessage)) {
	LogResp = fn
}

// WithToken helps to initialize your
// VK object with token.
func WithToken(token string) VK {
	vk := VK{}
	vk.AccessToken = token
	vk.Version = version
	vk.Account = Account{&vk}
	vk.Board = Board{&vk}
	vk.Fave = Fave{&vk}
	vk.Likes = Likes{&vk}
	vk.Photos = Photos{&vk}
	vk.Status = Status{&vk}
	vk.Upload = Upload{}
	vk.Wall = Wall{&vk}
	vk.User = User{&vk}
	vk.Friends = Friends{&vk}
	vk.Group = Group{&vk}
	vk.Apps = Apps{&vk}
	return vk
}

// WithAuth helps to initialize your VK object
// with signing in by login, password, client id and scope
// Scope must be a string like "friends,wall"
func WithAuth(login, password, clientID, scope string) (VK, error) {
	u := fmt.Sprintf(authURL, clientID, scope, version)
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	
	resp, err := client.Get(u)
	if err != nil {
		return VK{}, err
	}
	defer resp.Body.Close()
	
	args, u := parseForm(resp.Body)
	
	args.Add("email", login)
	args.Add("pass", password)
	
	resp, err = client.PostForm(u, args)
	if err != nil {
		return VK{}, err
	}
	
	if resp.Request.URL.Path != "/blank.html" {
		args, u := parseForm(resp.Body)
		resp, err = client.PostForm(u, args)
		if err != nil {
			return VK{}, err
		}
		defer resp.Body.Close()
		
		if resp.Request.URL.Path != "/blank.html" {
			return VK{}, errors.New("can't log in")
		}
	}
	
	urlArgs, err := url.ParseQuery(resp.Request.URL.Fragment)
	if err != nil {
		return VK{}, err
	}
	
	return WithToken(urlArgs["access_token"][0]), nil
}

func parseForm(body io.ReadCloser) (url.Values, string) {
	tokenizer := html.NewTokenizer(body)
	
	u := ""
	formData := map[string]string{}
	
	end := false
	for !end {
		tag := tokenizer.Next()
		
		switch tag {
		case html.ErrorToken:
			end = true
		case html.StartTagToken:
			switch token := tokenizer.Token(); token.Data {
			case "form":
				for _, attr := range token.Attr {
					if attr.Key == "action" {
						u = attr.Val
					}
				}
			case "input":
				if token.Attr[1].Val == "_origin" {
					formData["_origin"] = token.Attr[2].Val
				}
				if token.Attr[1].Val == "to" {
					formData["to"] = token.Attr[2].Val
				}
			}
		case html.SelfClosingTagToken:
			switch token := tokenizer.Token(); token.Data {
			case "input":
				if token.Attr[1].Val == "ip_h" {
					formData["ip_h"] = token.Attr[2].Val
				}
				if token.Attr[1].Val == "lg_h" {
					formData["lg_h"] = token.Attr[2].Val
				}
			}
		}
	}
	
	args := url.Values{}
	
	for key, val := range formData {
		args.Add(key, val)
	}
	
	return args, u
}

// Request provides access to VK API methods.
func (vk *VK) Request(method string, params map[string]string) ([]byte, error) {
	u, err := url.Parse(apiURL + method)
	if err != nil {
		return nil, err
	}
	
	query := url.Values{}
	for k, v := range params {
		query.Set(k, v)
	}
	query.Set("access_token", vk.AccessToken)
	if params["Version"] == "" {
		query.Set("v", vk.Version)
		params["v"] = vk.Version
	}
	u.RawQuery = query.Encode()
	
	//Попытки получить данные
	tryCount := 0
try:
	tryCount++
	
	targetTime := vk.LastTimeReq.Add(vkReqEveryMs)
	if time.Now().Before(targetTime) {
		time.Sleep(targetTime.Sub(time.Now()))
	}
	logReqDelay := time.Since(vk.LastTimeReq) //Сколько мы ждали с предыдущего вызова
	//fmt.Printf("[DURATION] vkGet %v\n", logReqDelay)
	vk.LastTimeReq = time.Now()
	
	//Основной Запрос в ВК
	cl := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Panicln("[ERR http.NewRequest]", err)
	}
	//req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.167 Safari/537.36")
	//req.Header.Add("accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	//req.Header.Add("accept-encoding","gzip, deflate, br")
	//req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	
	APDEX_vk_start := time.Now() //Тут счиатем чистое время VK на запрос
	resp, err := cl.Do(req)      //---------========  VK  ==========---------------
	//resp, vkErr := http.Get(u.String())
	APDEX_vk_do := time.Since(APDEX_vk_start)
	
	if err != nil {
		log.Println("[ERR vkErr] tryCount:", tryCount, err)
		//read tcp 192.168.100.10:52059->87.240.129.179:443: wsarecv: An existing connection was forcibly closed by the remote host.
		//if strings.Contains(err.Error(),"(Client.Timeout exceeded while awaiting headers)") && tryCount <= 10 {
		if tryCount <= 10 {
			time.Sleep(10 * time.Second)
			goto try
		}
		return nil, err
	}
	defer resp.Body.Close()
	
	//fmt.Printf("APDEX_vk %s %v ReqDelay:%v Token:%s\n", method, APDEX_vk, logReqDelay, vk.AccessToken)
	
	//Log - APDEX
	pt := influxdb.Point{
		Name: "apdex_vk",
		Tags: []influxdb.Tag{{
			Key:   "Method",
			Value: strings.Replace(method, ".", "_", -1),
		}},
		Fields: map[string]interface{}{
			"APDEX_vk": APDEX_vk_do.Seconds() * 1000,
			"ReqDelay": logReqDelay.Seconds() * 1000,
		},
		Time: time.Now(),
	}
	go pt.WriteTo(logT)
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var handler struct {
		Error    *Error
		Response json.RawMessage
	}
	err = json.Unmarshal(body, &handler)
	
	if handler.Error != nil {
		return nil, handler.Error
	}
	
	if LogResp != nil {
		//LogResp(method, params, &handler.Response)
		go LogResp(method, params, &handler.Response)
	}
	
	return handler.Response, nil
}
