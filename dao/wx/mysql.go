package wx

import (
	"commons/tools/utils"
	"errors"
	"strings"
)

func (d Dao) insertOne(beans ...interface{}) error {
	_, err := d.engine.Insert(beans)
	return err
}

func (d Dao) insertMany(beans ...interface{}) error {
	if beans == nil {
		return errors.New("model is nil")
	}
	session := d.engine.NewSession()
	defer session.Close()
	tx, err := session.BeginTrans()
	if err != nil {
		return err
	}
	for _, bean := range beans {
		_, err := tx.Session().Insert(bean)
		if err != nil {
			defer tx.RollbackTrans()
			return err
		}
	}
	err = tx.CommitTrans()
	return err
}

func (d Dao) exist(bean ...interface{}) bool {
	exist, err := d.engine.Exist(bean...)

	if exist && utils.CheckError(err, exist) {
		return true
	}
	return false
}
func (d Dao) delete(id int64, bean interface{}) (int64, error) {
	return d.engine.Id(id).Delete(bean)
}

//要支持模糊查询
func (d Dao) findOne(bean interface{}, table, orderBy string, query []string, values []interface{}, ps, pn int) (interface{}, int64) {
	var count int64
	var err error
	if len(query) == 0 {
		count, err = d.engine.Table(table).Limit(ps, ps*(pn-1)).OrderBy(orderBy).FindAndCount(bean)
		if utils.CheckError(err, count) {
			return bean, count
		}
		return nil, 0

	} else {
		count, err = d.engine.Table(table).Where(strings.Join(query, " "), values...).Limit(ps, ps*(pn-1)).OrderBy(orderBy).FindAndCount(bean)
		if utils.CheckError(err, count) {
			return bean, count
		}
		return nil, 0
	}
}

func (d Dao) updateMap(table string, m map[string]interface{}, cols, query []string, values []interface{}) (int64, error) {
	//affected, err := engine.Table(new(User)).Id(id).Update(map[string]interface{}{"age":0})
	return d.engine.Table(table).Where(strings.Join(query, ""), values...).Cols(cols...).Update(m)
}
func (d Dao) updateStruct(bean interface{}, cols, query []string, values []interface{}) (int64, error) {
	if query == nil {
		return d.engine.Cols(cols...).Update(bean)
	}
	return d.engine.Where(strings.Join(query, ""), values...).Cols(cols...).Update(bean)
}

//2表联查
func (d Dao) join2Table(bean interface{}, table, alias, cols, orderBy string, ps, pn int, query []string, values []interface{}, join [][3]interface{}) (interface{}, int64) {
	//mmm := [][3]interface{}{{
	//	"join", "type", "group.id = user.group_id"},
	//	{"join", []string{"group", "g"}, "group.id = user.group_id"},}
	var count int64
	var err error
	session := d.engine.Table(table).Alias(alias).Select(cols)
	for _, v := range join {
		session.Join(v[0].(string), v[1], v[2].(string))
	}
	count, err = session.Where(strings.Join(query, " "), values...).Limit(ps, ps*(pn-1)).OrderBy(orderBy).FindAndCount(bean)
	if utils.CheckError(err, count) {
		return bean, count
	}
	//count, err := d.Engine.Table(table).Join(joinOperator, tableName, condition, args).Where(strings.Join(query, " "), values...).FindAndCount(bean)
	return bean, count
}

//级联删除
func (d Dao) delete2Table(beans [][2]interface{}) error {
	/**
	数据结构如下
	[][2]{{id,bean},{id,bean}}
	*/
	session := d.engine.NewSession()
	defer session.Close()
	tx, err := session.BeginTrans()
	if err != nil {
		return err
	}
	for _, bean := range beans {
		_, err := tx.Session().ID(bean[0]).Delete(&bean)
		if err != nil {
			return tx.RollbackTrans()
		}
	}
	return nil
}

func (d Dao) getOne(bean interface{}, cols ...string) interface{} {
	get, err := d.engine.Cols(cols...).Get(bean)
	if utils.CheckError(err, get) && get {
		return bean
	}
	return nil
}
