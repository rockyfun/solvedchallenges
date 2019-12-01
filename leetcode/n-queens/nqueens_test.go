package nqueens

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	type pair struct {
		n int
		b [][]string
	}
	table := []*pair{
		{1, [][]string{{"Q"}}},
		{2, nil},
		{3, nil},
		{4, [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}},
		{5, [][]string{
			{"Q....", "..Q..", "....Q", ".Q...", "...Q."},
			{"Q....", "...Q.", ".Q...", "....Q", "..Q.."},
			{".Q...", "...Q.", "Q....", "..Q..", "....Q"},
			{".Q...", "....Q", "..Q..", "Q....", "...Q."},
			{"..Q..", "Q....", "...Q.", ".Q...", "....Q"},
			{"..Q..", "....Q", ".Q...", "...Q.", "Q...."},
			{"...Q.", "Q....", "..Q..", "....Q", ".Q..."},
			{"...Q.", ".Q...", "....Q", "..Q..", "Q...."},
			{"....Q", ".Q...", "...Q.", "Q....", "..Q.."},
			{"....Q", "..Q..", "Q....", "...Q.", ".Q..."},
		}},
		{6, [][]string{
			{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."},
			{"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."},
			{"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."},
			{"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."},
		}},
		{7, [][]string{
			{"Q......", "..Q....", "....Q..", "......Q", ".Q.....", "...Q...", ".....Q."},
			{"Q......", "...Q...", "......Q", "..Q....", ".....Q.", ".Q.....", "....Q.."},
			{"Q......", "....Q..", ".Q.....", ".....Q.", "..Q....", "......Q", "...Q..."},
			{"Q......", ".....Q.", "...Q...", ".Q.....", "......Q", "....Q..", "..Q...."},
			{".Q.....", "...Q...", "Q......", "......Q", "....Q..", "..Q....", ".....Q."},
			{".Q.....", "...Q...", ".....Q.", "Q......", "..Q....", "....Q..", "......Q"},
			{".Q.....", "....Q..", "Q......", "...Q...", "......Q", "..Q....", ".....Q."},
			{".Q.....", "....Q..", "..Q....", "Q......", "......Q", "...Q...", ".....Q."},
			{".Q.....", "....Q..", "......Q", "...Q...", "Q......", "..Q....", ".....Q."},
			{".Q.....", ".....Q.", "..Q....", "......Q", "...Q...", "Q......", "....Q.."},
			{".Q.....", "......Q", "....Q..", "..Q....", "Q......", ".....Q.", "...Q..."},
			{"..Q....", "Q......", ".....Q.", ".Q.....", "....Q..", "......Q", "...Q..."},
			{"..Q....", "Q......", ".....Q.", "...Q...", ".Q.....", "......Q", "....Q.."},
			{"..Q....", "....Q..", "......Q", ".Q.....", "...Q...", ".....Q.", "Q......"},
			{"..Q....", ".....Q.", ".Q.....", "....Q..", "Q......", "...Q...", "......Q"},
			{"..Q....", "......Q", ".Q.....", "...Q...", ".....Q.", "Q......", "....Q.."},
			{"..Q....", "......Q", "...Q...", "Q......", "....Q..", ".Q.....", ".....Q."},
			{"...Q...", "Q......", "..Q....", ".....Q.", ".Q.....", "......Q", "....Q.."},
			{"...Q...", "Q......", "....Q..", ".Q.....", ".....Q.", "..Q....", "......Q"},
			{"...Q...", ".Q.....", "......Q", "....Q..", "..Q....", "Q......", ".....Q."},
			{"...Q...", ".....Q.", "Q......", "..Q....", "....Q..", "......Q", ".Q....."},
			{"...Q...", "......Q", "..Q....", ".....Q.", ".Q.....", "....Q..", "Q......"},
			{"...Q...", "......Q", "....Q..", ".Q.....", ".....Q.", "Q......", "..Q...."},
			{"....Q..", "Q......", "...Q...", "......Q", "..Q....", ".....Q.", ".Q....."},
			{"....Q..", "Q......", ".....Q.", "...Q...", ".Q.....", "......Q", "..Q...."},
			{"....Q..", ".Q.....", ".....Q.", "..Q....", "......Q", "...Q...", "Q......"},
			{"....Q..", "..Q....", "Q......", ".....Q.", "...Q...", ".Q.....", "......Q"},
			{"....Q..", "......Q", ".Q.....", "...Q...", ".....Q.", "Q......", "..Q...."},
			{"....Q..", "......Q", ".Q.....", ".....Q.", "..Q....", "Q......", "...Q..."},
			{".....Q.", "Q......", "..Q....", "....Q..", "......Q", ".Q.....", "...Q..."},
			{".....Q.", ".Q.....", "....Q..", "Q......", "...Q...", "......Q", "..Q...."},
			{".....Q.", "..Q....", "Q......", "...Q...", "......Q", "....Q..", ".Q....."},
			{".....Q.", "..Q....", "....Q..", "......Q", "Q......", "...Q...", ".Q....."},
			{".....Q.", "..Q....", "......Q", "...Q...", "Q......", "....Q..", ".Q....."},
			{".....Q.", "...Q...", ".Q.....", "......Q", "....Q..", "..Q....", "Q......"},
			{".....Q.", "...Q...", "......Q", "Q......", "..Q....", "....Q..", ".Q....."},
			{"......Q", ".Q.....", "...Q...", ".....Q.", "Q......", "..Q....", "....Q.."},
			{"......Q", "..Q....", ".....Q.", ".Q.....", "....Q..", "Q......", "...Q..."},
			{"......Q", "...Q...", "Q......", "....Q..", ".Q.....", ".....Q.", "..Q...."},
			{"......Q", "....Q..", "..Q....", "Q......", ".....Q.", "...Q...", ".Q....."},
		}},
		// {8, [][]string{
		// 	{"Q.......", "....Q...", ".......Q", ".....Q..", "..Q.....", "......Q.", ".Q......", "...Q...."},
		// 	{"Q.......", ".....Q..", ".......Q", "..Q.....", "......Q.", "...Q....", ".Q......", "....Q..."},
		// 	{"Q.......", "......Q.", "...Q....", ".....Q..", ".......Q", ".Q......", "....Q...", "..Q....."},
		// 	{"Q.......", "......Q.", "....Q...", ".......Q", ".Q......", "...Q....", ".....Q..", "..Q....."},
		// 	{".Q......", "...Q....", ".....Q..", ".......Q", "..Q.....", "Q.......", "......Q.", "....Q..."},
		// 	{".Q......", "....Q...", "......Q.", "Q.......", "..Q.....", ".......Q", ".....Q..", "...Q...."},
		// 	{".Q......", "....Q...", "......Q.", "...Q....", "Q.......", ".......Q", ".....Q..", "..Q....."},
		// 	{".Q......", ".....Q..", "Q.......", "......Q.", "...Q....", ".......Q", "..Q.....", "....Q..."},
		// 	{".Q......", ".....Q..", ".......Q", "..Q.....", "Q.......", "...Q....", "......Q.", "....Q..."},
		// 	{".Q......", "......Q.", "..Q.....", ".....Q..", ".......Q", "....Q...", "Q.......", "...Q...."},
		// 	{".Q......", "......Q.", "....Q...", ".......Q", "Q.......", "...Q....", ".....Q..", "..Q....."},
		// 	{".Q......", ".......Q", ".....Q..", "Q.......", "..Q.....", "....Q...", "......Q.", "...Q...."},
		// 	{"..Q.....", "Q.......", "......Q.", "....Q...", ".......Q", ".Q......", "...Q....", ".....Q.."},
		// 	{"..Q.....", "....Q...", ".Q......", ".......Q", "Q.......", "......Q.", "...Q....", ".....Q.."},
		// 	{"..Q.....", "....Q...", ".Q......", ".......Q", ".....Q..", "...Q....", "......Q.", "Q......."},
		// 	{"..Q.....", "....Q...", "......Q.", "Q.......", "...Q....", ".Q......", ".......Q", ".....Q.."},
		// 	{"..Q.....", "....Q...", ".......Q", "...Q....", "Q.......", "......Q.", ".Q......", ".....Q.."},
		// 	{"..Q.....", ".....Q..", ".Q......", "....Q...", ".......Q", "Q.......", "......Q.", "...Q...."},
		// 	{"..Q.....", ".....Q..", ".Q......", "......Q.", "Q.......", "...Q....", ".......Q", "....Q..."},
		// 	{"..Q.....", ".....Q..", ".Q......", "......Q.", "....Q...", "Q.......", ".......Q", "...Q...."},
		// 	{"..Q.....", ".....Q..", "...Q....", "Q.......", ".......Q", "....Q...", "......Q.", ".Q......"},
		// 	{"..Q.....", ".....Q..", "...Q....", ".Q......", ".......Q", "....Q...", "......Q.", "Q......."},
		// 	{"..Q.....", ".....Q..", ".......Q", "Q.......", "...Q....", "......Q.", "....Q...", ".Q......"},
		// 	{"..Q.....", ".....Q..", ".......Q", "Q.......", "....Q...", "......Q.", ".Q......", "...Q...."},
		// 	{"..Q.....", ".....Q..", ".......Q", ".Q......", "...Q....", "Q.......", "......Q.", "....Q..."},
		// 	{"..Q.....", "......Q.", ".Q......", ".......Q", "....Q...", "Q.......", "...Q....", ".....Q.."},
		// 	{"..Q.....", "......Q.", ".Q......", ".......Q", ".....Q..", "...Q....", "Q.......", "....Q..."},
		// 	{"..Q.....", ".......Q", "...Q....", "......Q.", "Q.......", ".....Q..", ".Q......", "....Q..."},
		// 	{"...Q....", "Q.......", "....Q...", ".......Q", ".Q......", "......Q.", "..Q.....", ".....Q.."},
		// 	{"...Q....", "Q.......", "....Q...", ".......Q", ".....Q..", "..Q.....", "......Q.", ".Q......"},
		// 	{"...Q....", ".Q......", "....Q...", ".......Q", ".....Q..", "Q.......", "..Q.....", "......Q."},
		// 	{"...Q....", ".Q......", "......Q.", "..Q.....", ".....Q..", ".......Q", "Q.......", "....Q..."},
		// 	{"...Q....", ".Q......", "......Q.", "..Q.....", ".....Q..", ".......Q", "....Q...", "Q......."},
		// 	{"...Q....", ".Q......", "......Q.", "....Q...", "Q.......", ".......Q", ".....Q..", "..Q....."},
		// 	{"...Q....", ".Q......", ".......Q", "....Q...", "......Q.", "Q.......", "..Q.....", ".....Q.."},
		// 	{"...Q....", ".Q......", ".......Q", ".....Q..", "Q.......", "..Q.....", "....Q...", "......Q."},
		// 	{"...Q....", ".....Q..", "Q.......", "....Q...", ".Q......", ".......Q", "..Q.....", "......Q."},
		// 	{"...Q....", ".....Q..", ".......Q", ".Q......", "......Q.", "Q.......", "..Q.....", "....Q..."},
		// 	{"...Q....", ".....Q..", ".......Q", "..Q.....", "Q.......", "......Q.", "....Q...", ".Q......"},
		// 	{"...Q....", "......Q.", "Q.......", ".......Q", "....Q...", ".Q......", ".....Q..", "..Q....."},
		// 	{"...Q....", "......Q.", "..Q.....", ".......Q", ".Q......", "....Q...", "Q.......", ".....Q.."},
		// 	{"...Q....", "......Q.", "....Q...", ".Q......", ".....Q..", "Q.......", "..Q.....", ".......Q"},
		// 	{"...Q....", "......Q.", "....Q...", "..Q.....", "Q.......", ".....Q..", ".......Q", ".Q......"},
		// 	{"...Q....", ".......Q", "Q.......", "..Q.....", ".....Q..", ".Q......", "......Q.", "....Q..."},
		// 	{"...Q....", ".......Q", "Q.......", "....Q...", "......Q.", ".Q......", ".....Q..", "..Q....."},
		// 	{"...Q....", ".......Q", "....Q...", "..Q.....", "Q.......", "......Q.", ".Q......", ".....Q.."},
		// 	{"....Q...", "Q.......", "...Q....", ".....Q..", ".......Q", ".Q......", "......Q.", "..Q....."},
		// 	{"....Q...", "Q.......", ".......Q", "...Q....", ".Q......", "......Q.", "..Q.....", ".....Q.."},
		// 	{"....Q...", "Q.......", ".......Q", ".....Q..", "..Q.....", "......Q.", ".Q......", "...Q...."},
		// 	{"....Q...", ".Q......", "...Q....", ".....Q..", ".......Q", "..Q.....", "Q.......", "......Q."},
		// 	{"....Q...", ".Q......", "...Q....", "......Q.", "..Q.....", ".......Q", ".....Q..", "Q......."},
		// 	{"....Q...", ".Q......", ".....Q..", "Q.......", "......Q.", "...Q....", ".......Q", "..Q....."},
		// 	{"....Q...", ".Q......", ".......Q", "Q.......", "...Q....", "......Q.", "..Q.....", ".....Q.."},
		// 	{"....Q...", "..Q.....", "Q.......", ".....Q..", ".......Q", ".Q......", "...Q....", "......Q."},
		// 	{"....Q...", "..Q.....", "Q.......", "......Q.", ".Q......", ".......Q", ".....Q..", "...Q...."},
		// 	{"....Q...", "..Q.....", ".......Q", "...Q....", "......Q.", "Q.......", ".....Q..", ".Q......"},
		// 	{"....Q...", "......Q.", "Q.......", "..Q.....", ".......Q", ".....Q..", "...Q....", ".Q......"},
		// 	{"....Q...", "......Q.", "Q.......", "...Q....", ".Q......", ".......Q", ".....Q..", "..Q....."},
		// 	{"....Q...", "......Q.", ".Q......", "...Q....", ".......Q", "Q.......", "..Q.....", ".....Q.."},
		// 	{"....Q...", "......Q.", ".Q......", ".....Q..", "..Q.....", "Q.......", "...Q....", ".......Q"},
		// 	{"....Q...", "......Q.", ".Q......", ".....Q..", "..Q.....", "Q.......", ".......Q", "...Q...."},
		// 	{"....Q...", "......Q.", "...Q....", "Q.......", "..Q.....", ".......Q", ".....Q..", ".Q......"},
		// 	{"....Q...", ".......Q", "...Q....", "Q.......", "..Q.....", ".....Q..", ".Q......", "......Q."},
		// 	{"....Q...", ".......Q", "...Q....", "Q.......", "......Q.", ".Q......", ".....Q..", "..Q....."},
		// 	{".....Q..", "Q.......", "....Q...", ".Q......", ".......Q", "..Q.....", "......Q.", "...Q...."},
		// 	{".....Q..", ".Q......", "......Q.", "Q.......", "..Q.....", "....Q...", ".......Q", "...Q...."},
		// 	{".....Q..", ".Q......", "......Q.", "Q.......", "...Q....", ".......Q", "....Q...", "..Q....."},
		// 	{".....Q..", "..Q.....", "Q.......", "......Q.", "....Q...", ".......Q", ".Q......", "...Q...."},
		// 	{".....Q..", "..Q.....", "Q.......", ".......Q", "...Q....", ".Q......", "......Q.", "....Q..."},
		// 	{".....Q..", "..Q.....", "Q.......", ".......Q", "....Q...", ".Q......", "...Q....", "......Q."},
		// 	{".....Q..", "..Q.....", "....Q...", "......Q.", "Q.......", "...Q....", ".Q......", ".......Q"},
		// 	{".....Q..", "..Q.....", "....Q...", ".......Q", "Q.......", "...Q....", ".Q......", "......Q."},
		// 	{".....Q..", "..Q.....", "......Q.", ".Q......", "...Q....", ".......Q", "Q.......", "....Q..."},
		// 	{".....Q..", "..Q.....", "......Q.", ".Q......", ".......Q", "....Q...", "Q.......", "...Q...."},
		// 	{".....Q..", "..Q.....", "......Q.", "...Q....", "Q.......", ".......Q", ".Q......", "....Q..."},
		// 	{".....Q..", "...Q....", "Q.......", "....Q...", ".......Q", ".Q......", "......Q.", "..Q....."},
		// 	{".....Q..", "...Q....", ".Q......", ".......Q", "....Q...", "......Q.", "Q.......", "..Q....."},
		// 	{".....Q..", "...Q....", "......Q.", "Q.......", "..Q.....", "....Q...", ".Q......", ".......Q"},
		// 	{".....Q..", "...Q....", "......Q.", "Q.......", ".......Q", ".Q......", "....Q...", "..Q....."},
		// 	{".....Q..", ".......Q", ".Q......", "...Q....", "Q.......", "......Q.", "....Q...", "..Q....."},
		// 	{"......Q.", "Q.......", "..Q.....", ".......Q", ".....Q..", "...Q....", ".Q......", "....Q..."},
		// 	{"......Q.", ".Q......", "...Q....", "Q.......", ".......Q", "....Q...", "..Q.....", ".....Q.."},
		// 	{"......Q.", ".Q......", ".....Q..", "..Q.....", "Q.......", "...Q....", ".......Q", "....Q..."},
		// 	{"......Q.", "..Q.....", "Q.......", ".....Q..", ".......Q", "....Q...", ".Q......", "...Q...."},
		// 	{"......Q.", "..Q.....", ".......Q", ".Q......", "....Q...", "Q.......", ".....Q..", "...Q...."},
		// 	{"......Q.", "...Q....", ".Q......", "....Q...", ".......Q", "Q.......", "..Q.....", ".....Q.."},
		// 	{"......Q.", "...Q....", ".Q......", ".......Q", ".....Q..", "Q.......", "..Q.....", "....Q..."},
		// 	{"......Q.", "....Q...", "..Q.....", "Q.......", ".....Q..", ".......Q", ".Q......", "...Q...."},
		// 	{".......Q", ".Q......", "...Q....", "Q.......", "......Q.", "....Q...", "..Q.....", ".....Q.."},
		// 	{".......Q", ".Q......", "....Q...", "..Q.....", "Q.......", "......Q.", "...Q....", ".....Q.."},
		// 	{".......Q", "..Q.....", "Q.......", ".....Q..", ".Q......", "....Q...", "......Q.", "...Q...."},
		// 	{".......Q", "...Q....", "Q.......", "..Q.....", ".....Q..", ".Q......", "......Q.", "....Q..."},
		// }},
	}

	for _, r := range table {

		res := solveNQueens(r.n)

		if !reflect.DeepEqual(r.b, res) {
			t.Logf("(%d) got: \n%s", r.n, printSol(res))
			t.Logf("(%d) want: \n%s", r.n, printSol(r.b))
			t.FailNow()
		}
	}
}

