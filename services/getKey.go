package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/solrac97gr/cryptoAPI/async"
	cacheDb "github.com/solrac97gr/cryptoAPI/cache"
	"github.com/solrac97gr/cryptoAPI/models"
)

func fetchKey() string {
	colorResponse := new(models.ColorResponse)
	response, _ := http.Get("https://random-data-api.com/api/color/random_color")
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &colorResponse)
	return colorResponse.GetUID()
}

func GetKey() string {
	var usedTimes string
	var uuidStored string
	if uuid, found := cacheDb.CacheDb.Get("uid"); found {
		if usedTimesStored, found := cacheDb.CacheDb.Get("usedTimes"); found {
			usedTimes = usedTimesStored.(string)
			usedTimesInt, _ := strconv.Atoi(usedTimes)
			if usedTimesInt < 3 {
				uuidStored = uuid.(string)
				usedTimesInt++
				usedTimesSting := strconv.Itoa(usedTimesInt)
				cacheDb.CacheDb.Set("usedTimes", usedTimesSting, -1)
			} else {
				future := async.Exec(func() interface{} {
					return fetchKey()
				})
				val := future.Await()
				uuidStored = val.(string)
				cacheDb.CacheDb.Set("uid", val, -1)
				cacheDb.CacheDb.Set("usedTimes", "1", -1)
			}
		}
	} else {
		future := async.Exec(func() interface{} {
			return fetchKey()
		})
		val := future.Await()
		uuidStored = val.(string)
		cacheDb.CacheDb.Set("uid", val, -1)
		cacheDb.CacheDb.Set("usedTimes", "1", -1)
	}
	return uuidStored
}
