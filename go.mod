module github.com/tmulkern/xk6-virtualwebauthn

go 1.19

//Need this till PR is merged
replace github.com/descope/virtualwebauthn => github.com/tmulkern/virtualwebauthn v0.0.0-20230711154158-df5eae37d129

require (
	github.com/descope/virtualwebauthn v1.0.3-0.20230613131733-5740254dc414
	github.com/dop251/goja v0.0.0-20230304130813-e2f543bf4b4c
	go.k6.io/k6 v0.38.0
)

require (
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fxamacker/cbor/v2 v2.4.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.4-0.20211119122758-180fcef48034+incompatible // indirect
	github.com/google/pprof v0.0.0-20230207041349-798e818bf904 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/serenize/snaker v0.0.0-20201027110005-a7ad2135616e // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/spf13/afero v1.1.2 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	golang.org/x/time v0.0.0-20220224211638-0e9765cccd65 // indirect
	gopkg.in/guregu/null.v3 v3.3.0 // indirect
)
