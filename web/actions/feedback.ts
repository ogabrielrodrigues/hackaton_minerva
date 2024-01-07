'use server'

import { api_url } from '@/config/api'
import { GetEmployee } from '@/hooks/get_employee'
import { GetToken } from '@/hooks/get_token'
import { revalidatePath } from 'next/cache'

export async function feedbackAction(_: any, fd: FormData) {
  try {
    const { employee } = await GetEmployee()
    const { token } = await GetToken()

    const body = {
      employee_registry: employee?.registry,
      content: fd.get('content')?.toString(),
    }

    const response = await fetch(`${api_url}/feedback`, {
      method: 'post',
      body: JSON.stringify(body),
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    if (response.status != 201) throw new Error()

    revalidatePath('/')
    return { feedback: 'success' }
  } catch (err) {
    return { feedback: 'error' }
  }
}
