import { React, useEffect, useState } from "react";
import { SubMenu, MenuItem } from "react-pro-sidebar";

const ChooseStockMenu = (props) => {
  const [stocksList, setStocksList] = useState();

  const setId = async (e) => {
    e.preventDefault();

    props.setStockId(parseInt(e.currentTarget.id));
    props.setStockName(e.currentTarget.innerText);
  };

  useEffect(() => {
    const stocksList = props.stocks.map((stock) => (
      <MenuItem
        className={
          stock.id === props.stockId ? "grey-background" : "default-background"
        }
        onClick={setId}
        id={stock.id}
        key={stock.id}
      >
        {stock.name}
      </MenuItem>
    ));

    setStocksList(stocksList);
  }, [props.stocks, props.stockId]);

  return <SubMenu title="Choose Stock">{stocksList}</SubMenu>;
};

export default ChooseStockMenu;
