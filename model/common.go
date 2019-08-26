package model

type ResponseResult struct {
	Code    int  `json:"code"`
	Content interface{} `json:"content"`
	Msg     string      `json:"msg"`
} 
