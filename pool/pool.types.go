package pool

import "time"

type pool struct {
	ID       string `json:"id"` // id do pool
	Name     string `json:"name"`
	ClientID string `json:"clientid"`
	Ips      []IP   `json:"ips"`
}

// TODO: avaliar se o score deve ser por par ip+cliente/alvo em vez de global por ip
// (mesmo ip pode ser bom pro cliente A e banido no alvo do cliente B)

type Lease struct {
	ID         string // o leaseId
	ClientIDd  *string
	IPID       string
	AcquiredAt *time.Time
	ClosedAt   *time.Time // nil = aberto
	ExpiresAt  float64    // Tempo de expiração obrigatória, em segundos.
}
