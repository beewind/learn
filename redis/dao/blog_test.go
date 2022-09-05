package dao

import "testing"

func TestBlog(t *testing.T) {
	t.Log(QueryHotBlog(1, 1))
}
