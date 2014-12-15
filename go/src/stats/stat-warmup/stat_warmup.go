package main

import (
	"fmt"
	"strings"
	"log"
	"strconv"
	"bufio"
	"os"
	"sort"
	"errors"
	"math"
)

type Stats struct {
	Mean float64
	Mode int64
	Median float64
	StandardDev float64
	MaxConfInterval float64
	MinConfInterval float64
}

func createStats() *Stats {
	return &Stats {
		Mean: 0.0,
		Mode: 0,
		Median: 0.0,
		StandardDev: 0.0,
		MaxConfInterval: 0.0,
		MinConfInterval: 0.0,	
	}
}

func (stats * Stats) Print() string {
	return fmt.Sprintf(
		"%.1f\n%.1f\n%d\n%.1f\n%.1f %.1f",
		stats.Mean,
		stats.Median,
		stats.Mode,		
		stats.StandardDev,
		stats.MaxConfInterval,
		stats.MinConfInterval)
}
func (stats *Stats) FindMedian(data []int) float64 {
	var length int = len(data)

	if length % 2 == 0 {
		return float64(data[length/2-1]+data[length/2]) / 2.0		
	} else {
		return float64(data[length/2]) / 2.0
	}
	
}
func (stats *Stats) FindMode(counts map[int]int) (int64, error) {

	//log.Println("Counts:", counts)

	makeIntArray := func() []int {
		return make([]int,0)
	}

	var data  = makeIntArray()
	var freqs = makeIntArray()
	var maxFreq int = -1
	
	for key,val := range counts {

		switch {
		case maxFreq == val: freqs = append(freqs, key)
		case maxFreq  < val:
			maxFreq = val
			freqs = makeIntArray()
			freqs = append(freqs, key)
		}

		data = append(data, key)
				
		//log.Println(key, val, "MaxFreq:",maxFreq,"Freqs:",freqs)		
	}

	if len(freqs) > 0 {
		// Found at lest one max freq
		sort.Ints(freqs)
		return int64(freqs[0]), nil
	} else {
		return -1, errors.New("No max freq found, all equal")
	}
	
	
}

func calculate(data []int) *Stats {

	var stats *Stats = createStats()
	counts := make(map[int]int)
	subMeanData := make([]float64, len(data))
	
	if len(data) == 0 {
		return stats
	}

	var lengthf float64 = float64(len(data))
	var sum float64  = 0.0

	// Sort data
	sort.Ints(data)

	// Initial sum and counts
	for _,v := range data {
		val := float64(v)
		sum += val
		// Record counts
		if foundVal, ok := counts[v]; ok {
			counts[v] = foundVal + 1 //1+ Occurs
		} else {
			counts[v] = 1 // First Occurs
		}
		
	}
	
	stats.Mean = sum / lengthf
	stats.Median = stats.FindMedian(data)
	if smode, err := stats.FindMode(counts); err != nil {
		stats.Mode = int64(data[0])
	} else {
		stats.Mode = smode
	}

	
	var subMeanSum float64 = 0.0
	for _,v := range data {
		val := float64(v)
		res := val - stats.Mean
		res2 := res*res
		subMeanSum += res2		
	}
	stats.StandardDev = math.Sqrt(subMeanSum/lengthf)

	log.Println(subMeanData)
	
	return stats
}

func main() {
	testCases, err := parseHackerRank()

	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Stats: %+v\n", calculate(testCases))
	var stats *Stats = calculate(testCases)
	fmt.Println(stats.Print())
}


//////////////////////////////////////////////////////
// Reading/Parsing ///////////////////////////////////

func ParseArray(ref []int) error {
	// Read array line
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		return err
	}
	// Separate each value
	strs := strings.Split(string(line)," ")

	// For each value, split
	for i,v := range strs {
		if res, err := strconv.Atoi(v); err != nil {
			return err
		} else {
			ref[i] = res 
		}
	}

	return nil	
}

func ParseInt(ref *int) error {
	_, err := fmt.Scanf("%d", ref)
	return err
}

func parseHackerRank() ([]int, error) {
	var testCount int
	if err := ParseInt(&testCount); err != nil {
		return nil, err
	}
	// Alloc memory for test cases and fill
	testCases := make([]int, testCount)
	ParseArray(testCases)
	return testCases, nil
}
