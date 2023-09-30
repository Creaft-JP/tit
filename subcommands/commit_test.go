package subcommands

import (
	"context"
	"database/sql"
	"github.com/Creaft-JP/tit/db/local/ent"
	"github.com/Creaft-JP/tit/db/local/ent/committedfile"
	"github.com/Creaft-JP/tit/db/local/ent/titcommit"
	e "github.com/Creaft-JP/tit/error"
	"github.com/Creaft-JP/tit/test"
	"github.com/morikuni/failure"
	"go.uber.org/multierr"
	"testing"
)

func TestCommitNewFile(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	setUpTestingRepository(t, client, ctx)
	if _, err := client.StagedFile.Create().
		SetPath("file3").
		SetContent("Goodbye, TitHub!!").Save(ctx); err != nil {
		t.Fatal(err)
	}
	if err := commit("Create file3.", client, ctx); err != nil {
		t.Fatal(err)
	}
	count, err := client.TitCommit.Query().Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 6 {
		t.Fatalf("number of commits should be 6, but got %d", count)
	}
	c6, err := client.TitCommit.Query().Where(titcommit.Message("Create file3.")).Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	section, err := c6.QuerySection().Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	page, err := section.QueryPage().Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if page.Pathname != "/committed" || section.Slug != "last-section" {
		t.Fatalf("committed section should be /committed#last-section, but got %s#%s", page.Pathname, section.Slug)
	}
	if c6.Number != 2 {
		t.Errorf("commit number should be 2, but got %d", c6.Number)
	}
	gf3, err := c6.QueryFiles().Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if gf3.Path != "file3" {
		t.Errorf("committed file path should be \"file3\", but got %s", gf3.Path)
	}
	if gf3.Content != "Goodbye, TitHub!!" {
		t.Errorf("committed file content should be \"Goodbye, TitHub!!\", but got %s", gf3.Content)
	}
	stg, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stg) != 0 {
		t.Errorf("the stage should be empty, but there's %d, file(s)", len(stg))
	}
}
func TestCommitExistingFile(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	setUpTestingRepository(t, client, ctx)
	if _, err := client.StagedFile.Create().
		SetPath("file1").
		SetContent("Goodbye, TitHub!!").Save(ctx); err != nil {
		t.Fatal(err)
	}
	if err := commit("Rewrite file1.", client, ctx); err != nil {
		t.Fatal(err)
	}
	count, err := client.TitCommit.Query().Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 6 {
		t.Fatalf("number of commits should be 6, but got %d", count)
	}
	c6, err := client.TitCommit.Query().Where(titcommit.Message("Rewrite file1.")).Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	section, err := c6.QuerySection().Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	page, err := section.QueryPage().Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if page.Pathname != "/committed" || section.Slug != "last-section" {
		t.Fatalf("committed section should be /committed#last-section, but got %s#%s", page.Pathname, section.Slug)
	}
	if c6.Number != 2 {
		t.Errorf("commit number should be 2, but got %d", c6.Number)
	}
	f1e, err := c6.QueryFiles().Only(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if f1e.Path != "file1" {
		t.Errorf("committed file path should be \"file1\", but got %s", f1e.Path)
	}
	if f1e.Content != "Goodbye, TitHub!!" {
		t.Errorf("committed file content should be \"Goodbye, TitHub!!\", but got %s", f1e.Content)
	}
	cfc, err := client.CommittedFile.Query().
		Where(committedfile.Path("file1")).
		Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if cfc != 5 {
		t.Errorf("number of commits referring file1 should be 5, but got %d", cfc)
	}
	stg, err := client.StagedFile.Query().All(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(stg) != 0 {
		t.Errorf("the stage should be empty, but there's %d file(s)", len(stg))
	}
}
func TestCommitWithEmptyMessage(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	setUpTestingRepository(t, client, ctx)
	if _, err := client.StagedFile.Create().
		SetPath("file3").
		SetContent("Goodbye, TitHub!!").Save(ctx); err != nil {
		t.Fatal(err)
	}
	err := commit("", client, ctx)
	if err == nil {
		t.Fatal("an error should be thrown, but none were")
	}
	if !failure.Is(err, e.Operation) {
		t.Fatal("the thrown error should be operation failure, but not")
	}
	wmes := "aborting commit due to empty commit message"
	gmes, _ := failure.MessageOf(err)
	if gmes != wmes {
		t.Errorf("failure message should be \"%s\", but got \"%s\"", wmes, gmes)
	}
	gstg, err := client.StagedFile.Query().Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if gstg != 1 {
		t.Fatalf("number of staged files should be 1, but got %d", gstg)
	}
	gcom, err := client.TitCommit.Query().Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if gcom != 5 {
		t.Fatalf("number of commits should be 5, but got %d", gcom)
	}
	gcomf, err := client.CommittedFile.Query().Where(committedfile.Path("file3")).Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if gcomf != 0 {
		t.Fatalf("file3 should not be committed, but has been %d time(s)", gcom)
	}
}
func TestCommitWithSectionOfLastPageNotExisting(t *testing.T) {
	client, _, ctx := test.SetUp(t)
	defer test.TearDown(t, client)
	setUpTestingRepository(t, client, ctx)
	if _, err := client.Page.Create().
		SetPathname("/true-committed").
		SetNumber(3).
		SetTitle("True Last Page").
		SetOverviewSentence("in case of this test, this page is to be committed").Save(ctx); err != nil {
		t.Fatal(err)
	}
	if _, err := client.StagedFile.Create().
		SetPath("file3").
		SetContent("Goodbye, TitHub!!").Save(ctx); err != nil {
		t.Fatal(err)
	}
	err := commit("Create file3.", client, ctx)
	if err == nil {
		t.Fatal("an error should be thrown, but none were")
	}
	if !failure.Is(err, e.Operation) {
		t.Fatal("thrown error should be operation failure, but not")
	}
	gmes, _ := failure.MessageOf(err)
	wmes := "please start first section of page /true-committed"
	if gmes != wmes {
		t.Errorf("message of thrown error should be \"%s\", but got \"%s\"", wmes, gmes)
	}
	gstg, err := client.StagedFile.Query().Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if gstg != 1 {
		t.Errorf("number of staged files should be 1, but got %d", gstg)
	}
	gcom, err := client.TitCommit.Query().Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if gcom != 5 {
		t.Errorf("number of commits should be 5, but got %d", gcom)
	}
	gcomf, err := client.CommittedFile.Query().Where(committedfile.Path("file3")).Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if gcomf != 0 {
		t.Errorf("file3 shouldn't be committed, but has been %d times", gcomf)
	}
}
func setUpTestingRepository(t *testing.T, cl *ent.Client, ctx context.Context) {
	tx1, err := cl.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}

	f1a, err := tx1.CommittedFile.Create().
		SetPath("file1").
		SetContent("Hello, world!!").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}
	c1, err := tx1.TitCommit.Create().
		SetNumber(1).
		SetMessage("Write \"Hello, world!!\" to file1.").
		AddFiles(f1a).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}

	f1b, err := tx1.CommittedFile.Create().
		SetPath("file1").
		SetContent("Hello, TitHub!!").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}
	f2a, err := tx1.CommittedFile.Create().
		SetPath("file2").
		SetContent("Goodbye, world!!").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}
	c2, err := tx1.TitCommit.Create().
		SetNumber(2).
		SetMessage("Change file1 and create file2.").
		AddFiles(f1b, f2a).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}

	s1, err := tx1.Section.Create().
		SetSlug("create-files").
		SetNumber(1).
		SetTitle("Create Files").
		SetOverviewSentence("create 2 file2").
		AddCommits(c1, c2).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}

	f1c, err := tx1.CommittedFile.Create().
		SetPath("file1").
		SetContent("").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}
	f2b, err := tx1.CommittedFile.Create().
		SetPath("file2").
		SetContent("").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}
	c3, err := tx1.TitCommit.Create().
		SetNumber(1).
		SetMessage("Remove all contents of file1 and file2.").
		AddFiles(f1c, f2b).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}

	s2, err := tx1.Section.Create().
		SetSlug("remove-files").
		SetNumber(2).
		SetTitle("Remove Files").
		SetOverviewSentence("remove all contents").
		AddCommits(c3).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}

	if _, err := tx1.Page.Create().
		SetPathname("/").
		SetNumber(1).
		SetTitle("First Page").
		SetOverviewSentence("create and remove files").
		AddSections(s1, s2).Save(ctx); err != nil {
		t.Fatal(multierr.Append(err, tx1.Rollback()))
	}

	if err := tx1.Commit(); err != nil {
		t.Fatal(err)
	}

	tx2, err := cl.BeginTx(ctx, &sql.TxOptions{})

	f1d, err := tx2.CommittedFile.Create().
		SetPath("file1").
		SetContent("").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx2.Rollback()))
	}
	f2c, err := tx2.CommittedFile.Create().
		SetPath("file2").
		SetContent("").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx2.Rollback()))
	}
	c4, err := tx2.TitCommit.Create().
		SetNumber(1).
		SetMessage("There's no change from previous page.").
		AddFiles(f1d, f2c).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx2.Rollback()))
	}

	s3, err := tx2.Section.Create().
		SetSlug("no-change").
		SetNumber(1).
		SetTitle("No Change").
		SetOverviewSentence("there's no change").
		AddCommits(c4).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx2.Rollback()))
	}

	c5, err := tx2.TitCommit.Create().
		SetNumber(1).
		SetMessage("Do nothing.").Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx2.Rollback()))
	}

	s4, err := tx2.Section.Create().
		SetSlug("last-section").
		SetNumber(2).
		SetTitle("Last Section").
		SetOverviewSentence("section to be committed").
		AddCommits(c5).Save(ctx)
	if err != nil {
		t.Fatal(multierr.Append(err, tx2.Rollback()))
	}

	if _, err := tx2.Page.Create().
		SetPathname("/committed").
		SetNumber(2).
		SetTitle("Last Page").
		SetOverviewSentence("page to be committed").
		AddSections(s3, s4).Save(ctx); err != nil {
		t.Fatal(multierr.Append(err, tx2.Rollback()))
	}

	if err := tx2.Commit(); err != nil {
		t.Fatal(err)
	}
}
