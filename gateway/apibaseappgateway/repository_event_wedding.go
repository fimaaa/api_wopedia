package apibaseappgateway

import (
	"backend_base_app/domain/entity"
	"backend_base_app/gateway"
	"backend_base_app/shared/dbhelpers"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateEventWeddingDataRepo interface {
	CreateEventWeddingData(ctx context.Context, obj entity.EventWeddingData) error
	FindOneEventWeddingDataById(ctx context.Context, id string) (*entity.EventWeddingData, error)
	FindAllEventWeddingData(ctx context.Context, req entity.BaseReqFind) ([]*entity.EventWeddingData, int64, error)
	UpdateEventWeddingData(ctx context.Context, editEventWeddingData entity.EditEventWeddingData) (*entity.EventWeddingData, error)
	DeleteEventWeddingData(ctx context.Context, id string) (bool, error)
}

type eventWeddingCollection struct {
	*mongo.Collection
}

func (r GatewayApiBaseApp) getEventWeddingCollection() eventWeddingCollection {
	return eventWeddingCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionEventWedding),
	}
}

func getFilterEventWeddingKeyword(obj entity.FindEventWeddingData) primitive.M {
	//====== execute query using transaction ======
	//count the existing users
	keywordFilter := make([]bson.M, 0)

	switch {
	case obj.MinTotalInvited != nil && obj.MaxTotalInvited != nil:
		{
			keyword := bson.M{
				"total_invited": bson.M{
					"$gte": obj.MinTotalInvited,
					"$lte": obj.MaxTotalInvited,
				},
			}
			// Append the date filter to the keywordFilter slice
			keywordFilter = append(keywordFilter, keyword)
		}
	case obj.MaxTotalInvited != nil:
		{
			keyword := bson.M{
				"total_invited": bson.M{
					"$lte": obj.MaxTotalInvited,
				},
			}
			// Append the date filter to the keywordFilter slice
			keywordFilter = append(keywordFilter, keyword)
		}
	case obj.MinTotalInvited != nil:
		{
			keyword := bson.M{
				"total_invited": bson.M{
					"$gte": obj.MinTotalInvited,
				},
			}
			// Append the date filter to the keywordFilter slice
			keywordFilter = append(keywordFilter, keyword)
		}
	}

	list, err := util.GenerateMongoFilter(obj)
	if err == nil {
		keywordFilter = append(keywordFilter, list...)
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

func (coll eventWeddingCollection) GetTotalEventWedding(ctx context.Context, obj entity.FindEventWeddingData, onlySimiliar bool) (int64, error) {
	criteria := getFilterEventWeddingKeyword(obj)

	countOpts := options.CountOptions{}
	return coll.CountDocuments(ctx, criteria, &countOpts)
}

func (r GatewayApiBaseApp) CreateEventWeddingData(ctx context.Context, obj entity.EventWeddingData) error {
	var err error

	randomId := util.GenerateTimeUuidWithoutDash()
	id, err := entity.NewEventWeddingDataID(randomId)
	if err != nil {
		return err
	}
	obj.ID = id

	eventWeddingCollection := &eventWeddingCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionEventWedding),
	}

	err = dbhelpers.WithTransaction(ctx, r.MongoWithTransactionImpl, func(dbCtx context.Context) error {
		info, err := eventWeddingCollection.InsertOne(ctx, obj)
		log.Info(ctx, "info >>> ", info)
		return err
	})

	return err
}

func (r GatewayApiBaseApp) FindOneEventWeddingDataById(ctx context.Context, id string) (*entity.EventWeddingData, error) {
	var (
		resultEventWeddingData *entity.EventWeddingData
		err                    error
	)

	coll := r.getEventWeddingCollection()
	resCol := coll.FindOne(ctx, bson.M{"id": id})
	err = resCol.Decode(&resultEventWeddingData)
	if err != nil {
		log.Error(ctx, err.Error())

		// error return if data not found
		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("organizer wedding data not found")
		}

		return nil, err
	}
	return resultEventWeddingData, err
}

func (r GatewayApiBaseApp) FindAllEventWeddingData(ctx context.Context, req entity.BaseReqFind) ([]*entity.EventWeddingData, int64, error) {
	var (
		err  error
		objs []*entity.EventWeddingData
	)

	coll := r.getEventWeddingCollection()

	findData, _ := req.Value.(entity.FindEventWeddingData)
	criteria := getFilterEventWeddingKeyword(findData)

	findOpts := gateway.BaseReqFindToOptOption(req)

	cursor, err := coll.Find(ctx, criteria, &findOpts)
	if err != nil {
		return nil, 0, err
	}

	if err := cursor.All(ctx, &objs); err != nil {
		return nil, 0, err
	}

	//counting
	count, err := coll.GetTotalEventWedding(ctx, findData, true)

	return objs, count, err
}

func (r GatewayApiBaseApp) UpdateEventWeddingData(ctx context.Context, editEventWeddingData entity.EditEventWeddingData) (*entity.EventWeddingData, error) {
	log.Info(ctx, "called")

	editEventWeddingData.UpdatedAt = time.Now().Local().UTC()

	editBson := util.StructToBSONM(editEventWeddingData)

	info, err := r.MongoWithTransactionImpl.UpdateByCustomId(ctx, r.database, entity.CollectionEventWedding, editEventWeddingData.ID.String(), editBson)
	log.Info(ctx, "info >>> ", info)
	if err != nil {
		log.Info(ctx, "error >>> "+err.Error())
	}
	dataResult, ok := info.(entity.EventWeddingData)
	if !ok {
		return nil, fmt.Errorf("organizer Transform")
	}

	return &dataResult, err
}

func (r GatewayApiBaseApp) DeleteEventWeddingData(ctx context.Context, id string) (bool, error) {
	var (
		err error
	)

	coll := r.getEventWeddingCollection()
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

func (r GatewayApiBaseApp) GetTotalInvitedEventWedding(ctx context.Context, id string) (int, error) {
	coll := r.getEventWeddingCollection()
	projection := bson.M{"total_invited": 1, "_id": 0} // Projection to include only total_invited field
	resCol := coll.FindOne(ctx, bson.M{"id": id}, options.FindOne().SetProjection(projection))

	// Decode the result
	var result struct {
		TotalInvited *int `bson:"total_invited"`
	}
	if err := resCol.Decode(&result); err != nil {
		return -1, err
	}
	// Access the totalInvited field
	totalInvited := result.TotalInvited
	return *totalInvited, nil
}
