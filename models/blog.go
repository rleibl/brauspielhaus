
package models

import (
	"encoding/json"
	"io/ioutil"
)

type BlogEntry struct {
	Id      int    `json:id`
	Date    string `json:date`
	Title   string `json:title`
	Content string `json:content`
	Author  string `json:author`
}

// Assume all Blog entries are in the same file
func LoadBlogEntriesFromJson(filename string) []BlogEntry {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return BlogFromJson(content)
}

func BlogFromJson(b []byte) []BlogEntry {

	var be []BlogEntry

	err := json.Unmarshal(b, &be)
	if err != nil {
		panic(err)
	}

	return be
}
