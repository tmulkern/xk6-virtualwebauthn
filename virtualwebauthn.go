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
	"go.k6.io/k6/js/modules"
)

type VirtualWebAuthn struct {
	vu modules.VU
}

func newVirtualWebAuthn(vu modules.VU) *VirtualWebAuthn {
	return &VirtualWebAuthn{vu: vu}
}

func (v *VirtualWebAuthn) CreateAssertionResponse(rpObj goja.Value, auth Authenticator, credObj goja.Value, options AssertionOptions) string {

	var cred = parseCredentialObjectToPortalCredential(credObj)
	var rp = parseRelyingPartyObjectToRelyingPartyx(rpObj)
	return virtualwebauthnx.CreateAssertionResponse(rp, auth.createAuthenticatorx(), cred.createCredentialx(), options.createAssertionOptionsx())
}

func (v *VirtualWebAuthn) CreateAttestationResponse(rpObj goja.Value, auth Authenticator, credObj goja.Value, options AttestationOptions) string {

	var cred = parseCredentialObjectToPortalCredential(credObj)
	var rp = parseRelyingPartyObjectToRelyingPartyx(rpObj)
	return virtualwebauthnx.CreateAttestationResponse(rp, auth.createAuthenticatorx(), cred.createCredentialx(), options.createAttestationOptionsx())
}

func (v *VirtualWebAuthn) ParseAssertionOptions(str string) (assertionOptions *AssertionOptions) {

	assertionOptionsx, err := virtualwebauthnx.ParseAssertionOptions(str)

	if err != nil {
		panic("Problem parsing AssertionOptions JSON string")
	}

	var assertionOptionsToReturn = createAssertionOptionsFromAssertionOptionsx(v.vu.Runtime(), *assertionOptionsx)
	return &assertionOptionsToReturn
}

func (v *VirtualWebAuthn) ParseAttestationOptions(str string) (attestationOptions *AttestationOptions) {

	attestationOptionsx, err := virtualwebauthnx.ParseAttestationOptions(str)

	if err != nil {
		panic("Problem parsing AttestationOptions JSON string")
	}

	var attestationOptionsToReturn = createAttestationOptionsFromAttestationOptionsx(v.vu.Runtime(), *attestationOptionsx)
	return &attestationOptionsToReturn
}

func (v *VirtualWebAuthn) NewAuthenticator() *Authenticator {

	var authenticatorx = virtualwebauthnx.NewAuthenticator()
	var authenticator = createAuthenticatorFromAuthenticatorx(v.vu.Runtime(), authenticatorx)
	authenticator.runtime = v.vu.Runtime()
	return &authenticator
}

func (v *VirtualWebAuthn) NewAuthenticatorWithOptions(optionsObj goja.Value) *Authenticator {

	var options = parseAuthenticatorOptionsObjectToAuthenticatorOptions(v.vu.Runtime(), optionsObj)

	var authenticatorx = virtualwebauthnx.NewAuthenticatorWithOptions(options.createAuthenticatorOptionsx())

	var authenticator = createAuthenticatorFromAuthenticatorx(v.vu.Runtime(), authenticatorx)
	authenticator.runtime = v.vu.Runtime()
	return &authenticator
}

func parseRelyingPartyObjectToRelyingPartyx(rpObj goja.Value) virtualwebauthnx.RelyingParty {

	relyingParty := virtualwebauthnx.RelyingParty{}

	var rpMap = rpObj.Export()

	if rpData, ok := rpMap.(map[string]interface{}); ok {

		for rpK, rpV := range rpData {
			switch rpK {
			case "ID":
				val, ok := rpV.(string)
				if !ok {
					panic("Parsing RelyingParty: 'ID' field must be of type string")
				}
				relyingParty.ID = val
			case "Name":
				val, ok := rpV.(string)
				if !ok {
					panic("Parsing RelyingParty: 'Name' field must be of type string")
				}
				relyingParty.Name = val
			case "Origin":
				val, ok := rpV.(string)
				if !ok {
					panic("Parsing RelyingParty: 'Origin' field must be of type string")
				}
				relyingParty.Origin = val
			}
		}
	}

	return relyingParty
}
