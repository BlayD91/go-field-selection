package main

import "time"

type ImageGroup struct {
	Id   int
	Name string
}

type Role struct {
	Id     int
	Name   string
	Images []Image
}

type Image struct {
	Id          int
	Url         string
	Size        float64
	ImageGroups []ImageGroup
}

type User struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	Password  string
	Roles     []Role
	Image     Image
	Birthday  time.Time
}


