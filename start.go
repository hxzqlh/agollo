package agollo

import (
	"github.com/zouyx/agollo/v2/agcache"
	. "github.com/zouyx/agollo/v2/component/log"
)

func init() {
	//init config
	initFileConfig()

	initDefaultConfig()
}

//InitCustomConfig init config by custom
func InitCustomConfig(loadAppConfig func() (*AppConfig, error)) {

	initConfig(loadAppConfig)

	initDefaultConfig()
}

//start apollo
func Start() error {
	return startAgollo()
}

//SetLogger 设置自定义logger组件
func SetLogger(loggerInterface LoggerInterface) {
	if loggerInterface != nil {
		initLogger(loggerInterface)
	}
}

//SetCache 设置自定义cache组件
func SetCache(cacheFactory *agcache.DefaultCacheFactory) {
	if cacheFactory != nil {
		initConfigCache(cacheFactory)
	}
}

//StartWithLogger 通过自定义logger启动agollo
func StartWithLogger(loggerInterface LoggerInterface) error {
	SetLogger(loggerInterface)
	return startAgollo()
}

//StartWithCache 通过自定义cache启动agollo
func StartWithCache(cacheFactory *agcache.DefaultCacheFactory) error {
	SetCache(cacheFactory)
	return startAgollo()
}

func startAgollo() error {
	//init server ip list
	go initServerIpList()
	//first sync
	if err := notifySyncConfigServices(); err != nil {
		return err
	}
	Logger.Debug("init notifySyncConfigServices finished")

	//start long poll sync config
	go StartRefreshConfig(&NotifyConfigComponent{})

	Logger.Info("agollo start finished ! ")

	return nil
}
