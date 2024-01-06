'use server'

import { api_url } from '@/config/api'
import { GetToken } from '@/hooks/get_token'
import { revalidatePath } from 'next/cache'

export async function answerAction(_: any, fd: FormData) {
  try {
    const body = {
      feedback_id: fd.get('feedback_id')?.toString(),
      content: fd.get('answer')?.toString(),
    }

    const { token } = await GetToken()

    const response = await fetch(`${api_url}/answer/reply`, {
      method: 'post',
      body: JSON.stringify(body),
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    if (response.status != 200) throw new Error()

    revalidatePath('/')
    return { answer: 'success' }
  } catch (err) {
    return { answer: 'error' }
  }
}
