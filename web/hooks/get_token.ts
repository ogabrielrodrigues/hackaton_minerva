import { cookies } from 'next/headers'

export async function GetToken() {
  const token = cookies().get('auth_token')?.value

  if (!token) {
    return { token: null }
  }

  return { token }
}
