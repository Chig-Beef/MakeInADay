package main

func rectCollide(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return !(x1+w1 < x2 || x1 > x2+w2 || y1+h1 < y2 || y1 > y2+h2)
}
