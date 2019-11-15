package wx

import (
	"commons/tools/utils"
	"fmt"
	"time"
)

/**
es相关的简单操作
*/

const (
	ArticleIdQueue = "articleIdQueue"
)

func (d Dao) set(key string, value interface{}, expireTime int64) bool {
	res, err := d.rdx.Set(key, value, time.Duration(expireTime)).Result()
	if utils.CheckErrorArgs(res, err) {
		return true
	}

	return false
}

func (d Dao) get(key string) []byte {
	bytes, err := d.rdx.Get(key).Bytes()
	if utils.CheckErrorArgs("get redis task", err) {
		return bytes
	}
	return nil
}

/**
set集合
*/
func (d Dao) sAdd(members ...interface{}) (memList []string) {

	for _, member := range members {
		if d.rdx.SIsMember(ArticleIdQueue, member).Val() {
			continue
		}
		err := d.rdx.SAdd(ArticleIdQueue, members...).Err()
		if utils.CheckErrorArgs(fmt.Sprintf("add queue id is %s", member), err) {
			memList = append(memList, member.(string))
		}
	}
	return
}

/**
判断一个元素是否属于集合
*/
func (d Dao) sisMember(key string, member interface{}) bool {
	result, err := d.rdx.SIsMember(key, member).Result()
	if utils.CheckErrorArgs(result, err) {
		return true
	}

	return false
}

func (d Dao) setQueue(key string, members ...interface{}) bool {
	err := d.rdx.SAdd(key, members...).Err()
	if utils.CheckErrorArgs("add queue rdx", err) {
		return true
	}
	return false
}

func (d Dao) sPop(key string, ps int64) []string {
	result, err := d.rdx.SPopN(key, ps).Result()
	if utils.CheckErrorArgs(result, err) {
		return result
	}
	return nil
}
