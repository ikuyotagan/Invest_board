import {useEffect, useState, React} from "react";
import { MenuItem, SubMenu } from "react-pro-sidebar";

const ChooseValueMenu = (props) => {
  const [valueList, setValueList] = useState();

  const setValue = async (e) => {
    e.preventDefault();

    props.setValue(e.currentTarget.innerText);
  };

  useEffect(() => {
    const valueList = (
      <div>
        <MenuItem
          className={
            "Open Price" == props.value
              ? "grey-background"
              : "default-background"
          }
          onClick={setValue}
        >
          Open Price
        </MenuItem>
        <MenuItem
          className={
            "Close Price" == props.value
              ? "grey-background"
              : "default-background"
          }
          onClick={setValue}
        >
          Close Price
        </MenuItem>
        <MenuItem
          className={
            "Highest Price" == props.value
              ? "grey-background"
              : "default-background"
          }
          onClick={setValue}
        >
          Highest Price
        </MenuItem>
        <MenuItem
          className={
            "Lowest Price" == props.value
              ? "grey-background"
              : "default-background"
          }
          onClick={setValue}
        >
          Lowest Price
        </MenuItem>
        <MenuItem
          className={
            "Stock Value" == props.value
              ? "grey-background"
              : "default-background"
          }
          onClick={setValue}
        >
          Stock Value
        </MenuItem>
      </div>
    );

    setValueList(valueList);
  }, [props.value]);

  return <SubMenu title="Choose Value">{valueList}</SubMenu>;
};

export default ChooseValueMenu;
