import React from "react";
import { useNavigate } from "react-router-dom";
import "./Home.css";
// import {imgdd }from  "./src/assets/kashmir-bg.jpg'"

const Home = () => {
  const navigate = useNavigate();

  return (
    <div className="home-container ">
      <div className="overlay"></div>
      <div className="home-content">
        <h1 className="home-title">Welcome to Meeras</h1>
        <p className="home-subtitle">
          Discover and book traditional Kashmiri experiences — from shikara rides, heritage venues, to local crafts and artisans.
        </p>
        <div className="home-buttons">
          <button onClick={() => navigate("/login")} className="home-button">
            Login
          </button>
          <button onClick={() => navigate("/signup")} className="home-button signup">
            Sign Up
          </button>
        </div>
        <div className="home-about">
          <h2>About Meeras</h2>
          <p>
            Meeras is a Kashmir-based platform where tradition meets technology. Whether you’re planning a birthday on Dal Lake, 
            booking a local cafe, or shopping authentic hand-crafted shawls and carpets — Meeras brings Kashmir’s beauty to your fingertips.
          </p>
        </div>
      </div>
    </div>
  );
};

export default Home;
