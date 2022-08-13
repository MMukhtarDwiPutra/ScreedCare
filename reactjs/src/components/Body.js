import React, { useState, useEffect } from "react";
import Cardbody from "./Cardbody";
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

const Body = () => {
  const [items, setItem] = useState([]);
  useEffect(() => {
    fetch("http://192.168.1.7:8080/api/product/")
      .then((res) => res.json())
      .then((result) => {
        setItem(result.data);
      });
  }, []);
  return (
    <div>
      <div className="container">
        <div className="col-12 text-center" style={{marginTop: "150px"}}>
          <h1>TOP Rekomendasi Masker</h1>
        </div>
        <Row xs="auto">
          {items.map((item) => (
            <Col my-5><Cardbody image={item.foto} namabarang={item.nama_barang} hargabarang={item.harga_str} /></Col>
          ))}
        </Row>
      </div>
    </div>
  )
}

export default Body;
