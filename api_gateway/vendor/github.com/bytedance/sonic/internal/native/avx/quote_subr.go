// +build !noasm !appengine
// Code generated by asm2asm, DO NOT EDIT.

package avx

import (
	`github.com/bytedance/sonic/loader`
)

const (
    _entry__quote = 48
)

const (
    _stack__quote = 56
)

const (
    _size__quote = 1696
)

var (
    _pcsp__quote = [][2]uint32{
        {1, 0},
        {4, 8},
        {6, 16},
        {8, 24},
        {10, 32},
        {12, 40},
        {13, 48},
        {1649, 56},
        {1653, 48},
        {1654, 40},
        {1656, 32},
        {1658, 24},
        {1660, 16},
        {1662, 8},
        {1663, 0},
        {1690, 56},
    }
)

var _cfunc_quote = []loader.CFunc{
    {"_quote_entry", 0,  _entry__quote, 0, nil},
    {"_quote", _entry__quote, _size__quote, _stack__quote, _pcsp__quote},
}
