package service

func (svc *Service) Test(key, value string) error {
	return svc.pool.RedisTest(key, value)
}
