package db

import (
	"context"
	"database/sql"
	"fmt"

	"elasticsearch-capable-web-app/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var ErrNoRecord = fmt.Errorf("no matching record found")

func (db Database) CreatePost(post *models.Post) error {
	post.UserID = defaultUser.ID
	return post.Insert(context.Background(), db.conn, boil.Infer())
}

func (db Database) UpdatePost(postId int, post models.Post) error {
	_, err := models.Posts(
		models.PostWhere.ID.EQ(postId),
	).UpdateAll(
		context.Background(),
		db.conn,
		models.M{
			models.PostColumns.UserID: defaultUser.ID,
			models.PostColumns.Title:  post.Title,
			models.PostColumns.Body:   post.Body,
		},
	)
	if err == sql.ErrNoRows {
		return ErrNoRecord
	}
	return err
}

func (db Database) DeletePost(postId int) error {
	_, err := models.Posts(
		models.PostWhere.ID.EQ(postId),
	).UpdateAll(
		context.Background(),
		db.conn,
		models.M{
			models.PostColumns.UserID:  defaultUser.ID,
			models.PostColumns.Deleted: true,
		},
	)
	if err == sql.ErrNoRows {
		return ErrNoRecord
	}
	return err
}

func (db Database) GetPostById(postId int) (*models.Post, error) {
	user, err := models.Posts(
		models.PostWhere.ID.EQ(postId),
		models.PostWhere.Deleted.EQ(false),
	).One(context.Background(), db.conn)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoRecord
		}
		return nil, err
	}
	return user, nil
}

func (db Database) GetPosts() ([]*models.Post, error) {
	posts, err := models.Posts(
		models.PostWhere.Deleted.EQ(false),
	).All(context.Background(), db.conn)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	}
	return posts, nil
}
