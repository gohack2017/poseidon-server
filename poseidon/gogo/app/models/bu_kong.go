package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type _BuKong struct{}

var (
	BuKong *_BuKong

	bukongCollection = "poseidon_bukong"
	bukongIndexes    = []mgo.Index{}
)

type BukongModel struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	URI          string        `bson:"uri" json:"uri"`
	Name         string        `bson:"name" json:"name"`   //填报人姓名
	Phone        string        `bson:"phone" json:"phone"` //填报人电话
	MonitorClass string        `bson:"class" json:"class"`
	CreatedAt    time.Time     `bson:"created_at" json:"-"`
	UpdatedAt    time.Time     `bson:"updated_at" json:"-"`
	isNewRecord  bool          `bson:"-"`
}

func NewBukongModel(uri, name, phone, monitorClass string) *BukongModel {
	return &BukongModel{
		ID:           bson.NewObjectId(),
		Name:         name,
		Phone:        phone,
		URI:          uri,
		MonitorClass: monitorClass,

		isNewRecord: true,
	}
}

func (bukong *BukongModel) Save() (err error) {
	if !bukong.ID.Valid() || bukong.Name == "" || bukong.Phone == "" {
		return ErrInvalidArgs
	}

	BuKong.Query(func(c *mgo.Collection) {
		t := time.Now()
		if bukong.IsNewRecord() {
			bukong.CreatedAt = t
			bukong.UpdatedAt = t

			if err = c.Insert(bukong); err == nil {
				bukong.isNewRecord = false
			}
		} else {
			settings := bson.M{
				"name":  bukong.Name,
				"phone": bukong.Phone,
				"uri":   bukong.URI,
				"class": bukong.MonitorClass,
			}

			err = c.UpdateId(bukong.ID, bson.M{
				"$set": settings,
			})
		}
	})
	return
}

func (_ *_BuKong) Find(id string) (res *BukongModel, err error) {
	if !bson.IsObjectIdHex(id) {
		err = ErrInvalidID
		return
	}

	BuKong.Query(func(c *mgo.Collection) {
		query := bson.M{
			"_id": bson.ObjectId(id),
		}

		err = c.Find(query).One(&res)
	})

	return
}

func (_ *_BuKong) Delete(id string) (err error) {
	// if !bson.IsObjectIdHex(id) {
	// 	err = ErrInvalidID
	// 	return
	// }

	BuKong.Query(func(c *mgo.Collection) {
		err = c.RemoveId(bson.ObjectIdHex(id))
	})

	return
}

func (_ *_BuKong) All(limit int, marker string) (result []*BukongModel, err error) {
	limit = Helper.ModifyLimit(limit)
	BuKong.Query(func(c *mgo.Collection) {
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

func (bukong *BukongModel) IsNewRecord() bool {
	return bukong.isNewRecord
}

func (_ *_BuKong) Query(query func(c *mgo.Collection)) {
	Model().Query(bukongCollection, bukongIndexes, query)
}
