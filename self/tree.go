package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fullTree := FullTreeInit(arr)

	middleTreePrint(fullTree)
}

func firstTreePrint(root *Tree) {
	q := newExQueue()
	q.lPush(root)
	for q.length > 0 {

		t := q.lPop()
		fmt.Println(t.data)

		if t.right != nil {
			q.lPush(t.right)
		}

		if t.left != nil {
			q.lPush(t.left)
		}
	}
}

func middleTreePrint(root *Tree) {
	q := newExQueue()

	for root != nil || q.length > 0 {

		if root != nil {
			q.lPush(root)
			root = root.left

		} else {
			root = q.lGet()
			fmt.Println(root.data)
			q.lPop()

			root = root.right
		}

	}
}

type Tree struct {
	data  int
	left  *Tree
	right *Tree
}

func newTree(n int) *Tree {
	return &Tree{n, nil, nil}
}

func FullTreeInit(arr []int) *Tree {
	q := newExQueue()
	var root *Tree
	for _, v := range arr {

		t := q.lPop()
		if t == nil {
			t = newTree(v)
			q.lPush(t)

			root = t
			continue
		}
		if t.left == nil {
			l := newTree(v)
			t.left = l

			q.lPush(t)
			q.rPush(l)
		} else if t.right == nil {
			r := newTree(v)
			t.right = r

			q.rPush(r)
		}
	}

	return root
}

type exQueue struct {
	buf    []*Tree
	length int
}

func newExQueue() *exQueue {
	return &exQueue{make([]*Tree, 0), 0}
}

func (q *exQueue) rPush(e *Tree) {
	q.buf = append(q.buf, e)
	q.length += 1
}

func (q *exQueue) rPushs(elements ...*Tree) {
	q.buf = append(q.buf, elements...)
	q.length += len(elements)
}

func (q *exQueue) lPush(e *Tree) {
	q.buf = append([]*Tree{e}, q.buf...)
	q.length += 1
}

func (q *exQueue) rPop() *Tree {

	if q.length == 0 {
		return nil
	}

	n := q.buf[q.length-1]

	q.buf = q.buf[:q.length-1]
	q.length -= 1

	return n
}

func (q *exQueue) lPop() *Tree {

	if q.length == 0 {
		return nil
	}

	n := q.buf[0]

	q.buf = q.buf[1:]
	q.length -= 1

	return n
}

func (q *exQueue) lGet() *Tree {
	if q.length == 0 {
		return nil
	}

	return q.buf[0]
}

func (q *exQueue) rGet() *Tree {
	if q.length == 0 {
		return nil
	}

	return q.buf[q.length-1]
}
