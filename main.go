package main

import (
	"fmt"
	//"github.com/Pattern-Recognition3/gomnist"
	"github.com/Pattern-Recognition3/decisionTree"
	//"os"
)

func main() {

	//taking user input
	fmt.Println("Enter number of trees, samples and feature space...")
	var t,s,fa int
	//fmt.Scanf("%d %d %d", &t,&s,&fa)
	t = 100
	s = 2000
	fa = 30


	decisionTree.TrainMNISTData(t,s,fa)



}
