import { cookies } from 'next/headers'

export async function GetFilter() {
  const filter = cookies().get('feedback_filter')?.value

  if (!filter) {
    return { filter: null }
  }

  return { filter }
}
