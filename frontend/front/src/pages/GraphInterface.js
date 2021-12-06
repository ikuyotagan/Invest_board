import { React, useState, useEffect } from "react";
import Sidebar from "../components/Sidebar";
import Chart from "../components/Chart";

const GraphInterface = (props) => {
  const [stocks, setStocks] = useState();
  const [chartData, setchartData] = useState();
  const [stockName, setStockName] = useState();

  useEffect(() => {
    const fetchData = async () => {
      const result = await fetch(props.api + "/private/stocks", {
        credentials: "include",
      });

      if (result.ok) {
        const data = await result.json();

        setStockName(data[0].name);

        setStocks(data);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      <Sidebar
        stocks={stocks}
        setchartData={setchartData}
        setStockName={setStockName}
        stockName={stockName}
        api={props.api}
      />
      <Chart chartData={chartData} stockName={stockName} />
    </div>
  );
};
export default GraphInterface;
