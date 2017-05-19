import  Button from "react-bootstrap/lib/Button";
import  React from 'react';
import ReactDOM from 'react-dom'
import { Navbar, Nav, NavItem, NavDropdown, MenuItem, Modal} from 'react-bootstrap/lib';
import { Grid, Col, Row, Table} from 'react-bootstrap/lib';
import axios from 'axios';
import { FaBars ,FaBolt ,FaCheckSquareO, FaSearch, FaUmbrella, FaUserSecret, FaWechat, FaFilePdfO,
FaPieChart, FaSignOut, FaCheck, FaBuildingO, FaClockO, FaExchange, FaFlag, FaEye, FaFileText } from 'react-icons/lib/fa';

var Menu = React.createClass({

render (){
 return ( 
          <Navbar>
           <Navbar.Header>
      <Navbar.Brand>
        <a href="#">Vdesk</a>
      </Navbar.Brand>
    </Navbar.Header>
     <Nav pullRight>
        <NavItem eventKey={1} href="#">Sair</NavItem>
        <NavItem eventKey={2} href="#"></NavItem>
      </Nav>
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
                <MenuRight/>
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

var MenuRight = React.createClass({
  render(){
    return(
  <Nav bsStyle="pills" stacked activeKey={2}>
    <NavItem eventKey={1} href="/home"><FaPieChart/> Graphs</NavItem>
    <NavItem eventKey={1} href="/home"><FaSignOut/> Sair</NavItem>
  </Nav>
      );
  }
})

const columnStyle = {
  columnWidth : '500px'
}
var TicketsTable = React.createClass({
getInitialState(){
      return {
          ticketsPendentes: [],
          ticketsUsuario: []

      }
},
componentWillMount: function(){
  var that = this;
   axios.get("/ticketsPendentes").then(function(response){
    that.setState({ticketsPendentes: response.data})
  });
    axios.get("/ticketsUsuario").then(function(response){
    that.setState({ticketsUsuario: response.data})
  });
},
render(){
return (
  <div>
		<Table striped bordered>
    <thead>
      <tr>
        <th>Id</th>
        <th>Descricao</th>
        <th>Aberto em</th>
        <th>Solicitante</th>
        <th>Status</th>
        <th>Sla</th>
        <th>Departamento</th>
        <th>Problema</th>
         <th style={columnStyle}>Acoes</th>
      </tr>
    </thead>
    <tbody>
     {this.state.ticketsPendentes.map(function(s){

      return (

        <tr>
          <td>{s.Id}</td>
          <td>{s.Descricao}</td>
          <td>{s.AbertoEm}</td>
          <td>{s.Solicitante_id.Nome}</td>
          <td>{s.Status_id.Valor}</td>
          <td>{s.Sla_id.Valor}</td>
          <td>{s.Departamento_id.Nome}</td>
          <td>{s.Problema_id.Valor}</td>
          <td style={columnStyle}><Button bsStyle="success">Pegar</Button><DialogTicket type={1} id={s.Id} descricao={s.Descricao} abertura={s.AbertoEm} status={s.Status_id.Valor} sla={s.Sla_id.Valor} departamento={s.Departamento_id.Nome}/></td>
        </tr>

      );

     })}
      
     </tbody>
  </Table>
  <h4>System table</h4>
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
     {this.state.ticketsUsuario.map(function(s){

      return (

        <tr>
       <td>{s.Id}</td>
          <td>{s.Descricao}</td>
          <td>{s.AbertoEm}</td>
          <td>{s.Solicitante_id.Nome}</td>
          <td>{s.Status_id.Valor}</td>
          <td>{s.Sla_id.Valor}</td>
          <td>{s.Departamento_id.Nome}</td>
          <td>{s.Problema_id.Valor}</td>
          <td><DialogTicket type={2} id={s.Id} descricao={s.Descricao} abertura={s.AbertoEm} status={s.Status_id.Valor} sla={s.Sla_id.Valor} departamento={s.Departamento_id.Nome}/>
          <Button bsStyle="danger">Fechar</Button></td>
        </tr>

      );

     })}
      
     </tbody>
  </Table>
  </div>
		);
}
});


var DialogTicket = React.createClass({
  getInitialState(){
    return {showModal : false};
  },
  close(){
    this.setState({showModal:false});
      this.setState({ disabled: true});

  },
  open(){
       this.setState({showModal:true});
   },
  botoes(){
    const p = this.props.type;

    if (this.props.type == 1){
return <Button bsStyle="success" onClick={this.close} >Pegar</Button>;
       }else{
return <Button bsStyle="warning" onClick={this.close} >Encerrar</Button>;
       }
  },
 
  render (){
      return (
        <div>
        <Button bsStyle="primary" onClick={this.open}><FaEye/></Button>
          <Modal show={this.state.showModal} onHide={this.close}>
          <Modal.Header> Ticket => {this.props.id} </Modal.Header>
          <Grid>
            <Row className="show-grid">  
              <Col md={1}>
               <FaCheck/> => {this.props.id}
              </Col>
               <Col md={2}>
              <FaBuildingO /> =>                {this.props.departamento}
                            </Col>
               <Col md={3}>
               <FaClockO /> => {this.props.abertura}
              </Col>
            </Row>

             <Row className="show-grid"> 
              <Col md={2}>
               <FaExchange /> => {this.props.status}
              </Col>
               <Col md={2}>
             <FaFlag/> => {this.props.sla} 
              </Col>
           
            </Row>

             <Row className="show-grid" > 
               <Col md={4}>
               <br/> <FaFileText/> => {this.props.descricao}
              </Col>
            </Row>
          </Grid>
                 <Modal.Footer> 
                 <Button bsStyle="danger" onClick={this.close} >Fechar</Button>
                  <Button bsStyle="danger" onClick={this.input} >Editar</Button>
                  {this.botoes()}
                 </Modal.Footer> 
          </Modal>
          
          </div>
        );
  }

});

ReactDOM.render(<Layout />, document.getElementById("layout"));
