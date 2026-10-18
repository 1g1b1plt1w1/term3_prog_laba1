package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command int

const (
	PRINT Command = iota

	MCREATE
	MPUSH
	MPUSHIDX
	MGET
	MDELIDX
	MSETIDX
	MLENGTH
	MREAD

	FPUSH
	FPUSHHEAD
	FPUSHBEFORE
	FPUSHAFTER
	FPOP
	FPOPHEAD
	FPOPBEFORE
	FPOPAFTER
	FDELVAL
	FFINDVAL
	FREADHT
	FREADTH

	LPUSHHEAD
	LPUSH
	LPUSHBEFORE
	LPUSHAFTER
	LPOPHEAD
	LPOP
	LPOPBEFORE
	LPOPAFTER
	LDELVAL
	LFINDVAL
	LREADHT
	LREADTH

	SPUSH
	SPOP
	SREAD

	QPUSH
	QPOP
	QREAD
)

var commandMap = map[string]Command{
	"PRINT": PRINT,

	"MCREATE":  MCREATE,
	"MPUSH":    MPUSH,
	"MPUSHIDX": MPUSHIDX,
	"MGET":     MGET,
	"MGETIDX":  MGET,
	"MDELIDX":  MDELIDX,
	"MSETIDX":  MSETIDX,
	"MLENGTH":  MLENGTH,
	"MREAD":    MREAD,

	"FPUSH":       FPUSH,
	"FPUSHTAIL":   FPUSH,
	"FPUSHHEAD":   FPUSHHEAD,
	"FPUSHBEFORE": FPUSHBEFORE,
	"FPUSHAFTER":  FPUSHAFTER,
	"FPOP":        FPOP,
	"FPOPTAIL":    FPOP,
	"FPOPHEAD":    FPOPHEAD,
	"FPOPBEFORE":  FPOPBEFORE,
	"FPOPAFTER":   FPOPAFTER,
	"FDELVAL":     FDELVAL,
	"FFINDVAL":    FFINDVAL,
	"FREADHT":     FREADHT,
	"FREADTH":     FREADTH,

	"LPUSHHEAD":   LPUSHHEAD,
	"LPUSH":       LPUSH,
	"LPUSHTAIL":   LPUSH,
	"LPUSHBEFORE": LPUSHBEFORE,
	"LPUSHAFTER":  LPUSHAFTER,
	"LPOPHEAD":    LPOPHEAD,
	"LPOP":        LPOP,
	"LPOPTAIL":    LPOP,
	"LPOPBEFORE":  LPOPBEFORE,
	"LPOPAFTER":   LPOPAFTER,
	"LDELVAL":     LDELVAL,
	"LFINDVAL":    LFINDVAL,
	"LREADHT":     LREADHT,
	"LREADTH":     LREADTH,

	"SPUSH": SPUSH,
	"SPOP":  SPOP,
	"SREAD": SREAD,

	"QPUSH": QPUSH,
	"QPOP":  QPOP,
	"QREAD": QREAD,
}

type Array struct {
	data []string
	cap  int
}

func NewArray(capacity int) Array {
	if capacity < 16 {
		capacity = 16
	}
	return Array{
		data: make([]string, 0, capacity),
		cap:  capacity,
	}
}

func (a *Array) PushBack(v string) {
	a.data = append(a.data, v)
}

func (a *Array) Insert(idx int, v string) {
	if idx < 0 || idx > len(a.data) {
		return
	}
	a.data = append(a.data, "")
	copy(a.data[idx+1:], a.data[idx:])
	a.data[idx] = v
}

func (a *Array) Get(idx int) (string, error) {
	if idx < 0 || idx >= len(a.data) {
		return "", fmt.Errorf("Array_get: index out of range")
	}
	return a.data[idx], nil
}

func (a *Array) Set(idx int, v string) {
	if idx < 0 || idx >= len(a.data) {
		return
	}
	a.data[idx] = v
}

func (a *Array) Erase(idx int) {
	if idx < 0 || idx >= len(a.data) {
		return
	}
	copy(a.data[idx:], a.data[idx+1:])
	a.data = a.data[:len(a.data)-1]
}

func (a *Array) Length() int {
	return len(a.data)
}

func (a *Array) Read() string {
	var sb strings.Builder
	for i := 0; i < len(a.data); i++ {
		sb.WriteString(" ")
		sb.WriteString(a.data[i])
	}
	return sb.String()
}

type OneNode struct {
	data string
	next *OneNode
}

type OneLinkedList struct {
	head *OneNode
	tail *OneNode
}

