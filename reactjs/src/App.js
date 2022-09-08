
import Navigation from "./components/Navigation.js";
import Maincontent from "./components/Maincontent.js";
import MetaTags from "react-meta-tags";

function App() {
  return (
    <>
      <MetaTags>
      <meta name="google-site-verification" content="VIUZDNIlJdjy1bIIb2wi8wYKIxpCx_12OZHM7eUikCE" />
      </MetaTags>
      <Navigation />
      <Maincontent />
    </>
  );
}

export default App;
