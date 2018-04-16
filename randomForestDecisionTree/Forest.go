//Note : this code is taken from [https://github.com/fxsjy/RF.go] for generating random forest decision tree, which provides Ensembles of decision trees
//slight change of code is done to solve problem 1
package RF

import (
"math"
"time"
"math/rand"
"fmt"
"os"
"encoding/json"
"sync"
)
type Forest struct{
	Trees []*Tree
}

/*
Using random forest to construct decision tree
Ex : 0 1 2 3 4 5 6 7 8 9 is divided into
 0 1 2 3 4  | 5 6 7 8 9 which is further divided into
0 1 2 |  3 4 and so on till we get single digit as a tree
 */
func BuildForest(inputs [][]interface{},labels []string, treesAmount, samplesAmount, selectedFeatureAmount int) *Forest{
	rand.Seed(time.Now().UnixNano())
	forest := &Forest{}
	forest.Trees = make([]*Tree,treesAmount)
	done_flag := make(chan bool)
	prog_counter := 0
	mutex := &sync.Mutex{}
	for i:=0;i<treesAmount;i++{
		go func(x int){
			fmt.Printf("tree # %v...\n", x)
			forest.Trees[x] = BuildTree(inputs,labels,samplesAmount,selectedFeatureAmount)
			mutex.Lock()
			prog_counter+=1
			mutex.Unlock()
			done_flag <- true
		}(i)
	}

	for i:=1;i<=treesAmount;i++{
		<-done_flag
	}

	fmt.Println("Analysis completed done.")
	return forest
}

func DefaultForest(inputs [][]interface{},labels []string, treesAmount int) *Forest{
	m := int( math.Sqrt( float64( len(inputs[0]) ) ) )
	n := int( math.Sqrt( float64( len(inputs) ) )  )
	return BuildForest(inputs,labels, treesAmount,n,m)
}

func (self *Forest) Predicate(input []interface{}) string{
	counter := make(map[string]float64)
	for i:=0;i<len(self.Trees);i++{
		tree_counter := PredicateTree(self.Trees[i],input)
		total := 0.0
		for _,v := range tree_counter{
			total += float64(v)
		}
		for k,v := range tree_counter{
			counter[k] += float64(v) / total
		}
	}

	max_c := 0.0
	max_label := ""
	for k,v := range counter{
		if v>=max_c{
			max_c = v
			max_label = k
		}
	}
	return max_label
}

func DumpForest(forest *Forest, fileName string){
	out_f, err:=os.OpenFile(fileName,os.O_CREATE | os.O_RDWR,0777)
	if err!=nil{
		panic("failed to create "+fileName)
	}
	defer out_f.Close()
	encoder := json.NewEncoder(out_f)
	encoder.Encode(forest)
}

func LoadForest(fileName string) *Forest{
	in_f ,err := os.Open(fileName)
	if err!=nil{
		panic("failed to open "+fileName)
	}
	defer in_f.Close()
	decoder := json.NewDecoder(in_f)
	forest := &Forest{}
	decoder.Decode(forest)
	return forest
}


