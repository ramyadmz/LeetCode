package lc2402

import (
	"container/heap"
	"sort"
)

// MeetingRoom represents a meeting room.
type MeetingRoom struct {
	id            int // Room ID
	meetings      int // Number of meetings held in the room
	occupiedUntil int // Time until the room is occupied
}

// AvailableRoomsHeap is a heap for available meeting rooms.
type AvailableRoomsHeap []*MeetingRoom

func (h AvailableRoomsHeap) Len() int      { return len(h) }
func (h AvailableRoomsHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h AvailableRoomsHeap) Less(i, j int) bool {
	return h[i].id < h[j].id
}
func (h *AvailableRoomsHeap) Push(x any) { *h = append(*h, x.(*MeetingRoom)) }
func (h *AvailableRoomsHeap) Pop() any {
	n := len(*h) - 1
	item := (*h)[n]
	(*h)[n] = nil
	*h = (*h)[:n]
	return item
}

// UsedRoomsHeap is a heap for used meeting rooms.
type UsedRoomsHeap []*MeetingRoom

func (h UsedRoomsHeap) Len() int      { return len(h) }
func (h UsedRoomsHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h UsedRoomsHeap) Less(i, j int) bool {
	if h[i].occupiedUntil == h[j].occupiedUntil {
		return h[i].id < h[j].id
	}
	return h[i].occupiedUntil < h[j].occupiedUntil
}
func (h *UsedRoomsHeap) Push(x any) { *h = append(*h, x.(*MeetingRoom)) }
func (h *UsedRoomsHeap) Pop() any {
	n := len(*h) - 1
	item := (*h)[n]
	(*h)[n] = nil
	*h = (*h)[:n]
	return item
}

// NewMeetingRoom creates a new meeting room with the given ID.
func NewMeetingRoom(id int) *MeetingRoom {
	return &MeetingRoom{id: id}
}

// Implement the heap.Interface methods for AvailableRoomsHeap and UsedRoomsHeap here...

// AllocateRoom allocates a room for a meeting.
func (room *MeetingRoom) AllocateRoom(startTime, endTime int) {
	room.meetings++
	room.occupiedUntil = endTime
}

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	availableRooms := make(AvailableRoomsHeap, n)
	for i := 0; i < n; i++ {
		availableRooms[i] = NewMeetingRoom(i)
	}
	heap.Init(&availableRooms)

	usedRooms := make(UsedRoomsHeap, 0)
	heap.Init(&usedRooms)

	var mostMeetingsRoom *MeetingRoom
	for _, meeting := range meetings {
		// Check if any used rooms are available again
		for len(usedRooms) > 0 && usedRooms[0].occupiedUntil <= meeting[0] {
			room := heap.Pop(&usedRooms).(*MeetingRoom)
			heap.Push(&availableRooms, room)
		}

		// Check available rooms first, then delay if necessary
		if len(availableRooms) > 0 {
			mostMeetingsRoom = heap.Pop(&availableRooms).(*MeetingRoom)
		} else {
			mostMeetingsRoom = heap.Pop(&usedRooms).(*MeetingRoom)
			delay := mostMeetingsRoom.occupiedUntil - meeting[0]
			meeting[0] += delay
			meeting[1] += delay
		}

		mostMeetingsRoom.AllocateRoom(meeting[0], meeting[1])
		heap.Push(&usedRooms, mostMeetingsRoom)
	}

	mostUsed := -1
	maxMeetings := -1
	for _, room := range usedRooms {
		if room.meetings > maxMeetings || (room.meetings == maxMeetings && room.id < mostUsed) {
			maxMeetings = room.meetings
			mostUsed = room.id
		}
	}

	return mostUsed
}
