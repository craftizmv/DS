package arrays

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getTheGroups' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY queryType
 *  3. INTEGER_ARRAY students1
 *  4. INTEGER_ARRAY students2
 */

func getTheGroups(n int32, queryType []string, students1 []int32, students2 []int32) []int32 {
	// Write your code here
	// we need to find the max cluster size which we can form using the friends query type, and we can count distinct type.
	// n = distint number for persons

	// Add error cases

	maxFriendClusterMap := make([]map[int32]struct{}, 0)

	for i := range queryType {
		query := queryType[i]

		if query == "Friend" {

			if i == 0 {
				// seed val
				maxFriendClusterMap = append(maxFriendClusterMap, map[int32]struct{}{
					students2[i]: {},
					students1[i]: {},
				})
			}

			for i2 := range maxFriendClusterMap {
				fCluster := maxFriendClusterMap[i2]

				// If that student exist in the map, add its friend
				if _, ok := fCluster[students1[i]]; ok {
					fCluster[students2[i]] = struct{}{}
				}

				// If that student exist in the map, add its friend
				if _, ok := fCluster[students2[i]]; ok {
					fCluster[students1[i]] = struct{}{}
				}
			}

		}

		if query == "Total" {
			// Here we have to calculate the max cluster size and return the answer.
			var maxClusterSize int32 = 0
			for i3 := range maxFriendClusterMap {
				fCluster := maxFriendClusterMap[i3]
				cLen := len(fCluster)
				if int32(cLen) > maxClusterSize {
					maxClusterSize = int32(cLen)
				}
			}

			// distinct groups
			var dGroup = n - maxClusterSize
			val := math.Max(float64(maxClusterSize+dGroup), float64(maxClusterSize))

			return []int32{int32(val)}
		}
	}

	return []int32{}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	queryTypeCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queryType []string

	for i := 0; i < int(queryTypeCount); i++ {
		queryTypeItem := readLine(reader)
		queryType = append(queryType, queryTypeItem)
	}

	students1Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var students1 []int32

	for i := 0; i < int(students1Count); i++ {
		students1ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		students1Item := int32(students1ItemTemp)
		students1 = append(students1, students1Item)
	}

	students2Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var students2 []int32

	for i := 0; i < int(students2Count); i++ {
		students2ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		students2Item := int32(students2ItemTemp)
		students2 = append(students2, students2Item)
	}

	result := getTheGroups(n, queryType, students1, students2)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
