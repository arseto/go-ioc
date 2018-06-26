package ioc

type instantiator func() interface{}

type objectContainer struct {
	singleton bool
	inst      instantiator
	objCache  interface{}
}

type ioc struct {
	container map[string]*objectContainer
}

func new() *ioc {
	return &ioc{
		container: make(map[string]*objectContainer),
	}
}

func (c *ioc) singleton(key string, inst instantiator) {
	c.localBind(key, inst, true)
}

func (c *ioc) bind(key string, inst instantiator) {
	c.localBind(key, inst, false)
}

func (c *ioc) localBind(key string, inst instantiator, singleton bool) {
	c.container[key] = &objectContainer{
		singleton: singleton,
		inst:      inst,
		objCache:  nil,
	}
}

func (c *ioc) make(key string) interface{} {
	single := c.container[key].singleton

	if !single || (c.container[key].objCache == nil) {
		c.container[key].objCache = c.container[key].inst()
	}

	return c.container[key].objCache
}

func (c *ioc) isRegistered(key string) bool {
	_, ok := c.container[key]
	return ok
}

var i *ioc

func init() {
	i = new()
}

//Singleton will register object for single instance,
//any call in subsequent `Make` method will return same instance.
func Singleton(key string, inst instantiator) {
	i.singleton(key, inst)
}

//Bind will register object to always create new instance,
//any call in subsequent `Make` method will return new instance.
func Bind(key string, inst instantiator) {
	i.bind(key, inst)
}

//IsRegistered checks if a key is already registered in container
func IsRegistered(key string) bool {
	return i.isRegistered(key)
}

//Make calls object from container, whether it's singleton or not
func Make(key string) interface{} {
	return i.make(key)
}
