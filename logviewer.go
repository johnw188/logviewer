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
	"errors"
	"time"
)

const HelpText string = "[h] and [l] to navigate, [t] to show timestamps, [q] to exit"

// A Line represents a single line of log content
type Line struct {
	Log       string
	Timestamp time.Time
}

// A Viewer is used to display log data from one or more sources
type Viewer struct {
	// Name for the viewer, displayed in the help text at the top of the screen
	feeds         []*Feed
	tabpane       *extra.Tabpane
	title         *ui.Par
	initialized   bool
	showTimestamp bool
}

type Feed struct {
	name       string
	logs       []*Line
	maxHistory int
	par        *ui.Par
	viewer     *Viewer
}

// NewViewer initializes a Viewer with a given name
func NewViewer(name string) *Viewer {
	viewer := &Viewer{
		feeds:         []*Feed{},
		tabpane:       extra.NewTabpane(),
		title:         ui.NewPar("[" + name + "](fg-green,fg-bold)   " + HelpText),
		initialized:   false,
		showTimestamp: false,
	}
	viewer.tabpane.Border = true
	viewer.title.Height = 1
	viewer.title.Border = false

	return viewer
}

func newLogFeed(name string, maxHistory int) *Feed {
	feed := Feed{
		name:       name,
		logs:       []*Line{},
		maxHistory: maxHistory,
	}
	return &feed
}

// AddLogFeed adds a new Feed to the Viewer, returning the index of the Feed for use in
// future calls to AddLogLine
func (v*Viewer) AddLogFeed(name string, maxHistory int) *Feed {
	feed := newLogFeed(name, maxHistory)
	feed.viewer = v
	v.feeds = append(v.feeds, feed)

	feed.par = ui.NewPar(" ")
	feed.par.Border = false

	tab := extra.NewTab(feed.name)
	tab.AddBlocks(feed.par)

	v.tabpane.SetTabs(append(v.tabpane.Tabs, *tab)...)
	v.render()
	return feed
}

// RemoveLogFeed removes a given Feed from the Viewer
func (v*Viewer) RemoveLogFeed(f *Feed) error {
	for i := range v.feeds {
		if v.feeds[i] == f {
			// TODO this probably leaks
			tabs := v.tabpane.Tabs
			v.tabpane.SetTabs(append(tabs[:i], tabs[i+1:]...)...)
			v.render()
			return nil
		}
	}
	return errors.New("Feed not found")
}

// AddLogLine adds a Line to a Feed,
func (f*Feed) AddLogLine(l *Line) {
	f.logs = append(f.logs, l)
	// TODO circular queue with size of maxHistory
	if len(f.logs) > f.maxHistory {
		over := len(f.logs) - f.maxHistory
		f.logs = f.logs[over:]
	}
	f.updatePar()
	f.viewer.render()
}

func (f*Feed) updatePar() {
	h := f.par.GetHeight()

	var text bytes.Buffer
	logs := f.logs
	if h < len(f.logs) {
		logs = logs[len(f.logs)-h:]
	}
	for i, line := range logs {
		if f.viewer.showTimestamp {
			text.WriteString("[")
			text.WriteString(line.Timestamp.Format("2006-01-02T15:04:05-0700"))
			text.WriteString("] ")
		}
		text.WriteString(line.Log)
		if i != len(logs)-1 {
			text.WriteString("\n")
		}
	}
	f.par.Text = text.String()
}

func (v*Viewer) titleString(width int) string {
	var title bytes.Buffer
	for i, feed := range v.feeds {
		title.WriteString("[")
		title.WriteString(strconv.Itoa(i))
		title.WriteString("]")
		if len(feed.name)+title.Len() > width-3 {

		}
		title.WriteString(feed.name)
		title.WriteString(" ")
		if title.Len() > width-5 {
			title.WriteString("...")
			break
		}
	}
	return title.String()
}

func (v*Viewer) render() {
	if v.initialized {
		v.updateDimensions()
		ui.Body.Align()
		ui.Render(ui.Body)
	}
}

func (v*Viewer) updateFeedDimensions(w, h int) {
	for _, f := range v.feeds {
		f.par.Width = w
		f.par.Height = h
	}
}

// Display starts the viewer. This call will block until the log display is exited by the user
func (v*Viewer) Display() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()
	v.initialized = true

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, v.title)),
		ui.NewRow(
			ui.NewCol(12, 0, v.tabpane)))

	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		v.render()
	})

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/l", func(ui.Event) {
		v.tabpane.SetActiveRight()
		v.render()
	})

	ui.Handle("/sys/kbd/h", func(ui.Event) {
		v.tabpane.SetActiveLeft()
		v.render()
	})

	ui.Handle("/sys/kbd/t", func(ui.Event) {
		v.showTimestamp = !v.showTimestamp
		for _, f := range v.feeds {
			f.updatePar()
		}
		v.render()
	})

	v.render()
	ui.Loop()
}

func (v*Viewer) updateDimensions() {
	w := ui.TermWidth()
	h := ui.TermHeight() - v.tabpane.Height
	ui.Body.Width = w
	v.updateFeedDimensions(w, h)
}
