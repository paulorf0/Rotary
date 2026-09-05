package pool

type HTTPStatus int

var (
	StatusOK           HTTPStatus = 200 // WARN: Bloqueio por captcha pode produzir 200 também
	StatusForbidden    HTTPStatus = 403 // ban do antibot no alvo (Cloudflare, PerimeterX, etc)
	StatusUnauthorized HTTPStatus = 401 // às vezes usado como "challenge wall" por antibot, não auth real
)

type Report struct {
	HTTPStatus  HTTPStatus
	TargetURL   string
	RequestTime float32
	Rotate      bool // Se true, mantém lógica de rotação. Se false, devolve o mesmo IP anterior.
}
