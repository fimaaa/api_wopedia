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

type CreateOrganizerWeddingDataRepo interface {
	CreateOrganizerWeddingData(ctx context.Context, obj entity.OrganizerWeddingData) error
	FindOneOrganizerWeddingDataById(ctx context.Context, id string) (*entity.OrganizerWeddingData, error)
	FindAllOrganizerWeddingData(ctx context.Context, req entity.BaseReqFind) ([]*entity.OrganizerWeddingData, int64, error)
	UpdateOrganizerWeddingData(ctx context.Context, organizerWeddingData entity.EditOrganizerWeddingData) (*entity.OrganizerWeddingData, error)
	DeleteOrganizerWeddingData(ctx context.Context, id string) (bool, error)
}

type organizerWeddingCollection struct {
	*mongo.Collection
}

func (r GatewayApiBaseApp) getOrganizerWeddingCollection() organizerWeddingCollection {
	return organizerWeddingCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionOrganizerWedding),
	}
}

func getFilterOrganizerWeddingKeyword(
	obj entity.FindOrganizerWeddingData,
	onlySimiliar bool,
) primitive.M {
	//====== execute query using transaction ======
	//count the existing users
	keywordFilter := make([]bson.M, 0)

	if obj.Name != nil && *obj.Name != "" {
		keyword := bson.M{"name": obj.Name}
		if onlySimiliar {
			keyword = bson.M{"name": primitive.Regex{Pattern: string(*obj.Name), Options: "i"}}
		}
		keywordFilter = append(keywordFilter, keyword)
	}
	if obj.Phone != nil && *obj.Phone != "" {
		keyword := bson.M{"phone": obj.Phone}
		keywordFilter = append(keywordFilter, keyword)
	}
	if obj.PIC != nil && *obj.PIC != "" {
		keyword := bson.M{"pic": obj.PIC}
		keywordFilter = append(keywordFilter, keyword)
	}
	if obj.MinRangePrice != nil {
		keyword := bson.M{"min_range_price": bson.M{"$gte": obj.MinRangePrice}}
		keywordFilter = append(keywordFilter, keyword)
	}

	if obj.MaxRangePrice != nil {
		keyword := bson.M{"min_range_price": bson.M{"$lte": obj.MaxRangePrice}}
		keywordFilter = append(keywordFilter, keyword)
	}
	if obj.CreatedAtFrom != nil && !obj.CreatedAtFrom.IsZero() {
		createdAtTo := time.Now()
		if obj.CreatedAtTo != nil && !obj.CreatedAtTo.IsZero() {
			createdAtTo = *obj.CreatedAtTo
		}
		keyword := bson.M{
			"created_at": bson.M{
				"$gte": obj.CreatedAtFrom, // Greater than or equal to date_from
				"$lte": createdAtTo,       // Less than or equal to date_to
			},
		}
		// Append the date filter to the keywordFilter slice
		keywordFilter = append(keywordFilter, keyword)
	}
	if obj.UpdatedAtFrom != nil && !obj.UpdatedAtFrom.IsZero() {
		updatedAtTo := time.Now()
		if obj.UpdatedAtTo != nil && !obj.UpdatedAtTo.IsZero() {
			updatedAtTo = *obj.UpdatedAtTo
		}
		keyword := bson.M{
			"updated_at": bson.M{
				"$gte": obj.CreatedAtFrom, // Greater than or equal to date_from
				"$lte": updatedAtTo,       // Less than or equal to date_to
			},
		}

		// Append the date filter to the keywordFilter slice
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

func (coll organizerWeddingCollection) GetTotalOrganizerWedding(ctx context.Context, obj entity.FindOrganizerWeddingData, onlySimiliar bool) (int64, error) {
	criteria := getFilterOrganizerWeddingKeyword(obj, onlySimiliar)

	countOpts := options.CountOptions{}
	return coll.CountDocuments(ctx, criteria, &countOpts)
}

func (r GatewayApiBaseApp) CreateOrganizerWeddingData(ctx context.Context, obj entity.OrganizerWeddingData) error {
	var err error

	randomId := util.GenerateTimeUuidWithoutDash()
	id, err := entity.NewOrganizerWeddingDataID(randomId)
	if err != nil {
		return err
	}
	obj.ID = id

	organizerWeddingCollection := &organizerWeddingCollection{
		r.MongoWithTransactionImpl.MongoClient.Database(r.database).Collection(entity.CollectionOrganizerWedding),
	}

	err = dbhelpers.WithTransaction(ctx, r.MongoWithTransactionImpl, func(dbCtx context.Context) error {
		info, err := organizerWeddingCollection.InsertOne(ctx, obj)
		log.Info(ctx, "info >>> ", info)
		return err
	})

	return err
}

func (r GatewayApiBaseApp) FindOneOrganizerWeddingDataById(ctx context.Context, id string) (*entity.OrganizerWeddingData, error) {
	var (
		resultOrganizerWeddingData *entity.OrganizerWeddingData
		err                        error
	)

	coll := r.getOrganizerWeddingCollection()
	resCol := coll.FindOne(ctx, bson.M{"id": id})
	err = resCol.Decode(&resultOrganizerWeddingData)
	if err != nil {
		log.Error(ctx, err.Error())

		// error return if data not found
		if err.Error() == "mongo: no documents in result" {
			return nil, fmt.Errorf("organizer wedding data not found")
		}

		return nil, err
	}
	return resultOrganizerWeddingData, err
}

func (r GatewayApiBaseApp) FindAllOrganizerWeddingData(ctx context.Context, req entity.BaseReqFind) ([]*entity.OrganizerWeddingData, int64, error) {
	var (
		err  error
		objs []*entity.OrganizerWeddingData
	)

	coll := r.getOrganizerWeddingCollection()

	findData, _ := req.Value.(entity.FindOrganizerWeddingData)
	criteria := getFilterOrganizerWeddingKeyword(findData, true)

	findOpts := gateway.BaseReqFindToOptOption(req)

	cursor, err := coll.Find(ctx, criteria, &findOpts)
	if err != nil {
		return nil, 0, err
	}

	if err := cursor.All(ctx, &objs); err != nil {
		return nil, 0, err
	}

	//counting
	count, err := coll.GetTotalOrganizerWedding(ctx, findData, true)

	return objs, count, err
}

func (r GatewayApiBaseApp) UpdateOrganizerWeddingData(ctx context.Context, OrganizerWeddingData entity.EditOrganizerWeddingData) (*entity.OrganizerWeddingData, error) {
	log.Info(ctx, "called")

	OrganizerWeddingData.UpdatedAt = time.Now().Local().UTC()

	editBson := util.StructToBSONM(OrganizerWeddingData)

	info, err := r.MongoWithTransactionImpl.UpdateByCustomId(ctx, r.database, entity.CollectionOrganizerWedding, OrganizerWeddingData.ID.String(), editBson)
	log.Info(ctx, "info >>> ", info)
	if err != nil {
		log.Info(ctx, "error >>> "+err.Error())
	}
	dataResult, ok := info.(entity.OrganizerWeddingData)
	if !ok {
		return nil, fmt.Errorf("organizer Transform")
	}

	return &dataResult, err
}

func (r GatewayApiBaseApp) DeleteOrganizerWeddingData(ctx context.Context, id string) (bool, error) {
	var (
		err error
	)

	coll := r.getOrganizerWeddingCollection()
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
