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

package ropcg

import (
	"log"
	"net/http"

	"github.com/bsankar/openauthd/utils"

	"golang.org/x/crypto/bcrypt"
)

func HandleResourceOwnerPasswordCredGrant(resp http.ResponseWriter, req *http.Request) {

	err := VerifyCredentials(req.FormValue("username"), req.FormValue("password"))

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			resp.WriteHeader(http.StatusUnauthorized)
			resp.Write([]byte("UnAuthorized"))
			return

		} else {
			resp.WriteHeader(http.StatusInternalServerError)
			
			if f, ok := resp.(http.Flusher); ok {
				f.Flush()
			}

			return

		}
	}

	utils.SendAccessTokenResponse(resp)
	return

}

func VerifyCredentials(username string, password string) (err error) {
	data, err := ReadPrincipalRecord(username)
	if err != nil {
		log.Println("Error Reading DB Data ", err.Error())
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data[2]), []byte(password))

	if err != nil {

		return err
	}

	return nil
}
