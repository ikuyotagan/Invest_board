import { forwardRef, React } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import { MenuItem, SubMenu } from "react-pro-sidebar";
import {Button} from "react-bootstrap"

const ChoosePeriodSubMenu = (props) => {
  const CustomInput = forwardRef(({ value, onClick }, ref) => (
    <Button style={{color: "#dee2e6"}} variant="secondary" onClick={onClick} ref={ref}>
      {value}
    </Button>
  ));
    
  return (
    <SubMenu title="Choose Period">
      <MenuItem>
        <DatePicker
          selected={props.startDate}
          onChange={(date) => props.setStartDate(date)}
          timeInputLabel="Time:"
          dateFormat="MM/dd/yyyy h:mm aa"
          showTimeInput
          customInput={<CustomInput />}
        />
      </MenuItem>
      <MenuItem>
        <DatePicker
          selected={props.endDate}
          onChange={(date) => props.setEndDate(date)}
          timeInputLabel="Time:"
          dateFormat="MM/dd/yyyy h:mm aa"
          showTimeInput
          customInput={<CustomInput />}
        />
      </MenuItem>
    </SubMenu>
  );
};

export default ChoosePeriodSubMenu;
