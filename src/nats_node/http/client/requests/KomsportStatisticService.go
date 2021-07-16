package request

import (
	"encoding/json"
	"fmt"
	"komsport/http/model"
	"komsport/utils/logger"
	"strconv"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	LvlDB        *leveldb.DB
	ChConnection *sqlx.DB
	err          error
	Mutex        sync.Mutex
	Condition    *sync.Cond
)

func init() {
	fmt.Println("Init Worker...")

	LvlDB, err = leveldb.OpenFile("/tmp/komsport.db", nil)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	ChConnection, err = sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?compress=truedebug=true")
	if err != nil {
		logger.Logger.PrintError(err)
	}
	if err := ChConnection.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Println("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			logger.Logger.PrintError(err)
		}
		return
	}

	_, err := ChConnection.Exec(`
			CREATE TABLE IF NOT EXISTS clickhouse.statistic (
				  Date Date DEFAULT toDate(DT),
				  DT DateTime('Europe/Moscow'),
				  EventDate Date DEFAULT toDate(EventTime),
				  EventTime DateTime('Europe/Moscow'),
					Firstname String,
					Lastname String,
					Middlename String,
					ChangeDate Date,
					ChangeReason String,
					BirthDate Date,
					BirthPlace String,
					Gender String,
					DisableCategory String,
					DisableFrom Date,
					DisableTill Date,
					SportRank String,
					SportType String,
					SportDiscipline String,
					SportOrganization String,
					QualificationCategory String,
					Disqualification String,
					SportAchivements String,
					SportActivities String,
					SportObjects String
				)
				ENGINE = MergeTree
				PARTITION BY toYYYYMM(Date)
				ORDER BY (Date, EventDate)
				SETTINGS index_granularity = 8192
			`)

	if err != nil {
		logger.Logger.PrintError(err)
	}

	go InsertIntoClickHouse()
}

func SendKomsportStatistic() {
	var b []byte //getFromNats
	var m *model.Model = &model.Model{}
	err := json.Unmarshal(b, m)
	if err != nil {
		logger.Logger.PrintError(err)
	}

	batch := new(leveldb.Batch)
	timeInBytes := []byte(strconv.Itoa(time.Now().Nanosecond()))
	batch.Put(timeInBytes, b)

	err = LvlDB.Write(batch, nil)
	if err != nil {
		logger.Logger.PrintError(err)
	}
}

func InsertIntoClickHouse() {
	for range time.Tick(time.Second * 30) {
		entities := make([]model.Model, 0)
		batch := &leveldb.Batch{}

		iter := LvlDB.NewIterator(nil, nil)
		for iter.Next() {
			var m model.Model
			err = json.Unmarshal(iter.Value(), &m)
			if err != nil {
				logger.Logger.PrintError(err)
			} else {
				entities = append(entities, m)
				batch.Delete(iter.Key())
			}
		}

		err = statisticInsert(entities)
		if err == nil {
			err = LvlDB.Write(batch, nil)
			if err != nil {
				logger.Logger.PrintError(err)
			}
		}
	}
}

func statisticInsert(entities []model.Model) error {
	tx, err := ChConnection.Begin()
	if nil != err {
		return err
	}

	stmt, err := tx.Prepare(insertStatistic)
	if nil != err {
		return err
	}
	defer stmt.Close()

	for _, entry := range entities {
		_, err := stmt.Exec(
			time.Now(),
			time.Now(),
			time.Now(),
			time.Now(),
			entry.Firstname,
			entry.Lastname,
			entry.Middlename,
			entry.ChangeDate,
			entry.ChangeReason,
			entry.BirtDate,
			entry.BirthPlace,
			entry.Gender,
			entry.DisableCategory,
			entry.DisableFrom,
			entry.DisableTill,
			entry.SportRank,
			entry.SportType,
			entry.SportDiscipline,
			entry.SportOrganization,
			entry.QualificationCategory,
			entry.Disqualifications,
			entry.Achivements,
			entry.Activities,
			entry.SportObjects,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

const (
	insertStatistic = `
				INSERT INTO
								clickhouse.statistic (
														Date,
														DT,
														EventDate,
														EventTime,
														Firstname,
														Lastname,
														Middlename,
														ChangeDate,
														ChangeReason,
														BirthDate,
														BirthPlace,
														Gender,
														DisableCategory,
														DisableFrom,
														DisableTill,
														SportRank,
														SportType,
														SportDiscipline,
														SportOrganization,
														QualificationCategory,
														Disqualification,
														SportAchivements,
														SportActivities,
														SportObjects
							 )
							 VALUES
											(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
			 `
)
