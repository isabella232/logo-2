package main

import (
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo"
)

var fabiofont = `<font id="fabio" horiz-adv-x="1024" >
  <font-face
    font-family="fabio"
    font-weight="400"
    font-stretch="normal"
    units-per-em="2048"
    panose-1="2 11 6 7 3 5 2 6 2 4"
    ascent="1606"
    descent="-442"
    x-height="949"
    cap-height="1365"
    bbox="-337 -512 2943 1983"
    underline-thickness="130"
    underline-position="-236"
    unicode-range="U+0020-F002"
  />
    <glyph glyph-name="F" unicode="F" horiz-adv-x="1384" d="M102 0v1365h1282v-160h-1100v-432h1051v-160h-1051v-613h-182z" />
    <glyph glyph-name="a" unicode="a" horiz-adv-x="1289" d="M246 664h-156q-2 26 -2 41q0 149 112 203t420 54q242 0 348 -22t166 -84q50 -50 65 -118t15 -238v-500h-149l3 111h-8q-56 -78 -134.5 -100.5t-296.5 -22.5q-249 0 -346 12t-147 50q-50 36 -69.5 86t-19.5 140q0 110 29.5 169t99.5 88q76 33 355 33q292 0 382 -17.5
t139 -84.5h13v102q0 125 -45 180q-40 49 -121.5 66.5t-266.5 17.5q-243 0 -314.5 -28.5t-71.5 -124.5v-13zM630 433q-185 0 -259 -6.5t-105 -24.5q-59 -35 -59 -129q0 -93 70 -126q57 -27 384 -27q239 0 316 31t77 126q0 92 -86 124t-338 32z" />
    <glyph glyph-name="b" unicode="b" horiz-adv-x="1379" d="M89 1365h150v-530h9q51 76 135.5 101t286.5 25q235 0 334 -15t162 -60q79 -57 113 -143q39 -99 39 -282q0 -202 -60 -303q-79 -132 -285 -159q-85 -11 -277 -11q-227 0 -312.5 25t-138.5 108h-8v-121h-148v1365zM673 829q-149 0 -225.5 -12t-118.5 -43
q-52 -36 -73.5 -105t-21.5 -197q0 -169 36 -233q42 -73 134.5 -96.5t343.5 -23.5q131 0 201 13t118 46t69.5 102t21.5 191q0 163 -38 235q-40 75 -127 99t-320 24z" />
    <glyph glyph-name="f" unicode="f" horiz-adv-x="829" d="M-14 821v128h212q4 141 17.5 202t52.5 105q49 56 129.5 78.5t257.5 28.5l173 5v-128l-161 -5q-129 -4 -185.5 -16.5t-85.5 -44.5q-25 -27 -32.5 -67t-10.5 -158h475v-128h-480v-821h-150v821h-212z" />
    <glyph glyph-name="i" unicode="i" horiz-adv-x="328" d="M89 949h150v-949h-150v949zM89 1365h150v-178h-150v178z" />
    <glyph glyph-name="o" unicode="o" horiz-adv-x="1373" d="M645 961q224 0 330.5 -13.5t166.5 -49.5q90 -54 130 -153.5t40 -270.5q0 -190 -48 -289q-35 -75 -93.5 -117t-153.5 -61q-92 -19 -306 -19q-202 0 -305 13.5t-167 48.5q-95 51 -136 150.5t-41 278.5q0 200 59 303q60 106 166 142t358 37zM669 829q-149 0 -225 -11
t-121 -40q-53 -34 -77 -107.5t-24 -200.5q0 -144 32 -215.5t110 -101.5q86 -33 335 -33q162 0 238.5 12t121.5 45q50 35 71.5 105t21.5 197q0 150 -34 220q-38 78 -126.5 104t-322.5 26z" />
    <hkern u1="F" u2="o" k="-5" />
    <hkern u1="F" u2="i" k="-21" />
    <hkern u1="F" u2="a" k="2" />
    <hkern u1="F" u2="O" k="-10" />
    <hkern u1="F" u2="A" k="113" />
    <hkern u1="f" u2="o" k="15" />
    <hkern u1="f" u2="i" k="-14" />
    <hkern u1="f" u2="f" k="-66" />
    <hkern u1="f" u2="a" k="-2" />
  </font>
`

func main() {
	var (
		// box dimensions
		bW   = 80     // box width
		bH   = bW / 2 // box height
		bRad = 10     // box corner radius
		bPad = 16     // box padding

		// other dimensions
		botY    = bH + 4*bPad      // top of bottom row
		centerX = bW + bPad + bW/2 // canvas center

		// canvas dimensions
		width  = 749 // by try and error
		height = botY + bH

		// colors
		darkBlue  = "#4294d3"
		lightBlue = "#21bdf0"

		// styles
		darkBoxStyle  = "fill:" + darkBlue
		lightBoxStyle = "fill:" + lightBlue
		lineStyle     = "stroke-width:6;stroke-linecap:round;stroke:" + lightBlue
		textStyle     = "text-anchor:start;font-size:160px;font-family:fabio;fill:" + darkBlue
	)

	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Title("Fabio")

	// use this to determine the width
	// canvas.Rect(0, 0, width, height, "fill:red")

	// font
	canvas.Def()
	fmt.Fprintf(canvas.Writer, fabiofont)
	canvas.DefEnd()

	// fabio
	tX := 3*bW + 4*bPad
	tY := botY + bH
	canvas.Text(tX, tY, "Fabio", textStyle)

	// boxes

	// box draws a box
	box := func(x, y int, style string) {
		canvas.Roundrect(x, y, bW, bH, bRad, bRad, style)
	}

	// top row box
	box(bW+bPad, 0, darkBoxStyle)

	// bottom row
	box(0, botY, darkBoxStyle)
	box(bW+bPad, botY, darkBoxStyle)
	box(2*bW+2*bPad, botY, lightBoxStyle)

	// lines

	// center vertical row
	cW, cH := 0, 6*bPad/2
	cX1, cY1 := centerX, bH+bPad/2
	cX2, cY2 := cX1+cW, cY1+cH
	canvas.Line(cX1, cY1, cX2, cY2, lineStyle)

	// left vertical row
	lvW, lvH := 0, 2*bPad/2
	lvX1, lvY1 := bW/2, bH+5*bPad/2
	lvX2, lvY2 := lvX1+lvW, lvY1+lvH
	canvas.Line(lvX1, lvY1, lvX2, lvY2, lineStyle)

	// right vertical row
	rvW, rvH := 0, lvH
	rvX1, rvY1 := 2*bW+2*bPad+bW/2, lvY1
	rvX2, rvY2 := rvX1+rvW, rvY1+rvH
	canvas.Line(rvX1, rvY1, rvX2, rvY2, lineStyle)

	// top cross row
	tcW, tcH := bW+2*bPad, 0
	tcX1, tcY1 := bW, bH+3*bPad/2
	tcX2, tcY2 := tcX1+tcW, tcY1+tcH
	canvas.Line(tcX1, tcY1, tcX2, tcY2, lineStyle)

	// diag left
	canvas.Line(tcX1, tcY1, lvX1, lvY1, lineStyle)

	// diag right
	canvas.Line(tcX2, tcY2, rvX1, rvY1, lineStyle)

	canvas.End()
}
