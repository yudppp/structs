package structs

import (
	"encoding/json"
	"testing"
)

func TestNewExample(t *testing.T) {

	type User struct {
		Name string `example:"ichiro"`
	}

	type Simple struct {
		ID      *int   `example:"1"`
		Name    string `example:"hello"`
		Active  bool   `example:"true"`
		User    User
		Tags    *[]int `example:"1,3"`
		Friends []*User
	}

	example := NewExample(Simple{}).(Simple)
	b, _ := json.Marshal(example)
	expect := `{"ID":1,"Name":"hello","Active":true,"User":{"Name":"ichiro"},"Tags":[1,3],"Friends":[{"Name":"ichiro"}]}`
	if expect != string(b) {
		t.Errorf("failed test exmaple 1")
	}

	type DataCategory struct {
		Name string `json:"name" example:"aws"`
		Num  int    `json:"num" example:"123"`
	}

	type DataPost struct {
		AcceptComment bool   `json:"accept_comment" example:"true"`
		Status        int    `json:"status" example:"1"`
		Tags          []int  `json:"tags" example:"1,2,4"`
		Title         string `json:"title" example:"hello world"`
	}

	type Data struct {
		Categories      []DataCategory `json:"categories"`
		Comments        []interface{}  `json:"comments"`
		Post            DataPost       `json:"post"`
		ProfileImageURL string         `json:"profile_image_url" example:"http://blog.yudppp.com/img/profile.gif"`
		URL             string         `json:"url" example:"http://blog.yudppp.com"`
	}

	example2 := NewExample(Data{}).(Data)
	b, _ = json.Marshal(example2)
	expect = `{"categories":[{"name":"aws","num":123}],"comments":[],"post":{"accept_comment":true,"status":1,"tags":[1,2,4],"title":"hello world"},"profile_image_url":"http://blog.yudppp.com/img/profile.gif","url":"http://blog.yudppp.com"}`
	if expect != string(b) {
		t.Errorf("failed test exmaple 1")
	}

}
