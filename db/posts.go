package db

import (
	"database/sql"
	"fmt"

	"elasticsearch-capable-web-app/models"
)

var ErrNoRecord = fmt.Errorf("no matching record found")

func (db Database) CreatePost(post *models.Post) error {
	query := `INSERT INTO posts_store (title, body) VALUES ($1, $2) RETURNING id`
	err := db.conn.QueryRow(query, post.Title, post.Body).Scan(&post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) modifyPost(
	postId uint,
	delete bool,
	post models.Post,
) error {
	query := `INSERT INTO posts_store (id, title, body, is_deleted) VALUES ($1, $2, $3, $4)`
	err := db.conn.QueryRow(query, postId, post.Title, post.Body, delete).Err()
	if err != nil {
		return err
	}
	return nil
}

func (db Database) UpdatePost(postId uint, post models.Post) error {
	_, err := db.GetPostById(postId)
	if err != nil {
		return err
	}
	err = db.modifyPost(postId, false, post)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) DeletePost(postId uint) error {
	post, err := db.GetPostById(postId)
	if err != nil {
		return err
	}
	err = db.modifyPost(postId, true, post)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) GetPostById(postId uint) (models.Post, error) {
	post := models.Post{}
	query := "SELECT id, title, body FROM posts WHERE id = $1"
	row := db.conn.QueryRow(query, postId)
	switch err := row.Scan(&post.ID, &post.Title, &post.Body); err {
	case sql.ErrNoRows:
		return post, ErrNoRecord
	default:
		return post, err
	}
}

func (db Database) GetPosts() ([]models.Post, error) {
	var list []models.Post
	query := "SELECT id, title, body FROM posts ORDER BY id DESC"
	rows, err := db.conn.Query(query)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			return list, err
		}
		list = append(list, post)
	}
	return list, nil
}
