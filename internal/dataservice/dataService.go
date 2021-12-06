package dataservice

import "wsc2017/domain/model"

type CacheDataInterface interface {
	Get(key string) ([]byte, string)
	Store(key string, value []byte) error
}

type UserDataInterface interface {
	Find()
}

type PlaceDataInterface interface {
	Create(place *model.Place)
	Update(id string, place *model.Place)
	Delete(id string)
}

type ScheduleDataInterface interface {
}

type TokenDataInterface interface {
}

type RouteHistoryDataInterface interface {
}
