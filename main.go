package main

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile("./db", o)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//
	//ok, err := db.Has([]byte("art-10000"), nil)
	//fmt.Println(ok, err)

	//err = db.Put(str2byt("key"), str2byt("val"), nil)
	//if err != nil {
	//	log.Fatal(false)
	//}
	//
	//val, err := db.Get(str2byt("key"), nil)
	//
	//if errors.Is(err, leveldb.ErrNotFound) {
	//	fmt.Println("key not found")
	//	return
	//}
	//
	//fmt.Println("value is :", string(val))

	//cate := []string{"golang", "javascript"}
	//
	//for i := 0; i < 10000; i++ {
	//	art := Article{
	//		ID:      strconv.Itoa(i),
	//		CateId:  cate[rand.Intn(len(cate))],
	//		Title:   fmt.Sprintf("Create or open a database: %d", i),
	//		Content: content,
	//	}
	//	data, _ := json.Marshal(art)
	//	db.Put(str2byt("art-"+art.ID), data, nil)
	//}
	//
	//for _, c := range cate {
	//	cate := Category{
	//		ID:   c,
	//		Name: c,
	//	}
	//	data, _ := json.Marshal(cate)
	//	db.Put(str2byt("cate-"+c), data, nil)
	//}

	// page + pageSize
	//  3 页
	// 1003
	//pageSize := 2
	//lastId := "1003" // 1003
	//
	//// art-1003
	//// art-1004
	//
	//// 'a'  > 'a'
	//// 'b'  < 'b'
	//iter := db.NewIterator(util.BytesPrefix([]byte("art-"+lastId)), nil)
	//for iter.Next() {
	//	if pageSize == 0 {
	//		break
	//	}
	//	key := iter.Key()
	//	value := iter.Value()
	//	fmt.Println(string(key), string(value))
	//	pageSize--
	//}

	// https://www.xxx.com/article?id=2
	// https://www.xxx.com/article?id=5  --> 无效!
	// https://www.xxx.com/article?id=6
	// https://www.xxx.com/article?id=4
	/*
			----------------
			|0 0 0 0 0 1 1 0|
			----------------
			----------------
			|0 0 0 0 0 1 1 00 0 0 0 0 1 1 00 0 0 0 0 1 1 00 0 0 0 0 1 1 0|
			----------------

		核心：如果过滤器反回值不存在，那么这个值一定不存在，否则，可能存在!
	*/

}

func str2byt(s string) []byte {
	return []byte(s)
}

type Article struct {
	ID      string `json:"id"`
	CateId  string `json:"cate_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var content = "This is an implementation of the [LevelDB key/value database](https://github.com/google/leveldb) in the [Go programming language](https://go.dev).\n\n[![Build Status](https://app.travis-ci.com/syndtr/goleveldb.svg?branch=master)](https://app.travis-ci.com/syndtr/goleveldb)\n\nInstallation\n-----------\n\n\tgo get github.com/syndtr/goleveldb/leveldb\n\nRequirements"
