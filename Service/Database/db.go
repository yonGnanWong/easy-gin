package Database

type redisConn interface {
	Get()
	Set()
	SetNx()
	Delete()
}
