package virtualwebauthn

import ( // nolint:gci
	// nolint
	// nolint
	virtualwebauthnx "github.com/descope/virtualwebauthn"
	"github.com/dop251/goja"
)

type Key struct {
	Type            virtualwebauthnx.KeyType `json:"type"`
	Pkcs8SigningKey goja.ArrayBuffer         `json:"signingKey"`
}
