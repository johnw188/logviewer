// Copyright 2017 John Welsh <john.welsh@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

// Package logviewer provides a terminal UI for viewing log streams from multiple sources
package logviewer

import (
	ui "github.com/gizak/termui"
	"bytes"
	"strconv"
	"github.com/gizak/termui/extra"
)

// LogLevel enables filtering/coloring
type LogLevel uint16

// 4 log levels by default
const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// A LogLine represents a single line of log content
type LogLine struct {
	Log string
	LogLevel LogLevel
}

type logFeed struct {
	name string
	logs []*LogLine
	maxHistory int
	par *ui.Par
}

type LogViewer struct {
	feeds []*logFeed
	tabpane *extra.Tabpane
	uiInit bool
}

func NewLogViewer() *LogViewer {
	viewer := &LogViewer{
		feeds:   []*logFeed{},
		tabpane: extra.NewTabpane(),
		uiInit:  false,
	}
	viewer.tabpane.Border = false
	return viewer
}

func newLogFeed(name string, maxHistory int) *logFeed {
	feed := logFeed{
		name:       name,
		logs:       []*LogLine{},
		maxHistory: maxHistory,
	}
	return &feed
}

func (lv* LogViewer) AddLogFeed(name string, maxHistory int) int {
	feed := newLogFeed(name, maxHistory)
	lv.feeds = append(lv.feeds, feed)

	feed.par = ui.NewPar(" ")
	feed.par.Border = false

	tab := extra.NewTab(feed.name)
	tab.AddBlocks(feed.par)

	lv.tabpane.SetTabs(append(lv.tabpane.Tabs, *tab)...)
	lv.render()
	return len(lv.feeds) - 1
}

func (lv* LogViewer) RemoveLogFeed(idx int) {
	tabs := lv.tabpane.Tabs
	lv.tabpane.SetTabs(append(tabs[:idx],tabs[idx+1:]...)...)
	lv.render()
}

func (lv* LogViewer) AddLogLine(l *LogLine, idx int) {
	f := lv.feeds[idx]
	f.logs = append(f.logs, l)
	if len(f.logs) > f.maxHistory {
		over := len(f.logs) - f.maxHistory
		f.logs = f.logs[over:]
	}
	f.updatePar()
	lv.render()
}

func (f* logFeed) updatePar() {
	h := f.par.GetHeight()

	var text bytes.Buffer
	logs := f.logs
	if h < len(f.logs) {
		logs = logs[len(f.logs)-h:]
	}
	for i, line := range logs {
		text.WriteString(line.Log)
		if i != len(logs) - 1 {
			text.WriteString("\n")
		}
	}
	f.par.Text = text.String()
}

func (lv* LogViewer) titleString(width int) string {
	var title bytes.Buffer
	for i, feed := range lv.feeds {
		title.WriteString("[")
		title.WriteString(strconv.Itoa(i))
		title.WriteString("]")
		if len(feed.name) + title.Len() > width - 3 {

		}
		title.WriteString(feed.name)
		title.WriteString(" ")
		if title.Len() > width - 5 {
			title.WriteString("...")
			break
		}
	}
	return title.String()
}

func (lv* LogViewer) render() {
	if lv.uiInit {
		lv.updateDimensions()
		ui.Body.Align()
		ui.Render(ui.Body)
	}
}

func (lv* LogViewer) updateFeedDimensions(w, h int) {
	for _, f := range lv.feeds {
		f.par.Width = w
		f.par.Height = h
	}
}

func (lv* LogViewer) Display() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()
	lv.uiInit = true

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, lv.tabpane)))

	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		lv.render()
	})

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/l", func(ui.Event) {
		lv.tabpane.SetActiveRight()
		lv.render()
	})

	ui.Handle("/sys/kbd/h", func(ui.Event) {
		lv.tabpane.SetActiveLeft()
		lv.render()
	})

	lv.render()
	ui.Loop()
}

func (lv* LogViewer) updateDimensions() {
	w := ui.TermWidth()
	h := ui.TermHeight() - lv.tabpane.Height
	ui.Body.Width = w
	lv.updateFeedDimensions(w, h)
}


