import { useState, React, useEffect } from "react";
import { Line } from "react-chartjs-2";
import { CDBContainer } from "cdbreact";

const Chart = (props) => {
  const [labels, setLabels] = useState([]);
  const [data, setData] = useState([]);

  useEffect(() => {
    let cLabels = [];
    let cData = [];
    for (let i = 0; i < props.chartData.length; i++) {
      cLabels.push(props.chartData[i].time);
      cData.push(props.chartData[i].openPrice);
    }
    setLabels(cLabels);
    setData(cData);
  }, [props.chartData]);

  useEffect(() => {
    setChart({
      labels: labels,
      datasets: [
        {
          label: props.stockName,
          backgroundColor: "rgba(194, 116, 161, 0.5)",
          borderColor: "rgb(194, 116, 161)",
          data: data,
        },
      ],
    });
  }, [data, props.stockName]);

  const [chart, setChart] = useState({
    labels: labels,
    datasets: [
      {
        label: props.stockName,
        backgroundColor: "rgba(194, 116, 161, 0.5)",
        borderColor: "rgb(194, 116, 161)",
        data: data,
      },
    ],
  });

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
