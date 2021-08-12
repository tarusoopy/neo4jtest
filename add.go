package main

import (
	"fmt"
	"github.com/jmcvetta/neoism"
)

func AddNode() {
	db, _ := neoism.Connect("http://localhost:7474/db/data")

	// ルートnodeを作成
	rootnode, _ := db.CreateNode(neoism.Props{"word": "テスト", "url": "http://www.test.com"})
	//defer newnode.Delete()

	// リレーション用ノードを作成
	relnode1, _ := db.CreateNode(neoism.Props{"word": "テスト1", "url": "http://www.test1.com"})
	relnode2, _ := db.CreateNode(neoism.Props{"word": "テスト2", "url": "http://www.test2.com"})
	relnode3, _ := db.CreateNode(neoism.Props{"word": "テスト3", "url": "http://www.test3.com"})
	relnode4, _ := db.CreateNode(neoism.Props{"word": "テスト4", "url": "http://www.test4.com"})
	relnode5, _ := db.CreateNode(neoism.Props{"word": "テスト5", "url": "http://www.test5.com"})

	fmt.Println(newnode.Id())

	// ルートノードにリレーション用ノードを関連付ける
	newnode.Relate("ConnectTo", relnode1.Id(), neoism.Props{})
	newnode.Relate("ConnectTo", relnode2.Id(), neoism.Props{})
	newnode.Relate("ConnectTo", relnode3.Id(), neoism.Props{})
	newnode.Relate("ConnectTo", relnode4.Id(), neoism.Props{})
	newnode.Relate("ConnectTo", relnode5.Id(), neoism.Props{})

}
