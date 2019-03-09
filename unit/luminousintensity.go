// Code generated by "go generate gonum.org/v1/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
)

// Candela represents a luminous intensity in candela.
type LuminousIntensity float64

const (
	Yottacandela LuminousIntensity = 1e24
	Zettacandela LuminousIntensity = 1e21
	Exacandela   LuminousIntensity = 1e18
	Petacandela  LuminousIntensity = 1e15
	Teracandela  LuminousIntensity = 1e12
	Gigacandela  LuminousIntensity = 1e9
	Megacandela  LuminousIntensity = 1e6
	Kilocandela  LuminousIntensity = 1e3
	Hectocandela LuminousIntensity = 1e2
	Decacandela  LuminousIntensity = 1e1
	Candela      LuminousIntensity = 1.0
	Decicandela  LuminousIntensity = 1e-1
	Centicandela LuminousIntensity = 1e-2
	Millicandela LuminousIntensity = 1e-3
	Microcandela LuminousIntensity = 1e-6
	Nanocandela  LuminousIntensity = 1e-9
	Picocandela  LuminousIntensity = 1e-12
	Femtocandela LuminousIntensity = 1e-15
	Attocandela  LuminousIntensity = 1e-18
	Zeptocandela LuminousIntensity = 1e-21
	Yoctocandela LuminousIntensity = 1e-24
)

// Unit converts the LuminousIntensity to a *Unit
func (j LuminousIntensity) Unit() *Unit {
	return New(float64(j), Dimensions{
		LuminousIntensityDim: 1,
	})
}

// LuminousIntensity allows LuminousIntensity to implement a LuminousIntensityer interface
func (j LuminousIntensity) LuminousIntensity() LuminousIntensity {
	return j
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (j *LuminousIntensity) From(u Uniter) error {
	if !DimensionsMatch(u, Candela) {
		*j = LuminousIntensity(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*j = LuminousIntensity(u.Unit().Value())
	return nil
}

func (j LuminousIntensity) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", j, float64(j))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " cd"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-len(unit)), p, float64(j))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(j))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-len(unit)), float64(j))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(j))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g cd)", c, j, float64(j))
	}
}