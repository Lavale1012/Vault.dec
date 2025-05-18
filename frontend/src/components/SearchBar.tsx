const SearchBar = () => {
  return (
    <div className=" w-100 h-8 bg-slate-700 flex items-center justify-center rounded-md">
      <input
        type="text"
        placeholder="Search..."
        className="bg-slate-800 text-white w-full h-full rounded-md px-2 focus:outline-none"
      />
    </div>
  );
};

export default SearchBar;
