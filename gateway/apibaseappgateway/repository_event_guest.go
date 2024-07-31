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

type CreateEventGuestDataRepo interface {
	CreateEventGuestData(ctx context.Context, obj entity.EventGuestData) error
	FindOneEventGuestDataById(ctx context.Context, id string) (*entity.EventGuestData, error)
	FindAllEventGuestData(ctx context.Context, req entity.BaseReqFind) ([]*entity.EventGuestData, int64, error)
	UpdateEventGuestData(ctx context.Context, IDEvent string, editEventGuestData entity.EditEventGuestData) (*entity.EventGuestData, error)
	DeleteEventGuestData(ctx context.Context, id string) (bool, error)
}

type eventGuestCollection struct {
	*mongo.Collection
}

func (r GatewayApiBaseApp) getEventGuestCollection() eventGuestCollection {
	return eventGuestCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionEventGuest),
	}
}

func getFilterEventGuestKeyword(obj entity.FindEventGuestData) primitive.M {
	//====== execute query using transaction ======
	//count the existing users
	keywordFilter := make([]bson.M, 0)

	switch {
	case obj.MinQty != nil && obj.MaxQty != nil:
		{
			keyword := bson.M{
				"total_invited": bson.M{
					"$gte": obj.MinQty,
					"$lte": obj.MaxQty,
				},
			}
			// Append the date filter to the keywordFilter slice
			keywordFilter = append(keywordFilter, keyword)
		}
	case obj.MaxQty != nil:
		{
			keyword := bson.M{
				"total_invited": bson.M{
					"$lte": obj.MaxQty,
				},
			}
			// Append the date filter to the keywordFilter slice
			keywordFilter = append(keywordFilter, keyword)
		}
	case obj.MinQty != nil:
		{
			keyword := bson.M{
				"total_invited": bson.M{
					"$gte": obj.MinQty,
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

func (coll eventGuestCollection) GetTotalEventGuest(ctx context.Context, obj entity.FindEventGuestData, onlySimiliar bool) (int64, error) {
	criteria := getFilterEventGuestKeyword(obj)

	countOpts := options.CountOptions{}
	return coll.CountDocuments(ctx, criteria, &countOpts)
}

func (r GatewayApiBaseApp) CreateEventGuestData(ctx context.Context, obj entity.EventGuestData) error {
	var err error

	totalInvited, err := r.GetTotalInvitedEventWedding(ctx, obj.IDEvent)
	if err != nil {
		return err
	}
	totalAlreadyInvited, err := r.getTotalQtyAlreadyInvited(ctx, obj.IDEvent, nil)
	if err != nil {
		return err
	}

	if totalAlreadyInvited+obj.Qty > totalInvited {
		return CustomError{message: ("Guest Qty already Reach quote, qty available is " + strconv.Itoa(totalInvited-totalAlreadyInvited))}
	}

	randomId := util.GenerateTimeUuidWithoutDash()
	id, err := entity.NewEventGuestDataID(randomId)
	if err != nil {
		return err
	}
	obj.ID = id

	coll := r.getEventGuestCollection()
	err = dbhelpers.WithTransaction(ctx, r.MongoWithTransactionImpl, func(dbCtx context.Context) error {
		info, err := coll.InsertOne(ctx, obj)
		log.Info(ctx, "info >>> ", info)
		return err
	})

	return err
}

func (r GatewayApiBaseApp) FindOneEventGuestDataById(ctx context.Context, id string) (*entity.EventGuestData, error) {
	var (
		resultEventGuestData *entity.EventGuestData
		err                  error
	)

	coll := r.getEventGuestCollection()
	resCol := coll.FindOne(ctx, bson.M{"id": id})
	err = resCol.Decode(&resultEventGuestData)
	if err != nil {
		log.Error(ctx, err.Error())

		// error return if data not found
		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("organizer wedding data not found")
		}

		return nil, err
	}
	return resultEventGuestData, err
}

func (r GatewayApiBaseApp) FindAllEventGuestData(ctx context.Context, req entity.BaseReqFind) ([]*entity.EventGuestData, int64, error) {
	var (
		err  error
		objs []*entity.EventGuestData
	)

	coll := r.getEventGuestCollection()

	findData, _ := req.Value.(entity.FindEventGuestData)
	criteria := getFilterEventGuestKeyword(findData)

	findOpts := gateway.BaseReqFindToOptOption(req)

	cursor, err := coll.Find(ctx, criteria, &findOpts)
	if err != nil {
		return nil, 0, err
	}

	if err := cursor.All(ctx, &objs); err != nil {
		return nil, 0, err
	}

	//counting
	count, err := coll.GetTotalEventGuest(ctx, findData, true)

	return objs, count, err
}

func (r GatewayApiBaseApp) UpdateEventGuestData(ctx context.Context, IDEvent string, editEventGuestData entity.EditEventGuestData) (*entity.EventGuestData, error) {
	log.Info(ctx, "called")

	if editEventGuestData.Qty != nil {
		totalInvited, err := r.GetTotalInvitedEventWedding(ctx, IDEvent)
		if err != nil {
			return nil, err
		}
		totalAlreadyInvited, err := r.getTotalQtyAlreadyInvited(ctx, IDEvent, (*string)(&editEventGuestData.ID))
		if err != nil {
			return nil, err
		}

		if totalAlreadyInvited+*editEventGuestData.Qty > totalInvited {
			return nil, CustomError{message: ("Guest Qty already Reach quote, qty available is " + strconv.Itoa(totalInvited-totalAlreadyInvited))}
		}

	}

	editEventGuestData.UpdatedAt = time.Now().Local().UTC()

	editBson := util.StructToBSONM(editEventGuestData)

	info, err := r.MongoWithTransactionImpl.UpdateByCustomId(ctx, r.database, entity.CollectionEventGuest, editEventGuestData.ID.String(), editBson)
	log.Info(ctx, "info >>> ", info)
	if err != nil {
		log.Info(ctx, "error >>> "+err.Error())
	}
	dataResult, ok := info.(entity.EventGuestData)
	if !ok {
		return nil, fmt.Errorf("organizer Transform")
	}

	return &dataResult, err
}

func (r GatewayApiBaseApp) DeleteEventGuestData(ctx context.Context, id string) (bool, error) {
	var (
		err error
	)

	coll := r.getEventGuestCollection()
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

func (r GatewayApiBaseApp) GetTotalInvitedEventGuest(ctx context.Context, id string) (int, error) {
	coll := r.getEventGuestCollection()
	projection := bson.M{"qty": 1, "_id": 0} // Projection to include only total_invited field
	resCol := coll.FindOne(ctx, bson.M{"id": id}, options.FindOne().SetProjection(projection))

	// Decode the result
	var result struct {
		TotalInvited *int `bson:"qty"`
	}
	if err := resCol.Decode(&result); err != nil {
		return -1, err
	}
	// Access the totalInvited field
	totalInvited := result.TotalInvited
	return *totalInvited, nil
}

func (r GatewayApiBaseApp) getTotalQtyAlreadyInvited(ctx context.Context, eventID string, excludeID *string) (int, error) {
	coll := r.getEventGuestCollection()

	// Create the match stage
	matchConditions := bson.D{{Key: "id_event", Value: eventID}}
	if excludeID != nil && *excludeID != "" {
		matchConditions = append(matchConditions, bson.E{Key: "_id", Value: bson.D{{Key: "$ne", Value: *excludeID}}})
	}

	matchStage := bson.D{{Key: "$match", Value: matchConditions}}
	groupStage := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: nil},
		{Key: "totalQty", Value: bson.D{{Key: "$sum", Value: "$qty"}}},
	}}}

	pipeline := mongo.Pipeline{matchStage, groupStage}
	cursor, err := coll.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result struct {
		TotalQty int `bson:"totalQty"`
	}
	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, err
		}
	}

	return result.TotalQty, nil
}
