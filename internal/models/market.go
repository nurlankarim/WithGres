package models

type Market struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
type MarketView struct {
//	Id:					int				 `json:"id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	Items       []ItemEdit `json:"items"`
}
type MarketEdit struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	ItemIds     []int  `json:"items"`
}

func (m *Market) MapToView(items []ItemEdit) MarketView {
	return MarketView{
		//Id:          m.Id,
		Name:        m.Name,
		Address:     m.Address,
		PhoneNumber: m.PhoneNumber,
		Items:       items,
	}
}
