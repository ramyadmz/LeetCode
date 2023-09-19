package lc359

type LogEntry struct {
	message string
	time    int
}
type LogNode struct {
	value LogEntry
	next  *LogNode
	prev  *LogNode
}

type Logger struct {
	head, tail *LogNode
	logMap     map[string]*LogNode
}

func NewLogger() Logger {
	return Logger{
		head:   nil,
		tail:   nil,
		logMap: make(map[string]*LogNode),
	}
}
func (this *Logger) Add(message string, time int) {
	log := &LogEntry{message, time}
	node := &LogNode{
		value: *log,
		next:  nil,
		prev:  this.tail,
	}
	if this.head == nil {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.logMap[log.message] = node
}

func (this *Logger) Upsert(message string, time int) {
	this.Delete(message)
	this.Add(message, time)
}

func (this *Logger) Delete(message string) {
	node, ok := this.logMap[message]
	if !ok {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if this.head == node {
		this.head = node.next
	}
	if this.tail == node {
		this.tail = node.prev
	}
	delete(this.logMap, message)
}

func (this *Logger) CleanUp(time int) {
	for this.head != nil && this.head.value.time < time-10 {
		this.Delete(this.head.value.message)
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	logNode, exist := this.logMap[message]
	if !exist || timestamp-logNode.value.time >= 10 {
		this.Upsert(message, timestamp)
		this.CleanUp(timestamp)
		return true
	}
	return false
}
