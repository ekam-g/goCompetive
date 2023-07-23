package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

/*
 * Complete the 'reverse' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts INTEGER_SINGLY_LINKED_LIST llist as parameter.
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	var values []int32
	for ; llist != nil; llist = llist.next {
		values = append(values, llist.data)
	}
	fmt.Println(values)
	head := &SinglyLinkedListNode{}
	iterPointer := head
	for x := len(values) - 1; x >= 0; x-- {
		fmt.Println(values[x])
		iterPointer.data = values[x]
		if 0 != x {
			iterPointer.next = &SinglyLinkedListNode{}
			iterPointer = iterPointer.next
		}
	}
	return head
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	testsTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tests := int32(testsTemp)

	for testsItr := 0; testsItr < int(tests); testsItr++ {
		llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		llist := SinglyLinkedList{}
		for i := 0; i < int(llistCount); i++ {
			llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
			checkError(err)
			llistItem := int32(llistItemTemp)
			llist.insertNodeIntoSinglyLinkedList(llistItem)
		}

		llist1 := reverse(llist.head)

		printSinglyLinkedList(llist1, " ", writer)
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}
