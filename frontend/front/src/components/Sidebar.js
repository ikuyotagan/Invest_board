import { React, useState, useEffect } from "react";
import { ProSidebar, Menu } from "react-pro-sidebar";
import "./sidebar.scss";
import ChooseStockMenu from "./ChooseStockMenu";
import ChooseValueMenu from "./ChooseValueMenu";
import ChoosePeriodMenu from "./ChoosePeriodMenu";
import { Button } from "react-bootstrap";

const Sidebar = (props) => {
  const [stocks, setStocks] = useState(props.stocks);
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(new Date());
  const [stockId, setStockId] = useState(1);
  const [value, setValue] = useState("Open Price");

  useEffect(() => {
    setStocks(props.stocks)
    console.log(stockId);
    console.log(value)
    console.log(startDate);
    console.log(endDate);
  }, [props.stocks, stockId, value, startDate, endDate]);

  const submit = async (e) => {
    e.preventDefault();

    const response = await fetch("http://localhost:8080/private/candels", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        start: startDate,
        end: endDate,
        stock_id: stockId,
      }),
    });

    const content = await response.json();

    if (content.ok) {
      props.setchartData(content);
    }
  };

  return (
    <div className="Sidebar">
      <ProSidebar>
        <Menu>
          <div
            style={{
              padding: "20px",
              marginTop: "-15px",
              marginLeft: "-8px",
              color: "white",
              fontSize: "20px",
            }}
          >
            Stocks
          </div>
          <div
            style={{
              marginLeft: "-8px",
            }}
          >
            <ChooseStockMenu stocks={stocks} setStockId={setStockId} stockId={stockId}/>
            <ChoosePeriodMenu
              startDate={startDate}
              endDate={endDate}
              setStartDate={setStartDate}
              setEndDate={setEndDate}
            />
            <ChooseValueMenu value={value} setValue={setValue} />
            <Button
              style={{ marginLeft: "52px", marginTop: "10px" }}
              onClick={submit}
            >
              See the Graph
            </Button>
          </div>
        </Menu>
      </ProSidebar>
    </div>
  );
};

export default Sidebar;
