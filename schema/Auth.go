package schema

type SignUpJsonSchema struct {
	Role     []string `json:"role" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Username string   `json:"username" binding:"required"`
	Password string   `json:"password" binding:"required"`
}
type SignInJsonSchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInLDAPJsonSchema struct {
	Domain   string `json:"domain"`
	Ldap     bool   `json:"ldap"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LDAPData struct {
	Data LDAP `json:"data"`
}

type LDAP struct {
	Samaccountname             string   `json:"samaccountname"`
	Fullname                   string   `json:"fullname"`
	Givenname                  string   `json:"givenname"`
	Surname                    string   `json:"surname"`
	Email                      string   `json:"email"`
	Phone                      string   `json:"phone"`
	Department                 string   `json:"department"`
	Manager                    string   `json:"manager"`
	Desc                       []string `json:"desc"`
	Title                      string   `json:"title"`
	Company                    string   `json:"company"`
	Employeenumber             string   `json:"employeenumber"`
	Physicaldeliveryofficename string   `json:"physicaldeliveryofficename"`
	Province                   string   `json:"province"`
	Domainname                 string   `json:"domainName"`
}
type LDAPLOGIN struct {
	AccountName    string   `json:"account_name"`
	Fullname       string   `json:"fullname"`
	GivenName      string   `json:"given_name"`
	Surname        string   `json:"surname"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Department     string   `json:"department"`
	Manager        string   `json:"manager"`
	Desc           []string `json:"desc"`
	Title          string   `json:"title"`
	Company        string   `json:"company"`
	EmployeeNumber string   `json:"nik"`
	OfficeName     string   `json:"office_name"`
	Province       string   `json:"province"`
	DomainName     string   `json:"domain_name"`
}
