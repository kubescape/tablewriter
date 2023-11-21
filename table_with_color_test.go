// Copyright 2014 Oleku Konko All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// This module is a Table Writer  API for the Go Programming Language.
// The protocols were written in pure Go and works on windows and unix systems

package tablewriter

import (
	"os"
	"testing"
)

func TestSetHeaderColorNonTTY(t *testing.T) {
	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}

	os.Stdout = nil
	table := NewWriter(os.Stdout)

	table.SetHeader([]string{"Name", "Sign", "Rating"})
	want := table.headerParams
	table.SetHeaderColor(Colors{Bold, FgHiYellowColor}, Colors{Bold, FgHiYellowColor}, Colors{Bold, FgHiYellowColor})
	table.AppendBulk(data)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.Render()

	// The color codes are not added in case of non TTY output.
	got := table.headerParams

	checkEqual(t, got, want, "SetHeaderColor when TTY is not attached failed")
}

func TestSetColumnColorNonTTY(t *testing.T) {
	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}

	os.Stdout = nil
	table := NewWriter(os.Stdout)

	want := table.columnsParams
	table.SetColumnColor(Colors{Bold, FgHiYellowColor}, Colors{Bold, FgHiYellowColor}, Colors{Bold, FgHiYellowColor})
	table.AppendBulk(data)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.Render()

	// The color codes are not added in case of non TTY output.
	got := table.columnsParams

	checkEqual(t, got, want, "SetColumnColor when TTY is not attached failed")
}

func TestSetFooterColorNonTTY(t *testing.T) {
	data := [][]string{
		{"Regular", "regular line", "1"},
		{"Thick", "particularly thick line", "2"},
		{"Double", "double line", "3"},
	}

	// Set stdout to nil to simulate not being a terminal
	os.Stdout = nil

	table := NewWriter(os.Stdout)
	table.SetFooter([]string{"Constant", "Meaning", "Seq"})
	want := table.footerParams
	table.SetFooterColor(Colors{Bold, FgHiYellowColor}, Colors{Bold, FgHiYellowColor}, Colors{Bold, FgHiYellowColor})
	table.AppendBulk(data)

	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.Render()

	// The color codes are added in case of TTY output.
	got := table.footerParams

	checkEqual(t, got, want, "SetHeaderColor when TTY is not attached failed")
}
