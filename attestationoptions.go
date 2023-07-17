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

type AttestationOptions struct {
	Challenge          goja.ArrayBuffer `json:"challenge,omitempty"`
	ExcludeCredentials []string         `json:"excludeCredentials,omitempty"`
	RelyingPartyID     string           `json:"rpId,omitempty"`
	RelyingPartyName   string           `json:"rpName,omitempty"`
	UserID             string           `json:"user,omitempty"`
	UserName           string           `json:"userName,omitempty"`
	UserDisplayName    string           `json:"userDisplayName,omitempty"`
}

func (o *AttestationOptions) createAttestationOptionsx() virtualwebauthnx.AttestationOptions {

	var optionsx = virtualwebauthnx.AttestationOptions{
		Challenge:          o.Challenge.Bytes(),
		ExcludeCredentials: o.ExcludeCredentials,
		RelyingPartyID:     o.RelyingPartyID,
		RelyingPartyName:   o.RelyingPartyName,
		UserID:             o.UserID,
		UserName:           o.UserName,
		UserDisplayName:    o.UserDisplayName,
	}

	return optionsx
}

func createAttestationOptionsFromAttestationOptionsx(r *goja.Runtime, optionsx virtualwebauthnx.AttestationOptions) AttestationOptions {

	var options = AttestationOptions{
		Challenge:          r.NewArrayBuffer(optionsx.Challenge),
		ExcludeCredentials: optionsx.ExcludeCredentials,
		RelyingPartyID:     optionsx.RelyingPartyID,
		RelyingPartyName:   optionsx.RelyingPartyName,
		UserID:             optionsx.UserID,
		UserName:           optionsx.UserName,
		UserDisplayName:    optionsx.UserDisplayName,
	}

	return options
}
