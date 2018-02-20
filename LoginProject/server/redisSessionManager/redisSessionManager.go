package redisSessionManager

import (
	"LoginProject/server/api/common/constants"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func Init() {
	logginghelper.LogInfo("Inside redisSessionManager:: Initializing redis Client")
	client = redis.NewClient(&redis.Options{
		Addr:     constants.REDIS_CONNECTION_ADDR,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if nil != err {
		logginghelper.LogInfo("Redis inti error:", err)
	}
	logginghelper.LogInfo(pong)
}

func Set(key string, value string) error {
	logginghelper.LogInfo("Inside redisSessionManager::Setting session")

	err := client.Set(key, value, constants.DEFAULT_SESSION_EXPIRE_TIME).Err()
	if nil != err {
		logginghelper.LogError("Failed to set Key", err)
		return err
	}
	return nil
}

func getActiveSessionCount(pattern string) (int, error) {
	logginghelper.LogInfo("Inside redisSessionManager:: getActiveSessionCount")
	pattern = pattern + "*"
	keys, err := client.Keys(pattern).Result()

	if err != nil {
		logginghelper.LogError("session count redis nil", err)
		return -1, err
	}
	return len(keys), nil
}

func IsSessionLimitReached(username string) bool {
	logginghelper.LogInfo("Inside redisSessionManager:: IsSessionLimitReached")
	activeSessions, err := getActiveSessionCount(username)
	if nil != err {
		logginghelper.LogError("IsSessionLimitReached", err)
		return true
	}

	if -1 != activeSessions && activeSessions >= constants.MAX_SESSION_ALLOWED {
		return true
	}
	return false
}

func Get(key string) (string, error) {
	logginghelper.LogInfo("Inside redisSessionManager:: Get")
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		logginghelper.LogInfo("key does not exists")
		return "", err
	} else if err != nil {
		logginghelper.LogError(err)
		return "", err
	} else {
		return val, nil
	}
}

func Del(keys ...string) error {
	logginghelper.LogInfo("Inside redisSessionManager:: Del")
	err := client.Del(keys...).Err()
	if err == redis.Nil {
		logginghelper.LogInfo("key does not exists")
	}
	if nil != err {
		logginghelper.LogInfo("Failed to delete Session")
		return err
	}
	return nil
}

func SlideSession(key string) error {
	logginghelper.LogInfo("Inside redisSessionManager:: SlideSession")
	err := client.Expire(key, constants.DEFAULT_SESSION_EXPIRE_TIME).Err()
	if nil != err {
		logginghelper.LogError("Failed to Extend Key TTL")
		return err
	}
	return nil
}
