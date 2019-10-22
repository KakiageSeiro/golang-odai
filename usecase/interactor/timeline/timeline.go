package timeline

import (
	"context"
	"golang-odai/domain"
	"golang-odai/usecase/repository"
)

type TimeLine interface {
	GetLatestList(ctx context.Context) ([]domain.Post, error)
}

type timeLine struct {
	repoPost repository.Post
	repoUser repository.User
}

func New(repoPost repository.Post, repoUser repository.User) TimeLine {
	return &timeLine{
		repoPost: repoPost,
		repoUser: repoUser,
	}
}

func (tl *timeLine) GetLatestList(ctx context.Context) ([]domain.Post, error) {
	posts, err := tl.repoPost.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if err := tl.fetch(ctx, posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (tl *timeLine) fetch(ctx context.Context, ps []domain.Post) error {
	var uIDs []uint32
	for _, post := range ps {
		uIDs = append(uIDs, post.UserID)
	}

	us, err := tl.repoUser.FindByIDs(ctx, uIDs...)
	if err != nil {
		return err
	}

	usMap := make(map[uint32]domain.User)
	for _, u := range us {
		usMap[u.ID] = u
	}

	for idx, post := range ps {
		u := usMap[post.UserID]
		post.User = &u
		ps[idx] = post
	}

	return nil
}
