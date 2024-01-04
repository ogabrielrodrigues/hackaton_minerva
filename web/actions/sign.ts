'use server'

import { api_url } from '@/config/api'

export async function signAction(_: any, fd: FormData) {
  const body = {
    name: fd.get('name')?.toString(),
    email: fd.get('email')?.toString(),
    unit: fd.get('unit')?.toString(),
    sector: fd.get('sector')?.toString(),
    administrator: false,
    password: fd.get('password')?.toString(),
  }

  try {
    const response = await fetch(`${api_url}/employee`, {
      method: 'post',
      body: JSON.stringify(body),
    })

    if (response.status != 201) throw new Error()

    return { success: true }
  } catch (err) {
    return { success: false }
  }
}
