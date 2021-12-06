import { useEffect, useState, React } from "react";

import PersonalSidebar from "../components/PersonalSidebar";
import Chart from "../components/Chart";

const PersonalGraphInterface = (props) => {
  const [stocks, setStocks] = useState();
  const [chartData, setchartData] = useState();
  const [stockName, setStockName] = useState();

  useEffect(() => {
    const fetchData = async () => {
      const result = await fetch(props.api + "/private/tinkoff/personal_stocks", {
        credentials: "include",
      });

      if (result.ok) {
        const data = await result.json();

        let stockList = [];

        for (let i = 0; i < data.length; i++) {
          const stock = {
            id: data[i].stock_id,
            name: data[i].stock_name,
          };
          stockList.push(stock);
        }

        setStockName(stockList[0].name);

        setStocks(stockList);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      <PersonalSidebar
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
export default PersonalGraphInterface;
