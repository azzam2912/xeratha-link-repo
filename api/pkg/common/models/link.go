package models

type Link struct {
	Id				int		`json:"id"`
	Title			string	`json:"title"`
	LinkAddress		string	`json:"linkAddress"`
	Tags			string 	`json:"tags"`
}