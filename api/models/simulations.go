package models

type Simulations struct{
	Id string
	UserId string
	Status string `json:"status"`
	DailyRate int `json:"daily_rate"`
	ContractLength int `json:"contract_length"`
	MonthlyCharges int `json:"monthly_charges"`
}