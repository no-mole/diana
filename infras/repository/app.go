package repository

import (
	"context"
	"diana/core"
	"diana/core/app"
	"github.com/no-mole/neptune/database"
)

type AppRepository interface {
	Store(app *app.App) (*app.App, error)
	Delete(appId string) error
	Fetch(appId string) (*app.App, error)
}

type appRepository struct {
	db *database.BaseDb
}

func New(ctx context.Context) AppRepository {
	appDb := &database.BaseDb{}
	appDb.SetEngine(ctx, "app")
	return &appRepository{
		db: appDb,
	}
}

func (repo *appRepository) Store(app *app.App) (*app.App, error) {
	if err := repo.storeSubjects(app.AppId, app.GetAddedSubject()); err != nil {
		return nil, err
	}
	if err := repo.removeSubjects(app.AppId, app.GetRemovedSubject()); err != nil {
		return nil, err
	}

	if err := repo.storeObjects(app.AppId, app.GetAddedObjects()); err != nil {
		return nil, err
	}
	if err := repo.removeObjects(app.AppId, app.GetRemovedObjects()); err != nil {
		return nil, err
	}

	repo.db.Create(app)
	return nil, nil
}

func (repo *appRepository) Delete(appId string) error {
	return nil
}

func (repo *appRepository) convertSubjects(appId string, subjects []core.Subject) []*AppSubjectRef {
	appSubjectRefs := make([]*AppSubjectRef, 0)
	for _, subject := range subjects {
		appSubjectRefs = append(appSubjectRefs, &AppSubjectRef{
			subjectId:   subject.GetId(),
			subjectType: int32(subject.GetType()),
			appId:       appId,
		})
	}
	return appSubjectRefs
}

func (repo *appRepository) convertObjects(appId string, objs []core.Object) []*AppObjectRef {
	appObjectRefs := make([]*AppObjectRef, 0)
	for _, obj := range objs {
		appObjectRefs = append(appObjectRefs, &AppObjectRef{
			objectId:   obj.GetId(),
			objectType: int32(obj.GetType()),
			appId:      appId,
		})
	}
	return appObjectRefs
}

func (repo *appRepository) storeSubjects(appId string, subjects []core.Subject) error {
	if len(subjects) == 0 {
		return nil
	}
	appSubjectRefs := repo.convertSubjects(appId, subjects)
	return repo.db.Create(appSubjectRefs).Error
}

func (repo *appRepository) removeSubjects(appId string, subjects []core.Subject) error {
	if len(subjects) == 0 {
		return nil
	}
	appSubjectRefs := repo.convertSubjects(appId, subjects)
	return repo.db.Delete(appSubjectRefs).Error
}

func (repo *appRepository) storeObjects(appId string, objs []core.Object) error {
	if len(objs) == 0 {
		return nil
	}

	appObjectRefs := repo.convertObjects(appId, objs)
	return repo.db.Create(appObjectRefs).Error
}

func (repo *appRepository) removeObjects(appId string, objs []core.Object) error {
	if len(objs) == 0 {
		return nil
	}
	appObjectRefs := repo.convertObjects(appId, objs)
	return repo.db.Delete(appObjectRefs).Error
}

func (repo *appRepository) Fetch(appId string) (*app.App, error) {
	return nil, nil
}
