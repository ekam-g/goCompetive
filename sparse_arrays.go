package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
There is a collection of input strings and a collection of query strings. For each query string, determine how many times it occurs in the list of input strings. Return an array of the results.

Example


There are instances of '', of '' and of ''. For each query, add an element to the return array,

.

Function Description

Complete the function matchingStrings in the editor below. The function must return an array of integers representing the frequency of occurrence of each query string in stringList.

matchingStrings has the following parameters:

    string stringList[n] - an array of strings to search
    string queries[q] - an array of query strings

Returns

    int[q]: an array of results for each query

Input Format

The first line contains and integer
, the size of .
Each of the next lines contains a string .
The next line contains , the size of .
Each of the next lines contains a string

.

Constraints
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY stringList
 *  2. STRING_ARRAY queries
*/

func matchingStrings(stringList []string, queries []string) []int32 {
	// Write your code here
	total := make([]int32, len(queries))
	for _, word := range stringList {
		for n, query := range queries {
			if word == query {
				total[n]++
			}
		}
	}
	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stringListCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var stringList []string

	for i := 0; i < int(stringListCount); i++ {
		stringListItem := readLine(reader)
		stringList = append(stringList, stringListItem)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(stringList, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}
