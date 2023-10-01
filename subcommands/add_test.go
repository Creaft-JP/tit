package subcommands

import (
	"github.com/Creaft-JP/tit/db/local/ent/stagedfile"
	"github.com/Creaft-JP/tit/test"
	"github.com/morikuni/failure"
	"os"
	"path/filepath"
	"testing"
)

func TestFirstAdd(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	if err := os.WriteFile("new-file", []byte("Hello, world!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := Add([]string{"new-file"}, client, ctx); err != nil {
		t.Fatal(err)
	}
	got, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 1
	if len(got) != wlen {
		t.Fatalf("number of staged files shouled be %d, but got %d", wlen, len(got))
	}
	wpath := "new-file"
	wcon := "Hello, world!!"
	if got[0].Path != wpath {
		t.Errorf("path of the staged file should be \"%s\", but got \"%s\"", wpath, got[0].Path)
	}
	if got[0].Content != wcon {
		t.Errorf("content of the staged file should be \"%s\", but got \"%s\"", wcon, got[0].Content)
	}
}
func TestSecondAdd(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	if err := os.WriteFile("new-file-1", []byte("Hello, world!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("new-file-2", []byte("Hello, tit!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := Add([]string{"new-file-1", "new-file-2"}, client, ctx); err != nil {
		t.Fatal(err)
	}
	got, err := client.StagedFile.Query().Order(stagedfile.ByPath()).All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 2
	if len(got) != wlen {
		t.Fatalf("number of staged files should be %d, but got %d", wlen, len(got))
	}
	wpath1 := "new-file-1"
	wcon1 := "Hello, world!!"
	wpath2 := "new-file-2"
	wcon2 := "Hello, tit!!"
	if got[0].Path != wpath1 {
		t.Errorf("path of the staged file should be \"%s\", but got \"%s\"", wpath1, got[0].Path)
	}
	if got[0].Content != wcon1 {
		t.Errorf("content of the staged file should be \"%s\", but got \"%s\"", wcon1, got[0].Content)
	}
	if got[1].Path != wpath2 {
		t.Errorf("path of the staged file should be \"%s\", but got \"%s\"", wpath2, got[1].Path)
	}
	if got[1].Content != wcon2 {
		t.Errorf("content of the staged file should be \"%s\", but got \"%s\"", wcon2, got[1].Content)
	}
}
func TestOverwritingAdd(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	if err := os.WriteFile("new-file", []byte("Hello, world!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := Add([]string{"new-file"}, client, ctx); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("new-file", []byte("Hello, tit!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := Add([]string{"new-file"}, client, ctx); err != nil {
		t.Fatal(err)
	}
	got, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 1
	if len(got) != wlen {
		t.Fatalf("number of staged files should be %d, but got %d", wlen, len(got))
	}
	wpath := "new-file"
	wcon := "Hello, tit!!"
	if got[0].Path != wpath {
		t.Errorf("path of the staged file should be \"%s\", but got \"%s\"", wpath, got[0].Path)
	}
	if got[0].Content != wcon {
		t.Errorf("content of the staged file should be \"%s\", but got \"%s\"", wcon, got[0].Content)
	}
}
func TestTrickyOverwritingAdd(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	if err := os.WriteFile("new-file", []byte("Hello, world!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := Add([]string{"new-file"}, client, ctx); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("new-file", []byte("Hello, tit!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir("tmp", 0755); err != nil {
		t.Fatal(err)
	}
	if err := Add([]string{filepath.Join("tmp", "..", "new-file")}, client, ctx); err != nil {
		t.Fatal(err)
	}
	got, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 1
	if len(got) != wlen {
		t.Fatalf("number of staged files should be %d, but got %d", wlen, len(got))
	}
	wpath := "new-file"
	wcon := "Hello, tit!!"
	if got[0].Path != wpath {
		t.Errorf("path of the staged file should be \"%s\", but got \"%s\"", wpath, got[0].Path)
	}
	if got[0].Content != wcon {
		t.Errorf("content of the staged file should be \"%s\", but got \"%s\"", wcon, got[0].Content)
	}
}
func TestAddNonExistent(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	err := Add([]string{"new-file"}, client, ctx)
	if err == nil {
		t.Fatal("an error should be thrown, but none were")
	}
	gmes, isflr := failure.MessageOf(err)
	if !isflr {
		t.Fatal("a failure should be thrown, but the throws error is not a failure")
	}
	wmes := "path new-file doesn't exist"
	if gmes != wmes {
		t.Errorf("the error message should be \"%s\", but got \"%s\"", wmes, gmes)
	}
	grec, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 0
	if len(grec) != wlen {
		t.Fatalf("number of staged files should be %d, but got %d", wlen, len(grec))
	}
}
func TestAddDirectory(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	if err := os.Mkdir("new-directory", 0755); err != nil {
		t.Fatal(err)
	}
	err := Add([]string{"new-directory"}, client, ctx)
	if err == nil {
		t.Fatal("an error should be thrown, but none were")
	}
	gmes, isflr := failure.MessageOf(err)
	if !isflr {
		t.Fatal("a failure should be thrown, but the throws error is not a failure")
	}
	wmes := "path new-directory is a directory"
	if gmes != wmes {
		t.Errorf("the error message should be \"%s\", but got \"%s\"", wmes, gmes)
	}
	grec, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 0
	if len(grec) != wlen {
		t.Fatalf("number of staged files should be %d, but got %d", wlen, len(grec))
	}
}
func TestAddOutOfRepository(t *testing.T) {
	client, ctx := test.SetUpInner(t)
	defer test.TearDown(t, client)
	path := filepath.Join("..", "new-file")
	if err := os.WriteFile(path, []byte("Hello, world!!"), 0755); err != nil {
		t.Fatal(err)
	}
	err := Add([]string{path}, client, ctx)
	if err == nil {
		t.Fatal("an error should be thrown, but none were")
	}
	gmes, isflr := failure.MessageOf(err)
	if !isflr {
		t.Fatal("a failure should be thrown, but the throws error is not a failure")
	}
	wmes := "can't add a file out of tit repository"
	if gmes != wmes {
		t.Errorf("the error message should be \"%s\", but got \"%s\"", wmes, gmes)
	}
	grec, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 0
	if len(grec) != wlen {
		t.Fatalf("number of staged files should be %d, but got %d", wlen, len(grec))
	}
}
func TestTrickyAddOutOfRepository(t *testing.T) {
	client, ctx := test.SetUpInner(t)
	defer test.TearDown(t, client)
	if err := os.WriteFile(filepath.Join("..", "new-file"), []byte("Hello, world!!"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir("tmp", 0755); err != nil {
		t.Fatal(err)
	}
	err := Add([]string{filepath.Join("tmp", "..", "..", "new-file")}, client, ctx)
	if err == nil {
		t.Fatal("an error should be thrown, but none were")
	}
	gmes, isflr := failure.MessageOf(err)
	if !isflr {
		t.Fatal("a failure should be thrown, but the throws error is not a failure")
	}
	wmes := "can't add a file out of tit repository"
	if gmes != wmes {
		t.Errorf("the error message should be \"%s\", but got \"%s\"", wmes, gmes)
	}
	grec, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	wlen := 0
	if len(grec) != wlen {
		t.Fatalf("number of staged files should be %d, but got %d", wlen, len(grec))
	}
}
