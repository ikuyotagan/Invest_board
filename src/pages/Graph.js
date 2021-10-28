import React, { useEffect, useState } from "react";
import { Line } from "react-chartjs-2";

const Graph = () => {
  const [chartData, setData] = useState({});

  const fetchData = async () => {
    const result = await fetch("http://localhost:8080/api/tinkoff");
    const data = await result.json();

    setData(data);
  };

  useEffect(() => {
    fetchData();
  }, []);

  const tinkoff = () => {
    const labels = [];
    const data = [];

    for (let i = 0; i < chartData.length; i++) {
      data.push(chartData[i])
      labels.push(i.toString());
    }

    return {
      labels,
      datasets: [
        {
          label: "Ахуеть",
          data,
          borderWidth: 1,
        },
      ],
    };
  };

  const options = {
    scales: {
      y: {
        beginAtZero: true,
      },
    },
  };

  return (
    <form>
      <Line data={tinkoff()} options={options} />
    </form>
  );
};
export default Graph;
