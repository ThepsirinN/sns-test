package repositories

import (
	"context"
	"errors"
	"sns-barko/constant"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"
	"strconv"
)

func (r *repoV1) CreateFriendRequest(ctx context.Context, model models.Friend) error {
	tx := r.db.WithContext(ctx).Create(&model)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *repoV1) GetAllFriendRequest(ctx context.Context, id int32, model *[]entities.GetAllFriendRequestResponse) error {
	tx := r.db.WithContext(ctx).Raw(`SELECT
	f.id,
	f.source_id,
	f.destination_id dest_id,
	f.status,
	u1.email source_email,
	u1.first_name source_first_name,
	u1.last_name source_last_name,
	u1.img_profile source_profile_image,
	u2.email dest_email,
	u2.first_name dest_first_name,
	u2.last_name dest_last_name,
	u2.img_profile dest_profile_image
FROM
	friend f
	INNER JOIN USER u1 ON f.source_id = u1.id
	INNER JOIN USER u2 ON f.destination_id = u2.id
WHERE
	(f.source_id = ? OR f.destination_id = ?) AND
	f.status = ?`, id, id, constant.FRIEND_STATUS_PENDING).Find(model)
	if err := tx.Error; err != nil {
		return err
	}

	if tx.RowsAffected == 0 {
		return errors.New("not found data from user " + strconv.Itoa(int(id)))
	}

	return nil
}

func (r *repoV1) ListFriend(ctx context.Context, id int32, model *[]entities.ListFriendQuery) error {
	tx := r.db.WithContext(ctx).Raw(`SELECT
	f.id,
	f.source_id,
	f.destination_id dest_id,
	u1.email source_email,
	u1.first_name source_first_name,
	u1.last_name source_last_name,
	u1.img_profile source_profile_image,
	u2.email dest_email,
	u2.first_name dest_first_name,
	u2.last_name dest_last_name,
	u2.img_profile dest_profile_image
FROM
	friend f
	INNER JOIN USER u1 ON f.source_id = u1.id
	INNER JOIN USER u2 ON f.destination_id = u2.id
WHERE
	(f.source_id = ? OR f.destination_id = ?) AND
	f.status = ?`, id, id, constant.FRIEND_STATUS_SUCCESS).Find(model)
	if err := tx.Error; err != nil {
		return err
	}

	if tx.RowsAffected == 0 {
		return errors.New("not found data from user " + strconv.Itoa(int(id)))
	}

	return nil
}

func (r *repoV1) UpdateFriendRequestStatus(ctx context.Context, model models.Friend) error {
	tx := r.db.Updates(model)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("cannot update data " + strconv.Itoa(int(model.Id)))
	}

	return nil
}

func (r *repoV1) DeleteFriend(ctx context.Context, model models.Friend) error {
	tx := r.db.Where(models.Friend{Id: model.Id, SourceId: model.SourceId, DestId: model.DestId}).Delete(model)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("cannot delete data " + strconv.Itoa(int(model.Id)))
	}

	return nil
}