func NewOneLinkedList() OneLinkedList {
	return OneLinkedList{}
}

func (l *OneLinkedList) PushFront(v string) {
	n := &OneNode{data: v}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	n.next = l.head
	l.head = n
}

func (l *OneLinkedList) PushBack(v string) {
	n := &OneNode{data: v}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.next = n
	l.tail = n
}

func (l *OneLinkedList) PushAfter(prevVal, v string) {
	if l.head == nil {
		l.PushBack(v)
		return
	}
	prev := l.head
	for prev != nil && prev.data != prevVal {
		prev = prev.next
	}
	if prev == nil {
		return
	}
	n := &OneNode{data: v, next: prev.next}
	prev.next = n
	if prev == l.tail {
		l.tail = n
	}
}

func (l *OneLinkedList) PushBefore(nextVal, v string) {
	if l.head == nil {
		return
	}
	if l.head.data == nextVal {
		l.PushFront(v)
		return
	}
	prev := l.head
	for prev != nil && prev.next != nil && prev.next.data != nextVal {
		prev = prev.next
	}
	if prev == nil || prev.next == nil {
		return
	}
	n := &OneNode{data: v, next: prev.next}
	prev.next = n
}

func (l *OneLinkedList) PopFront() string {
	if l.head == nil {
		return ""
	}
	val := l.head.data
	l.head = l.head.next
	if l.head == nil {
		l.tail = nil
	}
	return val
}

func (l *OneLinkedList) PopBack() string {
	if l.head == nil {
		return ""
	}
	val := l.tail.data
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return val
	}
	cur := l.head
	for cur.next != l.tail {
		cur = cur.next
	}
	cur.next = nil
	l.tail = cur
	return val
}

func (l *OneLinkedList) PopAfter(prevVal string) string {
	if l.head == nil {
		return ""
	}
	prev := l.head
	for prev != nil && prev.data != prevVal {
		prev = prev.next
	}
	if prev == nil || prev.next == nil {
		return ""
	}
	del := prev.next
	val := del.data
	prev.next = del.next
	if del == l.tail {
		l.tail = prev
	}
	return val
}

func (l *OneLinkedList) PopBefore(nextVal string) string {
	if l.head == nil {
		return ""
	}
	if l.head.data == nextVal {
		return ""
	}
	if l.head.next != nil && l.head.next.data == nextVal {
		return l.PopFront()
	}
	prev := l.head
	for prev != nil && prev.next != nil && prev.next.next != nil && prev.next.next.data != nextVal {
		prev = prev.next
	}
	if prev == nil || prev.next == nil || prev.next.next == nil {
		return ""
	}
	del := prev.next
	val := del.data
	prev.next = del.next
	return val
}

func (l *OneLinkedList) DelByValue(val string) {
	if l.head == nil {
		return
	}
	if l.head.data == val {
		l.PopFront()
		return
	}
	cur := l.head
	for cur.next != nil && cur.next.data != val {
		cur = cur.next
	}
	if cur.next != nil {
		temp := cur.next
		cur.next = temp.next
		if temp == l.tail {
			l.tail = cur
		}
	}
}

func (l *OneLinkedList) FindByValue(val string) int {
	if l.head == nil {
		return -1
	}
	count := 0
	cur := l.head
	for cur != nil && cur.data != val {
		cur = cur.next
		count++
	}
	if cur != nil {
		return count
	}
	return -1
}

func (l *OneLinkedList) ReadHT() string {
	if l.head == nil {
		return ""
	}
	var sb strings.Builder
	cur := l.head
	for cur != l.tail {
		sb.WriteString(" ")
		sb.WriteString(cur.data)
		cur = cur.next
	}
	sb.WriteString(" ")
	sb.WriteString(cur.data)
	return sb.String()
}

func (l *OneLinkedList) ReadTH() string {
	if l.head == nil {
		return ""
	}
	var temp []string
	cur := l.head
	for cur != l.tail {
		temp = append(temp, cur.data)
		cur = cur.next
	}
	temp = append(temp, cur.data)
	var sb strings.Builder
	for i := len(temp) - 1; i >= 0; i-- {
		sb.WriteString(" ")
		sb.WriteString(temp[i])
	}
	return sb.String()
}

