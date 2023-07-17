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

type Authenticator struct {
	Options     AuthenticatorOptions `json:"options"`
	Aaguid      goja.ArrayBuffer     `json:"aaguid"`
	Credentials []PortableCredential `json:"credentials,omitempty"`
	runtime     *goja.Runtime
}

type AuthenticatorOptions struct {
	UserHandle      goja.ArrayBuffer
	UserNotPresent  bool
	UserNotVerified bool
	Runtime         *goja.Runtime
}

func (a *Authenticator) AddCredential(credObj goja.Value) {

	var credentialx = parseCredentialObjectToPortalCredential(credObj)
	a.Credentials = append(a.Credentials, credentialx)
}

func (a *Authenticator) FindAllowedCredential(options AssertionOptions) *PortableCredential {

	var authenticatorx = a.createAuthenticatorx()
	var optionsx = options.createAssertionOptionsx()

	var credentialx = authenticatorx.FindAllowedCredential(optionsx)
	var credential = createPortalCredentialFromCredentialx(a.runtime, *credentialx)
	return &credential
}

func (a *Authenticator) createAuthenticatorx() virtualwebauthnx.Authenticator {

	authx := virtualwebauthnx.Authenticator{
		Options: a.Options.createAuthenticatorOptionsx(),
	}

	copy(authx.Aaguid[:], a.Aaguid.Bytes())

	for _, cred := range a.Credentials {
		var credx = cred.createCredentialx()
		authx.Credentials = append(authx.Credentials, credx)
	}

	return authx
}

func (o *AuthenticatorOptions) createAuthenticatorOptionsx() virtualwebauthnx.AuthenticatorOptions {

	var optionsx = virtualwebauthnx.AuthenticatorOptions{
		UserHandle:      o.UserHandle.Bytes(),
		UserNotPresent:  o.UserNotPresent,
		UserNotVerified: o.UserNotVerified,
	}

	return optionsx
}

func createAuthenticatorFromAuthenticatorx(r *goja.Runtime, authx virtualwebauthnx.Authenticator) Authenticator {

	var auth = Authenticator{
		Options: createAuthenticatorOptionsFromAuthenticatorxOptions(r, authx.Options),
		Aaguid:  r.NewArrayBuffer(authx.Aaguid[:]),
		runtime: r,
	}

	for _, credx := range authx.Credentials {
		var cred = createPortalCredentialFromCredentialx(r, credx)
		auth.Credentials = append(auth.Credentials, cred)
	}

	return auth
}

func createAuthenticatorOptionsFromAuthenticatorxOptions(r *goja.Runtime, optionsx virtualwebauthnx.AuthenticatorOptions) AuthenticatorOptions {

	var options = AuthenticatorOptions{
		UserHandle:      r.NewArrayBuffer(optionsx.UserHandle),
		UserNotPresent:  optionsx.UserNotPresent,
		UserNotVerified: optionsx.UserNotVerified,
	}

	return options
}

func parseAuthenticatorOptionsObjectToAuthenticatorOptions(r *goja.Runtime, optionsObj goja.Value) AuthenticatorOptions {

	options := AuthenticatorOptions{}
	var optionsMap = optionsObj.Export()

	var hasOptionsAUserHandle = false

	if optionsData, ok := optionsMap.(map[string]interface{}); ok {

		for optionsK, optionsV := range optionsData {
			switch optionsK {
			case "UserHandle":
				val, ok := optionsV.(goja.ArrayBuffer)
				if !ok {
					panic("Parsing AuthenticatorOptions: 'UserHandle' field must be of type ArrayBuffer")
				}
				options.UserHandle = val
				hasOptionsAUserHandle = true
			case "UserNotPresent":
				val, ok := optionsV.(bool)
				if !ok {
					panic("Parsing AuthenticatorOptions: 'UserNotPresent' field must be of type bool")
				}
				options.UserNotPresent = val
			case "UserNotVerified":
				val, ok := optionsV.(bool)
				if !ok {
					panic("Parsing AuthenticatorOptions: 'UserNotVerified' field must be of type bool")
				}
				options.UserNotVerified = val
			}
		}
	}

	if !hasOptionsAUserHandle {
		options.UserHandle = r.NewArrayBuffer(make([]byte, 0))
	}

	return options
}
