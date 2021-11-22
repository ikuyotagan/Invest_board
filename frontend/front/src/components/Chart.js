import {useState, React, useEffect} from "react";
import { Line } from "react-chartjs-2";
import { CDBContainer } from "cdbreact";

const Chart = (props) => {
  const [data, setData] = useState({
    labels: [
      "Eating",
      "Drinking",
      "Sleeping",
      "Designing",
      "Coding",
      "Cycling",
      "Running",
    ],
    datasets: [
      {
        label: "Tinkoff",
        backgroundColor: "rgba(194, 116, 161, 0.5)",
        borderColor: "rgb(194, 116, 161)",
        data: [65, 59, 90, 81, 56, 55, 40],
      },
    ],
  });

  useEffect(() => {
    setData({
      labels: [
        "Eating",
        "Drinking",
        "Sleeping",
        "Designing",
        "Coding",
        "Cycling",
        "Running",
      ],
      datasets: [
        {
          label: "Tinkoff",
          backgroundColor: "rgba(194, 116, 161, 0.5)",
          borderColor: "rgb(194, 116, 161)",
          data: [65, 59, 90, 81, 56, 55, 40],
        },
      ],
    });
  }, [props.ChartData]);

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
        <h3 className="mt-5">Tinkoff</h3>
        <Line data={data} options={{ responsive: true }} />
      </CDBContainer>
    </div>
  );
};

export default Chart;
