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

package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte

func generateToken() (tokenString string, err error) {
	hmacSampleSecret = []byte(os.Getenv("TOKENSECRET"))
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err = token.SignedString(hmacSampleSecret)
	return
}

func SendAccessTokenResponse(resp http.ResponseWriter) {
	respJsonData := AuthResp{}
	gettoken, err := generateToken()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error Occured" + err.Error()))
		return
	}
	resp.Header().Set("Content-Type", `application/json`)
	respJsonData.AccessToken = gettoken
	respJsonData.RefreshToken = ""
	respJsonData.Expires = 3600
	respJsonData.TokenType = "jwt"
	out, err := json.Marshal(respJsonData)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Error Occured" + err.Error()))
	}
	resp.Write(out)
	return
}
