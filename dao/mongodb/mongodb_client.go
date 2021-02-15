package mongodb

import (
	"context"
	"fmt"
	"time"

	"go_schedule/util/config"
	"go_schedule/util/data_schema"
	"go_schedule/util/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var db string

func InitMongodb() error {
	clientSocket := config.Viper.GetString("mongodb.conn")
	clientOptions := options.Client().ApplyURI(clientSocket)
	clientOptions.SetConnectTimeout(config.Viper.GetDuration("mongodb.conn_timeout") * time.Second)
	clientOptions.SetSocketTimeout(config.Viper.GetDuration("mongodb.timeout") * time.Second)
	var err error
	// 连接到MongoDB
	mongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Errorf("mongodb connect fail, err:%+v, clienSocket:%s", err.Error(), clientSocket)
		return err
	}

	// 检查连接
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Errorf("mongodb ping error, error:%+v", err.Error())
		return err
	}

	db = config.Viper.GetString("mongodb.database")
	return nil
}

// InsertOneTask 插入一条doc
func InsertOneTask(ctx context.Context, task data_schema.TaskInfo) (id primitive.ObjectID, err error) {
	collection := mongoClient.Database(db).Collection("task")
	if collection == nil {
		log.Errorf("db.runoob.task is nil")
		return primitive.ObjectID{}, fmt.Errorf("collection is nil")
	}
	var result *mongo.InsertOneResult
	var i int
	for i = 0; i < 3; i++ {
		if result, err = collection.InsertOne(ctx, task); err == nil {
			break
		}
	}

	if i >= 3 {
		log.Errorf("insert task into mongo fail, task:%+v, error:%+v", task, err)
		return primitive.ObjectID{}, err
	}
	id = result.InsertedID.(primitive.ObjectID)
	return id, nil
}

// UpdateOneTask 更新一个文档
func UpdateOneTask(ctx context.Context, filter bson.D, update bson.D) (result *mongo.UpdateResult, err error) {
	collection := mongoClient.Database(db).Collection("task")
	if collection == nil {
		log.Errorf("db.runoob.task is nil")
		return nil, fmt.Errorf("collection is nil")
	}

	i := 0
	for i = 0; i < 3; i++ {
		if result, err = collection.UpdateOne(ctx, filter, update); err == nil {
			break
		}
	}

	if i >= 3 {
		log.Errorf("insert task into mongo fail, filter:%+v, task:%+v, error:%+v", filter, update, err)
		return nil, err
	}
	return result, err
}

// DeleteTask 删除任务
func DeleteTask(ctx context.Context, filter bson.D) (result *mongo.DeleteResult, err error) {
	collection := mongoClient.Database(db).Collection("task")
	if collection == nil {
		log.Errorf("db.runoob.task is nil")
		return nil, fmt.Errorf("collection is nil")
	}
	for i := 0; i < 3; i++ {
		if result, err = collection.DeleteMany(ctx, filter); err == nil {
			break
		}
	}
	if err != nil {
		log.Errorf("delete task from mongodb fail, filter:%+v, error:%+v", filter, err)
	}
	return result, err
}

// SearchTask 查询task
func SearchTask(ctx context.Context, filter bson.D) (tasks []data_schema.TaskInfo, err error) {
	collection := mongoClient.Database(db).Collection("task")
	if collection == nil {
		log.Errorf("db.runoob.task is nil")
		return nil, fmt.Errorf("collection is nil")
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Errorf("mongodb find task fail, filter:%+v, error:%+v", filter, err)
		return nil, err
	}
	tasks = make([]data_schema.TaskInfo, cursor.RemainingBatchLength())
	if err = cursor.All(ctx, &tasks); err != nil {
		log.Errorf("mongodb all function fail, filter:%+v, error:%+v", filter, err)
		return nil, err
	}
	return tasks, nil
}

// InsertOneSchedule 插入一条调度信息
func InsertOneSchedule(ctx context.Context, schedule data_schema.ScheduleHistory) (id string, err error) {
	collection := mongoClient.Database(db).Collection("schedule")
	if collection == nil {
		log.Errorf("db.runoob.schedule is nil")
		return "", fmt.Errorf("collection is nil")
	}
	var result *mongo.InsertOneResult
	var i int
	for i = 0; i < 3; i++ {
		if result, err = collection.InsertOne(ctx, schedule); err == nil {
			break
		}
	}

	if i >= 3 {
		log.Errorf("insert schedule into mongo fail, schedule:%+v, error:%+v", schedule, err)
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
