package ticket

import (
	_ "../../Db"
	_ "../departamento"
)

type Ticket struct {
	Id        int
	Descricao string
	AbertoEm  string
	// FechadoEm    string
	// CriadoPor    string
	// EditadoPor   string
	// Status       string
	// Solicitante  string
	// Sla          string
	// Solucao      string
	// Usuario      string
	// Departamento []d.Departamento
	// Problema     string
	// Tipo         string
}

func BuscaTodos() []Ticket {

	t := []Ticket{{1, "teste", "hoje"}, {2, "teste2", "ontem"}}
	// ticketsAll, _ := db.CONEXAO.Query("select t.id, t.descricao, t.abertoem, u.nome as criadopor,d.nome as departamento," +
	// 	" s.valor as sla, st.valor as status, p.valor as problema from ticket t  join usuario u on t.criadopor_id = u.id " +
	// 	"join sla s on t.sla_id = s.id join departamento d on t.departamento_id = d.id join ticketstatus st on t.status_id = st.id " +
	// 	"join tipoproblema p on t.problema_id = p.id")
	// defer ticketsAll.Close()
	// departamentosAll, _ := db.CONEXAO.Query("select id, nome from departamento")
	// defer departamentosAll.Close()
	// tickets := []Ticket{}
	// for ticketsAll.Next() {
	// 	ticket := Ticket{}
	// 	departamento := d.Departamento{}
	// 	ticketsAll.Scan(&ticket.Id, &ticket.Descricao, &ticket.AbertoEm, &ticket.CriadoPor, &ticket.Departamento, &ticket.Sla, &ticket.Status, &ticket.Problema)
	// 	departamentosAll.Scan(&departamento.Id, &departamento.Nome)
	// 	ticket.Departamento = append(ticket.Departamento, departamento)
	// 	tickets = append(tickets, ticket)
	// }
	return t
}

// func porId(id string) Ticket {
// 	u := Ticket{}
// 	err := db.CONEXAO.QueryRow("select id from Ticket where id = $1", id).Scan(&u.Id)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return u
// }
