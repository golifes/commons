package admin

import "commons/tools/utils"

/**
redis  set string
*/

const expireTime = 60

func (d Dao) set(key string, value interface{}) bool {
	result, err := d.rdx.Set(key, value, expireTime).Result()
	if utils.CheckError(err, result) {
		return true
	}

	return false
}

func (d Dao) get(key string) string {
	result, err := d.rdx.Get(key).Result()
	if utils.CheckError(err, result) {
		return result
	}
	return ""
}
