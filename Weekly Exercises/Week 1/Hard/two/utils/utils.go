package utils

type Counter interface {
	Inc(key string)
	Get(key string) int
}

type MapCounter struct {
	Counter map[string]int
}

func (c *MapCounter) Inc(key string) {
	if c.Counter == nil {
		c.Counter = make(map[string]int)
		c.Counter[key] = 1
		return
	}
	c.Counter[key]++
}

func (c *MapCounter) Get(key string) int {
	return c.Counter[key]
}

type FakeCounter struct {
	Counter map[string]int
}

func (c *FakeCounter) Inc(key string) {
	if c.Counter == nil {
		c.Counter = make(map[string]int)
		c.Counter[key] = 1
		return
	}
	c.Counter[key]++
}

func (c *FakeCounter) Get(key string) int {
	return c.Counter[key]
}

func RecordLogin(c Counter, userID string) {
	c.Inc("login: " + userID)
}