type Node struct {
	data string
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (l *LinkedList) PushFront(v string) {
	n := &Node{data: v, next: l.head}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.head.prev = n
	l.head = n
}

func (l *LinkedList) PushBack(v string) {
	n := &Node{data: v, prev: l.tail}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.next = n
	l.tail = n
}

func (l *LinkedList) PushBefore(nextEl, v string) {
	if l.head == nil {
		return
	}
	if nextEl == l.head.data {
		l.PushFront(v)
		return
	}
	cur := l.head
	for cur != nil && cur.data != nextEl {
		cur = cur.next
	}
	if cur == nil {
		return
	}
	prevEl := cur.prev
	newNode := &Node{data: v, next: cur, prev: prevEl}
	prevEl.next = newNode
	cur.prev = newNode
}

func (l *LinkedList) PushAfter(prevEl, v string) {
	if l.tail == nil {
		return
	}
	if prevEl == l.tail.data {
		l.PushBack(v)
		return
	}
	cur := l.head
	for cur != nil && cur.data != prevEl {
		cur = cur.next
	}
	if cur == nil {
		return
	}
	nextEl := cur.next
	newNode := &Node{data: v, next: nextEl, prev: cur}
	cur.next = newNode
	nextEl.prev = newNode
}

func (l *LinkedList) PopFront() string {
	if l.head == nil {
		return ""
	}
	val := l.head.data
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return val
	}
	l.head = l.head.next
	l.head.prev = nil
	return val
}

func (l *LinkedList) PopBack() string {
	if l.tail == nil {
		return ""
	}
	val := l.tail.data
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return val
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	return val
}

func (l *LinkedList) PopBefore(nextEl string) string {
	cur := l.head
	for cur != nil && cur.data != nextEl {
		cur = cur.next
	}
	if cur == nil || cur == l.head {
		return ""
	}
	del := cur.prev
	if del == l.head {
		return l.PopFront()
	}
	val := del.data
	prevEl := del.prev
	prevEl.next = cur
	cur.prev = prevEl
	return val
}

func (l *LinkedList) PopAfter(prevEl string) string {
	cur := l.head
	for cur != nil && cur.data != prevEl {
		cur = cur.next
	}
	if cur == nil || cur == l.tail {
		return ""
	}
	del := cur.next
	if del == l.tail {
		return l.PopBack()
	}
	val := del.data
	nextEl := del.next
	cur.next = nextEl
	nextEl.prev = cur
	return val
}

func (l *LinkedList) DelByValue(val string) {
	if l.head == nil {
		return
	}
	cur := l.head
	for cur != nil && cur.data != val {
		cur = cur.next
	}
	if cur == nil {
		return
	}
	if cur == l.head {
		l.PopFront()
		return
	}
	if cur == l.tail {
		l.PopBack()
		return
	}
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
}

func (l *LinkedList) Find(val string) int {
	cur := l.head
	count := 0
	for cur != nil && cur.data != val {
		cur = cur.next
		count++
	}
	if cur == nil {
		return -1
	}
	return count
}

func (l *LinkedList) ReadHT() string {
	var sb strings.Builder
	sb.WriteString(" ")
	cur := l.head
	for cur != nil {
		sb.WriteString(cur.data)
		sb.WriteString(" ")
		cur = cur.next
	}
	return strings.TrimRight(sb.String(), " ")
}

func (l *LinkedList) ReadTH() string {
	var sb strings.Builder
	sb.WriteString(" ")
	cur := l.tail
	for cur != nil {
		sb.WriteString(cur.data)
		sb.WriteString(" ")
		cur = cur.prev
	}
	return strings.TrimRight(sb.String(), " ")
}

type Stack struct {
	list LinkedList
}

func NewStack() Stack {
	return Stack{list: NewLinkedList()}
}

func (s *Stack) PushBack(v string) {
	s.list.PushBack(v)
}

func (s *Stack) Pop() string {
	return s.list.PopBack()
}

func (s *Stack) Read() string {
	return s.list.ReadTH()
}

type Queue struct {
	list LinkedList
}

func NewQueue() Queue {
	return Queue{list: NewLinkedList()}
}

func (q *Queue) PushBack(v string) {
	q.list.PushBack(v)
}

func (q *Queue) Pop() string {
	return q.list.PopFront()
}

func (q *Queue) Read() string {
	return q.list.ReadHT()
}

var (
	arrays    = map[string]Array{}
	onellists = map[string]OneLinkedList{}
	llists    = map[string]LinkedList{}
	queues    = map[string]Queue{}
	stacks    = map[string]Stack{}
)

func writeState(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Не удалось открыть файл при записи")
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for name, obj := range arrays {
		fmt.Fprintf(w, "Array %s%s\n", name, obj.Read())
	}
	for name, obj := range onellists {
		fmt.Fprintf(w, "OneLinkedList %s%s\n", name, obj.ReadHT())
	}
	for name, obj := range llists {
		fmt.Fprintf(w, "LinkedList %s%s\n", name, obj.ReadHT())
	}
	for name, obj := range queues {
		fmt.Fprintf(w, "Queue %s%s\n", name, obj.Read())
	}
	for name, obj := range stacks {
		fmt.Fprintf(w, "Stack %s%s\n", name, obj.list.ReadHT())
	}
	fmt.Fprintln(w)
	return w.Flush()
}

