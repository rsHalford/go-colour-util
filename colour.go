/*
Copyright © 2022 Richard Halford <richard@xhalford.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package colour

const (
	blackFore     = "\x1b[30m"
	redFore     = "\x1b[31m"
	greenFore   = "\x1b[32m"
	yellowFore  = "\x1b[33m"
	blueFore    = "\x1b[34m"
	magentaFore = "\x1b[35m"
	cyanFore    = "\x1b[36m"
	whiteFore    = "\x1b[37m"

	blackBack     = "\x1b[40m"
	redBack     = "\x1b[41m"
	greenBack   = "\x1b[42m"
	yellowBack  = "\x1b[43m"
	blueBack    = "\x1b[44m"
	magentaBack = "\x1b[45m"
	cyanBack    = "\x1b[46m"
	whiteBack    = "\x1b[47m"

	brightBlackFore     = "\x1b[90m"
	brightRedFore     = "\x1b[91m"
	brightGreenFore   = "\x1b[92m"
	brightYellowFore  = "\x1b[93m"
	brightBlueFore    = "\x1b[94m"
	brightMagentaFore = "\x1b[95m"
	brightCyanFore    = "\x1b[96m"
	brightWhiteFore    = "\x1b[97m"

	brightBlackBack     = "\x1b[100m"
	brightRedBack     = "\x1b[101m"
	brightGreenBack   = "\x1b[102m"
	brightYellowBack  = "\x1b[103m"
	brightBlueBack    = "\x1b[104m"
	brightMagentaBack = "\x1b[105m"
	brightCyanBack    = "\x1b[106m"
	brightWhiteBack    = "\x1b[107m"

	bold   = "\x1b[1m"
	faint  = "\x1b[2m"
	italic = "\x1b[3m"
	underline = "\x1b[4m"
	slowBlink = "\x1b[5m"
	rapidBlink = "\x1b[6m"
	strike = "\x1b[9m"

	reset = "\x1b[0m"
)

const version string = "v0.1.0"

func style(start, s string) string {
	return start + s + reset
}

func BlackFore(s string) string {
	return style(blackFore, s)
}

func RedFore(s string) string {
	return style(redFore, s)
}

func GreenFore(s string) string {
	return style(greenFore, s)
}

func YellowFore(s string) string {
	return style(yellowFore, s)
}

func BlueFore(s string) string {
	return style(blueFore, s)
}

func MagentaFore(s string) string {
	return style(magentaFore, s)
}

func CyanFore(s string) string {
	return style(cyanFore, s)
}

func WhiteFore(s string) string {
	return style(whiteFore, s)
}

func BlackBack(s string) string {
	return style(blackBack, s)
}

func RedBack(s string) string {
	return style(redBack, s)
}

func GreenBack(s string) string {
	return style(greenBack, s)
}

func YellowBack(s string) string {
	return style(yellowBack, s)
}

func BlueBack(s string) string {
	return style(blueBack, s)
}

func MagentaBack(s string) string {
	return style(magentaBack, s)
}

func CyanBack(s string) string {
	return style(cyanBack, s)
}

func WhiteBack(s string) string {
	return style(whiteBack, s)
}

func BrightBlackFore(s string) string {
	return style(brightBlackFore, s)
}

func BrightRedFore(s string) string {
	return style(brightRedFore, s)
}

func BrightGreenFore(s string) string {
	return style(brightGreenFore, s)
}

func BrightYellowFore(s string) string {
	return style(brightYellowFore, s)
}

func BrightBlueFore(s string) string {
	return style(brightBlueFore, s)
}

func BrightMagentaFore(s string) string {
	return style(brightMagentaFore, s)
}

func BrightCyanFore(s string) string {
	return style(brightCyanFore, s)
}

func BrightWhiteFore(s string) string {
	return style(brightWhiteFore, s)
}

func BrightBlackBack(s string) string {
	return style(brightBlackBack, s)
}

func BrightRedBack(s string) string {
	return style(brightRedBack, s)
}

func BrightGreenBack(s string) string {
	return style(brightGreenBack, s)
}

func BrightYellowBack(s string) string {
	return style(brightYellowBack, s)
}

func BrightBlueBack(s string) string {
	return style(brightBlueBack, s)
}

func BrightMagentaBack(s string) string {
	return style(brightMagentaBack, s)
}

func BrightCyanBack(s string) string {
	return style(brightCyanBack, s)
}

func BrightWhiteBack(s string) string {
	return style(brightWhiteBack, s)
}

func Bold(s string) string {
	return style(bold, s)
}

func Faint(s string) string {
	return style(faint, s)
}

func Italic(s string) string {
	return style(italic, s)
}

func Underline(s string) string {
	return style(underline, s)
}

func SlowBlink(s string) string {
	return style(slowBlink, s)
}

func RapidBlink(s string) string {
	return style(rapidBlink, s)
}

func Strike(s string) string {
	return style(strike, s)
}
