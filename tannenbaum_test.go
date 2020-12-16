package main

import (
	"testing"
)

func assertEquals(got string, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func Test_PrintChristmasTree(t *testing.T) {
	t.Run("a small tannenbaum with 1 line", func(t *testing.T) {
		got := PrintChristmasTree(1)
		want := "*"
		assertEquals(got, want, t)
	})
	t.Run("a not so small tannenbaum with 2 lines", func(t *testing.T) {
		got := PrintChristmasTree(2)
		want := " * \n" +
			"***"
		assertEquals(got, want, t)
	})
	t.Run("a tannenbaum with 3 lines", func(t *testing.T) {
		got := PrintChristmasTree(3)
		want := "  *  \n" +
			" *** \n" +
			"*****"
		assertEquals(got, want, t)
	})
	t.Run("a giant tannenbaum with 10 lines", func(t *testing.T) {
		got := PrintChristmasTree(10)
		want :=
			"         *         \n" +
				"        ***        \n" +
				"       *****       \n" +
				"      *******      \n" +
				"     *********     \n" +
				"    ***********    \n" +
				"   *************   \n" +
				"  ***************  \n" +
				" ***************** \n" +
				"*******************"
		assertEquals(got, want, t)
	})
}
