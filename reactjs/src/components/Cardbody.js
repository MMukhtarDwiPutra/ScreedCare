import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';

function Cardbody(props) {
  return (
    <Card style={{ width: '18rem' }}>
      <Card.Img variant="top" src={props.image}  />
      <Card.Body>
        <Card.Title>{props.namabarang}</Card.Title>
        <Card.Text>
          Some quick example text to build on the card title and make up the
          bulk of the card's content.
        </Card.Text>
      </Card.Body>
      <ListGroup className="list-group-flush">
        <ListGroup.Item>{props.hargabarang}</ListGroup.Item>
        <ListGroup.Item>4/5****</ListGroup.Item>
      </ListGroup>
      <Card.Body>
        <Card.Link href="#">Card Link</Card.Link>
        <Card.Link href="#">Another Link</Card.Link>
      </Card.Body>
    </Card>
  );
}

export default Cardbody;