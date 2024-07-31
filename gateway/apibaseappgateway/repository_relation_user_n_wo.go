package apibaseappgateway

import (
	"backend_base_app/domain/entity"
	"backend_base_app/gateway"
	"backend_base_app/shared/dbhelpers"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateRelationUserWoDataRepo interface {
	CreateRelationUserWoData(ctx context.Context, obj entity.RelationUserWoData) error
	FindOneRelationUserWoDataById(ctx context.Context, id string) (*entity.RelationUserWoData, error)
	FindAllRelationUserWoData(ctx context.Context, req entity.BaseReqFind) ([]*entity.RelationUserWoData, int64, error)
	DeleteOneRelationUserWoData(ctx context.Context, id string) (bool, error)
	DeleteRelationUserWoData(ctx context.Context, obj entity.FindRelationUserWoData) (bool, error)
}

type relationUserWoCollection struct {
	*mongo.Collection
}

func (r GatewayApiBaseApp) getRelationUserWoCollection() relationUserWoCollection {
	return relationUserWoCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionRelationUserWo),
	}
}

func getFilterRelationUserWoKeyword(obj entity.FindRelationUserWoData) primitive.M {
	//====== execute query using transaction ======
	//count the existing users
	keywordFilter := make([]bson.M, 0)

	if obj.IDUser != nil && *obj.IDUser != "" {
		keyword := bson.M{"id_user": obj.IDUser}
		keywordFilter = append(keywordFilter, keyword)
	}
	if obj.IDWeddingOrg != nil && *obj.IDWeddingOrg != "" {
		keyword := bson.M{"id_wedding_org": obj.IDWeddingOrg}
		keywordFilter = append(keywordFilter, keyword)
	}

	var allCriteria []bson.M
	var criteriaKeyword bson.M
	if len(keywordFilter) > 0 {
		criteriaKeyword = bson.M{"$or": keywordFilter}
		allCriteria = append(allCriteria, criteriaKeyword)
	}

	criteria := bson.M{}
	if len(allCriteria) > 0 {
		criteria = bson.M{"$and": allCriteria}
	}
	return criteria
}

func (coll relationUserWoCollection) GetTotalRelationUserWo(ctx context.Context, obj entity.FindRelationUserWoData) (int64, error) {
	criteria := getFilterRelationUserWoKeyword(obj)

	countOpts := options.CountOptions{}
	return coll.CountDocuments(ctx, criteria, &countOpts)
}

func (r GatewayApiBaseApp) CreateRelationUserWoData(ctx context.Context, obj entity.RelationUserWoData) error {
	var err error

	randomId := util.GenerateTimeUuidWithoutDash()
	id, err := entity.NewRelationUserWoDataID(randomId)
	if err != nil {
		return err
	}
	obj.ID = id

	relationUserWoCollection := &relationUserWoCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionRelationUserWo),
	}

	err = dbhelpers.WithTransaction(ctx, r.MongoWithTransactionImpl, func(dbCtx context.Context) error {
		info, err := relationUserWoCollection.InsertOne(ctx, obj)
		log.Info(ctx, "info >>> ", info)
		return err
	})

	return err
}

func (r GatewayApiBaseApp) FindOneRelationUserWoDataById(ctx context.Context, id string) (*entity.RelationUserWoData, error) {
	var (
		resultRelationUserWoData *entity.RelationUserWoData
		err                      error
	)

	coll := r.getRelationUserWoCollection()
	resCol := coll.FindOne(ctx, bson.M{"id": id})
	err = resCol.Decode(&resultRelationUserWoData)
	if err != nil {
		log.Error(ctx, err.Error())

		// error return if data not found
		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("organizer wedding data not found")
		}

		return nil, err
	}
	return resultRelationUserWoData, err
}

func (r GatewayApiBaseApp) FindAllRelationUserWoData(ctx context.Context, req entity.BaseReqFind) ([]*entity.RelationUserWoData, int64, error) {
	var (
		err  error
		objs []*entity.RelationUserWoData
	)

	coll := r.getRelationUserWoCollection()

	findData, _ := req.Value.(entity.FindRelationUserWoData)
	criteria := getFilterRelationUserWoKeyword(findData)

	findOpts := gateway.BaseReqFindToOptOption(req)
	cursor, err := coll.Find(ctx, criteria, &findOpts)
	if err != nil {
		return nil, 0, err
	}

	if err := cursor.All(ctx, &objs); err != nil {
		return nil, 0, err
	}

	//counting
	count, err := coll.GetTotalRelationUserWo(ctx, findData)

	return objs, count, err
}

func (r GatewayApiBaseApp) DeleteOneRelationUserWoData(ctx context.Context, id string) (bool, error) {
	var (
		err error
	)

	coll := r.getRelationUserWoCollection()
	resCol, err := coll.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		log.Error(ctx, err.Error())

		// error return if data not found
		if err.Error() == "mongo: no documents in result" {
			err = fmt.Errorf("organizer wedding data not found")
		}

		return false, err
	}

	return resCol.DeletedCount > 0, err
}

func (r GatewayApiBaseApp) DeleteRelationUserWoData(ctx context.Context, obj entity.FindRelationUserWoData) (bool, error) {
	var (
		err error
	)

	criteria := getFilterRelationUserWoKeyword(obj)

	coll := r.getRelationUserWoCollection()
	resCol, err := coll.DeleteMany(ctx, criteria)
	if err != nil {
		log.Error(ctx, err.Error())

		// error return if data not found
		if err.Error() == "mongo: no documents in result" {
			err = fmt.Errorf("organizer wedding data not found")
		}

		return false, err
	}

	return resCol.DeletedCount > 0, err
}
