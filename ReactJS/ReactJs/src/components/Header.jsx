function Header(props) {
  return (
    <>
      <h1>Welcome to {props.name}</h1>
      <nav>
        <a href="Home">Home</a>
        <a href="About">About</a>
        <a href="Contacts">Contacts</a>
      </nav>
    </>
  );
}

export default Header;
