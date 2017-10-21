package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type _Device struct{}

var (
	Device *_Device

	deviceCollection = "poseidon_device"
	deviceIndexes    = []mgo.Index{}
)

type DeviceModel struct {
	ID       bson.ObjectId `bson:"_id",json:"-"`
	Num      string        `bson:"num",json:"num"`
	Password string        `bson:"password",json:"password"`
	Address  string        `bson:"address",json:"address"`

	CreatedAt time.Time `bson:"created_at",json:"-"`
	UpdatedAt time.Time `bson:"updated_at",json:"-"`

	isNewRecord bool `bson:"-"`
}

func (device *DeviceModel) Save() (err error) {
	if !device.ID.Valid() || device.Password == "" || device.Address == "" {
		return ErrInvalidArgs
	}

	Device.Query(func(c *mgo.Collection) {
		t := time.Now()
		if device.IsNewRecord() {
			device.CreatedAt = t
			device.UpdatedAt = t

			if err = c.Insert(device); err == nil {
				device.isNewRecord = false
			}
		} else {
			settings := bson.M{
				"password":   device.Password,
				"updated_at": device.UpdatedAt,
			}

			err = c.UpdateId(device.ID, bson.M{
				"$set": settings,
			})
		}
	})
	return
}

func (_ *_Device) All(limit int, marker string) (result []*DeviceModel, err error) {
	limit = Helper.ModifyLimit(limit)
	Device.Query(func(c *mgo.Collection) {
		query := bson.M{}
		query["_id"] = bson.M{
			"$gte": bson.ObjectIdHex(marker),
		}

		err = c.Find(query).Limit(limit).All(&result)
	})

	return
}

// todo add num field
func NewDeviceModel(pass, addr string) *DeviceModel {
	return &DeviceModel{
		ID:          bson.NewObjectId(),
		Password:    pass,
		isNewRecord: true,
	}
}

func (device *DeviceModel) IsNewRecord() bool {
	return device.isNewRecord
}

func (_ *_Device) Query(query func(c *mgo.Collection)) {
	Model().Query(deviceCollection, deviceIndexes, query)
}
