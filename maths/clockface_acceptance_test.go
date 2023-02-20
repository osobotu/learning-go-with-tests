package clockface

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"
)

// func TestSecondHandAtMidnight(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 - 90}
// 	got := SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v, want %v", got, want)
// 	}
// }

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 + 90}
// 	got := SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v, wanted %v", got, want)
// 	}
// }

// func TestSecondHand(t *testing.T) {

// 	cases := []struct {
// 		time  time.Time
// 		point Point
// 	}{
// 		{time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC), Point{150, 150 - 90}},
// 		{time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC), Point{150, 150 + 90}},
// 		{time.Date(1337, time.January, 1, 0, 0, 45, 0, time.UTC), Point{150 - 90, 150}},
// 		{time.Date(1337, time.January, 1, 0, 0, 15, 0, time.UTC), Point{150 + 90, 150}},
// 	}

// 	for _, c := range cases {
// 		t.Run(testName(c.time), func(t *testing.T) {
// 			got := SecondHand(c.time)
// 			if !roughlyEqualPoint(got, c.point) {
// 				t.Fatalf("Wanted %v, but got %v", c.point, got)
// 			}
// 		})
// 	}

// }

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})

	}

}
