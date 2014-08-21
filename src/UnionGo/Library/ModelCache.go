package Library


import "reflect"

var ModelCache=&_modelCache{
	cache:make(map[string]reflect.Type),
}
// model info collection
type _modelCache struct {

	cache     map[string]reflect.Type

}
// get model info by table name
func (mc *_modelCache) Get(table string) (mi reflect.Type, ok bool) {
	mi, ok = mc.cache[table]
	return
}



// set model info to collection
func (mc *_modelCache) Set(table string, mi reflect.Type) reflect.Type {
	mii := mc.cache[table]
	mc.cache[table] = mi

	return mii
}
