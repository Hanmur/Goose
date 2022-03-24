package dao

//RedisTest 测试Redis的使用
func (pool *Pool) RedisTest(key string, val string) error {
	conn := pool.Pool.Get()
	defer func() {
		_ = conn.Close()
	}()

	_, err := conn.Do("Set", key, val)
	if err != nil {
		return err
	}

	return nil
}
