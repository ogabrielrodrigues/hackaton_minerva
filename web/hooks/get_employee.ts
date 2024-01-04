import { Employee } from '@/types/employee'
import { cookies } from 'next/headers'

export async function GetEmployee() {
  const employee = cookies().get('current_employee')?.value

  if (!employee) {
    return { employee: null }
  }

  return { employee: JSON.parse(employee) as Employee }
}
