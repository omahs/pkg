package server_test

import (
	"context"
	"math"
	"testing"

	api "github.com/wetware/ww/api/process"
	csp "github.com/wetware/ww/pkg/csp/server"
)

type testProc struct {
	pid   uint32
	alive bool
}

func (p *testProc) Kill(context.Context, api.Process_kill) error {
	p.alive = false
	return nil
}

func (p *testProc) Wait(ctx context.Context, call api.Process_wait) error {
	return nil
}

func testProcTree() csp.ProcTree {
	/*
	        0
	        |
	        1
	      /   \
	     2     10
	    / \   /
	   3   6 11
	    \   \
	     4   7
	    /   / \
	   5   8   9
	*/
	root := &csp.ProcNode{
		Pid: 0,
		Left: &csp.ProcNode{
			Pid: 1,
			Left: &csp.ProcNode{
				Pid: 2,
				Left: &csp.ProcNode{
					Pid: 3,
					Right: &csp.ProcNode{
						Pid: 4,
						Left: &csp.ProcNode{
							Pid: 5,
						},
					},
				},
				Right: &csp.ProcNode{
					Pid: 6,
					Right: &csp.ProcNode{
						Pid: 7,
						Left: &csp.ProcNode{
							Pid: 8,
						},
						Right: &csp.ProcNode{
							Pid: 9,
						},
					},
				},
			},
			Right: &csp.ProcNode{
				Pid: 10,
				Left: &csp.ProcNode{
					Pid: 11,
				},
			},
		},
	}

	procMap := make(map[uint32]api.Process_Server)
	for pid := uint32(0); pid <= 11; pid++ {
		procMap[pid] = &testProc{pid: pid, alive: true}
	}

	return csp.ProcTree{
		IDC:  csp.AtomicCounter{},
		PC:   csp.AtomicCounter{},
		Root: root,
		Map:  procMap,
	}
}

func TestProcTree_Find(t *testing.T) {
	pt := testProcTree()
	for i := uint32(0); i <= 11; i++ {
		n := pt.Find(i)
		if n == nil {
			t.Fatalf("failed to find node %d", i)
		}
		if n.Pid != i {
			t.Fatalf("found node %d instead of %d", n.Pid, i)
		}
	}
}

func TestProcTree_FindParent(t *testing.T) {
	// child, parent
	matches := [6][2]uint32{
		{8, 7},
		{9, 1},
		{11, 10},
		{3, 2},
		{4, 2},
		{5, 4},
	}
	pt := testProcTree()
	for _, match := range matches {
		c := match[0]
		p := pt.FindParent(c)
		if p == nil {
			t.Fatalf("nil parent for %d", c)
		}
		e := match[1]
		if p.Pid != e {
			t.Fatalf("found parent %d for %d but expected %d", p.Pid, c, e)
		}
	}
	c := uint32(math.MaxUint32)
	p := pt.FindParent(c)
	if p != nil {
		t.Fatalf("found parent %d for %d but expected no parent", p.Pid, c)
	}
}

func TestProcTree_Insert(t *testing.T) {
	// child, parent, branchof, 0=left 1=right
	matches := [4][4]uint32{
		{12, 5, 5, 0},
		{13, 12, 12, 0},
		{13, 1, 9, 1},
		{14, 7, 8, 1},
	}
	pt := testProcTree()
	for _, match := range matches {
		pid, ppid, expectedId, side := match[0], match[1], match[2], match[3]
		pt.Insert(pid, ppid)
		n := pt.Find(expectedId)
		if side == 0 {
			if n.Left == nil || n.Left.Pid != pid {
				t.Fatalf("failet to insert %d at %d (branch %s)", pid, ppid, n)
			}
		} else {
			if n.Right == nil || n.Right.Pid != pid {
				t.Fatalf("failet to insert %d at %d (branch %s)", pid, ppid, n)
			}
		}
	}

}

func TestProcTree_Pop(t *testing.T) {
	pt := testProcTree()
	parent := pt.FindParent(6)
	sibling := pt.Find(2)
	child := pt.Find(6)
	popped := pt.Pop(6)
	if popped.Pid != child.Pid {
		t.Fatalf("popped item with PID %d instead of %d", popped.Pid, child.Pid)
	}
	if sibling.Right.Pid != 7 {
		t.Fatalf("new right branch of %d should be 7, not %d", parent.Pid, sibling.Right.Pid)
	}
	// this test makes me dizzy
	parent = sibling.Right
	child = parent.Left
	if child.Pid != 8 {
		t.Fatalf("expected pid 8 got %d", child.Pid)
	}
	popped = pt.Pop(child.Pid)
	if popped.Pid != child.Pid {
		t.Fatalf("popped item with PID %d instead of %d", popped.Pid, child.Pid)
	}
	if parent.Left != nil {
		t.Fatalf("left branch of %d should be nil, not %d", sibling.Pid, sibling.Left.Pid)
	}
}

func TestProcTree_Kill(t *testing.T) {
	pt := testProcTree()

	mapCopy := make(map[uint32]api.Process_Server)
	for k, v := range pt.Map {
		mapCopy[k] = v
	}

	pt.Kill(1)
	if pt.Root.Left == nil || pt.Root.Left.Pid != 10 {
		t.Fatalf("expected to find 10 at the left of 0, found %s", pt.Root.Left)
	}
	killedProcs := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	aliveProcs := []uint32{10, 11}

	for _, pid := range killedProcs {
		if _, found := pt.Map[pid]; found {
			t.Fatalf("found process %d in map but it should have been deleted", pid)
		}
		p := mapCopy[pid]
		if tp, _ := p.(*testProc); tp.alive {
			t.Fatalf("killed process %d is still alive", pid)
		}
	}

	for _, pid := range aliveProcs {
		if _, found := pt.Map[pid]; !found {
			t.Fatalf("failed to find process %d in map", pid)
		}
		p := mapCopy[pid]
		if tp, _ := p.(*testProc); !tp.alive {
			t.Fatalf("killed process %d should still be alive", pid)
		}
	}
}
