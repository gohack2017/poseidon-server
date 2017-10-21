package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/dolab/session"
	"github.com/satori/go.uuid"
)

var (
	Session *_Session

	sessionCollection = "poseidon_session"
	sessionIndexes    = []mgo.Index{
		{
			Key:    []string{"sid"},
			Unique: true,
		},
	}
)

type SessionModel struct {
	ID      bson.ObjectId  `bson:"_id"`
	SID     string         `bson:"sid"`
	Salt    string         `bson:"salt"`
	Value   *session.Value `bson:"data"`
	Expires int            `bson:"expires"`
	Ctime   time.Time      `bson:"ctime"`
	Atime   time.Time      `bson:"atime"`

	isNewRecord bool `bson:"-"`
}

func NewSessionModel(sid string) *SessionModel {
	return &SessionModel{
		ID:          bson.NewObjectId(),
		SID:         sid,
		Salt:        uuid.NewV4().String(),
		Value:       session.NewValue(),
		isNewRecord: true,
	}
}

func (session *SessionModel) IsNewRecord() bool {
	return session.isNewRecord
}

func (session *SessionModel) Save() (err error) {
	if !session.ID.Valid() {
		return ErrInvalidID
	}

	Session.Query(func(c *mgo.Collection) {
		session.Atime = time.Now()

		if session.IsNewRecord() {
			session.Ctime = session.Atime

			err = c.Insert(session)
			if err == nil {
				session.isNewRecord = false
			}
		} else {
			update := bson.M{
				"$set": bson.M{
					"sid":   session.SID,
					"salt":  session.Salt,
					"data":  session.Value,
					"atime": session.Atime,
				},
			}

			err = c.UpdateId(session.ID, update)
		}
	})

	return
}

func (session *SessionModel) SessionID() string {
	return session.SID
}

func (session *SessionModel) SetValue(v *session.Value) (err error) {
	session.Value = v

	if v.IsChanged() {
		return session.Save()
	}

	return nil
}

func (session *SessionModel) GetValue() *session.Value {
	return session.Value
}

func (session *SessionModel) Touch() (err error) {
	if session.IsNewRecord() {
		return ErrNotPersisted
	}

	return session.Save()
}

type _Session struct{}

func (_ *_Session) FindBySID(sid string) (session *SessionModel, err error) {
	query := bson.M{
		"sid": sid,
	}

	Session.Query(func(c *mgo.Collection) {
		err = c.Find(query).One(&session)
	})

	return
}

// IMPL session.Provider.New
func (_ *_Session) New(sid string) (sto session.Storer, err error) {
	sess := NewSessionModel(sid)

	err = sess.Save()
	if err != nil {
		return
	}

	sto = sess

	return
}

// IMPLs session.Provider.Restore
func (_ *_Session) Restore(sid string) (sto session.Storer, err error) {
	sess, err := Session.FindBySID(sid)
	if err != nil {
		if err == mgo.ErrNotFound {
			err = session.ErrNotFound
		}

		return
	}

	sto = sess

	return
}

// IMPL session.Provider.Refresh
func (_ *_Session) Refresh(sid, newsid string) (sto session.Storer, err error) {
	var sess *SessionModel

	Session.Query(func(c *mgo.Collection) {
		sess, err = Session.FindBySID(sid)
		if err != nil {
			return
		}

		sess.SID = newsid
		sess.Salt = uuid.NewV4().String()
		sess.Atime = time.Now()

		err = sess.Save()
		if err == nil {
			sto = sess
		}
	})

	return
}

// IMPL session.Provider.Destroy
func (_ *_Session) Destroy(sid string) (err error) {
	query := bson.M{
		"sid": sid,
	}

	Session.Query(func(c *mgo.Collection) {
		err = c.Remove(query)
	})

	return
}

func (_ *_Session) Query(query func(c *mgo.Collection)) {
	Model().Query(sessionCollection, sessionIndexes, query)
}
