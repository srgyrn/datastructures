package single

import (
	"reflect"
	"testing"
)

func TestListNode_Delete(t *testing.T) {
	list := &ListNode{
		val: 5,
		next: &ListNode{
			val: 4,
			next: &ListNode{
				val: 3,
				next: &ListNode{
					val: 2,
					next: &ListNode{
						val:  1,
						next: nil,
					},
				},
			},
		},
	}

	type args struct {
		val int
	}
	tests := []struct {
		name string
		list *ListNode
		args args
		want *ListNode
	}{
		{
			name: "removes last inserted node",
			args: args{5},
			list: list,
			want: &ListNode{
				val: 4,
				next: &ListNode{
					val: 3,
					next: &ListNode{
						val: 2,
						next: &ListNode{
							val:  1,
							next: nil,
						},
					},
				},
			},
		},
		{
			name: "removes mid node",
			list: list,
			args: args{3},
			want: &ListNode{
				val: 4,
				next: &ListNode{
					val: 2,
					next: &ListNode{
						val:  1,
						next: nil,
					},
				},
			},
		},
		{
			name: "removes first node",
			list: list,
			args: args{1},
			want: &ListNode{
				val: 4,
				next: &ListNode{
					val: 2,
					next: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Delete(tt.args.val)

			if !reflect.DeepEqual(tt.list, tt.want) {
				t.Errorf("Delete() got: %v, want: %v", tt.list, tt.want)
			}
		})
	}
}

func TestListNode_Insert(t *testing.T) {
	type fields struct {
		val  int
		next *ListNode
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "valid insert",
			fields: fields{
				val:  1,
				next: nil,
			},
			args: args{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &ListNode{
				val:  tt.fields.val,
				next: tt.fields.next,
			}

			root.Insert(tt.args.val)
			want := &ListNode{
				val: tt.args.val,
				next: &ListNode{
					val:  tt.fields.val,
					next: nil,
				},
			}

			if !reflect.DeepEqual(want, root) {
				t.Errorf("Insert() want: %v, got: %v", want, root)
			}
		})
	}
}

func TestListNode_Size(t *testing.T) {
	tests := []struct {
		name   string
		fields []int
		want   int
	}{
		{
			name:   "single element",
			fields: []int{1},
			want:   1,
		},
		{
			name:   "multiple elements",
			fields: []int{1, 2, 3},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &ListNode{
				val:  tt.fields[0],
				next: nil,
			}

			if len(tt.fields) > 1 {
				for i := 1; i < len(tt.fields); i++ {
					newNode := &ListNode{
						val: tt.fields[i],
					}
					newNode.next = root
					root = newNode
				}
			}

			if got := root.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSingleLinkedList(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "returns new linked list",
			args: args{
				val: 1,
			},
			want: &ListNode{
				val:  1,
				next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingleLinkedList(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingleLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
