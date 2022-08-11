import React from "react";

function Footer() {
  return (
    <div>
      <section className="contact" id="contact">
        <div className="main-contact">
          <div className="contact-content">
            <li>
              <a href="#home">Home</a>
            </li>
            <li>
              <a href="#about">About Us</a>
            </li>
            <li>
              <a href="#contact">Contact</a>
            </li>
            <li>
              <a href="#blog">Blog</a>
            </li>
          </div>

          <div className="contact-content">
            <li>
              <a href="#">Shipping & Returns</a>
            </li>
            <li>
              <a href="#">Store Policy</a>
            </li>
            <li>
              <a href="#">Payment Methods</a>
            </li>
          </div>

          <div className="contact-content">
            <li>
              <a href="#">Costumer Service</a>
            </li>
            <li>
              <a href="#">+62 851-0000-0000</a>
            </li>
          </div>

          <div className="contact-content">
            <li>
              <a href="#">Twitter</a>
            </li>
            <li>
              <a href="#">Facebook</a>
            </li>
            <li>
              <a href="#">Instagram</a>
            </li>
          </div>
        </div>

        <div className="action">
          <form>
            <input type="email" name="email" placeholder="Your Email" required />
            <input type="submit" name="submit" value="Submit" required />
          </form>
        </div>
      </section>

      <div className="last">
        <p>screedCare</p>
      </div>
    </div>
  );
}

export default Footer;
