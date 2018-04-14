package gomnist

import (
	"fmt"
	"github.com/petar/GoMNIST"
	"math"
	"math/rand"
	"strconv"
	"time"
)

/*GetMeanAndSD - calculates mean and standard deviation of digits present in mnist training dataset
digitEntered 0 = mean, 1 = standard deviation

NOTE : After getting matrix from this function, the matrix needs to be passed in matlab console to get the image
Ref  :[https://www.mathworks.com/matlabcentral/answers/30784-how-to-convert-a-matrix-to-a-gray-scale-image]

Below code computes the image from the pixel matrix (for mean and standard deviation)
 >> m = [Mean or SD matrix]
 >> image(m);
*/
func GetMeanAndSD(digitEntered, typeOf int) {

	//loading dataset [Ref :https://github.com/petar/GoMNIST]
	traininingSet, _, err := GoMNIST.Load("./data")
	if err != nil {
		panic(err)
	}

	//imageData := make([]byte, len(traininingSet.Images[1111]))
	//imageData = traininingSet.Images[15344]
	////printing base value for first image
	//fmt.Println(ImageString(imageData, 28, 28))


	//getting all digit frequency so that they can be processed
	digitFrequency, _ := getDigitFrequency(traininingSet.Labels, traininingSet.Images)

	//iterating throught all digit frequency and calculating sum (which can be user later to calculate mean and standard deviation)
	sum := 0
	for i := range digitFrequency {
		sum += digitFrequency[i]
	}
	fmt.Println("digitFrequency -",digitFrequency)
	fmt.Println("Total image given in the data set  = ", sum)

	//take mean from each pixel form each image point
	//get all pixel sum from the images found in training set
	inputDigit := getAllPixelSum(traininingSet.Images)

	//based on the input argument pring mean od standard deviation
	if typeOf == 0 {
		fmt.Println("Calculating Mean for the image ")
		printMeanImage(inputDigit, digitFrequency[digitEntered])
	} else {
		fmt.Println("Calculating standard Deviation ")
		printStandardDeviationImage(inputDigit, digitFrequency[digitEntered])
	}
	return
}

//getAllPixelSum - returns all pixel of the matrix
func getAllPixelSum(currentImage []GoMNIST.RawImage) [][]uint8 {

	//2D matrix to store mean of each pixels of all occurences of zeroes in the data set
	allPixelSum := make([][]uint8, 28)
	for i := range allPixelSum {
		allPixelSum[i] = make([]uint8, 28)
	}

	//initializing all zero digit pixels to zero
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			allPixelSum[0][0] = 0
		}
	}

	//looping through the current image and filling resultant matrix
	for _, val := range currentImage {
		for i := 0; i < 28; i++ {
			for j := 0; j < 28; j++ {
				r, _, _, _ := val.At(i, j).RGBA()
				//applying mask to the pixel value
				//Reference : [https://stackoverflow.com/questions/10493411/what-is-bit-masking ,https://www.mathworks.com/help/images/ref/im2uint8.html]
				red := uint8(r >> 8)

				//get the sum (taking single pixel for the sum)
				allPixelSum[i][j] += red
			}
		}
	}

	return allPixelSum
}

//printMeanImage - prints mean image
//Ref : [https://stats.stackexchange.com/questions/20744/mean-and-standard-deviation-of-gaussian-distribution]
func printMeanImage(currentImageByte [][]uint8, numberOfImages int) interface{} {

	//looping through pixels and dividing with number of image possible for that digit
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			fmt.Print(currentImageByte[i][j]/byte(numberOfImages), " ")
		}
		fmt.Println()
	}
	return nil
}

//printStandardDeviationImage - prints Standard Deviation Image
func printStandardDeviationImage(currentImageByte [][]uint8, numberOfImages int) interface{} {
	//looping through pixels and calculating standard deviation by formula
	//Ref : [https://stats.stackexchange.com/questions/20744/mean-and-standard-deviation-of-gaussian-distribution ]
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			mean := currentImageByte[i][j] / byte(numberOfImages)

			asFloatMean := float64(mean)
			asFloatCurrentVal := float64(currentImageByte[i][j])

			temp := math.Pow(asFloatCurrentVal-asFloatMean, 2) / float64(numberOfImages)

			fmt.Print(math.Sqrt(temp), " ")
		}
		fmt.Println()
	}
	return nil
}

