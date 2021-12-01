import { useState, React, useEffect } from "react";
import { Line } from "react-chartjs-2";
import { CDBContainer } from "cdbreact";

const Chart = (props) => {
  const [chart, setChart] = useState({
    labels: "",
    datasets: [
      {
        label: props.stockName,
        backgroundColor: "rgba(194, 116, 161, 0.5)",
        borderColor: "rgb(194, 116, 161)",
        data: "",
      },
    ],
  });

  useEffect(() => {
    let Labels = [];
    let Data = [];
    for (let i = 0; i < props.chartData.length; i++) {
      Labels.push(props.chartData[i].time);
      Data.push(props.chartData[i].price);
    }
    setChart({
      labels: Labels,
      datasets: [
        {
          label: props.stockName,
          backgroundColor: "rgba(194, 116, 161, 0.5)",
          borderColor: "rgb(194, 116, 161)",
          data: Data,
        },
      ],
    });
  }, [props.chartData]);

  return (
    <div
      style={{
        marginTop: "-10px",
        height: "700px",
        position: "fixed",
        left: "246px",
        top: "50px",
        width: "800px",
      }}
    >
      <CDBContainer>
        <h3 className="mt-5">{props.stockName}</h3>
        <Line data={chart} options={{ responsive: true }} />
      </CDBContainer>
    </div>
  );
};

export default Chart;
