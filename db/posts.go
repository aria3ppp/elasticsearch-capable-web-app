package db

import (
	"database/sql"
	"fmt"

	"elasticsearch-capable-web-app/entity"
)

var ErrNoRecord = fmt.Errorf("no matching record found")

func (db Database) CreatePost(post *entity.Post) error {
	query := `INSERT INTO posts (title, body, contributed_by) VALUES ($1, $2, 1) RETURNING id`
	err := db.conn.QueryRow(query, post.Title, post.Body).Scan(&post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (db Database) UpdatePost(postId uint, post entity.Post) error {
	query := `UPDATE posts SET title = $1, body = $2 WHERE id = $3`
	err := db.conn.QueryRow(query, post.Title, post.Body, postId).Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNoRecord
		}
		return err
	}
	return nil
}

func (db Database) DeletePost(postId uint) error {
	query := `UPDATE posts SET deleted = TRUE WHERE id = $1`
	err := db.conn.QueryRow(query, postId).Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrNoRecord
		}
		return err
	}
	return nil
}

func (db Database) GetPostById(postId uint) (entity.Post, error) {
	post := entity.Post{}
	query := "SELECT id, title, body FROM posts WHERE NOT deleted AND id = $1"
	row := db.conn.QueryRow(query, postId)
	switch err := row.Scan(&post.ID, &post.Title, &post.Body); err {
	case sql.ErrNoRows:
		return post, ErrNoRecord
	default:
		return post, err
	}
}

func (db Database) GetPosts() ([]entity.Post, error) {
	var list []entity.Post
	query := "SELECT id, title, body FROM posts WHERE NOT deleted ORDER BY id DESC"
	rows, err := db.conn.Query(query)
	if err != nil {
		return list, err
	}
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil {
			return list, err
		}
		list = append(list, post)
	}
	return list, nil
}