//getDigitFrequency - calculates all digit frequency present in mnist image
func getDigitFrequency(imageLabels []GoMNIST.Label, images []GoMNIST.RawImage) ([]int, [][]GoMNIST.RawImage) {

	//digit array contains count of occurrence of each digit
	digitArray := make([]int, 10)

	//fill digit array
	for i := 0; i <= 9; i++ {
		digitArray[i] = i
	}

	//Create all zeroes, this array contains all the possible digit values from 0 to 9
	allZeroes := []GoMNIST.RawImage{}
	allOnes := []GoMNIST.RawImage{}
	allTwos := []GoMNIST.RawImage{}
	allThrees := []GoMNIST.RawImage{}
	allFours := []GoMNIST.RawImage{}
	allFives := []GoMNIST.RawImage{}
	allSixes := []GoMNIST.RawImage{}
	allSevens := []GoMNIST.RawImage{}
	allEighths := []GoMNIST.RawImage{}
	allNines := []GoMNIST.RawImage{}

	//return value
	allDigits := [][]GoMNIST.RawImage{}

	//looping through all labels and calculating the frequency of each digit
	for _, val := range imageLabels {
		//calculating number of digits
		if val == 0 {
			digitArray[val]++
			//concatenating all images present to form a feature
			allZeroes = append(allZeroes, images[val])
		} else if val == 1 {
			digitArray[val]++
			allOnes = append(allOnes, images[val])
		} else if val == 2 {
			digitArray[val]++
			allTwos = append(allTwos, images[val])
		} else if val == 3 {
			digitArray[val]++
			allThrees = append(allThrees, images[val])
		} else if val == 4 {
			digitArray[val]++
			allFours = append(allFours, images[val])
		} else if val == 5 {
			digitArray[val]++
			allFives = append(allFives, images[val])
		} else if val == 6 {
			digitArray[val]++
			allSixes = append(allSixes, images[val])
		} else if val == 7 {
			digitArray[val]++
			allSevens = append(allSevens, images[val])
		} else if val == 8 {
			digitArray[val]++
			allEighths = append(allEighths, images[val])
		} else if val == 9 {
			digitArray[val]++
			allNines = append(allNines, images[val])
		}
	}

	//appending all digit sum values to allDigits 2D array of images
	allDigits = append(allDigits, allZeroes, allOnes, allTwos, allThrees, allFours, allFives, allSixes, allSevens, allEighths, allNines)
	return digitArray, allDigits
}

//ImageString method which converts image byte to readable string format (which is to be processed by matlab)
func ImageString(buffer []byte, height, width int) (out string) {
	for i, y := 0, 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if buffer[i] > 128 {
				asInt := int(buffer[i])
				out += strconv.Itoa(asInt) + " "
			} else {
				out += "0 " //0 for black color
			}
			i++
		}
		out += "\n"
	}
	return
}

//Get01LossProbability - Gets loss probability using 0-1 loss function
func Get01LossProbability(position int) float64 {
	_, testingDataSet, err := GoMNIST.Load("./data")
	if err != nil {
		panic(err)
	}

	//getting DigitFrequency of testing dataset (("t10k-images-idx3-ubyte.gz"))
	digitFrequency, _ := getDigitFrequency(testingDataSet.Labels, testingDataSet.Images)

	totalImages := 0
	for _, val := range digitFrequency {
		totalImages += val
	}

	//loss probability result
	var lossProbabiltyResult []float64
	for _, val := range digitFrequency {
		x := float64(val) / float64(totalImages)
		lossProbabiltyResult = append(lossProbabiltyResult, float64(x)*100)
	}

	return lossProbabiltyResult[position] / 100
}

//BayesianDecisionClassification - function to do classification of imaget using bayesian decision rule
func BayesianDecisionClassification() {
	/*
		Algorithm :
			1. take random number (0 to 9)
			2. pass this to 10 discriminant functions
			3. get maximum value
			4. discriminant function with max is the expected number
			5. calculate accuracy (by comparing with expected label) by running test case multiple time (something like benchmark)
	*/

	//loadind testing data ("t10k-images-idx3-ubyte.gz")
	_, testingDataSet, err := GoMNIST.Load("./data")
	if err != nil {
		panic(err)
	}

	//calculating digit frequency for all testing data
	digitFrequency, _ := getDigitFrequency(testingDataSet.Labels, testingDataSet.Images)

	totalImages := 0
	for _, val := range digitFrequency {
		totalImages += val
	}

	//get random label
	allLabels := testingDataSet.Labels
	rand.Seed(time.Now().Unix())
	testingNumber := allLabels[rand.Intn(len(allLabels))]

	testResults := make([]float64, 10)

	//discriminant function for all digits from 0 to 9
	testResults[0] = getLabelsByDigit(0, testingNumber)
	testResults[1] = getLabelsByDigit(1, testingNumber)
	testResults[2] = getLabelsByDigit(2, testingNumber)
	testResults[3] = getLabelsByDigit(3, testingNumber)
	testResults[4] = getLabelsByDigit(4, testingNumber)
	testResults[5] = getLabelsByDigit(5, testingNumber)
	testResults[6] = getLabelsByDigit(6, testingNumber)
	testResults[7] = getLabelsByDigit(7, testingNumber)
	testResults[8] = getLabelsByDigit(8, testingNumber)
	testResults[9] = getLabelsByDigit(9, testingNumber)

	//running test case
	for i := 1; i <= 5; i++ {
		n := 10
		fmt.Println("Running ", i, " test case ", n, "times")
		runTestCases(n)
	}
}

