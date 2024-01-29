package domain

type Db interface {
	GetOrderDestinationLatLng(orderId int64) (loc Location, err error)
	GetOrderRestaurantLocation(orderId int64) (loc Location, err error)
}

type Map interface {
	GetPath(start Location, dest Location) ([]Location, error)
}

type Bus interface {
	Publish(topic string, msg interface{})
}
