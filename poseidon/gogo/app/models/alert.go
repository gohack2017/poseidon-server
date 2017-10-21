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
	alertIndexes    = []mgo.Index{
		{
			Key:    []string{"bukong_id", "device_id"},
			Unique: false,
		},
	}
)

type AlertModel struct {
	ID       bson.ObjectId `bson:"_id" json:"-"`
	BukongID bson.ObjectId `bson:"bukong_id" json:"-"`
	DeviceId bson.ObjectId `bson:"device_id" json:"-"`

	Address       string  `bson:"address" json:"address"`
	ScenePhotoUri string  `bson:"scene_photo" json:"scene_photo"`
	PhotoUri      string  `bson:"photo" json:"photo"`
	MonitorClass  string  `bson:"class" json:"class"`
	Score         float64 `bson:"score" json:"score"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"-"`

	isNewRecord bool `bson:"-"`
}

func NewAlertModel(addr, sphoto, photo, mclass, bukongId, deviceId string) *AlertModel {
	return &AlertModel{
		ID:            bson.NewObjectId(),
		DeviceId:      bson.ObjectIdHex(deviceId),
		BukongID:      bson.ObjectIdHex(bukongId),
		Address:       addr,
		ScenePhotoUri: sphoto,
		PhotoUri:      photo,
		MonitorClass:  mclass,
		isNewRecord:   true,
	}
}

func (_ *_Alert) FindByBukongAndDevice(bukongId, deviceId string) (res *AlertModel, err error) {
	if !bson.IsObjectIdHex(bukongId) || !bson.IsObjectIdHex(deviceId) {
		err = ErrInvalidID
		return
	}

	Alert.Query(func(c *mgo.Collection) {
		query := bson.M{
			"bukong_id": bson.ObjectIdHex(bukongId),
			"device_id": bson.ObjectIdHex(deviceId),
		}

		err = c.Find(query).One(&res)
	})

	return
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
