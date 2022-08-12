import React, { useState, useEffect } from "react";
import Cardbody from "./Cardbody";


const Body = () => {
  const [items, setItem] = useState([]);
  useEffect(() => {
    fetch("http://localhost:8080/api/product/")
      .then((res) => res.json())
      .then((result) => {
        setItem(result);
      });
  }, []);
  return (
    <div>
      <div className="container">
      <div className="row">
        <div className="col-12 text-center mt-5">
          <h1>TOP Rekomendasi Masker</h1>
        </div>
      </div>
      </div>
    </div>
  );
};

export default Body;
