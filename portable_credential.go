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

import (
	virtualwebauthnx "github.com/descope/virtualwebauthn"
	"github.com/dop251/goja"
)

// nolint:gci

// nolint

type PortableCredential struct {
	ID       goja.ArrayBuffer         `json:"id"`
	PKCS8Key goja.ArrayBuffer         `json:"key"`
	KeyType  virtualwebauthnx.KeyType `json:"keyType,omitempty"`
	Counter  uint32                   `json:"counter,omitempty"`
}

func (c *PortableCredential) IsAllowedForAssertion(options AssertionOptions) bool {
	var credentialx = c.createCredentialx()
	var optionsx = options.createAssertionOptionsx()
	return credentialx.IsAllowedForAssertion(optionsx)
}

func (c *PortableCredential) IsExcludedForAttestation(options AttestationOptions) bool {

	var credentialx = c.createCredentialx()
	var optionsx = options.createAttestationOptionsx()
	return credentialx.IsExcludedForAttestation(optionsx)
}

func (c *PortableCredential) createPortableCredentialx() virtualwebauthnx.PortableCredential {

	var credx = virtualwebauthnx.PortableCredential{
		ID:       c.ID.Bytes(),
		KeyType:  c.KeyType,
		PKCS8Key: c.PKCS8Key.Bytes(),
		Counter:  c.Counter,
	}
	return credx
}

func (c *PortableCredential) createCredentialx() virtualwebauthnx.Credential {
	var portableCredentialx = c.createPortableCredentialx()
	return portableCredentialx.ToCredential()
}

func createPortalCredentialFromCredentialx(r *goja.Runtime, credentialx virtualwebauthnx.Credential) PortableCredential {

	var portableCredentialx = credentialx.ExportToPortableCredential()
	portableCredential := PortableCredential{
		ID:       r.NewArrayBuffer(portableCredentialx.ID),
		PKCS8Key: r.NewArrayBuffer(portableCredentialx.PKCS8Key),
		KeyType:  portableCredentialx.KeyType,
		Counter:  portableCredentialx.Counter,
	}
	return portableCredential
}

func parseCredentialObjectToPortalCredential(credObj goja.Value) PortableCredential {

	credential := PortableCredential{}
	var credMap = credObj.Export()

	if credData, ok := credMap.(map[string]interface{}); ok {

		for credK, credV := range credData {
			switch credK {
			case "ID":
				val, ok := credV.(goja.ArrayBuffer)
				if !ok {
					panic("Parsing Credential: 'ID' field must be of type ArrayBuffer")
				}
				credential.ID = val
			case "Counter":
				val, ok := credV.(int64)
				if !ok {
					panic("Parsing Credential: 'Counter' field must be of type number")
				}
				credential.Counter = uint32(val)
			case "KeyType":
				val, ok := credV.(int64)
				if !ok {
					panic("Parsing Credential: 'KeyType' field must be of type number")
				}
				credential.KeyType = virtualwebauthnx.KeyType(val)

			case "PKCS8Key":
				val, ok := credV.(goja.ArrayBuffer)
				if !ok {
					panic("Parsing Credential: 'PKCS8Key' field must be of type ArrayBuffer")
				}
				credential.PKCS8Key = val
			}

		}
	}
	return credential
}
