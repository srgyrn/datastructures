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

	tree = &bst{}
	if tree.search(3) {
		t.Errorf("bst.search() failed")
	}
}

func Test_bst_remove(t *testing.T) {
	tests := []struct {
		name string
		b    *bst
		v    int
		want *bst
	}{
		{
			name: "remove leaf",
			b:    getBigBtree(),
			v:    77,
			want: &bst{
				root: &node{
					val: 71,
					left: &node{
						val: 29,
						left: &node{
							val: 6,
							right: &node{
								val: 7,
								right: &node{
									val:  20,
									left: &node{val: 18},
								},
							},
						},
						right: &node{
							val: 66,
							left: &node{
								val:   53,
								right: &node{val: 55},
							},
							right: &node{val: 68},
						},
					},
					right: &node{
						val:   83,
						left:  nil,
						right: &node{val: 87},
					},
				},
			},
		},
		{
			name: "remove leaf 2",
			b:    getBigBtree(),
			v:    55,
			want: &bst{
				root: &node{
					val: 71,
					left: &node{
						val: 29,
						left: &node{
							val: 6,
							right: &node{
								val: 7,
								right: &node{
									val:  20,
									left: &node{val: 18},
								},
							},
						},
						right: &node{
							val: 66,
							left: &node{
								val:   53,
								right: nil,
							},
							right: &node{val: 68},
						},
					},
					right: &node{
						val:   83,
						left:  &node{val: 77},
						right: &node{val: 87},
					},
				},
			},
		},
		{
			name: "remove node with one child",
			b:    getBigBtree(),
			v:    53,
			want: &bst{
				root: &node{
					val: 71,
					left: &node{
						val: 29,
						left: &node{
							val: 6,
							right: &node{
								val: 7,
								right: &node{
									val:  20,
									left: &node{val: 18},
								},
							},
						},
						right: &node{
							val: 66,
							left: &node{
								val: 55,
							},
							right: &node{val: 68},
						},
					},
					right: &node{
						val:   83,
						left:  &node{val: 77},
						right: &node{val: 87},
					},
				},
			},
		},
		{
			name: "remove node with right subtree",
			b:    getBigBtree(),
			v:    29,
			want: &bst{
				root: &node{
					val: 71,
					left: &node{
						val: 53,
						left: &node{
							val: 6,
							right: &node{
								val: 7,
								right: &node{
									val:  20,
									left: &node{val: 18},
								},
							},
						},
						right: &node{
							val: 66,
							left: &node{
								val: 55,
							},
							right: &node{val: 68},
						},
					},
					right: &node{
						val:   83,
						left:  &node{val: 77},
						right: &node{val: 87},
					},
				},
			},
		},
		{
			name: "remove child with no right subtree",
			b: &bst{
				root: &node{
					val: 71,
					left: &node{
						val: 29,
						left: &node{
							val: 6,
						},
					},
					right: &node{
						val: 83,
						left: &node{
							val: 77,
							left: &node{
								val: 74,
								left: &node{
									val: 73,
								},
								right: &node{
									val:  76,
									left: &node{val: 75},
								},
							},
						},
						right: &node{val: 87},
					},
				},
			},
			v: 77,
			want: &bst{
				root: &node{
					val: 71,
					left: &node{
						val: 29,
						left: &node{
							val: 6,
						},
					},
					right: &node{
						val: 83,
						left: &node{
							val: 74,
							left: &node{
								val: 73,
							},
							right: &node{
								val:  76,
								left: &node{val: 75},
							},
						},
						right: &node{val: 87},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.remove(tt.v)

			if !reflect.DeepEqual(tt.b, tt.want) {
				t.Errorf("b.remove() failed")
				return
			}
		})
	}
}

func getBigBtree() *bst {
	return &bst{
		root: &node{
			val: 71,
			left: &node{
				val: 29,
				left: &node{
					val: 6,
					right: &node{
						val: 7,
						right: &node{
							val:  20,
							left: &node{val: 18},
						},
					},
				},
				right: &node{
					val: 66,
					left: &node{
						val:   53,
						right: &node{val: 55},
					},
					right: &node{val: 68},
				},
			},
			right: &node{
				val:   83,
				left:  &node{val: 77},
				right: &node{val: 87},
			},
		},
	}
}
