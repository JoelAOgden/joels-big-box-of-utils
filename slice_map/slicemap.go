package slice_map

import "github.com/JoelAOgden/joels-big-box-of-utils/slice_util"

type valuePair[T any] struct {
	key      string
	valuePtr *T
}

// SliceMap :This was created to give all the positives (except speed) of a map_info without it being passed by reference which
// can be extremely annoying in golang, think of this more like an array list with an index hash or something
type SliceMap[T any] struct {
	keyValues []valuePair[T]
}

func New[T interface{}]() SliceMap[T] {
	keyValues := make([]valuePair[T], 0)
	return SliceMap[T]{
		keyValues: keyValues,
	}
}

func (sm SliceMap[T]) Duplicate() SliceMap[T] {
	newSm := New[T]()
	for _, v := range sm.keyValues {
		//if _, ok := v.(T); ok {
		newSm.Put(v.key, *v.valuePtr)
		//}
	}
	return newSm
}

func (sm SliceMap[T]) ToMap() (out map[string]T) {
	out = make(map[string]T)
	for i, v := range sm.keyValues {
		out[v.key] = *sm.keyValues[i].valuePtr
	}
	return out
}

func (sm SliceMap[T]) ContainsKey(key string) bool {
	for _, v := range sm.keyValues {
		if v.key == key {
			return true
		}
	}
	return false
}

func (sm *SliceMap[T]) Put(key string, value T) (replaced bool) {

	if !sm.ContainsKey(key) {
		newKVPair := valuePair[T]{
			key:      key,
			valuePtr: &value,
		}
		sm.keyValues = append(sm.keyValues, newKVPair)
		return false
	} else {
		ptr := sm.Get(key)
		*ptr = value
		return true
	}
}

// Replace will maintain the pointer
func (sm *SliceMap[T]) Replace(key string, value T) (ptr *T) {
	ptr = sm.Get(key)
	*ptr = value
	return ptr
}

func (sm SliceMap[T]) Get(key string) *T {

	if !sm.ContainsKey(key) {
		return nil
	}

	for i, v := range sm.keyValues {
		if v.key == key {
			pair := sm.keyValues[i]
			return pair.valuePtr
		}
	}
	return nil
}

func (sm SliceMap[T]) Values() (values []*T) {
	for _, v := range sm.keyValues {
		values = append(values, v.valuePtr)
	}
	return values
}

func (sm SliceMap[T]) Count() int {
	return len(sm.keyValues)
}

func (sm SliceMap[T]) Keys() (keys []string) {
	for _, v := range sm.keyValues {
		keys = append(keys, v.key)
	}
	return keys
}

func (sm SliceMap[T]) KeysAndValues() (keys []string, values []*T) {
	for _, v := range sm.keyValues {
		values = append(values, v.valuePtr)
		keys = append(keys, v.key)
	}
	return keys, values
}

func (sm *SliceMap[T]) Clear() {
	sm.keyValues = make([]valuePair[T], 0)
}

func (sm *SliceMap[T]) Remove(key string) {

	for i, v := range sm.keyValues {
		if v.key == key {
			sm.keyValues = slice_util.RemoveIndex(sm.keyValues, i)
		}
	}

}
