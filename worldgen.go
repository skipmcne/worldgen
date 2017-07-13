package main

import (
	"fmt"
	"math/rand"
    "time"
	)

var seed int64
var rg *rand.Rand
var world World
// Point within the map
type Point struct{
    X int64
    Y int64
}
// World defines information about the world
type World struct {
    ExtentX int64
    ExtentY int64
    
//    Radius int64
}

func main() {
	//var s  int64
    //s = 8181117722394100709
    s :=time.Now().UnixNano()
    var x,y float64
 	rg = rand.New(rand.NewSource(s))
    plates = Graph.new()

 // world is 40,000 Km x 40,000 Km
// treat it as a 400 x 400 map (each pixel is 100 Km x 100 Km)
    world.ExtentX = 400;
    world.ExtentY = 400;
//    world.Radius = 

    // pick centers for points
    // 90 plates max
    for (pl : 1: 90) {
        plates.points.add (world.extentX*rand.new(),world.extentY*rand.new());        
    } 
    // delauny triangulation
    
}