func readState(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Не удалось открыть файл при чтении")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		datatype := parts[0]
		name := parts[1]
		vals := parts[2:]

		switch datatype {
		case "Array":
			arr := NewArray(10)
			for _, v := range vals {
				arr.PushBack(v)
			}
			arrays[name] = arr
		case "OneLinkedList":
			onell := NewOneLinkedList()
			for _, v := range vals {
				onell.PushBack(v)
			}
			onellists[name] = onell
		case "LinkedList":
			ll := NewLinkedList()
			for _, v := range vals {
				ll.PushBack(v)
			}
			llists[name] = ll
		case "Queue":
			q := NewQueue()
			for _, v := range vals {
				q.PushBack(v)
			}
			queues[name] = q
		case "Stack":
			s := NewStack()
			for _, v := range vals {
				s.PushBack(v)
			}
			stacks[name] = s
		}
	}
	return scanner.Err()
}

func doCommand(name string, data []string, cmd Command) {
	switch cmd {
	case MCREATE:
		arrays[name] = NewArray(10)
	case MPUSH:
		arr := arrays[name]
		for _, v := range data {
			arr.PushBack(v)
		}
		arrays[name] = arr
	case MPUSHIDX:
		if len(data) >= 2 {
			idx, err := strconv.Atoi(data[0])
			if err == nil {
				arr := arrays[name]
				arr.Insert(idx, data[1])
				arrays[name] = arr
			}
		}
	case MGET:
		if len(data) >= 1 {
			idx, err := strconv.Atoi(data[0])
			if err == nil {
				arr := arrays[name]
				val, err := arr.Get(idx)
				if err == nil {
					fmt.Println(val)
				}
			}
		}
	case MDELIDX:
		if len(data) >= 1 {
			idx, err := strconv.Atoi(data[0])
			if err == nil {
				arr := arrays[name]
				arr.Erase(idx)
				arrays[name] = arr
			}
		}
	case MSETIDX:
		if len(data) >= 2 {
			idx, err := strconv.Atoi(data[0])
			if err == nil {
				arr := arrays[name]
				arr.Set(idx, data[1])
				arrays[name] = arr
			}
		}
	case MLENGTH:
		arr := arrays[name]
		fmt.Printf("Length: %d\n", arr.Length())
	case MREAD:
		arr := arrays[name]
		fmt.Println(arr.Read())

	case FPUSH:
		l := onellists[name]
		for _, v := range data {
			l.PushBack(v)
		}
		onellists[name] = l
	case FPUSHHEAD:
		l := onellists[name]
		for _, v := range data {
			l.PushFront(v)
		}
		onellists[name] = l
	case FPUSHBEFORE:
		if len(data) >= 2 {
			l := onellists[name]
			l.PushBefore(data[0], data[1])
			onellists[name] = l
		}
	case FPUSHAFTER:
		if len(data) >= 2 {
			l := onellists[name]
			l.PushAfter(data[0], data[1])
			onellists[name] = l
		}
	case FPOP:
		l := onellists[name]
		val := l.PopBack()
		onellists[name] = l
		if val != "" {
			fmt.Println(val)
		}
	case FPOPHEAD:
		l := onellists[name]
		val := l.PopFront()
		onellists[name] = l
		if val != "" {
			fmt.Println(val)
		}
	case FPOPAFTER:
		if len(data) >= 1 {
			l := onellists[name]
			val := l.PopAfter(data[0])
			onellists[name] = l
			if val != "" {
				fmt.Println(val)
			}
		}
	case FPOPBEFORE:
		if len(data) >= 1 {
			l := onellists[name]
			val := l.PopBefore(data[0])
			onellists[name] = l
			if val != "" {
				fmt.Println(val)
			}
		}
	case FDELVAL:
		if len(data) >= 1 {
			l := onellists[name]
			l.DelByValue(data[0])
			onellists[name] = l
		}
	case FFINDVAL:
		if len(data) >= 1 {
			l := onellists[name]
			fmt.Println(l.FindByValue(data[0]))
		}
	case FREADHT:
		l := onellists[name]
		fmt.Println(l.ReadHT())
	case FREADTH:
		l := onellists[name]
		fmt.Println(l.ReadTH())

	case LPUSHHEAD:
		l := llists[name]
		for _, v := range data {
			l.PushFront(v)
		}
		llists[name] = l
	case LPUSH:
		l := llists[name]
		for _, v := range data {
			l.PushBack(v)
		}
		llists[name] = l
	case LPOPHEAD:
		l := llists[name]
		val := l.PopFront()
		llists[name] = l
		if val != "" {
			fmt.Println(val)
		}
	case LPOP:
		l := llists[name]
		val := l.PopBack()
		llists[name] = l
		if val != "" {
			fmt.Println(val)
		}
	case LPOPBEFORE:
		if len(data) >= 1 {
			l := llists[name]
			val := l.PopBefore(data[0])
			llists[name] = l
			if val != "" {
				fmt.Println(val)
			}
		}
	case LPOPAFTER:
		if len(data) >= 1 {
			l := llists[name]
			val := l.PopAfter(data[0])
			llists[name] = l
			if val != "" {
				fmt.Println(val)
			}
		}
	case LPUSHBEFORE:
		if len(data) >= 2 {
			l := llists[name]
			l.PushBefore(data[0], data[1])
			llists[name] = l
		}
	case LPUSHAFTER:
		if len(data) >= 2 {
			l := llists[name]
			l.PushAfter(data[0], data[1])
			llists[name] = l
		}
	case LDELVAL:
		if len(data) >= 1 {
			l := llists[name]
			l.DelByValue(data[0])
			llists[name] = l
		}
	case LFINDVAL:
		if len(data) >= 1 {
			l := llists[name]
			fmt.Println(l.Find(data[0]))
		}
	case LREADHT:
		l := llists[name]
		fmt.Println(l.ReadHT())
	case LREADTH:
		l := llists[name]
		fmt.Println(l.ReadTH())

	case SPUSH:
		s := stacks[name]
		for _, v := range data {
			s.PushBack(v)
		}
		stacks[name] = s
	case SPOP:
		s := stacks[name]
		val := s.Pop()
		stacks[name] = s
		if val != "" {
			fmt.Println(val)
		}
	case SREAD:
		s := stacks[name]
		fmt.Println(s.Read())

	case QPUSH:
		q := queues[name]
		for _, v := range data {
			q.PushBack(v)
		}
		queues[name] = q
	case QPOP:
		q := queues[name]
		val := q.Pop()
		queues[name] = q
		if val != "" {
			fmt.Println(val)
		}
	case QREAD:
		q := queues[name]
		fmt.Println(q.Read())

	case PRINT:
		fmt.Println("=== PRINT ===")
		if arr, ok := arrays[name]; ok {
			fmt.Printf("Array %s:%s\n", name, arr.Read())
			return
		}
		if onell, ok := onellists[name]; ok {
			fmt.Printf("OneLinkedList %s (Head→Tail):%s\n", name, onell.ReadHT())
			fmt.Printf("OneLinkedList %s (Tail→Head):%s\n", name, onell.ReadTH())
			return
		}
		if ll, ok := llists[name]; ok {
			fmt.Printf("LinkedList %s (Head→Tail):%s\n", name, ll.ReadHT())
			fmt.Printf("LinkedList %s (Tail→Head):%s\n", name, ll.ReadTH())
			return
		}
		if q, ok := queues[name]; ok {
			fmt.Printf("Queue %s:%s\n", name, q.Read())
			return
		}
		if s, ok := stacks[name]; ok {
			fmt.Printf("Stack %s:%s\n", name, s.Read())
			return
		}
		fmt.Printf("Объект \"%s\" не найден.\n", name)

	default:
		fmt.Println("Неизвестная команда.")
	}
}

func parsing(filename, query string) error {
	if err := readState(filename); err != nil {
		return err
	}

	parts := strings.Fields(query)
	if len(parts) == 0 {
		return nil
	}

	commandStr := parts[0]
	cmd, ok := commandMap[commandStr]
	if !ok {
		fmt.Printf("Неизвестная команда: %s\n", commandStr)
		return nil
	}

	name := ""
	if len(parts) > 1 {
		name = parts[1]
	}
	var data []string
	if len(parts) > 2 {
		data = parts[2:]
	}

	doCommand(name, data, cmd)

	return writeState(filename)
}

func main() {
	if len(os.Args) != 5 {
		return
	}

	var filename, query string
	if os.Args[1] == "--file" {
		filename = os.Args[2]
	}
	if os.Args[3] == "--query" {
		query = os.Args[4]
	}

	if err := parsing(filename, query); err != nil {
		fmt.Println(err)
	}
}
