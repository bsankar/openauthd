/*
 Copyright 2022 Balaje Sankar

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"net/http"

	"github.com/bsankar/openauthd/implgrant"
	"github.com/bsankar/openauthd/ropcg"
)

func main() {
	_, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(privatekey.PublicKey)
	fs := http.FileServer(http.Dir("implgrant/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/authorize", authReqhandler)
	err = http.ListenAndServeTLS(":4443", "example.crt", "example.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func authReqhandler(resp http.ResponseWriter, req *http.Request) {

	err := req.ParseForm()

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error Occured" + err.Error()))
		return
	}
	for k, v := range req.Form {
		log.Println(k, v)
	}

	responsetype := req.FormValue("response_type")
	grant_type := req.FormValue("grant_type")
	if responsetype != "" {
		switch responsetype {
		case "code":
		case "token":
			implgrant.HandleImplicitGrantAuthorize(resp, req)
		}
	} else if grant_type != "" {
		granttype := req.FormValue("grant_type")

		switch granttype {
		case "password":
			ropcg.HandleResourceOwnerPasswordCredGrant(resp, req)
			return

		}
	}

}
