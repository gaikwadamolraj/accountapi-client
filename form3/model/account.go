package model

type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

type Data struct {
	Data interface{} `json:"data,omitempty"`
}

type Accounts struct {
	AccountData
}

func (a *AccountData) GetAccountID() string {
	return a.ID
}

func (a *AccountData) SetAccountID(accountId string) {
	a.ID = accountId
}

func (a *AccountData) GetOrgId() string {
	return a.OrganisationID
}

func (a *AccountData) SetOrgId(orgId string) {
	a.OrganisationID = orgId
}

func (a *AccountData) GetType() string {
	return a.Type
}

func (a *AccountData) SetType(t string) {
	a.Type = t
}

func (a *AccountData) GetVersion() int64 {
	if a.Version != nil {
		return *a.Version
	}
	return -1
}

func (a *AccountData) SetVersion(v int64) {
	a.Version = getint64Pointer(v)
}

func (a *AccountData) GetStatus() string {
	return *a.Attributes.Status
}

func (a *AccountData) SetCountry(country string) {
	a.Attributes.Country = setStringPointer(country)
}

func (a *AccountData) SetBaseCurrency(c string) {
	a.Attributes.BaseCurrency = c
}

func (a *AccountData) SetBankID(c string) {
	a.Attributes.BankID = c
}

func (a *AccountData) SetBankIDCode(c string) {
	a.Attributes.BankIDCode = c
}

func (a *AccountData) SetBic(c string) {
	a.Attributes.Bic = c
}

func (a *AccountData) SetIban(c string) {
	a.Attributes.Iban = c
}

func (a *AccountData) SetStatus(c string) {
	a.Attributes.Status = setStringPointer(c)
}

func (a *AccountData) SetName(name string) {
	a.Attributes.Name = []string{name}
}

func setStringPointer(s string) *string {
	return &s
}

func getint64Pointer(s int64) *int64 {
	return &s
}

func GetAccountModel() *AccountData {
	return &AccountData{
		ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		Type:           "accounts",
		Attributes: &AccountAttributes{
			Country:      setStringPointer("GB"),
			BaseCurrency: "GBP",
			BankID:       "400300",
			BankIDCode:   "GBDSC",
			Bic:          "NWBKGB22",
			Iban:         "GB11NWBK40030041426819",
			Name:         []string{"Amol Gaikwad"},
		},
	}
}
