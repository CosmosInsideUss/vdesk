import  Button from "react-bootstrap/lib/Button";
import  React from 'react';
import ReactDOM from 'react-dom'
import { Navbar, Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap/lib';
import { Grid, Col, Row, Table} from 'react-bootstrap/lib';
import axios from 'axios';
import { FaBars ,FaBolt ,FaCheckSquareO, FaSearch, FaUmbrella, FaUserSecret, FaWechat, FaFilePdfO,
FaPieChart, FaSignOut } from 'react-icons/lib/fa';

var Menu = React.createClass({

render (){
 return ( 
          <Navbar>
          </Navbar>
  		)
}
});
var Layout = React.createClass({
	render(){
			return(
				<Grid>
					<Row className="show-grid">
					  <Col md={12}>   
					  	  <Menu/>
					  </Col>
          </Row>
          <Row className="show-grid">
					  <Col md={3}>   
            
                <MenuLeft/>
            </Col>
            <Col md={9}>   
						
					  		<TicketsTable/>
					  </Col>
					  </Row>
				</Grid>
				);
	}
});
var MenuLeft = React.createClass({
	render(){
		return(
  <Nav bsStyle="pills" stacked activeKey={2}>
    <NavItem eventKey={1} href="/home"><FaBars/> Home</NavItem>
    <NavItem eventKey={2} href="/home"><FaBolt/> Tickets</NavItem>
    <NavItem eventKey={1} href="/home"><FaCheckSquareO/> Fechados</NavItem>
    <NavItem eventKey={1} href="/home"><FaSearch/> Pesquisar</NavItem>
    <NavItem eventKey={1} href="/home"><FaUmbrella/> Auditoria</NavItem>
    <NavItem eventKey={1} href="/home"><FaUserSecret/> Administrador</NavItem>
    <NavItem eventKey={1} href="/home"><FaWechat/> Usuarios</NavItem>
    <NavItem eventKey={1} href="/home"><FaFilePdfO /> Relatorios</NavItem>
    <NavItem eventKey={1} href="/home"><FaPieChart/> Graphs</NavItem>
    <NavItem eventKey={1} href="/home"><FaSignOut/> Sair</NavItem>
  </Nav>
			);
	}
})


var TicketsTable = React.createClass({
getInitialState(){
      return {
          ticketsPendentes: []

      }
},
componentWillMount: function(){
  var that = this;
   axios.get("/ticketsPendentes").then(function(response){
    that.setState({ticketsPendentes: response.data})
  });
},
render(){
return (
		<Table striped bordered condensed hover>
    <thead>
      <tr>
        <th>Id</th>
        <th>Descricao</th>
        <th>Aberto em</th>
        <th>Criado Por</th>
        <th>Status</th>
        <th>Sla</th>
        <th>Departamento</th>
        <th>Problema</th>
         <th>Acoes</th>
      </tr>
    </thead>
    <tbody>
     {this.state.ticketsPendentes.map(function(s){

      return (

        <tr>
          <td>{s.Id}</td>
          <td>{s.Descricao}</td>
          <td>{s.AbertoEm}</td>
          <td>{s.CriadoPor}</td>
          <td>{s.Status}</td>
          <td>{s.Sla}</td>
          <td>{s.Departamento}</td>
          <td>{s.Problema}</td>
          <td><Button>Ver</Button></td>
        </tr>

      );

     })}
      
     </tbody>
  </Table>
		);
}
});
ReactDOM.render(<Layout />, document.getElementById("layout"));
