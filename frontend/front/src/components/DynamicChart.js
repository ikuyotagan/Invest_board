import { useState, useEffect, React } from "react";
import { Line } from "react-chartjs-2";
import { CDBContainer } from "cdbreact";

const DynamicChart = () => {
  const [data, setData] = useState({
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
    datasets: [
      {
        label: 'My First dataset',
        fill: false,
        lineTension: 0.1,
        backgroundColor: 'rgba(194, 116, 161, 0.5)',
        borderColor: 'rgb(194, 116, 161)',
        borderCapStyle: 'butt',
        borderDash: [],
        borderDashOffset: 0.0,
        borderJoinStyle: 'miter',
        pointBorderColor: 'rgba(75,192,192,1)',
        pointBackgroundColor: '#fff',
        pointBorderWidth: 1,
        pointHoverRadius: 5,
        pointHoverBackgroundColor: 'rgba(71, 225, 167, 0.5)',
        pointHoverBorderColor: 'rgb(71, 225, 167)',
        pointHoverBorderWidth: 2,
        pointRadius: 1,
        pointHitRadius: 10,
        data: [65, 59, 80, 81, 56, 55, 40]
      }
    ]
  });
  
  useEffect(() => {
    setInterval(function(){
      var oldDataSet = data.datasets[0];
      var newData = [];
      for(var x=0; x< data.labels.length; x++){
        newData.push(Math.floor(Math.random() * 100));
      }
      var newDataSet = {
        ...oldDataSet
      };
      newDataSet.data = newData;
      var newState = {
        ...data,
        datasets: [newDataSet]
      };
      setData(newState);
    }, 5000);
  }, [])

  return (
    <CDBContainer>
      <h3 className="mt-5">Dynamicly Refreshed Line chart</h3>
      <Line data={data} options={{ responsive: true }} />
    </CDBContainer>
  );
}
export default DynamicChart;