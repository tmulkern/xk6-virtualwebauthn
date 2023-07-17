// MIT License
//
// Copyright (c) 2023 Tadhg Mulkern
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package virtualwebauthn

import ( // nolint:gci

	// nolint

	// nolint
	virtualwebauthnx "github.com/descope/virtualwebauthn"
	"github.com/dop251/goja"
)

type AssertionOptions struct {
	Challenge        goja.ArrayBuffer `json:"challenge,omitempty"`
	AllowCredentials []string         `json:"allowCredentials,omitempty"`
	RelyingPartyID   string           `json:"rpId,omitempty"`
}

func (o *AssertionOptions) createAssertionOptionsx() virtualwebauthnx.AssertionOptions {

	var optionsx = virtualwebauthnx.AssertionOptions{
		Challenge:        o.Challenge.Bytes(),
		AllowCredentials: o.AllowCredentials,
		RelyingPartyID:   o.RelyingPartyID,
	}

	return optionsx
}

func createAssertionOptionsFromAssertionOptionsx(r *goja.Runtime, optionsx virtualwebauthnx.AssertionOptions) AssertionOptions {

	options := AssertionOptions{
		Challenge:        r.NewArrayBuffer(optionsx.Challenge),
		AllowCredentials: optionsx.AllowCredentials,
		RelyingPartyID:   optionsx.RelyingPartyID,
	}

	return options
}
