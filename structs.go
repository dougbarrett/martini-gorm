package main

type Item struct {
	Id          int64  `form:"id"`
	Title       string `form:"title"`
	Description string `form:"description"`
	UserName    string `form:"user_name"`
}
