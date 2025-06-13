package main

const SCREEN_WIDTH = 640
const SCREEN_HEIGHT = SCREEN_WIDTH / 2

const VIRTUAL_WIDTH = 100
const VIRTUAL_HEIGHT = VIRTUAL_WIDTH / 2

const IS_RELEASE = false

const FPS = 60
const MS_PER_FRAME = 1000/FPS

const DFLT_LINE_WIDTH = float64(VIRTUAL_WIDTH) / float64(SCREEN_WIDTH)

const GRID_SIZE = 9
const SUB_GRID_SIZE = 3
const TILE_SIZE = float64(VIRTUAL_HEIGHT)/float64(GRID_SIZE)
