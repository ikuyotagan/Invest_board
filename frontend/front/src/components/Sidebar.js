import { React, useState, useEffect } from "react";
import { ProSidebar, Menu } from "react-pro-sidebar";
import "./sidebar.scss";
import ChooseStockMenu from "./ChooseStockMenu";
import ChooseValueMenu from "./ChooseValueMenu";
import ChoosePeriodMenu from "./ChoosePeriodMenu";
import { Button } from "react-bootstrap";

const Sidebar = (props) => {
  const [startDate, setStartDate] = useState(() => {
    let d = new Date();
    d.setFullYear(d.getFullYear() - 1);
    return d;
  });
  const [endDate, setEndDate] = useState(new Date());
  const [stockId, setStockId] = useState(1);
  const [value, setValue] = useState("Open Price");
  const [stockName, setStockName] = useState(props.stockName);

  useEffect(() => {
    submit();
  }, []);

  useEffect(() => {
    if (props.stockName !== "") {
      setStockName(props.stockName);
      console.log(")))");
    }
  }, [props.stockName]);

  const submit = async () => {
    const response = await fetch("/api/private/candels", {
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

    if (content.length > 0) {
      if (stockName !== "") {
        props.setStockName(stockName);
      }

      let listCandels = [];
      for (let i = 0; i < content.length; i++) {
        let candel;
        switch (value) {
          case "Open price":
            candel = {
              time: new Date(content[i].time).toLocaleString(),
              price: content[i].open_price,
            };
            break;
          case "Close Price":
            candel = {
              time: new Date(content[i].time).toLocaleString(),
              price: content[i].close_price,
            };
            break;
          case "Highest Price":
            candel = {
              time: new Date(content[i].time).toLocaleString(),
              price: content[i].highest_price,
            };
            break;
          case "Lowest Price":
            candel = {
              time: new Date(content[i].time).toLocaleString(),
              price: content[i].lowest_price,
            };
            break;
          case "Traiding Volume":
            candel = {
              time: new Date(content[i].time).toLocaleString(),
              price: content[i].volume,
            };
            break;
          default:
            candel = {
              time: new Date(content[i].time).toLocaleString(),
              price: content[i].open_price,
            };
            break;
        }
        listCandels.push(candel);
      }

      props.setchartData(listCandels);
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
            <ChooseStockMenu
              stocks={props.stocks}
              setStockId={setStockId}
              stockId={stockId}
              setStockName={setStockName}
            />
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
