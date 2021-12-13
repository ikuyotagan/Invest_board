import { useState, useEffect, React } from "react";
import { ProSidebar, Menu } from "react-pro-sidebar";
import "./sidebar.scss";
import ChooseStockMenu from "./ChooseStockMenu";
import ChooseValueMenu from "./ChooseValueMenu";
import ChoosePeriodMenu from "./ChoosePeriodMenu";
import { Button } from "react-bootstrap";

const PersonalSidebar = (props) => {
  const [startDate, setStartDate] = useState(() => {
    let d = new Date();
    d.setFullYear(d.getFullYear() - 1);
    return d;
  });
  const [endDate, setEndDate] = useState(new Date());
  const [stockId, setStockId] = useState(1);
  const [value, setValue] = useState("Open Price");
  const [stockName, setStockName] = useState();
  const [dynamicData, setDynamicData] = useState();
  const [stopId, setStopId] = useState();
  const [error, setError] = useState(" ");

  useEffect(() => {
    commonGraph();
  }, []);

  useEffect(() => {
    if (props.stockName !== undefined) {
      setStockName(props.stockName);
    }
  }, [props.stockName]);

  const stream = () => {
    props.setchartData([]);
    setDynamicData([]);
    const dynamicGraph = async () => {
      const response = await fetch(props.api + "/private/tinkoff/last_candle", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          id: stockId,
        }),
      });

      props.setStockName(stockName);

      if (response.status === 204) {
        clearInterval(stop);
        setError("stock is closed");
      } else {
        if (error !== " ") {
          setError(" ");
        }
        const content = await response.json();

        if (content !== undefined) {

          let candel;

          switch (value) {
            case "Open price":
              candel = {
                time: new Date(content.time).toLocaleString(),
                price: content.open_price,
              };
              break;
            case "Close Price":
              candel = {
                time: new Date(content.time).toLocaleString(),
                price: content.close_price,
              };
              break;
            case "Highest Price":
              candel = {
                time: new Date(content.time).toLocaleString(),
                price: content.highest_price,
              };
              break;
            case "Lowest Price":
              candel = {
                time: new Date(content.time).toLocaleString(),
                price: content.lowest_price,
              };
              break;
            case "Traiding Volume":
              candel = {
                time: new Date(content.time).toLocaleString(),
                price: content.volume,
              };
              break;
            default:
              candel = {
                time: new Date(content.time).toLocaleString(),
                price: content.open_price,
              };
              break;
          }

          props.setchartData((dynamicData) => [...dynamicData, candel]);

          setDynamicData((dynamicData) => [...dynamicData, candel]);
        }
      }
    };

    dynamicGraph();

    let stop = setInterval(dynamicGraph, 60000);
    setStopId(stop);
  };

  const commonGraph = async () => {
    const response = await fetch(props.api + "/private/candels", {
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
      if (stockName !== undefined) {
        props.setStockName(stockName);
      }
    }
  };

  let errorResponse = (
    <h1
      style={{
        position: "relative",
        color: "red",
        fontSize: "8vw",
        left: "330px",
        top: "-75%",
        zIndex: "4",
      }}
      className="h3 mb-3 fw-normal"
    >
      {error}
    </h1>
  );

  return (
    <div style={{ width: "100%", height: "100%" }}>
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
                onClick={() => {
                  clearInterval(stopId);
                  commonGraph();
                  setError(" ");
                }}
              >
                See the Graph
              </Button>
              <div
                style={{
                  padding: "20px",
                  marginTop: "20px",
                  marginLeft: "0px",
                  color: "white",
                  fontSize: "20px",
                }}
              >
                Real Time Stock
              </div>
              <ChooseStockMenu
                stocks={props.stocks}
                setStockId={setStockId}
                stockId={stockId}
                setStockName={setStockName}
              />
              <ChooseValueMenu value={value} setValue={setValue} />
              <Button
                style={{ marginLeft: "52px", marginTop: "10px" }}
                onClick={() => {
                  clearInterval(stopId);
                  stream();
                  setError(" ");
                }}
              >
                See the Graph
              </Button>
            </div>
          </Menu>
        </ProSidebar>
      </div>
      {errorResponse}
    </div>
  );
};

export default PersonalSidebar;
