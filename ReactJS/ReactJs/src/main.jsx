import { createRoot } from "react-dom/client";
import Header from "./components/Header.jsx";
import MainContent from "./components/Maincontent.jsx";
import Footer from "./components/Footer.jsx";

// import App from "./App.jsx";

createRoot(document.getElementById("root")).render(
  <>
    <Header name="My Website" />
    <MainContent />
    <Footer email="my@email.com"/>
  </>,
);
