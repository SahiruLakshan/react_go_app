import type {
  Employee,
  CreateEmployeeRequest,
} from "../types/employee";

const API_URL = "http://localhost:8080/api";

export async function getEmployees(): Promise<Employee[]> {
  const response = await fetch(`${API_URL}/employees`);

  if (!response.ok) {
    throw new Error("Failed to get employees");
  }

  return response.json();
}

export async function createEmployee(
  employee: CreateEmployeeRequest
): Promise<Employee> {
  const response = await fetch(
    `${API_URL}/employees`,
    {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(employee),
    }
  );

  if (!response.ok) {
    throw new Error("Failed to create employee");
  }

  return response.json();
}

export async function deleteEmployee(
  id: number
): Promise<void> {
  const response = await fetch(
    `${API_URL}/employees/${id}`,
    {
      method: "DELETE",
    }
  );

  if (!response.ok) {
    throw new Error("Failed to delete employee");
  }
}