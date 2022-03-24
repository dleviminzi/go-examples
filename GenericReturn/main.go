package main

type genReturn[T any] interface {
	fetch()T
}