export interface Employee {
  id: number;
  name: string;
  email: string;
  department: string;
  salary: number;
}

export interface CreateEmployeeRequest {
  name: string;
  email: string;
  department: string;
  salary: number;
}