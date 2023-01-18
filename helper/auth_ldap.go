package helper

import (
	"bytes"
	"ddtservice_agri/schema"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLDAP(context *gin.Context, input schema.SignInLDAPJsonSchema) (schema.LDAPLOGIN, string) {
	var client = &http.Client{}
	var ldapLogin schema.LDAPLOGIN
	u, err := json.Marshal(input)
	if err != nil {
		return ldapLogin, err.Error()
	}
	url := "http://202.158.14.235:9393/auth/ldap"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(u)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return ldapLogin, err.Error()
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return ldapLogin, err.Error()
	}
	defer resp.Body.Close()
	var result schema.LDAPData
	body, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Printf("response:%s", string(body))

	ldapLogin.AccountName = result.Data.Samaccountname
	ldapLogin.Company = result.Data.Company
	ldapLogin.Department = result.Data.Department
	ldapLogin.Desc = result.Data.Desc
	ldapLogin.Email = result.Data.Email
	ldapLogin.Fullname = result.Data.Fullname
	ldapLogin.Surname = result.Data.Surname
	ldapLogin.GivenName = result.Data.Givenname
	ldapLogin.Manager = result.Data.Manager
	ldapLogin.EmployeeNumber = result.Data.Employeenumber
	ldapLogin.OfficeName = result.Data.Physicaldeliveryofficename
	ldapLogin.Title = result.Data.Title
	ldapLogin.DomainName = result.Data.Domainname
	ldapLogin.Phone = result.Data.Phone
	ldapLogin.Province = result.Data.Province
	return ldapLogin, ""
}
