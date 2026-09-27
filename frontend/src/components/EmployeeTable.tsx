import type { Employee } from "../types/employee";

interface EmployeeTableProps {
  employees: Employee[];
  onDelete: (id: number) => Promise<void>;
}

function EmployeeTable({
  employees,
  onDelete,
}: EmployeeTableProps) {

  return (
    <table>

      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Email</th>
          <th>Department</th>
          <th>Salary</th>
          <th>Action</th>
        </tr>
      </thead>

      <tbody>

        {employees.map((employee) => (

          <tr key={employee.id}>

            <td>{employee.id}</td>

            <td>{employee.name}</td>

            <td>{employee.email}</td>

            <td>{employee.department}</td>

            <td>
              {employee.salary.toLocaleString()}
            </td>

            <td>
              <button
                onClick={() =>
                  onDelete(employee.id)
                }
              >
                Delete
              </button>
            </td>

          </tr>

        ))}

      </tbody>

    </table>
  );
}

export default EmployeeTable;