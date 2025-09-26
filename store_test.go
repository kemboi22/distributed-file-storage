package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpic"
	pathKey := CASPathTransformFunc(key)
	originalFileName := "1ff3521faa379e96fa2f013522a356cc934c23d2"
	expectedPathName := "1ff35/21faa/379e9/6fa2f/01352/2a356/cc934/c23d2"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.FileName != originalFileName {
		t.Errorf("have %s want %s ", pathKey.FileName, originalFileName)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsbestpic"
	data := []byte("Some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsbestpic"
	data := []byte("Some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}
	b, err := io.ReadAll(r)
	if err != nil {
		t.Error(err)
	}
	if string(b) != string(data) {
		t.Errorf("data received %s data expected %s", string(b), string(data))
	}
	s.Delete(key)
}
