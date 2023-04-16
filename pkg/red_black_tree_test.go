package main

import (
	"math/rand"
	"testing"
)

// test for red_black_tree.go
// go test -v
func TestInsert(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(1, "test")
	if tree.Root.Key != 1 && tree.Root.Val != "test" {
		t.Errorf("Insert failed")
	}
}

func TestSearch(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(1, "test")
	if tree.Search(1).Val != "test" {
		t.Errorf("Search failed")
	}
}

func TestDelete(t *testing.T) {
	tree := NewRBTree()
	tree.Insert(1, "test")
	tree.Delete(1)
	if tree.Root != nil {
		t.Errorf("Delete failed")
	}
}

// fuzz test for red_black_tree.go
// go test -v -fuzz=Fuzz
func TestRedBlackTree(t *testing.T) {
	tree := NewRBTree()
	for i := 0; i < 100; i++ {
		key := rand.Intn(100)
		tree.Insert(key, "test")
	}
	// for i := 0; i < 50; i++ {
	// 	key := rand.Intn(100)
	// 	tree.Delete(key)
	// }

}

// benchmark for red_black_tree.go
// go test -bench=. -benchmem
func BenchmarkInsert(b *testing.B) {
	tree := NewRBTree()
	for i := 0; i < b.N; i++ {
		tree.Insert(i, "test")
	}
}

func BenchmarkSearch(b *testing.B) {
	tree := NewRBTree()
	for i := 0; i < b.N; i++ {
		tree.Insert(i, "test")
	}
	for i := 0; i < b.N; i++ {
		tree.Search(i)
	}
}

func BenchmarkDelete(b *testing.B) {
	tree := NewRBTree()
	for i := 0; i < b.N; i++ {
		tree.Insert(i, "test")
	}
	for i := 0; i < b.N; i++ {
		tree.Delete(i)
	}
}
