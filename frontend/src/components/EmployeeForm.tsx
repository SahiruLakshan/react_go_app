import { useState } from "react";

import type { CreateEmployeeRequest } from "../types/employee";

interface EmployeeFormProps {
  onEmployeeAdded: (
    employee: CreateEmployeeRequest
  ) => Promise<void>;
}

function EmployeeForm({
  onEmployeeAdded,
}: EmployeeFormProps) {

  const [form, setForm] =
    useState<CreateEmployeeRequest>({
      name: "",
      email: "",
      department: "",
      salary: 0,
    });

  const handleChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {

    const { name, value } = event.target;

    setForm((previous) => ({
      ...previous,
      [name]:
        name === "salary"
          ? Number(value)
          : value,
    }));
  };

  const handleSubmit = async (
    event: React.FormEvent<HTMLFormElement>
  ) => {

    event.preventDefault();

    await onEmployeeAdded(form);

    setForm({
      name: "",
      email: "",
      department: "",
      salary: 0,
    });
  };

  return (
    <form onSubmit={handleSubmit}>

      <input
        name="name"
        placeholder="Name"
        value={form.name}
        onChange={handleChange}
        required
      />

      <input
        name="email"
        type="email"
        placeholder="Email"
        value={form.email}
        onChange={handleChange}
        required
      />

      <input
        name="department"
        placeholder="Department"
        value={form.department}
        onChange={handleChange}
      />

      <input
        name="salary"
        type="number"
        placeholder="Salary"
        value={form.salary || ""}
        onChange={handleChange}
      />

      <button type="submit">
        Add Employee
      </button>

    </form>
  );
}

export default EmployeeForm;