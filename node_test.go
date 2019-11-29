package zss

import "testing"

func tree1() *Node {
	f := &Node{Label: "f"}
	d := &Node{Label: "d"}
	a := &Node{Label: "a"}
	c := &Node{Label: "c"}
	b := &Node{Label: "b"}
	e := &Node{Label: "e"}
	f.AddKid(d.AddKid(a, false).AddKid(c.AddKid(b, false), false), false).AddKid(e, false)
	return f
}

func TestNodeContainsWithLabel(t *testing.T) {
	root := tree1()
	contains := root.ContainsWithLabel("f")
	if contains != true {
		t.Errorf("%s not contains itself\n", root.Label)
	}

	contains = root.ContainsWithLabel("d")
	if contains != true {
		t.Errorf("%s not contains d\n", root.Label)
	}

	contains = root.ContainsWithLabel("e")
	if contains != true {
		t.Errorf("%s not contains d\n", root.Label)
	}

	contains = root.ContainsWithLabel("a")
	if contains != true {
		t.Errorf("%s not contains a\n", root.Label)
	}

	contains = root.ContainsWithLabel("c")
	if contains != true {
		t.Errorf("%s not contains c\n", root.Label)
	}

	contains = root.ContainsWithLabel("b")
	if contains != true {
		t.Errorf("%s not contains b\n", root.Label)
	}

}
