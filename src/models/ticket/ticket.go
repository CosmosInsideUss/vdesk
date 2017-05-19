package ticket

import (
	db "../../Db"
	//"database/sql"
	departamento "../departamento"
	tipoproblema "../problema"
	sla "../sla"
	ticketstatus "../ticketStatus"
	usuario "../usuario"
	"time"
)

type Ticket struct {
	Id              int
	Descricao       string
	AbertoEm        time.Time
	FechadoEm       time.Time
	Status_id       ticketstatus.TicketStatus
	Solicitante_id  usuario.Usuario
	Sla_id          sla.Sla
	Solucao         string
	Usuario_id      usuario.Usuario
	Departamento_id departamento.Departamento
	Problema_id     tipoproblema.TipoProblema
}

func Pendentes() []Ticket {

	rows, err := db.Conexao.Query("select t.id,t.descricao, t.abertoem, " +
		"s.*, sol.id, sol.nome, " +
		"sla.id, sla.valor, " +
		"u.id, u.nome, " +
		"d.id, d.nome, " +
		"p.id, p.valor " +
		"from ticket t join ticketstatus s on t.status_id = s.id join usuario sol on t.solicitante_id = sol.id join slas sla on t.sla_id = sla.id join usuario u on t.usuario_id = u.id join departamento d on t.departamento_id = d.id join tipoproblema p on t.problema_id = p.id where u.id = 1")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tickets := []Ticket{}
	for rows.Next() {
		t := Ticket{}
		rows.Scan(&t.Id, &t.Descricao, &t.AbertoEm, &t.Status_id.Id, &t.Status_id.Valor, &t.Solicitante_id.Id, &t.Solicitante_id.Nome, &t.Sla_id.Id, &t.Sla_id.Valor, &t.Usuario_id.Id, &t.Usuario_id.Nome, &t.Departamento_id.Id, &t.Departamento_id.Nome, &t.Problema_id.Id, &t.Problema_id.Valor)
		tickets = append(tickets, t)
	}

	return tickets

}

func PorUsuario() []Ticket {

	rows, err := db.Conexao.Query("select t.id,t.descricao, t.abertoem, " +
		"s.*, sol.id, sol.nome, " +
		"sla.id, sla.valor, " +
		"u.id, u.nome, " +
		"d.id, d.nome, " +
		"p.id, p.valor " +
		"from ticket t join ticketstatus s on t.status_id = s.id join usuario sol on t.solicitante_id = sol.id join slas sla on t.sla_id = sla.id join usuario u on t.usuario_id = u.id join departamento d on t.departamento_id = d.id join tipoproblema p on t.problema_id = p.id where u.id = 2")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	tickets := []Ticket{}
	for rows.Next() {
		t := Ticket{}
		rows.Scan(&t.Id, &t.Descricao, &t.AbertoEm, &t.Status_id.Id, &t.Status_id.Valor, &t.Solicitante_id.Id, &t.Solicitante_id.Nome, &t.Sla_id.Id, &t.Sla_id.Valor, &t.Usuario_id.Id, &t.Usuario_id.Nome, &t.Departamento_id.Id, &t.Departamento_id.Nome, &t.Problema_id.Id, &t.Problema_id.Valor)
		tickets = append(tickets, t)
	}

	return tickets

}
