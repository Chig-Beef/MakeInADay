package main

type Tile struct {
	num int
	locked bool
	known bool
	canBe [9]bool
}
