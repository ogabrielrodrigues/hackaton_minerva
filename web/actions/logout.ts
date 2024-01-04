'use server'

import { revalidatePath } from 'next/cache'

import { cookies } from 'next/headers'

export async function logoutAction() {
  cookies().delete('auth_token')
  cookies().delete('current_employee')
  revalidatePath('/')
}
