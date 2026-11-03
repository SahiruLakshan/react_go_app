import { useEffect, useState } from "react";

import EmployeeForm from "./components/EmployeeForm";
import EmployeeTable from "./components/EmployeeTable";

import {
  getEmployees,
  createEmployee,
  deleteEmployee,
} from "./api/employeeApi";

import type {
  Employee,
  CreateEmployeeRequest,
} from "./types/employee";

function App() {

  const [employees, setEmployees] =
    useState<Employee[]>([]);

  const [loading, setLoading] =
    useState<boolean>(true);

  const [error, setError] =
    useState<string>("");

  const loadEmployees = async () => {

    try {

      setLoading(true);
      setError("");

      const data = await getEmployees();

      setEmployees(data);

    } catch (error) {

      console.error(error);

      setError(
        "Failed to load employees"
      );

    } finally {

      setLoading(false);
    }
  };

  useEffect(() => {
    loadEmployees();
  }, []);

  const handleEmployeeAdded = async (
    employee: CreateEmployeeRequest
  ): Promise<void> => {

    try {

      await createEmployee(employee);

      await loadEmployees();

    } catch (error) {

      console.error(error);

      setError(
        "Failed to create employee"
      );
    }
  };

  const handleDelete = async (
    id: number
  ): Promise<void> => {

    try {

      await deleteEmployee(id);

      await loadEmployees();

    } catch (error) {

      console.error(error);

      setError(
        "Failed to delete employee"
      );
    }
  };

  return (
    <div className="container">

      <h1>
        Employee Management Dashboard
      </h1>

      <EmployeeForm
        onEmployeeAdded={
          handleEmployeeAdded
        }
      />

      {error && (
        <p className="error">
          {error}
        </p>
      )}

      <hr />

      {loading ? (
        <p>Loading employees...</p>
      ) : (
        <EmployeeTable
          employees={employees}
          onDelete={handleDelete}
        />
      )}

    </div>
  );
}

export default App;