package ioc

type instantiator func() interface{}

type objectContainer struct {
	singleton bool
	inst      instantiator
	objCache  interface{}
}

type Ioc struct {
	container map[string]*objectContainer
}

func New() *Ioc {
	return &Ioc{
		container: make(map[string]*objectContainer),
	}
}

func (c *Ioc) Singleton(key string, inst instantiator) {
	c.localBind(key, inst, true)
}

func (c *Ioc) Bind(key string, inst instantiator) {
	c.localBind(key, inst, false)
}

func (c *Ioc) localBind(key string, inst instantiator, singleton bool) {
	c.container[key] = &objectContainer{
		singleton: singleton,
		inst:      inst,
		objCache:  nil,
	}
}

func (c *Ioc) Make(key string) interface{} {
	single := c.container[key].singleton

	if !single || (c.container[key].objCache == nil) {
		c.container[key].objCache = c.container[key].inst()
	}

	return c.container[key].objCache
}

func (c *Ioc) IsRegistered(key string) bool {
	_, ok := c.container[key]
	return ok
}

var i *Ioc

func init() {
	i = New()
}

func Singleton(key string, inst instantiator) {
	i.Singleton(key, inst)
}

func Bind(key string, inst instantiator) {
	i.Bind(key, inst)
}

func IsRegistered(key string) bool {
	return i.IsRegistered(key)
}

func Make(key string) interface{} {
	return i.Make(key)
}
