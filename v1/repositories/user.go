package repositories

import (
	"context"
	"errors"
	"fmt"
	"sns-barko/v1/models"
)

func (r *repoV1) AutoMigrate(ctx context.Context) error {
	return r.db.WithContext(ctx).AutoMigrate(&models.User{})
}

func (r *repoV1) CreateUser(ctx context.Context, model models.User) error {
	tx := r.db.WithContext(ctx).Create(&model)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *repoV1) ReadUsersByEmail(ctx context.Context, model *models.User) error {
	tx := r.db.WithContext(ctx).Where(models.User{Email: model.Email}).Where("deleted_at IS NULL").Limit(1).Find(model)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("not found data from " + model.Email)
	}

	return nil
}

func (r *repoV1) ReadUsersById(ctx context.Context, model *models.User) error {
	tx := r.db.WithContext(ctx).Where(models.User{Id: model.Id}).Where("deleted_at IS NULL").Limit(1).Find(model)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("not found data from " + string(model.Id))
	}

	return nil
}

// func (r *repoV1) ReadAllUsers(ctx context.Context) error {
// 	return nil
// }

func (r *repoV1) FindUsersByEmail(ctx context.Context, email string, userId int32, model *[]models.User) error {
	tx := r.db.WithContext(ctx).Model(model).Where("email LIKE ? AND id <> ?", fmt.Sprint("%"+email+"%"), userId).Find(&model)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("not found data from " + string(email))
	}
	return nil
}

func (r *repoV1) UpdateUsers(ctx context.Context, model models.User) error {
	tx := r.db.WithContext(ctx).Updates(model)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("cannot update data from " + string(model.Id))
	}

	return nil
}

func (r *repoV1) DeleteUserById(ctx context.Context, model models.User) error {
	tx := r.db.WithContext(ctx).Updates(model)
	if err := tx.Error; err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return errors.New("cannot delete data from " + string(model.Id))
	}

	return nil
}