func printSol(sol [][]string) string {
	if sol == nil {
		return fmt.Sprintf("<nil>\n")
	}
	if len(sol) == 0 {
		return "[]"
	}
	var str string
	for _, b := range sol {

		var rows []string
		for _, r := range b {
			rows = append(rows, fmt.Sprintf("%q", r))
		}
		str += "{" + strings.Join(rows, ", ") + "},"
		// for _, row := range b {
		// 	str += fmt.Sprintf("%s\n", row)
		// }
		str += fmt.Sprintf("\n")
	}
	return str
}

func TestCountQs(t *testing.T) {

	c := countQs([][]byte{
		[]byte(".Q.."),
		[]byte("...Q"),
		[]byte("Q..."),
		[]byte("...."),
	})

	if c != 3 {
		t.FailNow()
	}

	c = countQs([][]byte{
		[]byte(".Q.."),
		[]byte("...Q"),
		[]byte("Q..."),
		[]byte("..Q."),
	})

	if c != 4 {
		t.FailNow()
	}
	// t.Log(c)
}

func TestCheckQ(t *testing.T) {
	type pair struct {
		ex   bool
		r, c int
		b    [][]byte
	}
	table := []*pair{
		{true, 3, 2, [][]byte{
			[]byte(".Q.."),
			[]byte("...Q"),
			[]byte("Q..."),
			[]byte("...."),
		}},
		{true, 1, 3, [][]byte{
			[]byte(".Q.."),
			[]byte("...."),
			[]byte("Q..."),
			[]byte("...."),
		}},
		// colums
		{false, 0, 1, [][]byte{
			[]byte("...Q"),
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
		}},
		{false, 3, 3, [][]byte{
			[]byte("...Q"),
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
		}},
		// diagonal rights
		{false, 2, 1, [][]byte{
			[]byte("...Q"),
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
		}},
		{false, 1, 0, [][]byte{
			[]byte(".Q.."),
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
		}},
		{false, 1, 1, [][]byte{
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
			[]byte("...Q"),
		}},
		{false, 0, 2, [][]byte{
			[]byte("...."),
			[]byte("...Q"),
			[]byte("...."),
			[]byte("...Q"),
		}},
		// diagonal lefts
		{false, 2, 2, [][]byte{
			[]byte("...."),
			[]byte("...."),
			[]byte("...."),
			[]byte("...Q"),
		}},
		{false, 3, 3, [][]byte{
			[]byte("...."),
			[]byte(".Q.."),
			[]byte("...."),
			[]byte("...."),
		}},
	}

	for _, r := range table {
		res := checkQ(r.r, r.c, 4, r.b)
		if res != r.ex {
			t.Fatalf("didn't check with %d %d in board: %v", r.r, r.c, r.b)
		}
	}
}
