//
// Copyright (C) 2019 Dmitry Kolesnikov
//
// This file may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.
// https://github.com/fogfish/dynamo
//

package main

import (
	"fmt"
	"os"

	"github.com/fogfish/dynamo"
	"github.com/fogfish/iri"
)

type person struct {
	iri.ID
	Name    string `dynamodbav:"name,omitempty" json:"name,omitempty"`
	Age     int    `dynamodbav:"age,omitempty" json:"age,omitempty"`
	Address string `dynamodbav:"address,omitempty" json:"address,omitempty"`
}

func main() {
	db, err := dynamo.New(os.Args[1])
	if err != nil {
		panic(err)
	}

	examplePut(db)
	exampleGet(db)
	exampleUpdate(db)
	exampleMatch(db)
	exampleRemove(db)
}

const n = 5

func examplePut(db dynamo.KeyVal) {
	for i := 0; i < n; i++ {
		val := folk(i)
		err := db.Put(val)

		fmt.Println("=[ put ]=> ", either(err, val))
	}
}

func exampleGet(db dynamo.KeyVal) {
	for i := 0; i < n; i++ {
		val := &person{ID: id(i)}
		err := db.Get(val)

		fmt.Println("=[ get ]=> ", either(err, val))
	}
}

func exampleUpdate(db dynamo.KeyVal) {
	for i := 0; i < n; i++ {
		val := &person{ID: id(i), Address: "Viktoriastrasse 37, Berne, 3013"}
		err := db.Update(val)

		fmt.Println("=[ update ]=> ", either(err, val))
	}
}

func exampleMatch(db dynamo.KeyVal) {
	seq := db.Match(iri.New("test"))

	for seq.Tail() {
		val := &person{}
		err := seq.Head(val)
		fmt.Println("=[ match ]=> ", either(err, val))
	}

	if err := seq.Error(); err != nil {
		fmt.Println("=[ match ]=> ", err)
	}

}

func exampleRemove(db dynamo.KeyVal) {
	for i := 0; i < n; i++ {
		val := &person{ID: id(i)}
		err := db.Remove(val)

		fmt.Println("=[ remove ]=> ", either(err, val))
	}
}

func folk(x int) *person {
	return &person{id(x), "Verner Pleishner", 64, "Blumenstrasse 14, Berne, 3013"}
}

func id(x int) iri.ID {
	return iri.New("test:%v", x)
}

func either(e error, x interface{}) interface{} {
	if e != nil {
		return e
	}
	return x
}
