import React from "react";

const Cardbody = (props) => {
  return (
    <section className="home" id="home">
      <div className="container">
        <div className="box">
          <img src={props.image} />
          <h5>{props.namabarang} </h5>
          <h6>{props.hargabarang}</h6>
        </div>
      </div>
    </section>
  );
};

export default Cardbody;
