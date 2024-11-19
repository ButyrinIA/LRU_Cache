package datastructures

import (
	"reflect"
	"testing"
)

func TestLinkedList_AddToFront(t *testing.T) {
	tests := []struct {
		name           string
		initialNodes   []*Node
		nodeToAdd      *Node
		expectedValues []int
	}{
		{
			name:           "добавление в пустой список",
			initialNodes:   nil,
			nodeToAdd:      &Node{Key: 1, Value: 11},
			expectedValues: []int{11},
		},
		{
			name: "добавление в не пустой список",
			initialNodes: []*Node{
				{Key: 1, Value: 11},
				{Key: 2, Value: 22},
			},
			nodeToAdd:      &Node{Key: 3, Value: 33},
			expectedValues: []int{33, 22, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}

			//инициализируем список
			for _, node := range tt.initialNodes {
				l.AddToFront(&Node{Key: node.Key, Value: node.Value})
			}
			//добавляем node
			l.AddToFront(tt.nodeToAdd)

			//проверка
			values := []int{}
			current := l.Head
			for current != nil {
				values = append(values, current.Value)
				current = current.Next
			}
			if !reflect.DeepEqual(values, tt.expectedValues) {
				t.Errorf("Expected values %v, got %v", tt.expectedValues, values)
			}
		})
	}
}

func TestLinkedList_RemoveNode(t *testing.T) {
	tests := []struct {
		name           string
		initialNodes   []*Node
		nodeToRemove   *Node
		expectedValues []int
	}{
		{
			name:           "удаление из пустого списка",
			initialNodes:   nil,
			nodeToRemove:   &Node{Key: 2, Value: 22},
			expectedValues: []int{},
		},
		{
			name: "удаление из середины списка",
			initialNodes: []*Node{
				{Key: 1, Value: 11},
				{Key: 2, Value: 22},
				{Key: 3, Value: 33},
			},
			nodeToRemove:   &Node{Key: 2, Value: 22},
			expectedValues: []int{33, 11},
		},
		{
			name: "удаление из начала списка",
			initialNodes: []*Node{
				{Key: 1, Value: 11},
				{Key: 2, Value: 22},
				{Key: 3, Value: 33},
			},
			nodeToRemove:   &Node{Key: 3, Value: 33},
			expectedValues: []int{22, 11},
		},
		{
			name: "удаление из конца списка",
			initialNodes: []*Node{
				{Key: 1, Value: 11},
				{Key: 2, Value: 22},
				{Key: 3, Value: 33},
			},
			nodeToRemove:   &Node{Key: 1, Value: 11},
			expectedValues: []int{33, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}

			//инициализируем список
			nodes := map[int]*Node{}
			for _, node := range tt.initialNodes {
				newNode := &Node{Key: node.Key, Value: node.Value}
				nodes[node.Key] = newNode
				l.AddToFront(newNode)
			}
			//удаление node
			nodeToRemove, exists := nodes[tt.nodeToRemove.Key]
			if exists {
				l.RemoveNode(nodeToRemove)
			} else {
				l.RemoveNode(tt.nodeToRemove)
			}

			//проверка
			values := []int{}
			current := l.Head
			for current != nil {
				values = append(values, current.Value)
				current = current.Next
			}
			if !reflect.DeepEqual(values, tt.expectedValues) {
				t.Errorf("Expected values %v, got %v", tt.expectedValues, values)
			}
		})
	}
}

func TestLinkedList_RemoveTail(t *testing.T) {
	tests := []struct {
		name           string
		initialNodes   []*Node
		removeTail     *Node
		expectedValues []int
	}{

		{
			name:           "удаление из пустого списка",
			initialNodes:   nil,
			removeTail:     nil,
			expectedValues: []int{},
		},
		{
			name: "удаление из не пустого списка",
			initialNodes: []*Node{
				{Key: 1, Value: 11},
				{Key: 2, Value: 22},
				{Key: 3, Value: 33},
			},
			removeTail:     &Node{Key: 1, Value: 11},
			expectedValues: []int{33, 22},
		},
		{
			name: "удаление из списка с одним звеном",
			initialNodes: []*Node{
				{Key: 1, Value: 11},
			},
			removeTail:     &Node{Key: 1, Value: 11},
			expectedValues: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{}

			//инициализируем список
			for _, node := range tt.initialNodes {
				l.AddToFront(&Node{Key: node.Key, Value: node.Value})
			}
			//удаление node
			removedTail := l.RemoveTail()
			if removedTail != nil {
				removedTail.Prev = nil
			}

			//проверка
			if !reflect.DeepEqual(removedTail, tt.removeTail) {
				t.Errorf("Expected removed tail %v, got %v", tt.removeTail, removedTail)
			}

			values := []int{}
			curr := l.Head
			for curr != nil {
				values = append(values, curr.Value)
				curr = curr.Next
			}

			if !reflect.DeepEqual(values, tt.expectedValues) {
				t.Errorf("Expected values %v, got %v", tt.expectedValues, values)
			}
		})
	}
}
