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
    if (props.chartData !== undefined) {
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
    }
  }, [props.chartData, props.stockName]);

  return (
    <div
      style={{
        position: "relative",
        left: "245px",
        top: "-99%",
        width: "65%",
        height: "55%",
      }}
    >
      <CDBContainer style={{paddingTop: "2px", paddingLeft: "15px"}}>
        <h3>{props.stockName}</h3>
        <Line data={chart} options={{ responsive: true }} />
      </CDBContainer>
    </div>
  );
};

export default Chart;
