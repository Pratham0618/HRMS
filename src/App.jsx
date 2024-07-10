import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import Layout from './components/shared/Layout'
import Dashboard from './components/Dashboard';
import EmployeeList from './components/Employee-List';
import Departmentsection from './components/Department-Section';
import LeaveTypes from './components/Leave-Types';
import LeaveManagement from './components/Leave-Management';
import Settings from './components/Settings';
import Login from './components/Login';

function App() {
  return (
    <Router>
    <Routes>
      <Route index element={<Login />} />
      <Route path="/" element={<Layout />}>
      <Route path="dashboard" element={<Dashboard />}></Route>
      <Route path="employeelist" element={<EmployeeList />}></Route>
      <Route path="departmentsection" element={<Departmentsection />}></Route>
      <Route path="leavetypes" element={<LeaveTypes />}></Route>
      <Route path="leavemanagement" element={<LeaveManagement />}></Route>
      <Route path="settings" element={<Settings />}></Route>
      <Route path="/" element={<Login />}></Route>
      </Route>
    </Routes>
    </Router>
  );
}

export default App;