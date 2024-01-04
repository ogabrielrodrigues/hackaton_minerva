'use server'

import jwt from 'jsonwebtoken'
import { api_url, jwt_secret } from '@/config/api'
import { revalidatePath } from 'next/cache'

import { JwtAuthPayload } from '@/types/jwt'
import { Employee } from '@/types/employee'
import { cookies } from 'next/headers'

export async function loginAction(_: any, fd: FormData) {
  const body = {
    email: fd.get('email')?.toString(),
    password: fd.get('password')?.toString(),
  }

  try {
    const response = await fetch(`${api_url}/employee/auth`, {
      method: 'post',
      body: JSON.stringify(body),
    })

    const token = response.headers.get('Authorization')

    if (!token) throw new Error('token error')

    const payload = jwt.verify(token, jwt_secret, {
      complete: true,
    }).payload as JwtAuthPayload

    const result = await fetch(`${api_url}/employee/${payload.registry}`)

    const employee: Employee = await result.json()
    delete employee.feedbacks

    cookies().set({
      name: 'auth_token',
      value: token,
      path: '/',
      maxAge: 86400,
    })

    cookies().set({
      name: 'current_employee',
      value: JSON.stringify(employee),
      path: '/',
      maxAge: 86400,
    })

    revalidatePath('/')
    return { success: true }
  } catch (err) {
    return { success: false }
  }
}
