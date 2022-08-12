import React from "react";
import "./Styleabout.css";
function Navbar() {
  return (
    <div>
      <header>
        <a href="#" className="logo">
          screedCare
        </a>
        <div className="bx bx-menu" id="menu-icon"></div>
        <ul className="navbar">
          <li>
            <a href="#home">HOME</a>
          </li>
          <li>
            <a href="#about">ABOUT US</a>
          </li>
          <li>
            <a href="#contact">CONTACT</a>
          </li>
          <li>
            <a href="#blog">BLOG</a>
          </li>
        </ul>

        <div className="icons">
          <a href="#">
            <i className="bx bx-search"></i>
          </a>
        </div>
      </header>
    </div>
  );
}

export default Navbar;