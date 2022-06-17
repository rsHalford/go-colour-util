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
	blackFG   = "\x1b[30m"
	redFG     = "\x1b[31m"
	greenFG   = "\x1b[32m"
	yellowFG  = "\x1b[33m"
	blueFG    = "\x1b[34m"
	magentaFG = "\x1b[35m"
	cyanFG    = "\x1b[36m"
	whiteFG   = "\x1b[37m"

	blackBG   = "\x1b[40m"
	redBG     = "\x1b[41m"
	greenBG   = "\x1b[42m"
	yellowBG  = "\x1b[43m"
	blueBG    = "\x1b[44m"
	magentaBG = "\x1b[45m"
	cyanBG    = "\x1b[46m"
	whiteBG   = "\x1b[47m"

	blackFGB   = "\x1b[90m"
	redFGB     = "\x1b[91m"
	greenFGB   = "\x1b[92m"
	yellowFGB  = "\x1b[93m"
	blueFGB    = "\x1b[94m"
	magentaFGB = "\x1b[95m"
	cyanFGB    = "\x1b[96m"
	whiteFGB   = "\x1b[97m"

	blackBGB   = "\x1b[100m"
	redBGB     = "\x1b[101m"
	greenBGB   = "\x1b[102m"
	yellowBGB  = "\x1b[103m"
	blueBGB    = "\x1b[104m"
	magentaBGB = "\x1b[105m"
	cyanBGB    = "\x1b[106m"
	whiteBGB   = "\x1b[107m"

	bold       = "\x1b[1m"
	faint      = "\x1b[2m"
	italic     = "\x1b[3m"
	underline  = "\x1b[4m"
	slowBlink  = "\x1b[5m"
	rapidBlink = "\x1b[6m"
	strike     = "\x1b[9m"

	reset = "\x1b[0m"
)

const version string = "v0.1.0"

func style(start, s string) string {
	return start + s + reset
}

func BlaFG(s string) string {
	return style(blackFG, s)
}

func RedFG(s string) string {
	return style(redFG, s)
}

func GreFG(s string) string {
	return style(greenFG, s)
}

func YelFG(s string) string {
	return style(yellowFG, s)
}

func BluFG(s string) string {
	return style(blueFG, s)
}

func MagFG(s string) string {
	return style(magentaFG, s)
}

func CyaFG(s string) string {
	return style(cyanFG, s)
}

func WhiFG(s string) string {
	return style(whiteFG, s)
}

func BlackBG(s string) string {
	return style(blackBG, s)
}

func RedBG(s string) string {
	return style(redBG, s)
}

func GreBG(s string) string {
	return style(greenBG, s)
}

func YelBG(s string) string {
	return style(yellowBG, s)
}

func BluBG(s string) string {
	return style(blueBG, s)
}

func MagBG(s string) string {
	return style(magentaBG, s)
}

func CyaBG(s string) string {
	return style(cyanBG, s)
}

func WhiBG(s string) string {
	return style(whiteBG, s)
}

func BlaFGB(s string) string {
	return style(blackFGB, s)
}

func RedFGB(s string) string {
	return style(redFGB, s)
}

func GreFGB(s string) string {
	return style(greenFGB, s)
}

func YelFGB(s string) string {
	return style(yellowFGB, s)
}

func BluFGB(s string) string {
	return style(blueFGB, s)
}

func MagFGB(s string) string {
	return style(magentaFGB, s)
}

func CyaFGB(s string) string {
	return style(cyanFGB, s)
}

func WhiFGB(s string) string {
	return style(whiteFGB, s)
}

func BlaBGB(s string) string {
	return style(blackBGB, s)
}

func RedBGB(s string) string {
	return style(redBGB, s)
}

func GreBGB(s string) string {
	return style(greenBGB, s)
}

func YelBGB(s string) string {
	return style(yellowBGB, s)
}

func BluBGB(s string) string {
	return style(blueBGB, s)
}

func MagBGB(s string) string {
	return style(magentaBGB, s)
}

func CyaBGB(s string) string {
	return style(cyanBGB, s)
}

func WhiBGB(s string) string {
	return style(whiteBGB, s)
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
