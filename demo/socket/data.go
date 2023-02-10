package main

/**
* @program: work_space
*
* @description:
*
* @author: khr
*
* @create: 2023-02-09 10:06
**/
type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}
