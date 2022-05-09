package jwtpayload

import "github.com/0B1t322/Magic-Circle/models/role"

type Payload struct {
	ID          int       `json:"id"`
	Role        role.Role `json:"role"`
	InstituteID int      `json:"intstituteId"`
}
