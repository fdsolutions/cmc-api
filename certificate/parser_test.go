package certificate_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/fdsolutions/cmc-api/certificate"
)

const (
	emptyPEMcert  = ``
	singlePEMcert = `
-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==
-----END CERTIFICATE-----`

	multiPEMcerts = `
-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==
-----END CERTIFICATE-----


-----BEGIN CERTIFICATE-----
MIID5DCCAswCAQAwDQYJKoZIhvcNAQEEBQAwgbcxCzAJBgNVBAYTAkZSMQ8wDQYD
VQQIEwZGcmFuY2UxDjAMBgNVBAcTBVBhcmlzMRowGAYDVQQKExFGcmFuY2UgVGVs
ZWNvbSBTQTEWMBQGA1UECxMNT1JBTkdFIEZSQU5DRTEjMCEGA1UEAxMaYmVuY2gt
d2VibWFpbHNzbC5vcmFuZ2UuZnIxLjAsBgkqhkiG9w0BCQEWH29jYy1jZXJ0aWZp
Y2F0c0BsaXN0Lm9yYW5nZS5jb20wHhcNMTQwNDI0MTI0MDUwWhcNMTUwNDI0MTI0
MDUwWjCBtzELMAkGA1UEBhMCRlIxDzANBgNVBAgTBkZyYW5jZTEOMAwGA1UEBxMF
UGFyaXMxGjAYBgNVBAoTEUZyYW5jZSBUZWxlY29tIFNBMRYwFAYDVQQLEw1PUkFO
R0UgRlJBTkNFMSMwIQYDVQQDExpiZW5jaC13ZWJtYWlsc3NsLm9yYW5nZS5mcjEu
MCwGCSqGSIb3DQEJARYfb2NjLWNlcnRpZmljYXRzQGxpc3Qub3JhbmdlLmNvbTCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALbiOtd5/9EmqrhZX5Q25n82
8S4I+nawgQ+YBldGY1BY04Jk1aISI7Vh8ovBhmbEcx0Y7/ztf3B15Nv78bmWPL8w
RUr2jvJVS6MndbbfsHUqKuu4Xr1ob49odlc+q1q4Cukeptxj6uQcQRzcA+cGfKif
vpGyWQo/pkoge9UeHCxQWwoF+8Vy04Yr3jXB3BnJf4PCwkr25bDWELlP490hMFNt
NNVN8VwI2SiRh+q80orLsRa5AtMs5YFIlyij0IfCwdcQMLBQUNCu6zrCCc+BubLZ
v3/qI0fvtgfdJmZHc5Ov3mUQS5A6iF4onoAqLQw6HbqEjjdo88dQXgstjvjtYzMC
AwEAATANBgkqhkiG9w0BAQQFAAOCAQEAKivGKJOdmn6Fp1k7G5N4xr3g/aSdU8Im
yh5SQJwaaPw7HDPP4nj86qz0ZtnYo1sTlOovneL4vhrE3vTwp8/rncIL/ls1eAdE
mmsD9EIxoDiYy9n48qkLZLe55luojVJc85pAc44ZYwdbUM5O32VLnhnclGdeI209
NnofPbHka7TNgUaA1wFDcug03WbazejjWQxg1xgDbB+FmbkzTotwyBZSr1wlZYxW
GOXaYoYMO6p7uXhjvbcBsIuD9r9TWF+ivYzWqgHBFHgKYm1Ic+B0DH7K7O/3v7GE
Onn7YHPo80icGrfh4IBwrXnoXSkKkgzRJuW51Wz/ntSVKuSlPw4cCA==
-----END CERTIFICATE-----



-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
MIIFZzCCBE+gAwIBAgIQCQVGmgGOQcdoRUfxhmpXCDANBgkqhkiG9w0BAQUFADCB
tTELMAkGA1UEBhMCVVMxFzAVBgNVBAoTDlZlcmlTaWduLCBJbmMuMR8wHQYDVQQL
ExZWZXJpU2lnbiBUcnVzdCBOZXR3b3JrMTswOQYDVQQLEzJUZXJtcyBvZiB1c2Ug
YXQgaHR0cHM6Ly93d3cudmVyaXNpZ24uY29tL3JwYSAoYykxMDEvMC0GA1UEAxMm
VmVyaVNpZ24gQ2xhc3MgMyBTZWN1cmUgU2VydmVyIENBIC0gRzMwHhcNMTQxMTI0
MDAwMDAwWhcNMTUxMTI0MjM1OTU5WjB/MQswCQYDVQQGEwJGUjERMA8GA1UECAwI
QnJldGFnbmUxGDAWBgNVBAcMD0Nlc3Nvbi1TZXZpZ27DqTEPMA0GA1UECgwGT3Jh
bmdlMQ8wDQYDVQQLDAZPcmFuZ2UxITAfBgNVBAMMGHdlYm1haWwtc3RhdGljLm9y
YW5nZS5mcjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANEGzVJce3m7
7yDUO9eMWvp9vhl8KEUiOMo00VWm6jKz1QXga7X2nErUSVF/M69zgjx0yyebsEYA
vWVhDcVIURMDzFW1ggK9GcQ04+bZMb7I92Vnm8gwvnarlFCbAYAeHvH8Tfxie4mz
jD+Bp+VLby9eEBH12B5ddJ6/aSf9xpQhk1YH7u4ya4/6ltQ/vXihxK0HzpkEEl38
vUeRvaxvfTYH4kFG9EgpzLGwCgKZ8217v+3FueFaf0/1r0E5DGCHBOan635wOX+9
37a+XH7VRyab/il8sbgiBaA5UzftzXkZadhzE8MxsY7QQGFZANt1zXjjr3qSegCm
aAPRkozBp6cCAwEAAaOCAaYwggGiMFgGA1UdEQRRME+CEXdlYm1haWwub3Jhbmdl
LmZygg1qcy53YW5hZG9vLmZyghFzdGF0aWMud2FuYWRvby5mcoIYd2VibWFpbC1z
dGF0aWMub3JhbmdlLmZyMAkGA1UdEwQCMAAwDgYDVR0PAQH/BAQDAgWgMB0GA1Ud
JQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjBlBgNVHSAEXjBcMFoGCmCGSAGG+EUB
BzYwTDAjBggrBgEFBQcCARYXaHR0cHM6Ly9kLnN5bWNiLmNvbS9jcHMwJQYIKwYB
BQUHAgIwGRoXaHR0cHM6Ly9kLnN5bWNiLmNvbS9ycGEwHwYDVR0jBBgwFoAUDURc
FlNEwYJ+HSCrJfQBY9i+eaUwKwYDVR0fBCQwIjAgoB6gHIYaaHR0cDovL3NkLnN5
bWNiLmNvbS9zZC5jcmwwVwYIKwYBBQUHAQEESzBJMB8GCCsGAQUFBzABhhNodHRw
Oi8vc2Quc3ltY2QuY29tMCYGCCsGAQUFBzAChhpodHRwOi8vc2Quc3ltY2IuY29t
L3NkLmNydDANBgkqhkiG9w0BAQUFAAOCAQEAK1bLNLGBDYAJTPjZ+8wFh/QPMHza
Y9p8eaoOT3BkdtzU19gotOo4SVjaoQpJYhp+ds2f7+1fQVVrcuRj1CQXTOj9kn8c
NgOIZSuiAPsVVTDja9wlHVNAv80YE2WKlDKog6ArXMWrWBZOvA9fKf+WqvFq+7Qj
+O0SgxAuxWL13siKuBetCmAN/GsgOMkaeHmiyXd+d/6SDvGM64E+DIWxvSADP0F9
yNf+FW/sTf+YRtROHUFJgtMU0Hwx3LWxow3p0siTxb0Dlrg2AXw+mVQQIsJ+rAF2
6ZOFnifz9dr71D7lcgTiX6peoOGC3x44jIKAzbEgJ7CZAufXKXiNOxYFIw==
-----END CERTIFICATE-----`
)

var _ = Describe("Parser", func() {
	Describe(".Parse", func() {
		Context("With no PEM certificate content passed", func() {
			It("should return an error.", func() {
				_, err := Parse(emptyPEMcert)
				fmt.Printf("[DEBUG] : %#v", err)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("With only certificate passed", func() {
			It("should return an error.", func() {
				certs, err := Parse(singlePEMcert)
				Expect(err).NotTo(HaveOccurred())
				Expect(certs).To(HaveLen(1))
			})
		})
	})
})
