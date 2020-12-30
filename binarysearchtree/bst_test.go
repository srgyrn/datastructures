package binarysearchtree

import (
	"reflect"
	"testing"
)

func Test_bst_insert(t *testing.T) {
	type fields struct {
		root   *node
		length int
	}

	tests := []struct {
		name   string
		fields fields
		arg    int
		want   *bst
	}{
		{
			name:   "set root",
			fields: fields{},
			arg:    5,
			want:   &bst{root: &node{val: 5}},
		},
		{
			name:   "second level first child/bigger than root",
			fields: fields{root: &node{val: 5}},
			arg:    7,
			want: &bst{
				root: &node{
					val:   5,
					right: &node{val: 7},
				},
			},
		},
		{
			name:   "second level first child/smaller than root",
			fields: fields{root: &node{val: 5}},
			arg:    4,
			want: &bst{
				root: &node{
					val:  5,
					left: &node{val: 4},
				},
			},
		},
		{
			name: "third level child/bigger than child root",
			fields: fields{
				root: &node{
					val:   5,
					right: &node{val: 7},
				},
			},
			arg: 8,
			want: &bst{
				root: &node{
					val:   5,
					right: &node{val: 7, right: &node{val: 8}},
				},
			},
		},
		{
			name: "third level child/bigger than left child root, smaller than right child root",
			fields: fields{
				root: &node{
					val:   5,
					right: &node{val: 7},
				},
			},
			arg: 6,
			want: &bst{
				root: &node{
					val:   5,
					right: &node{val: 7, left: &node{val: 6}},
				},
			},
		},
		{
			name: "third level child/smaller than left child root",
			fields: fields{
				root: &node{
					val:   5,
					left:  &node{val: 4},
					right: &node{val: 7},
				},
			},
			arg: 3,
			want: &bst{
				root: &node{
					val:   5,
					left:  &node{val: 4, left: &node{val: 3}},
					right: &node{val: 7},
				},
			},
		},
		{
			name: "fourth level child",
			fields: fields{
				root: &node{
					val:   5,
					right: &node{val: 7, left: &node{val: 6}, right: &node{val: 9}},
				},
			},
			arg: 8,
			want: &bst{
				root: &node{
					val:   5,
					right: &node{val: 7, left: &node{val: 6}, right: &node{val: 9, left: &node{val: 8}}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bst{
				root: tt.fields.root,
			}
			b.insert(tt.arg)

			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("b.insert() got: %+v want: %+v", b, tt.want)
			}
		})
	}
}

func Test_bst_search(t *testing.T) {
	tree := &bst{
		root: &node{
			val:  5,
			left: &node{val: 4},
			right: &node{
				val:  7,
				left: &node{val: 6},
				right: &node{
					val:  9,
					left: &node{val: 8},
				},
			},
		},
	}

	if tree.search(3) {
		t.Errorf("bst.search() failed")
	}

	if !tree.search(8) {
		t.Errorf("bst.search() failed")
	}
}
