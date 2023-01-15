package schema

type RequestEstate struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	CompanyId   string `json:"company_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RequestEstateUpdate struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	CompanyId   string `json:"company_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsDeleted   bool   `json:"delete" binding:"required"`
	IsActive    bool   `json:"status" binding:"required"`
}

type RequestCompany struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Fax         string `json:"fax" binding:"required"`
	Sector      string `json:"sector" binding:"required"`
	Domain      string `json:"domain" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type RequestCompanyUpdate struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Fax         string `json:"fax" binding:"required"`
	Sector      string `json:"sector" binding:"required"`
	Domain      string `json:"domain" binding:"required"`
	Address     string `json:"address" binding:"required"`
	IsDeleted   bool   `json:"delete" binding:"required"`
	IsActive    bool   `json:"status" binding:"required"`
}

type RequestDivision struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	EstateId    string `json:"estate_id"`
}

type RequestDivisionUpdate struct {
	Code        string `json:"code" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	EstateId    string `json:"estate_id"`
	IsDeleted   bool   `json:"delete" binding:"required"`
	IsActive    bool   `json:"status" binding:"required"`
}

type RequestGang struct {
	Code       string `json:"code" binding:"required"`
	Name       string `json:"name" binding:"required"`
	DivisionId string `json:"division_id" binding:"required"`
}

type RequestGangUpdate struct {
	Code       string `json:"code" binding:"required"`
	Name       string `json:"name" binding:"required"`
	DivisionId string `json:"division_id" binding:"required"`
	IsDeleted  bool   `json:"delete" binding:"required"`
	IsActive   bool   `json:"status" binding:"required"`
}

type RequestApplication struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	UpdatedNote     string `json:"updated_note" binding:"required"`
	Version         string `json:"version" binding:"required"`
	AppPackage      string `json:"app_package" binding:"required"`
	AppPackageClass string `json:"app_package_class" binding:"required"`
	AssetApk        string `json:"asset_apk" binding:"required"`
	AssetIcon       string `json:"asset_icon" binding:"required"`
}

type RequestApplicationUpdate struct {
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	UpdatedNote     string `json:"updated_note" binding:"required"`
	Version         string `json:"version" binding:"required"`
	AppPackage      string `json:"app_package" binding:"required"`
	AppPackageClass string `json:"app_package_class" binding:"required"`
	AssetApk        string `json:"asset_apk" binding:"required"`
	AssetIcon       string `json:"asset_icon" binding:"required"`
	IsDeleted       bool   `json:"delete" binding:"required"`
	IsActive        bool   `json:"status" binding:"required"`
}

type RequestUser struct {
	Role      []string `json:"role" binding:"required"`
	Email     string   `json:"email" binding:"required"`
	Username  string   `json:"username" binding:"required"`
	Password  string   `json:"password" binding:"required"`
	CompanyId string   `json:"company_id" binding:"required"`
}
