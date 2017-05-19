package problema

import (
	usuario "../usuario"
)

type TipoProblema struct {
	Id         int
	Valor      string
	Cor        string
	CriadoPor  usuario.Usuario
	EditadoPor usuario.Usuario
	Criadoem   string
	Editadorem string
}
