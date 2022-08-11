import React from "react";
import Cardbody from "./Cardbody";
import Navbar from "./Navbar";
import card1 from "./../assets/img/3.jpg";
import card2 from "./../assets/img/b1.jfif";
import card3 from "./../assets/img/b2.jfif";

function Body() {
  return (
    <div>
      <Navbar />

      <div style={{ container }}>
        <div style={{ row }}>
          <div style={{ textAlign: "center", marginTop: "200", justifyContent: "center" }}>
            <h1>TOP Rekomendasi Masker</h1>
          </div>
        </div>

        <div style={{ row, justifyContent: "center" }}>
          <div style={{ col: "4" }}>
            <Cardbody image={card1} namabarang="SuryaProMerah" hargabarang="Rp.26.000" />
          </div>
        </div>
      </div>
    </div>
  );
}
export default Body;
