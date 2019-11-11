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

	for member := range members {
		if d.rdx.SIsMember(ArticleIdQueue, member).Val() {
			continue
		}
		err := d.rdx.SAdd(ArticleIdQueue, members).Err()
		if utils.CheckErrorArgs(fmt.Sprintf("add queue id is %s", member), err) {
			memList = append(memList, string(member))
		}
	}
	return
}
