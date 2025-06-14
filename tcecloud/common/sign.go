// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"sort"

	tchttp "github.com/hiyongliz/tcecloud-sdk-go/tcecloud/common/http"
)

const (
	SHA256 = "HmacSHA256"
	SHA1   = "HmacSHA1"
)

func Sign(s, secretKey, method string) string {
	hashed := hmac.New(sha1.New, []byte(secretKey))
	if method == SHA256 {
		hashed = hmac.New(sha256.New, []byte(secretKey))
	}
	hashed.Write([]byte(s))

	return base64.StdEncoding.EncodeToString(hashed.Sum(nil))
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

func signRequest(request tchttp.Request, credential *Credential, method string) (err error) {
	if method != SHA256 {
		method = SHA1
	}
	checkAuthParams(request, credential, method)
	s := getStringToSign(request)
	signature := Sign(s, credential.SecretKey, method)
	request.GetParams()["Signature"] = signature
	return
}

func checkAuthParams(request tchttp.Request, credential *Credential, method string) {
	params := request.GetParams()
	credentialParams := credential.GetCredentialParams()
	for key, value := range credentialParams {
		params[key] = value
	}
	params["SignatureMethod"] = method
	delete(params, "Signature")
}

func getStringToSign(request tchttp.Request) string {
	method := request.GetHttpMethod()
	domain := request.GetDomain()
	path := request.GetPath()

	var buf bytes.Buffer
	buf.WriteString(method)
	buf.WriteString(domain)
	buf.WriteString(path)
	buf.WriteString("?")

	params := request.GetParams()
	// sort params
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i := range keys {
		k := keys[i]
		// TODO: check if server side allows empty value in url.
		if params[k] == "" {
			continue
		}
		buf.WriteString(k)
		buf.WriteString("=")
		buf.WriteString(params[k])
		buf.WriteString("&")
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}
