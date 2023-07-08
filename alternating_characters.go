package main

/*
You are given a string containing characters and

only. Your task is to change it into a string such that there are no matching adjacent characters. To do this, you are allowed to delete zero or more characters in the string.

Your task is to find the minimum number of required deletions.

Example
Remove an at positions and to make in

deletions.

Function Description

Complete the alternatingCharacters function in the editor below.

alternatingCharacters has the following parameter(s):

    string s: a string

Returns

    int: the minimum number of deletions required

Input Format

The first line contains an integer
, the number of queries.
The next lines each contain a string

to analyze.

Constraints

Each string will consist only of characters and

    .

Sample Input

5
AAAA
BBBBB
ABABABAB
BABABA
AAABBB

Sample Output

3
4
0
0
4

Explanation

The characters marked red are the ones that can be deleted so that the string does not have matching adjacent characters.

image
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'alternatingCharacters' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func alternatingCharacters(s string) int32 {
	var oldChar rune
	var deleted int32
	for _, newChar := range s {
		if oldChar == newChar {
			deleted++
		}
		oldChar = newChar
	}
	return deleted
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := alternatingCharacters(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}
