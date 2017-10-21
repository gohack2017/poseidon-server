package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type _User struct{}

var (
	User *_User

	userCollection = "poseidon_user"
	userIndexes    = []mgo.Index{
		{
			Key:    []string{"email"},
			Unique: true,
		},
	}
)

type UserModel struct {
	ID        bson.ObjectId `bson:"_id"`
	Email     string        `bson:"email"`
	Salt      string        `bson:"salt"`
	Password  string        `bson:"password"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`

	isNewRecord bool `bson:"-"`
}

func NewUserModel(email, password string) *UserModel {
	return &UserModel{
		ID:          bson.NewObjectId(),
		Email:       email,
		Password:    password,
		isNewRecord: true,
	}
}

func (user *UserModel) Save() (err error) {
	if !user.ID.Valid() || user.Email == "" || user.Password == "" {
		return ErrInvalidArgs
	}

	User.Query(func(c *mgo.Collection) {
		t := time.Now()
		if user.IsNewRecord() {
			//encrypt password for new user
			user.Salt = uuid.NewV4().String()
			user.Password = User.MakeEncryptPassword(user.Salt, user.Password)
			user.CreatedAt = t
			user.UpdatedAt = t

			if err = c.Insert(user); err == nil {
				user.isNewRecord = false
			}
		} else {
			settings := bson.M{
				"email":      user.Email,
				"password":   user.Password,
				"updated_at": t,
			}

			err = c.UpdateId(user.ID, bson.M{
				"$set": settings,
			})
		}
	})
	return
}

func (_ *_User) Find(id string) (user *UserModel, err error) {
	if !bson.IsObjectIdHex(id) {
		err = ErrInvalidID
		return
	}

	User.Query(func(c *mgo.Collection) {
		query := bson.M{
			"_id": bson.ObjectIdHex(id),
		}

		err = c.Find(query).One(&user)
	})

	return
}

func (_ *_User) FindByEmail(email string) (user *UserModel, err error) {
	User.Query(func(c *mgo.Collection) {
		query := bson.M{
			"email": email,
		}

		err = c.Find(query).One(&user)
	})

	return
}

func (_ *_User) All(limit int, marker string) (result []*UserModel, err error) {
	limit = Helper.ModifyLimit(limit)
	User.Query(func(c *mgo.Collection) {
		query := bson.M{}
		query["_id"] = bson.M{
			"$gte": bson.ObjectIdHex(marker),
		}

		err = c.Find(query).Limit(limit).All(&result)
	})

	return
}

func (user *UserModel) IsValidPassword(password string) bool {
	// default to false for new record
	if user.IsNewRecord() {
		return false
	}

	return 0 == strings.Compare(user.Password, User.MakeEncryptPassword(user.Salt, password))
}

func (_ *_User) MakeEncryptPassword(cipher, password string) string {
	hash := hmac.New(sha256.New, []byte(cipher))
	hash.Write([]byte(password))

	encpass := hash.Sum(nil)
	return hex.EncodeToString(encpass)
}

func (user *UserModel) IsNewRecord() bool {
	return user.isNewRecord
}

func (_ *_User) Query(query func(c *mgo.Collection)) {
	Model().Query(userCollection, userIndexes, query)
}
