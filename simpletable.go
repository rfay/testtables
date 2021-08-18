package main

import (
    "os"

    "github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	simple()
	colored()
	light()
}
func simple() {
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
    t.AppendRows([]table.Row{
        {1, "Arya", "Stark", 3000},
        {20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
    })
    t.AppendSeparator()
    t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
    t.AppendFooter(table.Row{"", "", "Total", 10000})
    t.Render()
}
func colored() {
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
    t.AppendRows([]table.Row{
        {1, "Arya", "Stark", 3000},
        {20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
    })
    t.AppendSeparator()
    t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
    t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.SetStyle(table.StyleColoredBright)
    t.Render()

}
func light() {
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
    t.AppendRows([]table.Row{
        {1, "Arya", "Stark", 3000},
        {20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
    })
    t.AppendSeparator()
    t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
    t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.SetStyle(table.StyleLight)
    t.Render()

}