//runTestCases - runs test case t times for testing dataset
func runTestCases(t int) {

	_, testingDataSet, err := GoMNIST.Load("./data")
	if err != nil {
		panic(err)
	}

	//number of times a hit is found (used to calculate accuracy)
	hits := 0

	//looping for all test cases
	for i := 1; i <= t; i++ {
		fmt.Println("test case # ", i)
		//get random label
		allLabels := testingDataSet.Labels
		rand.Seed(time.Now().Unix())
		testingNumber := allLabels[rand.Intn(len(allLabels))]

		testingNumberInt := (testingNumber)

		testResults := make([]float64, 10)
		testResults[0] = getLabelsByDigit(0, testingNumber)
		testResults[1] = getLabelsByDigit(1, testingNumber)
		testResults[2] = getLabelsByDigit(2, testingNumber)
		testResults[3] = getLabelsByDigit(3, testingNumber)
		testResults[4] = getLabelsByDigit(4, testingNumber)
		testResults[5] = getLabelsByDigit(5, testingNumber)
		testResults[6] = getLabelsByDigit(6, testingNumber)
		testResults[7] = getLabelsByDigit(7, testingNumber)
		testResults[8] = getLabelsByDigit(8, testingNumber)
		testResults[9] = getLabelsByDigit(9, testingNumber)

		//getting maximum from the all discriminant functions)
		value, _ := getMax(testResults)
		if int(testingNumberInt) == value {
			hits++
		}
	}

	result := float32(t-hits) / float32(t) //calculating the accuracy fraction
	fmt.Println("Accuracy := ", result*100, " %")
}

//getMax - get maximum value from array
func getMax(array []float64) (int, float64) {
	max := array[0]
	var keyR int
	for key, val := range array {
		if val > max {
			max = val
			keyR = key
		}
	}
	return keyR, max
}

//getLabelsByDigit - returns label for a particular digit
func getLabelsByDigit(digit, testingNumber GoMNIST.Label) float64 {
	_, testingDataSet, err := GoMNIST.Load("./data")
	if err != nil {
		panic(err)
	}

	count := 0

	var keyArray []int
	for key, val := range testingDataSet.Labels {
		if val == digit {
			keyArray = append(keyArray, key)
			count++
		}
	}

	//loop through all images and append the labels found in keuArray to get all images corresponding to all keys
	//creating empty image set to store images corrresponding to the digit passed as parameter
	finalImageResult := GoMNIST.Set{}.Images

	i := 0
	for i < len(keyArray) {
		finalImageResult = append(finalImageResult, testingDataSet.Images[keyArray[i]])
		i++
	}

	//
	meanMatrix := getAllPixelSum(finalImageResult)
	var sum float64
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			sum += float64(meanMatrix[i][j])
		}
	}

	//determinant of the matrix
	matrixMod := math.Abs(sum)

	//mean
	meanVal := sum / float64(count)

	//variance
	var varianceSum float64
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			varianceSum += float64(math.Pow(float64(meanMatrix[i][j]), 2))
		}
	}
	variance := varianceSum / float64(count)

	//get Standard Deviation
	standardDeviation := math.Sqrt(variance)

	//calculating discriminant value formula
	correctDiscriminantValue := calculateDiscriminant(meanVal, standardDeviation, float64(digit), float64(testingNumber), matrixMod, Get01LossProbability(int(digit)))

	return correctDiscriminantValue
}

/*
calculateDiscriminant - discriminant function
we define our loss function is based on 0-1 loss function and discriminant function is based on Bayesian decision.
Since, mean and covariance are different, we assume that digits in all labels follow multivariate gaussian distribution.
	gi(x) = x^t * Wi + wi^t * x + wi0
	Wi = -1/2 * (sd)^-1
	wi = sd * mean
	wi0 = -1/2 * mean^t * (sd)^-1 * mean - 1/2 * ln(|matrix mod|) + ln(pi)
*/
func calculateDiscriminant(mean, sd, x, t, matrixmod, pi float64) float64 {
	term1 := math.Pow(x, t) * -1 / 2 * math.Pow(sd, -1)
	term2 := sd * mean
	term3 := -1/2*math.Pow(mean, t)*math.Pow(sd, -1)*mean*-1/2*math.Log2(matrixmod) + math.Log2(pi)
	return term1 + term2 + term3
}
