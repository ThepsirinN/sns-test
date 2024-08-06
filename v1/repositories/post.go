package repositories

import (
	"context"
	"errors"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"
	"strconv"
)

func (r *repoV1) CreatePost(ctx context.Context, model models.Post) error {
	tx := r.db.WithContext(ctx).Create(&model)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *repoV1) ListAllPostFromUser(ctx context.Context, userId int32, resp *[]entities.ListAllPostFromUserResponse) error {
	tx := r.db.WithContext(ctx).Raw(`SELECT 
	p.id,
	p.owner_id,
	u.first_name owner_first_name,
	u.last_name owner_last_name,
	u.img_profile owner_profile_img,
	p.post_data,
	p.post_img,
	p.comment,
	p.like,
	p.created_at,
	p.updated_at
FROM
	post p
	INNER JOIN USER u ON u.id = p.owner_id
WHERE
	p.owner_id = ?`, userId).Find(resp)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("not found data from " + strconv.Itoa(int(userId)))
	}
	return nil
}

func (r *repoV1) ReadPostByPostId(ctx context.Context, resp *entities.ReadPostByPostIdResponse) error {
	tx := r.db.WithContext(ctx).Raw(`SELECT
	p.id,
	p.owner_id,
	u.first_name owner_first_name,
	u.last_name owner_last_name,
	u.img_profile owner_profile_img,
	p.post_data,
	p.post_img,
	p.comment,
	p.like,
	p.created_at,
	p.updated_at
FROM
	post p
	INNER JOIN USER u ON u.id = p.owner_id
WHERE
	p.owner_id = ?
	AND p.id = ?
LIMIT 1;`, resp.OwnerId, resp.Id).Find(resp)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("not found data from " + strconv.Itoa(int(resp.Id)))
	}
	return nil
}

func (r *repoV1) UpdatePostData(ctx context.Context, model models.Post) error {
	tx := r.db.WithContext(ctx).Where(models.Post{Id: model.Id, OwnerId: model.OwnerId}).Updates(model)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *repoV1) DeletePost(ctx context.Context, model models.Post) error {
	tx := r.db.WithContext(ctx).Where(models.Post{Id: model.Id, OwnerId: model.OwnerId}).Delete(model)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
