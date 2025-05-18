import SearchBar from "./SearchBar";

const NavBar = () => {
  return (
    <div>
      <div className="bg-slate-500 p-6 flex items-center  justify-between  h-13">
        <div>VAILT.DEV</div>
        <div>
          <SearchBar />
        </div>

        <ul className="flex items-center justify-center gap-4 text-white">
          <li>HOME</li>
          <li>VAULTS</li>
          <li>EXPLORE</li>
        </ul>
        <div className=" p-4 bg-slate-800"></div>
      </div>
    </div>
  );
};

export default NavBar;
