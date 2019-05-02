package service

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	bolt "go.etcd.io/bbolt"
)

type DAO struct {
	db *bolt.DB
}

func New() *DAO {
	db, _ := getDB()
	return &DAO{db}
}

type Row struct {
	Id    uint64 `json:"id"`
	Thing string `json:"thing"`
	Stat  int    `json:"stat"`
}

func (dao *DAO) getAll(cls bool) ([]*Row, error) {

	rows := []*Row{}
	err := dao.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("taskBucket"))

		err := b.ForEach(func(k, v []byte) error {
			row := &Row{}
			err := json.Unmarshal(v, &row)
			if err != nil {
				return err
			}

			rows = append(rows, row)
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	if cls {
		closeDB(dao.db)
	}
	return rows, nil
}

func (dao *DAO) getByStat(s int, cls bool) ([]*Row, error) {
	rows, err := dao.getAll(cls)
	newRows := []*Row{}
	if err != nil {
		return nil, err
	}

	for _, v := range rows {
		if v.Stat == s {
			newRows = append(newRows, v)
		}
	}

	return newRows, nil
}

func (dao *DAO) GetList() ([]*Row, error) {
	return dao.getByStat(0, true)
}

func (dao *DAO) GetCompleted() ([]*Row, error) {
	return dao.getByStat(1, true)
}

func (dao *DAO) Add(thing string) error {
	defer closeDB(dao.db)

	return dao.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("taskBucket"))
		id, _ := b.NextSequence()
		row := &Row{id, thing, 0}

		buf, err := json.Marshal(row)
		if err != nil {
			return err
		}

		return b.Put(itob(id), buf)
	})
}

func (dao *DAO) Rm(num int) (*Row, error) {
	defer closeDB(dao.db)

	row, err := dao.getRowByNum(num)
	if err != nil {
		return nil, err
	}

	return row, dao.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("taskBucket"))

		return b.Delete(itob(row.Id))
	})
}

func (dao *DAO) Do(num int) (*Row, error) {
	defer closeDB(dao.db)

	row, err := dao.getRowByNum(num)
	if err != nil {
		return nil, err
	}

	row.Stat = 1
	buf, err := json.Marshal(row)
	if err != nil {
		return nil, err
	}

	return row, dao.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("taskBucket"))

		return b.Put(itob(row.Id), buf)
	})
}

func (dao *DAO) getRowByNum(num int) (*Row, error) {
	rows, err := dao.getByStat(0, false)
	if err != nil {
		return nil, err
	}

	if num > len(rows) {
		return nil, errors.New("out of range")
	}

	num -= 1
	row := rows[num]
	return row, nil
}

func getDB() (*bolt.DB, error) {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	// Start a writable transaction.
	tx, err := db.Begin(true)
	if err != nil {
		return nil, err
	}

	// Use the transaction...
	_, err = tx.CreateBucketIfNotExists([]byte("taskBucket"))
	if err != nil {
		return nil, err
	}

	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return db, nil
}

func closeDB(db *bolt.DB) error {
	return db.Close()
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
