package post

import "testing"

type Post struct {
	Name    string
	Address string
}

type Service interface {
	ListPosts() ([]*Post, error)
}

func ListPosts(svc Service) ([]*Post, error) {
	return svc.ListPosts()
}

type fakeService struct {
}

func (s *fakeService) ListPosts() ([]*Post, error) {
	posts := make([]*Post, 0)
	posts = append(posts, &Post{
		Name:    "colin",
		Address: "Shenzhen",
	})
	posts = append(posts, &Post{
		Name:    "alex",
		Address: "Beijing",
	})
	return posts, nil
}

func NewFakeService() Service {
	return &fakeService{}
}

func TestListPosts(t *testing.T) {
	fake := NewFakeService()
	if _, err := ListPosts(fake); err != nil {
		t.Fatal("list posts failed")
	}
}
