package main

import "fmt"

type mymap map[int]*myentry

type myentry struct {
	m mymap
	b bool
}

func (mm mymap) get(idx ...int) *myentry {
	if len(idx) == 0 {
		return nil
	}
	entry, ok := mm[idx[0]]
	if !ok {
		return nil
	} else if len(idx) == 1 {
		return entry
	}
	for i := 1; i < len(idx); i++ {
		if entry == nil || entry.m == nil {
			return nil
		}
		entry = entry.m[idx[i]]
	}
	return entry
}

func (mm mymap) setbool(v bool, idx ...int) {
	if len(idx) == 0 {
		return
	}
	if mm[idx[0]] == nil {
		mm[idx[0]] = &myentry{m: make(mymap), b: false}
	} else if mm[idx[0]].m == nil {
		mm[idx[0]].m = make(mymap)
	}
	if len(idx) == 1 {
		mm[idx[0]].b = v
		return
	}
	entry := mm[idx[0]]
	for i := 1; i < len(idx); i++ {
		if entry.m == nil {
			entry.m = make(mymap)
			entry.m[idx[i]] = &myentry{m: make(mymap), b: false}
		} else if entry.m[idx[i]] == nil {
			entry.m[idx[i]] = &myentry{m: make(mymap), b: false}
		}
		entry = entry.m[idx[i]]
	}
	entry.b = v
}

func (m mymap) getbool(idx ...int) bool {
	if val := m.get(idx...); val != nil {
		return val.b
	}
	return false
}

func (m mymap) getmap(idx ...int) mymap {
	if val := m.get(idx...); val != nil {
		return val.m
	}
	return nil
}

func main() {
	m := make(mymap)
	m.setbool(true, 0, 1, 6)
	fmt.Println("Should be false:", m.getbool(0, 1))
	fmt.Println("Should be true:", m.getbool(0, 1, 6))
	m.setbool(true, 0, 1, 6, 2)
	fmt.Println("Should be true:", m.getbool(0, 1, 6, 2))
	m.setbool(false, 0, 1, 6)
	fmt.Println("Should be false:", m.getbool(0, 1, 6))
	fmt.Println("Should contain key 6 only:", m.getmap(0, 1))
	m.setbool(true, 0, 1, 2)
	m.setbool(true, 0, 1, 9)
	fmt.Println("Should contain keys 2, 6, and 9:", m.getmap(0, 1))
	q := m.getmap(0, 1)
	fmt.Println("should be true:", q.getbool(9))
}
