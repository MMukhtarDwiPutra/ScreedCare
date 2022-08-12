import React from "react";
import Cardbody from "./Cardbody";
import Navbar from "./Navbar";


class Body extends React.Component {
 
  // Constructor 
  constructor(props) {
      super(props);
 
      this.state = {
          items: []
      };
  }
 
  // ComponentDidMount is used to
  // execute the code 
  componentDidMount() {
      fetch("http://localhost:8080/api/product/")
          .then((r) => r.json())
          .then((json) => {
              this.setState({
                  items: json
              });
          })
  }
  render() {
    const { items } = this.state;

    return (
      <div>
        <Navbar />
        <div>
          <div>
            <div style={{ textAlign: "center", marginTop: "200", justifyContent: "center" }}>
              <h1>TOP Rekomendasi Masker </h1>
            </div>
            </div>
              <div style={{ justifyContent: "center" }}>
                {
                  items.map((item) => (
                <div style={{ col: "4" }}>
                  <Cardbody image={item.foto} namabarang={item.nama_barang} hargabarang={item.harga_str} />
                </div>
                ))
              }
              </div>
        </div>
      </div>
    )
  }
}
export default Body;
