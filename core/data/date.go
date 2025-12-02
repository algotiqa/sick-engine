//=============================================================================
/*
Copyright © 2024 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package data

import "fmt"

//=============================================================================

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

//=============================================================================

func NewDate(year int, month int, day int) *Date {
	return &Date{
		Year : year,
		Month: month,
		Day  : day,
	}
}

//=============================================================================

func (d *Date) IsValid() bool {
	if d.Year < 0 {
		return false
	}

	if d.Month < 1 || d.Month > 12 {
		return false
	}

	if d.Day < 1 || d.Day > 31 {
		return false
	}

	return true
}

//=============================================================================

func (d *Date) String() string {
	return fmt.Sprintf("%4d-%2d-%2d", d.Year, d.Month, d.Day)
}

//=============================================================================
