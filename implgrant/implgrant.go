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

package implgrant

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type LoginPageData struct {
	UserName string
	Password string
}

type Templatedata struct {
	Client_id     string
	Response_type string
	Redirect_uri  string
	CredValid     string
}

func HandleImplicitGrantAuthorize(resp http.ResponseWriter, req *http.Request) {
	var err error
	err = VerifyRequest(*req)
	if err != nil {
		switch err {
		case ErrRedirectURIMismatch:
			resp.WriteHeader(http.StatusUnauthorized)
			resp.Write([]byte("error=invalid_request"))
			return
		case ErrImplGrantMissingClientID:
			resp.WriteHeader(http.StatusUnauthorized)
			resp.Write([]byte("error=invalid_request"))
			return
		}
	}

	templatedata := Templatedata{
		Client_id:     req.FormValue("client_id"),
		Response_type: req.FormValue("response_type"),
		Redirect_uri:  req.FormValue("redirect_uri"),
		CredValid:     "hidden",
	}
	//urldata := "/authorize?client_id=" + req.FormValue("client_id") + "&response_type=token&redirect_uri=" + req.FormValue("redirect_uri")
	ShowLoginPage(resp, req, templatedata)

}

func ShowLoginPage(resp http.ResponseWriter, req *http.Request, templatedata Templatedata) {
	var err error
	tmpl := template.Must(template.ParseFiles("implgrant/mLoginPage.html"))

	if req.Method == http.MethodGet {
		tmpl.Execute(resp, templatedata)
		return
	}

	submittedData := LoginPageData{
		UserName: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	if err = VerifyCredentials(submittedData); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			templatedata.CredValid = ""
			tmpl.Execute(resp, templatedata)
			return
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
			resp.Write([]byte("error=temporarily_unavailable"))
			return
		}
	}

	output, err := GetAppRegistration(req.FormValue("client_id"))
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		resp.Write([]byte("error=invalid_request"))
	}
	http.Redirect(resp, req, output[1], http.StatusSeeOther)
	return
}

func VerifyCredentials(formdata LoginPageData) (err error) {
	data, err := ReadPrincipalRecord(formdata.UserName)
	if err != nil {
		log.Println("Error Reading DB Data ", err.Error())
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(data[2]), []byte(formdata.Password))

	if err != nil {

		return err
	}

	return nil

}

func VerifyRequest(req http.Request) (err error) {

	if req.FormValue("client_id") == "" {
		return ErrImplGrantMissingClientID
	}
	registrationdata, err := GetAppRegistration(req.FormValue("client_id"))
	if err != nil {
		return err
	}

	redirecturi := req.FormValue("redirect_uri")

	if redirecturi == "" {
		return nil
	}

	if strings.TrimRight(registrationdata[1], `/`) != strings.TrimRight(redirecturi, `/`) {
		return ErrRedirectURIMismatch
	}

	return nil

}
