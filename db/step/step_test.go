package step

import (
	"encoding/json"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	data := []byte(`{
		"id":1,
		"reader_num":3,
		"tag_num":2,
		"start":"` + time.Now().Format(layout) + `",
		"is_remove":true
	 }`)
	step := New()
	json.Unmarshal(data, &step)
	Create(step)
}

func TestSelect(t *testing.T) {
	SelectAll()
}
