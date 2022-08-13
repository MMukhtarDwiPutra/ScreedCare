import React, { useState } from "react";
import "./Styleabout.css";
import Navbar from 'react-bootstrap/Navbar';
import NavDropdown from 'react-bootstrap/NavDropdown';

function Navigation() {
  const [isOpen = false, setIsOpen] = useState(false);
  const handleClick = () => {
    if (isOpen == false){
      setIsOpen(!isOpen);
    }else{
      setIsOpen(!isOpen);
    }
  }

  return (
    <div>
      <header>
        <a href="#" className="logo">
          screedCare
        </a>

        <div className="bx bx-menu" id="menu-icon" onClick={handleClick}>
        </div>
        {isOpen && (
        <div class="dropdown-content">
          <ul>
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
        </div>
        )}
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

export default Navigation;