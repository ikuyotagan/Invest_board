import React from "react";
import { ProSidebar, Menu } from "react-pro-sidebar";
import "./sidebar.scss";
import ChooseStockMenu from "./ChooseStockMenu";
import ChooseValueMenu from "./ChooseValueMenu";
import ChoosePeriodMenu from "./ChoosePeriodMenu";
import { Button } from "react-bootstrap";

const PersonalSidebar = () => {
  return (
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
            <ChooseStockMenu />
            <ChoosePeriodMenu />
            <ChooseValueMenu />
            <Button style={{ marginLeft: "52px", marginTop: "10px" }}>
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
            <ChooseStockMenu />
            <ChooseValueMenu />
            <Button style={{ marginLeft: "52px", marginTop: "10px" }}>
              See the Graph
            </Button>
          </div>
        </Menu>
      </ProSidebar>
    </div>
  );
};

export default PersonalSidebar;
