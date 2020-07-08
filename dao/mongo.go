package dao

import (
	"fmt"
	"reflect"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// 创建数据
func Create(table string, item interface{}) error {

	// 更新createdAt和updatedAt
	if item == nil {
		return nil
	}
	date := time.Now()
	if reflect.TypeOf(item).Kind() != reflect.Ptr {
		//fmt.Println("不是指针, 无法设置 createdAt 和 updatedAt ")
	} else {
		reflect.ValueOf(item).Elem().FieldByName("CreatedAt").Set(reflect.ValueOf(date))
		reflect.ValueOf(item).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(date))
	}
	s, c := GetSession(table)
	defer s.Close()

	err := c.Insert(item)
	return err

}

// 获取一条数据
func FindOne(table string, object, where interface{}) error {
	s, c := GetSession(table)
	defer s.Close()
	q := c.Find(where)
	err := q.One(object)
	if err == mgo.ErrNotFound {
		return nil
	}
	return err
}

// 获取最后一条数据
func FindOneBySort(table string, object, where interface{}, sortField string) error {
	s, c := GetSession(table)
	defer s.Close()
	q := c.Find(where).Sort(fmt.Sprintf("-%s", sortField)).Limit(1)
	err := q.One(object)
	if err == mgo.ErrNotFound {
		return nil
	}
	return err
}

// 获取全部数据
func Find(table string, objects, where interface{}) error {
	s, c := GetSession(table)
	defer s.Close()
	q := c.Find(where)
	err := q.All(objects)
	return err
}

func FindAndSort(table string, objects, where interface{}, sortField string, desc bool) error {
	s, c := GetSession(table)
	defer s.Close()
	sortStr := ""
	if desc {
		sortStr = fmt.Sprintf("-%s", sortField)
	} else {
		sortStr = fmt.Sprintf("%s", sortField)
	}
	q := c.Find(where).Sort(sortStr)
	err := q.All(objects)
	return err
}

// 数据统计
func Count(table string, where interface{}) (int, error) {
	s, c := GetSession(table)
	defer s.Close()
	count, err := c.Find(where).Count()
	return count, err
}

// 更新数据
func Update(table string, where, updatedItem interface{}) error {

	// 更新updatedAt
	date := time.Now()

	tp := reflect.TypeOf(updatedItem).String()

	if reflect.TypeOf(updatedItem).Kind() != reflect.Ptr {
		fmt.Println("不是指针, 无法更新updatedAt ")
	} else if tp != "*bson.M" {
		reflect.ValueOf(updatedItem).Elem().FieldByName("UpdatedAt").Set(reflect.ValueOf(date))
	}

	if tp == "bson.M" && updatedItem.(bson.M)["$set"] != nil {
		u := updatedItem.(bson.M)["$set"]
		u.(bson.M)["updatedAt"] = date
	}

	s, c := GetSession(table)
	defer s.Close()
	err := c.Update(where, updatedItem)
	return err

}

// 更新所有数据
func UpdateAll(table string, where, updatedItem interface{}) error {
	s, c := GetSession(table)
	defer s.Close()
	_, err := c.UpdateAll(where, bson.M{"$set": updatedItem})
	return err
}

// 删除数据
func Delete(table string, where interface{}, RemoveAll bool) error {
	s, c := GetSession(table)
	defer s.Close()
	var err error
	if RemoveAll {
		_, err = c.RemoveAll(where)
		return err
	}
	err = c.Remove(where)
	return err
}

// 聚合查询一条数据
func AggGet(table string, object interface{}, queries ...bson.M) error {
	s, c := GetSession(table)
	defer s.Close()
	err := c.Pipe(queries).One(object)
	return err
}

// 聚合查询全部数据
func AggGetAll(table string, objects interface{}, queries ...bson.M) error {
	s, c := GetSession(table)
	defer s.Close()
	err := c.Pipe(queries).All(objects)
	return err
}

// 分页查询
func FindByPage(table string, objects interface{}, where bson.M, page, limit int) error {
	var skip int
	if page == 1 {
		skip = 0
	} else {
		skip = (page - 1) * limit
	}
	s, c := GetSession(table)
	defer s.Close()
	err := c.Find(where).Limit(limit).Skip(skip).All(objects)
	return err
}

// 批量插入
func BulkInsert(table string, objects []interface{}) error {
	s, c := GetSession(table)
	defer s.Close()
	bulk := c.Bulk()
	bulk.Insert(objects...)
	_, err := bulk.Run()
	return err
}

// 检查是否存在
func IsExists(table string, where interface{}) (bool, error) {
	s, c := GetSession(table)
	defer s.Close()
	object := &bson.M{}
	q := c.Find(where)
	err := q.One(object)
	if err == mgo.ErrNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	if len((*object)) > 0 {
		return true, nil
	}
	return false, nil
}

// distinct
func Distinct(table, key string, objects interface{}, where interface{}) error {

	s, c := GetSession(table)
	defer s.Close()

	err := c.Find(where).Distinct(key, objects)

	return err

}
