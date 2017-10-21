package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type _Alert struct{}

var (
	Alert *_Alert

	alertCollection = "poseidon_alert"
	alertIndexes    = []mgo.Index{}
)

type AlertModel struct {
	ID            bson.ObjectId `bson:"_id"`
	Address       string        `bson:"address"`
	ScenePhotoUri string        `bson:"scene_photo"`
	PhotoUri      string        `bson:"photo"`
	MonitorClass  string        `bson:"class"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`

	isNewRecord bool `bson:"-"`
}

func NewAlertModel(addr, sphoto, photo, mclass string) *AlertModel {
	return &AlertModel{
		ID:            bson.NewObjectId(),
		Address:       addr,
		ScenePhotoUri: sphoto,
		PhotoUri:      photo,
		MonitorClass:  mclass,
		isNewRecord:   true,
	}
}

func (alert *AlertModel) Save() (err error) {
	if !alert.ID.Valid() {
		return ErrInvalidArgs
	}

	Alert.Query(func(c *mgo.Collection) {
		t := time.Now()
		if alert.IsNewRecord() {
			alert.CreatedAt = t
			alert.UpdatedAt = t

			if err = c.Insert(alert); err == nil {
				alert.isNewRecord = false
			}
		} else {
			settings := bson.M{
				"updated_at": t,
			}

			err = c.UpdateId(alert.ID, bson.M{
				"$set": settings,
			})
		}
	})
	return
}

func (_ *_Alert) All(limit int, marker string) (result []*AlertModel, err error) {
	limit = Helper.ModifyLimit(limit)
	Alert.Query(func(c *mgo.Collection) {
		query := bson.M{}
		if bson.IsObjectIdHex(marker) {
			query["_id"] = bson.M{
				"$gte": bson.ObjectIdHex(marker),
			}
		}

		err = c.Find(query).Limit(limit).All(&result)
	})

	return
}

func (alert *AlertModel) IsNewRecord() bool {
	return alert.isNewRecord
}

func (_ *_Alert) Query(query func(c *mgo.Collection)) {
	Model().Query(alertCollection, alertIndexes, query)
}
