package main

import (
	"container/list"
)

//基于以下几点要点分析
//a. 月内活动中
//1. 用户数量较大
//2. 月内更新频繁
//考虑使用hash来实现对user的快速索引
//
//b. 活动结束后
//3. 需要有序
//4. 需要前后10名玩家
//考虑使用线性结构存储，实现前后的快速索引，比如list

var (
	userList list.List		// 保存User地址
	userMap map[string]*User
)

type User struct {
	userID string
	score	int
	updateTime	int
}
func (u *User) UpdateSocre(delta int) bool{
	return  false
}

func findUserAndUpdate(userId string, delta int) *User{
	u := userMap[userId]
	u.UpdateSocre(delta)
	return nil
}

 // 活动结束后
func sortUser(){
	// 遍历userMap, 保存&User到userList

	// userList sort： 自定义排序算法
	userList.Len()
}

func getRank(userId string) []User{
	users := make([]User, 0)
	u := userMap[userId]
	users = append(users, *u)

	// for v = u.previow
	//   append( , *v)
	// for v = u.next
	//   append( , *v)

	return users
}