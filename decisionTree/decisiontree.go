package decisionTree

import(
	"fmt"
	"runtime"
	"flag"
	"io"
	"encoding/binary"
	"os"
	"github.com/Pattern-Recognition3/randomForestDecisionTree"
)

/*
./example_mnist
-si train-images-idx3-ubyte
-sl train-labels-idx1-ubyte
-ti t10k-images-idx3-ubyte
-tl t10k-labels-idx1-ubyte

 */
func TrainMNISTData(t,s,fa int){
	fmt.Println("Training MNIST training data set...")
	runtime.GOMAXPROCS(8)


	//take test image files from console
	testImageFile := flag.String("tif","","testImageFile")
	testLabelFile := flag.String("tlf","","testLabelFile")

	//take MNIST files from console
	trainImageFile := flag.String("trif","","trainImageFile")
	trainLabelile := flag.String("trlf","","trainLabelile")

	//parse the flag
	flag.Parse()


	//open label file : todo: error check
	labelOpenedFile,_ := os.Open(*trainLabelile)
	imageOpenedFile,_ := os.Open(*trainImageFile)


	//load training data by reading from label
	loadedTrainingData := readLabel(labelOpenedFile)

	//load image data by reading from train image file
	loadedImageData, width, height := readImage(imageOpenedFile)

	fmt.Println(len(loadedImageData),len(loadedImageData[0]),width,height)
	fmt.Println(len(loadedTrainingData),loadedTrainingData[0:10])


	//

	trainImageDataRows := len(loadedImageData)
	trainImageInput := make([][]interface{}, trainImageDataRows)
	for i:=0;i< trainImageDataRows;i++{
		trainImageInput[i] = make([]interface{},len(loadedImageData[i]))
		for j:=0;j<len(loadedImageData[i]);j++{
			trainImageInput[i][j] = float64(loadedImageData[i][j])
		}
	}

	targets := make([]string,len(loadedTrainingData))
	for i:=0;i<len(targets);i++{
		targets[i] = fmt.Sprintf("%d",int(loadedTrainingData[i]))
	}

	//todo: check performance with verying treesAmount, samplesAmount, selected feature amount
	randomForestData := RF.BuildForest(trainImageInput,targets,t,s,fa)


	//RF.DumpForest(forest,"rf.bin")
	var testLabelDatas []byte
	var testImageDatas [][]byte
	if *testLabelFile != "" && *testImageFile != "" {
		fmt.Println("Loading test data...")
		testLabelFilefile, _ := os.Open(*testLabelFile)
		testImageFilefile, _ := os.Open(*testLabelFile)
		testLabelDatas = readLabel(testLabelFilefile)
		testImageDatas, _, _ = readImage(testImageFilefile)
	}

	//test_inputs := prepareX(testImageDatas)
	//test_targets := prepareY(testLabelDatas)

	rows := len(testImageDatas)
	testInputs := make([][]interface{},rows)
	for i:=0;i<rows;i++ {
		testInputs[i] = make([]interface{}, len(testImageDatas[i]))
		for j := 0; j < len(testImageDatas[i]); j++ {
			testInputs[i][j] = float64(testImageDatas[i][j])
		}
	}

	testTargets := make([]string,len(testLabelDatas))
	for i:=0;i<len(testTargets);i++{
		testTargets[i] = fmt.Sprintf("%d",int(testLabelDatas[i]))
	}

	//calculatePerformance()
	correct_ct :=0
	for i,p := range(testInputs){
		y := randomForestData.Predicate(p)
		yy := testTargets[i]
		//fmt.Println(y,yy)
		if y == yy{
			correct_ct += 1
		}
	}
	fmt.Println("Accuracy :: ", float64(correct_ct)/ float64(len(testInputs)) ,"%")

}

//func calculatePerformance(testInputs interface{} ) {
//	correct_ct :=0
//	for i,p := range(testInputs){
//		y := forest.Predicate(p)
//		yy := test_targets[i]
//		//fmt.Println(y,yy)
//		if y == yy{
//			correct_ct += 1
//		}
//	}
//
//	fmt.Println("correct rate: ", float64(correct_ct)/ float64(len(test_inputs)), correct_ct,len(test_inputs))
//}

func readLabel(label io.Reader) (labelByte []byte){
	//initialize data
	data := [2]int32{}
	binary.Read(label,binary.BigEndian,data)


	//fill labelByte
	labelByte = make([]byte,data[1])
	label.Read(labelByte)
	return
}

func readImage (r io.Reader) (outputImages [][]byte, w, h int) {
	data := [4]int32{}
	binary.Read(r, binary.BigEndian, &data)


	outputImages = make([][]byte, data[1])
	w, h = int(data[2]), int(data[3])
	for i := 0; i < len(outputImages); i++ {
		outputImages[i] = make([]byte, w*h)
		r.Read(outputImages[i])
	}
	return
}