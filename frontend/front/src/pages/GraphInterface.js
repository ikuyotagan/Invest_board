import { React, useState, useEffect } from "react";
import Sidebar from "../components/Sidebar";
import Chart from "../components/Chart";
import { TimeScale } from "chart.js";

const GraphInterface = () => {
  const [stocks, setStocks] = useState([]);
  const [chartData, setchartData] = useState({});

  const fetchData = async () => {
    const result = await fetch("http://localhost:8080/private/stocks", {
      credentials: "include",
    });

    if (result.ok) {
      const data = await result.json();

      setStocks(data);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
      <Sidebar stocks={stocks} setchartData={setchartData} />
      <Chart chartData={chartData} />
    </div>
  );
};
export default GraphInterface;
