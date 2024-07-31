package apibaseappgateway

import (
	"backend_base_app/domain/entity"
	"backend_base_app/gateway"
	"backend_base_app/shared/dbhelpers"
	"backend_base_app/shared/log"
	"backend_base_app/shared/util"
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CreateEventTestimonyDataRepo interface {
	CreateEventTestimonyData(ctx context.Context, obj entity.EventTestimonyData, testimonyQty int) error
	FindOneEventTestimonyDataById(ctx context.Context, id string) (*entity.EventTestimonyData, error)
	FindAllEventTestimonyData(ctx context.Context, req entity.BaseReqFind) ([]*entity.EventTestimonyData, int64, error)
	UpdateEventTestimonyData(ctx context.Context, eventTestimonyData entity.EditEventTestimonyData) (*entity.EventTestimonyData, error)
	DeleteEventTestimonyData(ctx context.Context, id string) (bool, error)
}

type eventTestimonyCollection struct {
	*mongo.Collection
}

func (r GatewayApiBaseApp) getEventTestimonyCollection() eventTestimonyCollection {
	return eventTestimonyCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionEventTestimony),
	}
}

func getFilterEventTestimonyKeyword(
	obj entity.FindEventTestimonyData,
	onlySimiliar bool,
) primitive.M {
	keywordFilter := make([]bson.M, 0)

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

func (coll eventTestimonyCollection) GetTotalEventTestimony(ctx context.Context, obj entity.FindEventTestimonyData, onlySimiliar bool) (int64, error) {
	criteria := getFilterEventTestimonyKeyword(obj, onlySimiliar)

	countOpts := options.CountOptions{}
	return coll.CountDocuments(ctx, criteria, &countOpts)
}

func (r GatewayApiBaseApp) CreateEventTestimonyData(ctx context.Context, obj entity.EventTestimonyData, testimonyQty int) error {
	var err error

	eventGuest, err := r.FindOneEventGuestDataById(ctx, obj.IDGuest)
	if err != nil {
		return err
	}
	if !eventGuest.IsOverQtyAllowed {
		totalTestimony, err := r.getTotalGuestCount(ctx, obj.IDGuest)
		if err != nil {
			return err
		}
		finalTestimony := testimonyQty + int(totalTestimony)
		if finalTestimony > eventGuest.Qty {
			return CustomError{message: ("Guest Testimony Qty already Reach quote, qty available is " + strconv.Itoa(eventGuest.Qty-finalTestimony))}
		}
	}

	var docs []interface{}

	for i := 0; i < testimonyQty; i++ {
		randomId := util.GenerateTimeUuidWithoutDash()
		id, err := entity.NewEventTestimonyDataID(randomId)
		if err != nil {
			return err
		}
		newData := obj
		newData.ID = id
		docs = append(docs, newData)
	}

	eventTestimonyCollection := &eventTestimonyCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionEventTestimony),
	}

	err = dbhelpers.WithTransaction(ctx, r.MongoWithTransactionImpl, func(dbCtx context.Context) error {
		// Insert the documents using InsertMany
		res, err := eventTestimonyCollection.InsertMany(dbCtx, docs)
		if err != nil {
			log.Error(ctx, "Error inserting documents: ", err)
			return err
		}

		log.Info(ctx, "Inserted documents: ", res.InsertedIDs)
		return nil
	})

	return err
}

func (r GatewayApiBaseApp) FindOneEventTestimonyDataById(ctx context.Context, id string) (*entity.EventTestimonyData, error) {
	var (
		resultEventTestimonyData *entity.EventTestimonyData
		err                      error
	)

	coll := r.getEventTestimonyCollection()
	resCol := coll.FindOne(ctx, bson.M{"id": id})
	err = resCol.Decode(&resultEventTestimonyData)
	if err != nil {
		log.Error(ctx, err.Error())

		// error return if data not found
		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("organizer wedding data not found")
		}

		return nil, err
	}
	return resultEventTestimonyData, err
}

func (r GatewayApiBaseApp) FindAllEventTestimonyData(ctx context.Context, req entity.BaseReqFind) ([]*entity.EventTestimonyData, int64, error) {
	var (
		err  error
		objs []*entity.EventTestimonyData
	)

	coll := r.getEventTestimonyCollection()

	findData, _ := req.Value.(entity.FindEventTestimonyData)
	criteria := getFilterEventTestimonyKeyword(findData, true)

	findOpts := gateway.BaseReqFindToOptOption(req)

	cursor, err := coll.Find(ctx, criteria, &findOpts)
	if err != nil {
		return nil, 0, err
	}

	if err := cursor.All(ctx, &objs); err != nil {
		return nil, 0, err
	}

	//counting
	count, err := coll.GetTotalEventTestimony(ctx, findData, true)

	return objs, count, err
}

func (r GatewayApiBaseApp) UpdateEventTestimonyData(ctx context.Context, EventTestimonyData entity.EditEventTestimonyData) (*entity.EventTestimonyData, error) {
	log.Info(ctx, "called")

	EventTestimonyData.UpdatedAt = time.Now().Local().UTC()

	editBson := util.StructToBSONM(EventTestimonyData)

	info, err := r.MongoWithTransactionImpl.UpdateByCustomId(ctx, r.database, entity.CollectionEventTestimony, EventTestimonyData.ID.String(), editBson)
	log.Info(ctx, "info >>> ", info)
	if err != nil {
		log.Info(ctx, "error >>> "+err.Error())
	}
	dataResult, ok := info.(entity.EventTestimonyData)
	if !ok {
		return nil, fmt.Errorf("organizer Transform")
	}

	return &dataResult, err
}

func (r GatewayApiBaseApp) DeleteEventTestimonyData(ctx context.Context, id string) (bool, error) {
	var (
		err error
	)

	coll := r.getEventTestimonyCollection()
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

func (r GatewayApiBaseApp) getTotalGuestCount(ctx context.Context, idGuest string) (int64, error) {
	coll := r.getEventTestimonyCollection()
	filter := bson.M{"id_guest": idGuest}
	return coll.CountDocuments(ctx, filter)
}
