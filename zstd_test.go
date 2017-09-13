package main

import (
	"bytes"
	"testing"
)

func TestStoreAndLoad(t *testing.T) {
	z := ZSTDStore{}

	tc := "first time writing tests, heh"

	err := z.Store("fixtures/test.txt", []byte(tc))
	if err != nil {
		t.Error(err)
	}

	d, err := z.Load("fixtures/test.txt")
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(d, []byte(tc)) {
		t.Errorf("expected: \"%s\", got: \"%s\"", tc, d)
	}

}

func TestZSTDStore_Compress(t *testing.T) {
	z := ZSTDStore{}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		z       ZSTDStore
		args    args
		wantErr bool
	}{
		{"Compress fail", z, args{"fixtures/missingFile.txt"}, true},
		{"Compress success", z, args{"fixtures/raw.txt"}, false},
	}

	// defer func(){
	// 	z.
	// }()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.z.Compress(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ZSTDStore.Compress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestZSTDStore_Uncompress(t *testing.T) {
	z := ZSTDStore{}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		z       ZSTDStore
		args    args
		wantErr bool
	}{
		{"success", z, args{"fixtures/raw.txt.gz"}, false},
		{"fail", z, args{"fixtures/missing.txt.gz"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.z.Decompress(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ZSTDStore.Uncompress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
